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
	"strings"
	"testing"
)

func test_DiffMarkup_Prerequisites(t *testing.T) {
	t.Run("InternalMarkup", testSzTest_InternalMarkup)
	t.Run("MarkupForTerminal", testSzTest_MarkupForTerminal)
	t.Run("GotWant", testSzTest_GotWant)
}

func testSzTest_InternalMarkup(t *testing.T) {
	const area = "internal markup"
	const msg = "<-- MSG -->"
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

	got = markAsChg(msg, strings.ToLower(msg), DiffGot)
	wnt = markChgOn + msg + markChgOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsChg(msg, strings.ToLower(msg), DiffWant)
	wnt = markChgOn + strings.ToLower(msg) + markChgOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = markAsChg(msg, strings.ToLower(msg), DiffMerge)
	wnt = markDelOn + strings.ToLower(msg) + markDelOff +
		markSepOn + "/" + markSepOff +
		markInsOn + msg + markInsOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}

func testSzTest_MarkupForTerminal(t *testing.T) {
	const area = "markup for terminal"
	const s1 = "ABC"
	const s2 = "DEF"
	var got, wnt string

	got = resolveMarksForDisplay(markAsIns(s1))
	wnt = settingMarkInsOn + s1 + settingMarkInsOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsDel(s1))
	wnt = settingMarkDelOn + s1 + settingMarkDelOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(s1, s2, DiffGot))
	wnt = settingMarkChgOn + s1 + settingMarkChgOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(s1, s2, DiffWant))
	wnt = settingMarkChgOn + s2 + settingMarkChgOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsChg(s1, s2, DiffMerge))
	wnt = "" +
		settingMarkDelOn + s2 + settingMarkDelOff +
		settingMarkSepOn + "/" + settingMarkSepOff +
		settingMarkInsOn + s1 + settingMarkInsOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = resolveMarksForDisplay(markAsMsg(s1))
	wnt = "" +
		settingMarkMsgOn + s1 + settingMarkMsgOff
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}

func testSzTest_GotWant(t *testing.T) {
	const area = "gotWnt"
	var got, wnt string

	got = gotWntDiff("ABC", "ADC", 1)
	wnt = "" +
		g("A"+markAsChg("B", "D", DiffGot)+"C") +
		"\n" +
		w("A"+markAsChg("B", "D", DiffWant)+"C")

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = gotWntDiff("AB\n", "AC\n", 3)
	wnt = "" +
		g("\n"+markAsChg("AB\n", "AC\n", DiffGot)) +
		"\n" +
		w("\n"+markAsChg("AB\n", "AC\n", DiffWant))

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}
}
