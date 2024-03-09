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
	"strings"
	"testing"
)

const (
	minRun1 = 1
	minRun2 = 2
	minRun3 = 3
	minRun4 = 4
	minRun5 = 5
	minRun6 = 6
)

func testDiffPrerequisites(t *testing.T) {
	t.Run("BestNextRunString", testSzTestDiffBestNextRunString)
	t.Run("DiffString", testSzTestDiffString)
	t.Run("DiffSlice", testSzTestDiffSlice)
	t.Run("CompareSlice", testSzTestCompareSlices)
	t.Run("CompareSlicesWithPercent", testSzTestCompareSlicesWithPercent)
	t.Run("CompareArrays", testSzTestCompareArrays)
}

func checkRun(
	t *testing.T,
	str1, str2 string,
	minRun int,
	wntOldIdx, wntNewIdx, wntCount int,
) {
	t.Helper()

	gotOldIdx, gotNewIdx, gotCount := bestNextRunString(str1, str2, minRun)

	if gotOldIdx != wntOldIdx {
		t.Error(errGotWnt("old", gotOldIdx, wntOldIdx))
	}
	if gotNewIdx != wntNewIdx {
		t.Error(errGotWnt("new", gotNewIdx, wntNewIdx))
	}
	if gotCount != wntCount {
		t.Error(errGotWnt("count", gotCount, wntCount))
	}
}

func testSzTestDiffBestNextRunString(t *testing.T) {
	str1 := "MNO"
	str2 := "NOP"
	// Forward.
	checkRun(t, str1, str2, minRun1, 1, 0, 2)
	checkRun(t, str1, str2, minRun2, 1, 0, 2)
	checkRun(t, str1, str2, minRun3, 0, 0, 0)
	// Backward.
	checkRun(t, str2, str1, minRun1, 0, 1, 2)
	checkRun(t, str2, str1, minRun2, 0, 1, 2)
	checkRun(t, str2, str1, minRun3, 0, 0, 0)

	str1 = "HIJ"
	str2 = "IJKHIJ"
	// Forward.
	checkRun(t, str1, str2, minRun1, 0, 3, 3)
	checkRun(t, str1, str2, minRun2, 0, 3, 3)
	checkRun(t, str1, str2, minRun3, 0, 3, 3)
	checkRun(t, str1, str2, minRun4, 0, 0, 0)
	// Backward.
	checkRun(t, str2, str1, minRun1, 3, 0, 3)
	checkRun(t, str2, str1, minRun2, 3, 0, 3)
	checkRun(t, str2, str1, minRun3, 3, 0, 3)
	checkRun(t, str2, str1, minRun4, 0, 0, 0)

	str1 = "STUVWXY"
	str2 = "STUVTUVWXSTU"
	// Forward.
	checkRun(t, str1, str2, minRun1, 1, 4, 5)
	checkRun(t, str1, str2, minRun2, 1, 4, 5)
	checkRun(t, str1, str2, minRun3, 1, 4, 5)
	checkRun(t, str1, str2, minRun4, 1, 4, 5)
	checkRun(t, str1, str2, minRun5, 1, 4, 5)
	checkRun(t, str1, str2, minRun6, 0, 0, 0)
	// Backward.
	checkRun(t, str2, str1, minRun1, 4, 1, 5)
	checkRun(t, str2, str1, minRun2, 4, 1, 5)
	checkRun(t, str2, str1, minRun3, 4, 1, 5)
	checkRun(t, str2, str1, minRun4, 4, 1, 5)
	checkRun(t, str2, str1, minRun5, 4, 1, 5)
	checkRun(t, str2, str1, minRun6, 0, 0, 0)

	str1 = "ABCHG"
	str2 = "BCDEHG"
	// Forward.
	checkRun(t, str1, str2, minRun1, 1, 0, 2)
	checkRun(t, str1, str2, minRun2, 1, 0, 2)
	checkRun(t, str1, str2, minRun3, 0, 0, 0)
	// Backward.
	checkRun(t, str2, str1, minRun1, 0, 1, 2)
	checkRun(t, str2, str1, minRun2, 0, 1, 2)
	checkRun(t, str2, str1, minRun3, 0, 0, 0)
}

type tstDiffString struct {
	got        string
	wnt        string
	minRun     int
	expDiffGot string
	expDiffWnt string
	expDiffMrg string
}

func chkDiffString(
	t *testing.T,
	tst *tstDiffString,
) {
	t.Helper()

	resolvedGot := resolveMarksForDisplay(
		DiffString(tst.got, tst.wnt, DiffGot, tst.minRun),
	)
	got, err := freezeMarks(resolvedGot)
	if err != nil {
		t.Fatal(err)
	}

	if got != tst.expDiffGot {
		t.Error(
			errGotWnt(
				fmt.Sprint("DiffGot minRun: ", tst.minRun),
				got,
				tst.expDiffGot,
			),
		)
	}

	resolvedGot = resolveMarksForDisplay(
		DiffString(tst.got, tst.wnt, DiffWant, tst.minRun),
	)
	got, err = freezeMarks(resolvedGot)
	if err != nil {
		t.Fatal(err)
	}

	if got != tst.expDiffWnt {
		t.Error(
			errGotWnt(
				fmt.Sprint("DiffWant minRun: ", tst.minRun),
				got,
				tst.expDiffWnt,
			),
		)
	}

	resolvedGot = resolveMarksForDisplay(
		DiffString(tst.got, tst.wnt, DiffMerge, tst.minRun),
	)
	got, err = freezeMarks(resolvedGot)
	if err != nil {
		t.Fatal(err)
	}

	if got != tst.expDiffMrg {
		t.Error(
			errGotWnt(
				fmt.Sprint("DiffMerge minRun: ", tst.minRun),
				got,
				tst.expDiffMrg,
			),
		)
	}
}

func testSzTestDiffString(t *testing.T) {
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCD",
		minRun:     1,
		expDiffGot: "⨭A⨮BC",
		expDiffWnt: "BC⨴D⨵",
		expDiffMrg: "⨭A⨮BC⨴D⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCD",
		minRun:     2,
		expDiffGot: "⨭A⨮BC",
		expDiffWnt: "BC⨴D⨵",
		expDiffMrg: "⨭A⨮BC⨴D⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCD",
		minRun:     3,
		expDiffGot: "«ABC»",
		expDiffWnt: "«BCD»",
		expDiffMrg: "⨴BCD⨵⧚/⧛⨭ABC⨮",
	})

	chkDiffString(t, &tstDiffString{
		got:        "BCD",
		wnt:        "ABC",
		minRun:     1,
		expDiffGot: "BC⨭D⨮",
		expDiffWnt: "⨴A⨵BC",
		expDiffMrg: "⨴A⨵BC⨭D⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCD",
		wnt:        "ABC",
		minRun:     2,
		expDiffGot: "BC⨭D⨮",
		expDiffWnt: "⨴A⨵BC",
		expDiffMrg: "⨴A⨵BC⨭D⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCD",
		wnt:        "ABC",
		minRun:     3,
		expDiffGot: "«BCD»",
		expDiffWnt: "«ABC»",
		expDiffMrg: "⨴ABC⨵⧚/⧛⨭BCD⨮",
	})

	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCDABC",
		minRun:     1,
		expDiffGot: "ABC",
		expDiffWnt: "⨴BCD⨵ABC",
		expDiffMrg: "⨴BCD⨵ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCDABC",
		minRun:     2,
		expDiffGot: "ABC",
		expDiffWnt: "⨴BCD⨵ABC",
		expDiffMrg: "⨴BCD⨵ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCDABC",
		minRun:     3,
		expDiffGot: "ABC",
		expDiffWnt: "⨴BCD⨵ABC",
		expDiffMrg: "⨴BCD⨵ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABC",
		wnt:        "BCDABC",
		minRun:     4,
		expDiffGot: "«ABC»",
		expDiffWnt: "«BCDABC»",
		expDiffMrg: "⨴BCDABC⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCDABC",
		wnt:        "ABC",
		minRun:     1,
		expDiffGot: "⨭BCD⨮ABC",
		expDiffWnt: "ABC",
		expDiffMrg: "⨭BCD⨮ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCDABC",
		wnt:        "ABC",
		minRun:     2,
		expDiffGot: "⨭BCD⨮ABC",
		expDiffWnt: "ABC",
		expDiffMrg: "⨭BCD⨮ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCDABC",
		wnt:        "ABC",
		minRun:     3,
		expDiffGot: "⨭BCD⨮ABC",
		expDiffWnt: "ABC",
		expDiffMrg: "⨭BCD⨮ABC",
	})
	chkDiffString(t, &tstDiffString{
		got:        "BCDABC",
		wnt:        "ABC",
		minRun:     4,
		expDiffGot: "«BCDABC»",
		expDiffWnt: "«ABC»",
		expDiffMrg: "⨴ABC⨵⧚/⧛⨭BCDABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     1,
		expDiffGot: "ABCDEF«G»",
		expDiffWnt: "A⨴BCD⨵BCDEF«ABC»",
		expDiffMrg: "A⨴BCD⨵BCDEF⨴ABC⨵⧚/⧛⨭G⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     2,
		expDiffGot: "«A»BCDEF«G»",
		expDiffWnt: "«ABCD»BCDEF«ABC»",
		expDiffMrg: "⨴ABCD⨵⧚/⧛⨭A⨮BCDEF⨴ABC⨵⧚/⧛⨭G⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     3,
		expDiffGot: "«A»BCDEF«G»",
		expDiffWnt: "«ABCD»BCDEF«ABC»",
		expDiffMrg: "⨴ABCD⨵⧚/⧛⨭A⨮BCDEF⨴ABC⨵⧚/⧛⨭G⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     4,
		expDiffGot: "«A»BCDEF«G»",
		expDiffWnt: "«ABCD»BCDEF«ABC»",
		expDiffMrg: "⨴ABCD⨵⧚/⧛⨭A⨮BCDEF⨴ABC⨵⧚/⧛⨭G⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     5,
		expDiffGot: "«A»BCDEF«G»",
		expDiffWnt: "«ABCD»BCDEF«ABC»",
		expDiffMrg: "⨴ABCD⨵⧚/⧛⨭A⨮BCDEF⨴ABC⨵⧚/⧛⨭G⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFG",
		wnt:        "ABCDBCDEFABC",
		minRun:     6,
		expDiffGot: "«ABCDEFG»",
		expDiffWnt: "«ABCDBCDEFABC»",
		expDiffMrg: "⨴ABCDBCDEFABC⨵⧚/⧛⨭ABCDEFG⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     1,
		expDiffGot: "A⨭BCD⨮BCDEF«ABC»",
		expDiffWnt: "ABCDEF«G»",
		expDiffMrg: "A⨭BCD⨮BCDEF⨴G⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     2,
		expDiffGot: "«ABCD»BCDEF«ABC»",
		expDiffWnt: "«A»BCDEF«G»",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABCD⨮BCDEF⨴G⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     3,
		expDiffGot: "«ABCD»BCDEF«ABC»",
		expDiffWnt: "«A»BCDEF«G»",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABCD⨮BCDEF⨴G⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     4,
		expDiffGot: "«ABCD»BCDEF«ABC»",
		expDiffWnt: "«A»BCDEF«G»",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABCD⨮BCDEF⨴G⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     5,
		expDiffGot: "«ABCD»BCDEF«ABC»",
		expDiffWnt: "«A»BCDEF«G»",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABCD⨮BCDEF⨴G⨵⧚/⧛⨭ABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDBCDEFABC",
		wnt:        "ABCDEFG",
		minRun:     6,
		expDiffGot: "«ABCDBCDEFABC»",
		expDiffWnt: "«ABCDEFG»",
		expDiffMrg: "⨴ABCDEFG⨵⧚/⧛⨭ABCDBCDEFABC⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFMNOPQUVWXY",
		wnt:        "ABqBCDEqwNOPqv23stuVWXYz",
		minRun:     1,
		expDiffGot: "ABCDE«FM»NOP«QU»VWXY",
		expDiffWnt: "A⨴Bq⨵BCDE«qw»NOP«qv23stu»VWXY⨴z⨵",
		expDiffMrg: "A⨴Bq⨵BCDE⨴qw⨵⧚/⧛⨭FM⨮NOP⨴qv23stu⨵⧚/⧛⨭QU⨮VWXY⨴z⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFMNOPQUVWXY",
		wnt:        "ABqBCDEqwNOPqv23stuVWXYz",
		minRun:     2,
		expDiffGot: "«A»BCDE«FM»NOP«QU»VWXY",
		expDiffWnt: "«ABq»BCDE«qw»NOP«qv23stu»VWXY⨴z⨵",
		expDiffMrg: "⨴ABq⨵⧚/⧛⨭A⨮BCDE⨴qw⨵⧚/⧛⨭FM⨮NOP⨴qv23stu⨵⧚/⧛⨭QU⨮VWXY⨴z⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFMNOPQUVWXY",
		wnt:        "ABqBCDEqwNOPqv23stuVWXYz",
		minRun:     3,
		expDiffGot: "«A»BCDE«FM»NOP«QU»VWXY",
		expDiffWnt: "«ABq»BCDE«qw»NOP«qv23stu»VWXY⨴z⨵",
		expDiffMrg: "⨴ABq⨵⧚/⧛⨭A⨮BCDE⨴qw⨵⧚/⧛⨭FM⨮NOP⨴qv23stu⨵⧚/⧛⨭QU⨮VWXY⨴z⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFMNOPQUVWXY",
		wnt:        "ABqBCDEqwNOPqv23stuVWXYz",
		minRun:     4,
		expDiffGot: "«A»BCDE«FMNOPQU»VWXY",
		expDiffWnt: "«ABq»BCDE«qwNOPqv23stu»VWXY⨴z⨵",
		expDiffMrg: "⨴ABq⨵⧚/⧛⨭A⨮BCDE⨴qwNOPqv23stu⨵⧚/⧛⨭FMNOPQU⨮VWXY⨴z⨵",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABCDEFMNOPQUVWXY",
		wnt:        "ABqBCDEqwNOPqv23stuVWXYz",
		minRun:     5,
		expDiffGot: "«ABCDEFMNOPQUVWXY»",
		expDiffWnt: "«ABqBCDEqwNOPqv23stuVWXYz»",
		expDiffMrg: "⨴ABqBCDEqwNOPqv23stuVWXYz⨵⧚/⧛⨭ABCDEFMNOPQUVWXY⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABqBCDEqwNOPqv23stuVWXYz",
		wnt:        "ABCDEFMNOPQUVWXY",
		minRun:     1,
		expDiffGot: "A⨭Bq⨮BCDE«qw»NOP«qv23stu»VWXY⨭z⨮",
		expDiffWnt: "ABCDE«FM»NOP«QU»VWXY",
		expDiffMrg: "A⨭Bq⨮BCDE⨴FM⨵⧚/⧛⨭qw⨮NOP⨴QU⨵⧚/⧛⨭qv23stu⨮VWXY⨭z⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABqBCDEqwNOPqv23stuVWXYz",
		wnt:        "ABCDEFMNOPQUVWXY",
		minRun:     2,
		expDiffGot: "«ABq»BCDE«qw»NOP«qv23stu»VWXY⨭z⨮",
		expDiffWnt: "«A»BCDE«FM»NOP«QU»VWXY",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABq⨮BCDE⨴FM⨵⧚/⧛⨭qw⨮NOP⨴QU⨵⧚/⧛⨭qv23stu⨮VWXY⨭z⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABqBCDEqwNOPqv23stuVWXYz",
		wnt:        "ABCDEFMNOPQUVWXY",
		minRun:     3,
		expDiffGot: "«ABq»BCDE«qw»NOP«qv23stu»VWXY⨭z⨮",
		expDiffWnt: "«A»BCDE«FM»NOP«QU»VWXY",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABq⨮BCDE⨴FM⨵⧚/⧛⨭qw⨮NOP⨴QU⨵⧚/⧛⨭qv23stu⨮VWXY⨭z⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABqBCDEqwNOPqv23stuVWXYz",
		wnt:        "ABCDEFMNOPQUVWXY",
		minRun:     4,
		expDiffGot: "«ABq»BCDE«qwNOPqv23stu»VWXY⨭z⨮",
		expDiffWnt: "«A»BCDE«FMNOPQU»VWXY",
		expDiffMrg: "⨴A⨵⧚/⧛⨭ABq⨮BCDE⨴FMNOPQU⨵⧚/⧛⨭qwNOPqv23stu⨮VWXY⨭z⨮",
	})
	chkDiffString(t, &tstDiffString{
		got:        "ABqBCDEqwNOPqv23stuVWXYz",
		wnt:        "ABCDEFMNOPQUVWXY",
		minRun:     5,
		expDiffGot: "«ABqBCDEqwNOPqv23stuVWXYz»",
		expDiffWnt: "«ABCDEFMNOPQUVWXY»",
		expDiffMrg: "⨴ABCDEFMNOPQUVWXY⨵⧚/⧛⨭ABqBCDEqwNOPqv23stuVWXYz⨮",
	})

	chkDiffString(t, &tstDiffString{
		got:        "ABAC",
		wnt:        "AA",
		minRun:     1,
		expDiffGot: "A⨭B⨮A⨭C⨮",
		expDiffWnt: "AA",
		expDiffMrg: "A⨭B⨮A⨭C⨮",
	})

	chkDiffString(t, &tstDiffString{
		got:        "AA",
		wnt:        "ABAC",
		minRun:     1,
		expDiffGot: "AA",
		expDiffWnt: "A⨴B⨵A⨴C⨵",
		expDiffMrg: "A⨴B⨵A⨴C⨵",
	})

	chkDiffString(t, &tstDiffString{
		got:        "sameDIF1sameDIF2sameDIF3sameDIF4",
		wnt:        "samesamesamesame",
		minRun:     4,
		expDiffGot: "same⨭DIF1⨮same⨭DIF2⨮same⨭DIF3⨮same⨭DIF4⨮",
		expDiffWnt: "samesamesamesame",
		expDiffMrg: "same⨭DIF1⨮same⨭DIF2⨮same⨭DIF3⨮same⨭DIF4⨮",
	})

	chkDiffString(t, &tstDiffString{
		got:        "DEF1ABCDEF2ABCDEF3ABCDEF4ABC",
		wnt:        "ABCABCABCABC",
		minRun:     3,
		expDiffGot: "⨭DEF1⨮ABC⨭DEF2⨮ABC⨭DEF3⨮ABC⨭DEF4⨮ABC",
		expDiffWnt: "ABCABCABCABC",
		expDiffMrg: "⨭DEF1⨮ABC⨭DEF2⨮ABC⨭DEF3⨮ABC⨭DEF4⨮ABC",
	})
}

type tstDiffSlice struct {
	got        []string
	wnt        []string
	minRunA    int
	expChanged bool
	expSlice   []string
}

func stringify(v any) string {
	return fmt.Sprintf("%v", v)
}

//nolint:cyclop // Ok.
func chkDiffSlice(t *testing.T, tst *tstDiffSlice) {
	var changed bool
	var got []string
	var wnt []string
	t.Helper()

	diffResult := DiffSlice(
		tst.got, tst.wnt,
		newDiffLnFmt(len(tst.got), len(tst.wnt)),
		&changed,
		tst.minRunA,
		settingDiffChars,
		defaultCmpFunc[string],
	)

	for _, entry := range diffResult {
		for _, line := range strings.Split(entry, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				markedUpLine := resolveMarksForDisplay(line)
				iMarkedUpLine, err := freezeMarks(markedUpLine)
				if err != nil {
					t.Fatal(err)
				}
				got = append(got, iMarkedUpLine)
			}
		}
	}
	for _, entry := range tst.expSlice {
		for _, line := range strings.Split(entry, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				wnt = append(wnt, line)
			}
		}
	}
	if changed != tst.expChanged {
		t.Error(
			"invalid error changed:  got: ", changed, "  wnt: ", tst.expChanged,
		)
	}
	if len(got) != len(wnt) {
		t.Error(errGotWnt("slice length", len(got), len(wnt)))
	}
	if strings.Join(got, "\n") != strings.Join(wnt, "\n") {
		t.Error(
			errGotWnt(
				"slice",
				"\n"+strings.Join(got, "\n"),
				"\n"+strings.Join(wnt, "\n"),
			),
		)
	}
}

func testSzTestDiffSlice(t *testing.T) {
	chkDiffSlice(t, &tstDiffSlice{
		got:        nil,
		wnt:        nil,
		minRunA:    1,
		expChanged: false,
		expSlice:   nil,
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        nil,
		wnt:        []string{"W1"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"-:⨴0⨵ ⨴W1⨵",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        nil,
		wnt:        []string{"W1", "W2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"-:⨴0⨵ ⨴W1⨵",
			"-:⨴1⨵ ⨴W2⨵",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"G1"},
		wnt:        nil,
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"⨭0⨮:- ⨭G1⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"G1", "G2"},
		wnt:        nil,
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"⨭0⨮:- ⨭G1⨮",
			"⨭1⨮:- ⨭G2⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"B1"},
		wnt:        []string{"B1"},
		minRunA:    1,
		expChanged: false,
		expSlice: []string{
			"0:0 B1",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"B1", "B2"},
		wnt:        []string{"B1", "B2"},
		minRunA:    1,
		expChanged: false,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"G1"},
		wnt:        []string{"W1"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"«0»:«0» ⨴W1⨵⧚/⧛⨭G1⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"G1", "G2"},
		wnt:        []string{"W1", "W2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"«0»:«0» ⨴W1⨵⧚/⧛⨭G1⨮",
			"«1»:«1» ⨴W2⨵⧚/⧛⨭G2⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"B1", "G2"},
		wnt:        []string{"B1", "W2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"«1»:«1» ⨴W2⨵⧚/⧛⨭G2⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"G1", "B2"},
		wnt:        []string{"W1", "B2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"«0»:«0» ⨴W1⨵⧚/⧛⨭G1⨮",
			"1:1 B2",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"B1", "B2"},
		wnt:        []string{"W1", "B1", "B2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"-:⨴0⨵ ⨴W1⨵",
			"0:1 B1",
			"1:2 B2",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        []string{"B1", "B2", "G1"},
		wnt:        []string{"B1", "B2"},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
			"⨭2⨮:- ⨭G1⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("B1 B2 B3", " "),
		wnt:        strings.Split("B1 B2 W1 B3", " "),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
			"-:⨴2⨵ ⨴W1⨵",
			"2:3 B3",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("B1 B2 G1 B3", " "),
		wnt:        strings.Split("B1 B2 B3", " "),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
			"⨭2⨮:- ⨭G1⨮",
			"3:2 B3",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("B1 B2 G1 B3", " "),
		wnt:        strings.Split("B1 B2 W1 B3", " "),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
			"«2»:«2» ⨴W1⨵⧚/⧛⨭G1⨮",
			"3:3 B3",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("B1 B2 G1 B3 G2 B4", " "),
		wnt:        strings.Split("B1 B2 W1 B3 W2 W3 B4", " "),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"0:0 B1",
			"1:1 B2",
			"«2»:«2» ⨴W1⨵⧚/⧛⨭G1⨮",
			"3:3 B3",
			"«4»:«4» ⨴W2⨵⧚/⧛⨭G2⨮",
			"-:⨴5⨵ ⨴W3⨵",
			"5:6 B4",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABC", ""),
		wnt:        strings.Split("BCD", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"⨭0⨮:- ⨭A⨮",
			"1:0 B",
			"2:1 C",
			"-:⨴2⨵ ⨴D⨵",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("BCD", ""),
		wnt:        strings.Split("ABC", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"-:⨴0⨵ ⨴A⨵",
			"0:1 B",
			"1:2 C",
			"⨭2⨮:- ⨭D⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABC", ""),
		wnt:        strings.Split("BCDABC", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"-:⨴0⨵ ⨴B⨵",
			"-:⨴1⨵ ⨴C⨵",
			"-:⨴2⨵ ⨴D⨵",
			"0:3 A",
			"1:4 B",
			"2:5 C",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("BCDABC", ""),
		wnt:        strings.Split("ABC", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"⨭0⨮:- ⨭B⨮",
			"⨭1⨮:- ⨭C⨮",
			"⨭2⨮:- ⨭D⨮",
			"3:0 A",
			"4:1 B",
			"5:2 C",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABCDEFG", ""),
		wnt:        strings.Split("ABCDBCDEFABC", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"00:00 A",
			"--:⨴01⨵ ⨴B⨵",
			"--:⨴02⨵ ⨴C⨵",
			"--:⨴03⨵ ⨴D⨵",
			"01:04 B",
			"02:05 C",
			"03:06 D",
			"04:07 E",
			"05:08 F",
			"«06»:«09» ⨴A⨵⧚/⧛⨭G⨮",
			"--:⨴10⨵ ⨴B⨵",
			"--:⨴11⨵ ⨴C⨵",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABCDBCDEFABC", ""),
		wnt:        strings.Split("ABCDEFG", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"00:00 A",
			"⨭01⨮:-- ⨭B⨮",
			"⨭02⨮:-- ⨭C⨮",
			"⨭03⨮:-- ⨭D⨮",
			"04:01 B",
			"05:02 C",
			"06:03 D",
			"07:04 E",
			"08:05 F",
			"«09»:«06» ⨴G⨵⧚/⧛⨭A⨮",
			"⨭10⨮:-- ⨭B⨮",
			"⨭11⨮:-- ⨭C⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABCDEFMNOPQUVWXY", ""),
		wnt:        strings.Split("ABqBCDEqwNOPqv23stuVWXYz", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"00:00 A",
			"--:⨴01⨵ ⨴B⨵",
			"--:⨴02⨵ ⨴q⨵",
			"01:03 B",
			"02:04 C",
			"03:05 D",
			"04:06 E",
			"«05»:«07» ⨴q⨵⧚/⧛⨭F⨮",
			"«06»:«08» ⨴w⨵⧚/⧛⨭M⨮",
			"07:09 N",
			"08:10 O",
			"09:11 P",
			"«10»:«12» ⨴q⨵⧚/⧛⨭Q⨮",
			"«11»:«13» ⨴v⨵⧚/⧛⨭U⨮",
			"--:⨴14⨵ ⨴2⨵",
			"--:⨴15⨵ ⨴3⨵",
			"--:⨴16⨵ ⨴s⨵",
			"--:⨴17⨵ ⨴t⨵",
			"--:⨴18⨵ ⨴u⨵",
			"12:19 V",
			"13:20 W",
			"14:21 X",
			"15:22 Y",
			"--:⨴23⨵ ⨴z⨵",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got:        strings.Split("ABqBCDEqwNOPqv23stuVWXYz", ""),
		wnt:        strings.Split("ABCDEFMNOPQUVWXY", ""),
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"00:00 A",
			"⨭01⨮:-- ⨭B⨮",
			"⨭02⨮:-- ⨭q⨮",
			"03:01 B",
			"04:02 C",
			"05:03 D",
			"06:04 E",
			"«07»:«05» ⨴F⨵⧚/⧛⨭q⨮",
			"«08»:«06» ⨴M⨵⧚/⧛⨭w⨮",
			"09:07 N",
			"10:08 O",
			"11:09 P",
			"«12»:«10» ⨴Q⨵⧚/⧛⨭q⨮",
			"«13»:«11» ⨴U⨵⧚/⧛⨭v⨮",
			"⨭14⨮:-- ⨭2⨮",
			"⨭15⨮:-- ⨭3⨮",
			"⨭16⨮:-- ⨭s⨮",
			"⨭17⨮:-- ⨭t⨮",
			"⨭18⨮:-- ⨭u⨮",
			"19:12 V",
			"20:13 W",
			"21:14 X",
			"22:15 Y",
			"⨭23⨮:-- ⨭z⨮",
		},
	})

	chkDiffSlice(t, &tstDiffSlice{
		got: []string{
			"A line with a difference - 1",
			"A commonly repeated line",
			"A line with a difference - 2",
			"A commonly repeated line",
			"A line with a difference - 3",
			"A commonly repeated line",
			"A line with a difference - 4",
			"A commonly repeated line",
			"A line with a difference - 5",
			"A commonly repeated line",
			"A line with a difference - 6",
			"A commonly repeated line",
			"A line with a difference - 7",
			"A commonly repeated line",
			"A line with a difference - 8",
			"A commonly repeated line",
			"A line with a difference - 9",
			"A commonly repeated line",
			"A line with a difference - 10",
			"A commonly repeated line",
		},
		wnt: []string{
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
			"A commonly repeated line",
		},
		minRunA:    1,
		expChanged: true,
		expSlice: []string{
			"⨭00⨮:-- ⨭A line with a difference - 1⨮",
			"01:00 A commonly repeated line",
			"⨭02⨮:-- ⨭A line with a difference - 2⨮",
			"03:01 A commonly repeated line",
			"⨭04⨮:-- ⨭A line with a difference - 3⨮",
			"05:02 A commonly repeated line",
			"⨭06⨮:-- ⨭A line with a difference - 4⨮",
			"07:03 A commonly repeated line",
			"⨭08⨮:-- ⨭A line with a difference - 5⨮",
			"09:04 A commonly repeated line",
			"⨭10⨮:-- ⨭A line with a difference - 6⨮",
			"11:05 A commonly repeated line",
			"⨭12⨮:-- ⨭A line with a difference - 7⨮",
			"13:06 A commonly repeated line",
			"⨭14⨮:-- ⨭A line with a difference - 8⨮",
			"15:07 A commonly repeated line",
			"⨭16⨮:-- ⨭A line with a difference - 9⨮",
			"17:08 A commonly repeated line",
			"⨭18⨮:-- ⨭A line with a difference - 10⨮",
			"19:09 A commonly repeated line",
		},
	})
}

type tstCompareSlices struct {
	got      []string
	wnt      []string
	expSlice []string
}

func chkCompareSlices(t *testing.T, tst *tstCompareSlices) {
	var got []string
	var wnt []string
	t.Helper()

	diffResult := CompareSlices(
		"[[Test Title]]",
		tst.got,
		tst.wnt,
		settingDiffSlice,
		settingDiffChars,
		defaultCmpFunc[string],
		stringify,
	)

	for _, line := range strings.Split(diffResult, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			markedUpLine := resolveMarksForDisplay(line)
			iMarkedUpLine, err := freezeMarks(markedUpLine)
			if err != nil {
				t.Fatal(err)
			}
			got = append(got, iMarkedUpLine)
		}
	}
	for _, entry := range tst.expSlice {
		for _, line := range strings.Split(entry, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				wnt = append(wnt, line)
			}
		}
	}
	if len(got) != len(wnt) {
		t.Error(errGotWnt("slice length", len(got), len(wnt)))
	}
	if strings.Join(got, "\n") != strings.Join(wnt, "\n") {
		t.Error(
			errGotWnt(
				"slice",
				"\n"+strings.Join(got, "\n"),
				"\n"+strings.Join(wnt, "\n"),
			),
		)
	}
}

func testSzTestCompareSlices(t *testing.T) {
	chkCompareSlices(t, &tstCompareSlices{
		got:      strings.Split("ABC", ""),
		wnt:      strings.Split("ABC", ""),
		expSlice: nil,
	})

	chkCompareSlices(t, &tstCompareSlices{
		got: strings.Split("ABC", ""),
		wnt: strings.Split("BCD", ""),
		expSlice: []string{
			"[[Test Title]]: got (3 lines) - want (3 lines)",
			"⨭0⨮:- ⨭A⨮",
			"1:0 B",
			"2:1 C",
			"-:⨴2⨵ ⨴D⨵",
		},
	})
}

func testSzTestCompareSlicesWithPercent(t *testing.T) {
	chkCompareSlices(t, &tstCompareSlices{
		got:      strings.Split("ABC", ""),
		wnt:      strings.Split("ABC", ""),
		expSlice: nil,
	})

	chkCompareSlices(t, &tstCompareSlices{
		got: strings.Split("A%BC", ""),
		wnt: strings.Split("BCD", ""),
		expSlice: []string{
			"[[Test Title]]: got (4 lines) - want (3 lines)",
			"⨭0⨮:- ⨭A⨮",
			"⨭1⨮:- ⨭%⨮",
			"2:0 B",
			"3:1 C",
			"-:⨴2⨵ ⨴D⨵",
		},
	})
}

type tstCompareArrays struct {
	got      []string
	wnt      []string
	expSlice []string
}

func chkCompareArrays(t *testing.T, tst *tstCompareArrays) {
	var got []string
	var wnt []string
	t.Helper()

	diffResult := CompareArrays(
		tst.got,
		tst.wnt,
	)

	for _, line := range strings.Split(diffResult, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			markedUpLine := resolveMarksForDisplay(line)
			iMarkedUpLine, err := freezeMarks(markedUpLine)
			if err != nil {
				t.Fatal(err)
			}
			got = append(got, iMarkedUpLine)
		}
	}
	for _, entry := range tst.expSlice {
		for _, line := range strings.Split(entry, "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				wnt = append(wnt, line)
			}
		}
	}
	if len(got) != len(wnt) {
		t.Error(errGotWnt("slice length", len(got), len(wnt)))
	}
	if strings.Join(got, "\n") != strings.Join(wnt, "\n") {
		t.Error(
			errGotWnt(
				"slice",
				"\n"+strings.Join(got, "\n"),
				"\n"+strings.Join(wnt, "\n"),
			),
		)
	}
}

func testSzTestCompareArrays(t *testing.T) {
	chkCompareArrays(t, &tstCompareArrays{
		got:      strings.Split("ABC", ""),
		wnt:      strings.Split("ABC", ""),
		expSlice: nil,
	})
	chkCompareArrays(t, &tstCompareArrays{
		got: strings.Split("ABC", ""),
		wnt: strings.Split("AbC", ""),
		expSlice: []string{
			"0:0 A",
			"«1»:«1» ⨴b⨵⧚/⧛⨭B⨮",
			"2:2 C",
		},
	})
}
