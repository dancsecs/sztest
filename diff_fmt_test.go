/*
   Golang test helper library: sztest.
   Copyright (C) 2023-2025 Leslie Dancsecs

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
	"testing"
)

func testDiffFmt(t *testing.T) {
	t.Run("FormatLineNumber1", testSzTestDiffFmtFmtLnNbr1)
	t.Run("FormatLineNumber2", testSzTestDiffFmtFmtLnNbr2)
	t.Run("FormatLineNumber5", testSzTestDiffFmtFmtLnNbr5)
	t.Run("FormatSameLine", testSzTestDiffFmtFmtSameLine)
	t.Run("FormatChangedLine", testSzTestDiffFmtFmtChangedLine)
	t.Run("FormatJustGotLine", testSzTestDiffFmtFmtJustGotLine)
	t.Run("FormatJustWntLine", testSzTestDiffFmtFmtJustWntLine)
}

func testSzTestDiffFmtFmtLnNbr1(t *testing.T) {
	defaultFormat := newDiffLnFmt(0, 0)
	dFmt := newDiffLnFmt(1, 0)

	if defaultFormat.nbrWidth != dFmt.nbrWidth {
		t.Error(errGotWnt("-1", defaultFormat.nbrWidth, dFmt.nbrWidth))
	}

	got := dFmt.fmtLnNbr(-1)
	wnt := "-"

	if got != wnt {
		t.Error(errGotWnt("-1", got, wnt))
	}

	got = dFmt.fmtLnNbr(0)
	wnt = "0"

	if got != wnt {
		t.Error(errGotWnt("0", got, wnt))
	}

	got = dFmt.fmtLnNbr(1)
	wnt = "1"

	if got != wnt {
		t.Error(errGotWnt("1", got, wnt))
	}

	got = dFmt.fmtLnNbr(2)
	wnt = "2"

	if got != wnt {
		t.Error(errGotWnt("2", got, wnt))
	}

	got = dFmt.fmtLnNbr(22) // Exceeds width.
	wnt = "22"

	if got != wnt {
		t.Error(errGotWnt("22", got, wnt))
	}
}

func testSzTestDiffFmtFmtLnNbr2(t *testing.T) {
	dFmt := newDiffLnFmt(0, 10)

	got := dFmt.fmtLnNbr(-1)
	wnt := "--"

	if got != wnt {
		t.Error(errGotWnt("-1", got, wnt))
	}

	got = dFmt.fmtLnNbr(0)
	wnt = "00"

	if got != wnt {
		t.Error(errGotWnt("0", got, wnt))
	}

	got = dFmt.fmtLnNbr(1)
	wnt = "01"

	if got != wnt {
		t.Error(errGotWnt("1", got, wnt))
	}

	got = dFmt.fmtLnNbr(2)
	wnt = "02"

	if got != wnt {
		t.Error(errGotWnt("2", got, wnt))
	}

	got = dFmt.fmtLnNbr(222) // Exceeds width.
	wnt = "222"

	if got != wnt {
		t.Error(errGotWnt("222", got, wnt))
	}
}

func testSzTestDiffFmtFmtLnNbr5(t *testing.T) {
	dFmt := newDiffLnFmt(10000, 0)

	got := dFmt.fmtLnNbr(-1)
	wnt := "-----"

	if got != wnt {
		t.Error(errGotWnt("-1", got, wnt))
	}

	got = dFmt.fmtLnNbr(0)
	wnt = "00000"

	if got != wnt {
		t.Error(errGotWnt("0", got, wnt))
	}

	got = dFmt.fmtLnNbr(1)
	wnt = "00001"

	if got != wnt {
		t.Error(errGotWnt("1", got, wnt))
	}

	got = dFmt.fmtLnNbr(2)
	wnt = "00002"

	if got != wnt {
		t.Error(errGotWnt("2", got, wnt))
	}

	got = dFmt.fmtLnNbr(22222) // Exceeds width.
	wnt = "22222"

	if got != wnt {
		t.Error(errGotWnt("22222", got, wnt))
	}
}

func testSzTestDiffFmtFmtSameLine(t *testing.T) {
	dFmt := newDiffLnFmt(100, 0)

	got := dFmt.same(0, 0, "the line")
	wnt := "000:000 the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.same(223, 159, "the line")
	wnt = "223:159 the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	dFmt = dFmt.newOffset(8, 12)

	got = dFmt.same(0, 0, "the line")
	wnt = "008:012 the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.same(223, 159, "the line")
	wnt = "231:171 the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}
}

func testSzTestDiffFmtFmtChangedLine(t *testing.T) {
	dFmt := newDiffLnFmt(100, 0)

	got := dFmt.changed(0, 0, "the line")
	wnt := "" +
		markAsChg("000", "", diffGot) +
		":" +
		markAsChg("", "000", diffWant) +
		" the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.changed(223, 159, "the line")
	wnt = "" +
		markAsChg("223", "", diffGot) +
		":" +
		markAsChg("", "159", diffWant) +
		" the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	dFmt = dFmt.newOffset(8, 12)

	got = dFmt.changed(0, 0, "the line")
	wnt = "" +
		markAsChg("008", "", diffGot) +
		":" +
		markAsChg("", "012", diffWant) +
		" the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.changed(223, 159, "the line")
	wnt = "" +
		markAsChg("231", "", diffGot) +
		":" +
		markAsChg("", "171", diffWant) +
		" the line"

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}
}

func testSzTestDiffFmtFmtJustGotLine(t *testing.T) {
	dFmt := newDiffLnFmt(100, 0)

	got := dFmt.justGot(0, "the line")
	wnt := "" +
		markAsIns("000") +
		":" +
		"--- " +
		markAsIns("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.justGot(223, "the line")
	wnt = "" +
		markAsIns("223") +
		":" +
		"--- " +
		markAsIns("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	dFmt = dFmt.newOffset(8, 12)

	got = dFmt.justGot(0, "the line")
	wnt = "" +
		markAsIns("008") +
		":" +
		"--- " +
		markAsIns("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.justGot(223, "the line")
	wnt = "" +
		markAsIns("231") +
		":" +
		"--- " +
		markAsIns("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}
}

func testSzTestDiffFmtFmtJustWntLine(t *testing.T) {
	dFmt := newDiffLnFmt(100, 0)

	got := dFmt.justWnt(0, "the line")
	wnt := "" +
		"---" +
		":" +
		markAsDel("000") +
		" " +
		markAsDel("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.justWnt(159, "the line")
	wnt = "" +
		"---" +
		":" +
		markAsDel("159") +
		" " +
		markAsDel("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	dFmt = dFmt.newOffset(8, 12)

	got = dFmt.justWnt(0, "the line")
	wnt = "" +
		"---" +
		":" +
		markAsDel("012") +
		" " +
		markAsDel("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}

	got = dFmt.justWnt(159, "the line")
	wnt = "" +
		"---" +
		":" +
		markAsDel("171") +
		" " +
		markAsDel("the line")

	if got != wnt {
		t.Error(errGotWnt("the line", got, wnt))
	}
}
