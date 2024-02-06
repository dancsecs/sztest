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
	var differencesFound = false

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
	stringify func(any) string,
) string {
	var differencesFound = false

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
	var r []string
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
			r = append(r, dFmt.same(lnNbr, lnNbr, s))
		}
		return r
	}
	if lenGotSlice == 0 {
		for lnNbr, s := range wntSlice {
			r = append(r, dFmt.justWnt(lnNbr, s))
		}
		*changed = true
		return r
	}
	if lenWntSlice == 0 {
		for lnNbr, s := range gotSlice {
			r = append(r, dFmt.justGot(lnNbr, s))
		}
		*changed = true
		return r
	}

	// biggest matching window between slices
	o, n, c := bestNextRun(gotSlice, wntSlice, minRunSlice, cmp)

	if c == 0 {
		// nothing matched.  All changed
		*changed = true
		switch {
		case lenGotSlice < lenWntSlice:
			for i := 0; i < lenGotSlice; i++ {
				r = append(
					r,
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
				r = append(r, dFmt.justWnt(i, wntSlice[i]))
			}
		case lenGotSlice > lenWntSlice:
			for i := 0; i < lenWntSlice; i++ {
				r = append(
					r,
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
				r = append(r, dFmt.justGot(i, gotSlice[i]))
			}
		default: // lengths are equal
			for i := 0; i < lenGotSlice; i++ {
				r = append(
					r,
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
		return r
	}
	// Return a composition of the largest matching segment prefixed
	// by recursively calling DiffString with the two strings prefix
	// before the matching section and suffixed by the recursive call
	// to diff string once again.

	// Check lines before largest identical section

	r = append(r, DiffSlice(
		gotSlice[:o],
		wntSlice[:n],
		dFmt,
		changed,
		minRunSlice,
		minRunString,
		cmp,
	)...)

	// Add unchanged lines
	for i := 0; i < c; i++ {
		r = append(r, dFmt.same(i+o, i+n, gotSlice[i+o]))
	}

	// Check lines after largest identical section
	var endOld []T
	var endNew []T
	if o+c < lenGotSlice {
		endOld = gotSlice[o+c:]
	}
	if n+c < lenWntSlice {
		endNew = wntSlice[n+c:]
	}
	r = append(r, DiffSlice(
		endOld,
		endNew,
		dFmt.newOffset(o+c, n+c),
		changed,
		minRunSlice,
		minRunString,
		cmp,
	)...)
	return r
}

// DiffString checks two strings for differences.
func DiffString(gotStr, wntStr string, dType diffType, minRun int) string {
	if gotStr == wntStr {
		// unchanged
		return gotStr
	}
	switch dType {
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
	o, n, c := bestNextRunString(gotStr, wntStr, minRun)

	if c == 0 {
		// nothing matched.  All changed
		return markAsChg(gotStr, wntStr, dType)
	}
	// Return a composition of the largest matching segment prefixed and suffixed
	// by recursively calling DiffString with prefix and suffix strings
	// respectively.
	return "" +
		DiffString(gotStr[:o], wntStr[:n], dType, minRun) +
		gotStr[o:o+c] +
		DiffString(gotStr[o+c:], wntStr[n+c:], dType, minRun)
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
				c := 1
				cGotIdx := gotIdx + 1
				cWntIdx := wntIdx + 1
				for cGotIdx < gotMax &&
					cWntIdx < wntMax &&
					cmp(gotSlice[cGotIdx], wntSlice[cWntIdx]) {
					//
					c++
					cGotIdx++
					cWntIdx++
				}
				if c >= minRun {
					scores = append(scores, match{
						gotStart: gotIdx,
						wntStart: wntIdx,
						length:   c,
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
		m1 := scores[i]
		m2 := scores[j]

		if m1.length == m2.length {
			if m1.gotStart == m2.gotStart {
				return m1.wntStart <= m2.wntStart
			}
			return m1.gotStart <= m2.gotStart
		}
		return m1.length >= m2.length
	})

	if len(scores) == 0 {
		return 0, 0, 0
	}
	return scores[0].gotStart, scores[0].wntStart, scores[0].length
}
