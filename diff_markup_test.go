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
	"strings"
	"testing"
)

func testDiffMarkupPrerequisites(t *testing.T) {
	t.Run("InternalMarkup", testSzTestInternalMarkup)
	t.Run("MarkupForTerminal", testSzTestMarkupForTerminal)
	t.Run("GotWant", testSzTestGotWant)
}

func testSzTestInternalMarkup(t *testing.T) {
	const (
		area = "internal markup"
		msg  = "<-- MSG -->"
	)

	var got, wnt string

	got = w(msg)
	wnt = markWntOn + labelWant + ": " + markWntOff + msg

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = g(msg)
	wnt = markGotOn + labelGot + ": " + markGotOff + msg

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsIns(msg)
	wnt = markInsOn + msg + markInsOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsDel(msg)
	wnt = markDelOn + msg + markDelOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsChg(msg, strings.ToLower(msg), diffGot)
	wnt = markChgOn + msg + markChgOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsChg(msg, strings.ToLower(msg), diffWant)
	wnt = markChgOn + strings.ToLower(msg) + markChgOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsChg(msg, strings.ToLower(msg), diffMerge)
	wnt = markDelOn + strings.ToLower(msg) + markDelOff +
		markSepOn + "/" + markSepOff +
		markInsOn + msg + markInsOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}

func testSzTestMarkupForTerminal(t *testing.T) {
	const (
		area    = "markup for terminal"
		tstStr1 = "ABC"
		tstStr2 = "DEF"
	)

	var got, wnt string

	got = resolveMarksForDisplay(markAsIns(tstStr1))
	wnt = settingMarkInsOn + tstStr1 + settingMarkInsOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsDel(tstStr1))
	wnt = settingMarkDelOn + tstStr1 + settingMarkDelOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(tstStr1, tstStr2, diffGot))
	wnt = settingMarkChgOn + tstStr1 + settingMarkChgOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(tstStr1, tstStr2, diffWant))
	wnt = settingMarkChgOn + tstStr2 + settingMarkChgOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(tstStr1, tstStr2, diffMerge))
	wnt = "" +
		settingMarkDelOn + tstStr2 + settingMarkDelOff +
		settingMarkSepOn + "/" + settingMarkSepOff +
		settingMarkInsOn + tstStr1 + settingMarkInsOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsMsg(tstStr1))
	wnt = "" +
		settingMarkMsgOn + tstStr1 + settingMarkMsgOff

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}

func testSzTestGotWant(t *testing.T) {
	const area = "gotWnt"

	var got, wnt string

	got = gotWntDiff("ABC", "ADC", 1)
	wnt = "" +
		g("A"+markAsChg("B", "D", diffGot)+"C") +
		"\n" +
		w("A"+markAsChg("B", "D", diffWant)+"C")

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = gotWntDiff("AB\n", "AC\n", 3)
	wnt = "" +
		g("\n"+markAsChg("AB\n", "AC\n", diffGot)) +
		"\n" +
		w("\n"+markAsChg("AB\n", "AC\n", diffWant))

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}
