/*
   Golang test helper library: sztest.
   Copyright (C) 2023, 2024 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sztest

import (
	"fmt"
	"sort"
	"strings"
)

type diffType rune

// Represents the type of output displayed when diffing strings.
const (
	DiffWant  = diffType('W')
	DiffGot   = diffType('G')
	DiffMerge = diffType('M')
)

// Default compare function.
func defaultCmpFunc[T chkType](got, want T) bool {
	return got == want
}

// CompareArrays returns "" an empty string if there are no differences
// otherwise it returns a string outlining the differences.
func CompareArrays[T chkType](got, wnt []T) string {
	differencesFound := false

	ret := "" +
		strings.Join(
			DiffSlice(
				got,
				wnt,
				newDiffLnFmt(len(got), len(wnt)),
				&differencesFound,
				settingDiffSlice,
				settingDiffChars,
				defaultCmpFunc[T],
			),
			"\n",
		)

	if !differencesFound {
		ret = ""
	}

	return ret
}

// CompareSlices checks two slices for differences.
func CompareSlices[T chkType](
	title string,
	got, want []T,
	minRunSlice, minRunString int,
	cmp func(a, b T) bool,
	_ func(any) string,
) string {
	differencesFound := false

	ret := "" +
		strings.Join(
			DiffSlice(
				got,
				want,
				newDiffLnFmt(len(got), len(want)),
				&differencesFound,
				minRunSlice,
				minRunString,
				cmp,
			),
			"\n",
		)

	if differencesFound {
		if title != "" {
			title += ": "
		}
		ret = fmt.Sprint(
			title,
			"got (", len(got), " lines)",
			" - ",
			"want (", len(want), " lines)",
			"\n",
		) + ret

		return ret
	}

	return ""
}

// DiffSlice compares two Slices.
//
//nolint:funlen // ok
func DiffSlice[T chkType](
	gotSlice, wntSlice []T,
	dFmt *diffLnFmt,
	changed *bool,
	minRunSlice int,
	minRunString int,
	cmp func(a, b T) bool,
) []string {
	var result []string
	lenGotSlice := len(gotSlice)
	lenWntSlice := len(wntSlice)
	same := lenGotSlice == lenWntSlice

	for i := 0; i < lenGotSlice && same; i++ {
		if !cmp(gotSlice[i], wntSlice[i]) {
			same = false
		}
	}
	if same {
		// Just return a clean unmarked list with changed equal to false.
		for lnNbr, s := range gotSlice {
			result = append(result, dFmt.same(lnNbr, lnNbr, s))
		}

		return result
	}
	if lenGotSlice == 0 {
		for lnNbr, s := range wntSlice {
			result = append(result, dFmt.justWnt(lnNbr, s))
		}
		*changed = true

		return result
	}
	if lenWntSlice == 0 {
		for lnNbr, s := range gotSlice {
			result = append(result, dFmt.justGot(lnNbr, s))
		}
		*changed = true

		return result
	}

	// biggest matching window between slices
	gotIdx, wntInd, numLines := bestNextRun(
		gotSlice, wntSlice, minRunSlice, cmp,
	)

	if numLines == 0 {
		// nothing matched.  All changed
		*changed = true
		switch {
		case lenGotSlice < lenWntSlice:
			for i := 0; i < lenGotSlice; i++ {
				result = append(
					result,
					dFmt.changed(i, i,
						DiffString(
							fmt.Sprint(gotSlice[i]),
							fmt.Sprint(wntSlice[i]),
							DiffMerge,
							minRunString,
						)),
				)
			}
			for i := lenGotSlice; i < lenWntSlice; i++ {
				result = append(result, dFmt.justWnt(i, wntSlice[i]))
			}
		case lenGotSlice > lenWntSlice:
			for i := 0; i < lenWntSlice; i++ {
				result = append(
					result,
					dFmt.changed(i, i,
						DiffString(
							fmt.Sprint(gotSlice[i]),
							fmt.Sprint(wntSlice[i]),
							DiffMerge,
							minRunString,
						)),
				)
			}
			for i := lenWntSlice; i < lenGotSlice; i++ {
				result = append(result, dFmt.justGot(i, gotSlice[i]))
			}
		default: // lengths are equal
			for i := 0; i < lenGotSlice; i++ {
				result = append(
					result,
					dFmt.changed(i, i,
						DiffString(
							fmt.Sprint(gotSlice[i]),
							fmt.Sprint(wntSlice[i]),
							DiffMerge,
							minRunString,
						)),
				)
			}
		}

		return result
	}
	// Return a composition of the largest matching segment prefixed
	// by recursively calling DiffString with the two strings prefix
	// before the matching section and suffixed by the recursive call
	// to diff string once again.

	// Check lines before largest identical section

	result = append(result, DiffSlice(
		gotSlice[:gotIdx],
		wntSlice[:wntInd],
		dFmt,
		changed,
		minRunSlice,
		minRunString,
		cmp,
	)...)

	// Add unchanged lines
	for i := 0; i < numLines; i++ {
		result = append(result, dFmt.same(i+gotIdx, i+wntInd, gotSlice[i+gotIdx]))
	}

	// Check lines after largest identical section
	var endOld []T
	var endNew []T
	if gotIdx+numLines < lenGotSlice {
		endOld = gotSlice[gotIdx+numLines:]
	}
	if wntInd+numLines < lenWntSlice {
		endNew = wntSlice[wntInd+numLines:]
	}
	result = append(result, DiffSlice(
		endOld,
		endNew,
		dFmt.newOffset(gotIdx+numLines, wntInd+numLines),
		changed,
		minRunSlice,
		minRunString,
		cmp,
	)...)

	return result
}

// DiffString checks two strings for differences.
func DiffString(gotStr, wntStr string, dType diffType, minRun int) string {
	if gotStr == wntStr {
		// unchanged
		return gotStr
	}
	switch dType { //nolint:exhaustive // Default handles all other cases.
	case DiffWant:
		if wntStr == "" {
			// missing from section
			return wntStr
		}
		if gotStr == "" {
			// new to section
			return markAsDel(wntStr)
		}
	case DiffGot:
		if gotStr == "" {
			// new to section
			return gotStr
		}
		if wntStr == "" {
			// missing from section
			return markAsIns(gotStr)
		}
	default: // DiffMerge.
		if gotStr == "" {
			// new to section
			return markAsDel(wntStr)
		}
		if wntStr == "" {
			// missing from section
			return markAsIns(gotStr)
		}
	}
	// Biggest matching window between strings.
	gotIdx, wntIdx, numChars := bestNextRunString(gotStr, wntStr, minRun)

	if numChars == 0 {
		// nothing matched.  All changed
		return markAsChg(gotStr, wntStr, dType)
	}
	// Return a composition of the largest matching segment prefixed and suffixed
	// by recursively calling DiffString with prefix and suffix strings
	// respectively.
	return "" +
		DiffString(gotStr[:gotIdx], wntStr[:wntIdx], dType, minRun) +
		gotStr[gotIdx:gotIdx+numChars] +
		DiffString(gotStr[gotIdx+numChars:], wntStr[wntIdx+numChars:], dType, minRun)
}

func bestNextRunString(got, wnt string, minRun int) (int, int, int) {
	return bestNextRun(
		strings.Split(got, ""),
		strings.Split(wnt, ""),
		minRun,
		defaultCmpFunc[string],
	)
}

type match struct {
	gotStart int
	wntStart int
	length   int
}

func calculateScores[V chkType](
	gotSlice, wntSlice []V,
	minRun int,
	cmp func(a, b V) bool,
) []match {
	scores := []match{}
	gotMax := len(gotSlice)
	wntMax := len(wntSlice)
	for gotIdx := range gotSlice {
		for wntIdx := range wntSlice {
			if cmp(gotSlice[gotIdx], wntSlice[wntIdx]) {
				numProcessed := 1
				cGotIdx := gotIdx + 1
				cWntIdx := wntIdx + 1
				for cGotIdx < gotMax &&
					cWntIdx < wntMax &&
					cmp(gotSlice[cGotIdx], wntSlice[cWntIdx]) {
					//
					numProcessed++
					cGotIdx++
					cWntIdx++
				}
				if numProcessed >= minRun {
					scores = append(scores, match{
						gotStart: gotIdx,
						wntStart: wntIdx,
						length:   numProcessed,
					})
				}
			}
		}
	}

	return scores
}

func bestNextRun[V chkType](
	gotSlice, wntSlice []V,
	minRun int,
	cmp func(a, b V) bool,
) (int, int, int) {
	scores := calculateScores(gotSlice, wntSlice, minRun, cmp)

	sort.Slice(scores, func(i int, j int) bool {
		// return true if i < j
		iMatch := scores[i]
		jMatch := scores[j]

		if iMatch.length == jMatch.length {
			if iMatch.gotStart == jMatch.gotStart {
				return iMatch.wntStart <= jMatch.wntStart
			}

			return iMatch.gotStart <= jMatch.gotStart
		}

		return iMatch.length >= jMatch.length
	})

	if len(scores) == 0 {
		return 0, 0, 0
	}

	return scores[0].gotStart, scores[0].wntStart, scores[0].length
}
