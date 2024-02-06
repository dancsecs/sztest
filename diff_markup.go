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
)

type markupFunction func(string) string

// Codes to internally represent text display decorations.
const (
	markWntOn  = "|-|WntOn|-|"
	markWntOff = "|-|WntOff|-|"
	markGotOn  = "|-|GoTOn|-|"
	markGotOff = "|-|GoTOff|-|"
	markDelOn  = "|-|DeLOn|-|"
	markDelOff = "|-|DeLOff|-|"
	markInsOn  = "|-|InSOn|-|"
	markInsOff = "|-|InSOff|-|"
	markChgOn  = "|-|ChGOn|-|"
	markChgOff = "|-|ChGOff|-|"
	markSepOn  = "|-|SePOn|-|"
	markSepOff = "|-|SePOff|-|"
	markMsgOn  = "|-|MsGOn|-|"
	markMsgOff = "|-|MsGOff|-|"

	labelWant = "WNT"
	labelGot  = "GOT"
)

func w(msg string) string {
	return markWntOn + labelWant + ": " + markWntOff + msg
}

func g(msg string) string {
	return markGotOn + labelGot + ": " + markGotOff + msg
}

func markAsIns(s string) string {
	return markInsOn + s + markInsOff
}

func markAsDel(s string) string {
	return markDelOn + s + markDelOff
}

func markAsChg(got, wnt string, dType diffType) string {
	switch dType {
	case DiffGot:
		return markChgOn + got + markChgOff
	case DiffWant:
		return markChgOn + wnt + markChgOff
	default:
		return markDelOn + wnt + markDelOff + markSepOn +
			"/" + markSepOff + markInsOn + got + markInsOff
	}
}

func markAsMsg(s string) string {
	return markMsgOn + s + markMsgOff
}

func resolveMarksForDisplay(r string) string {
	r = strings.ReplaceAll(r, markDelOn, settingMarkDelOn)
	r = strings.ReplaceAll(r, markDelOff, settingMarkDelOff)
	r = strings.ReplaceAll(r, markInsOn, settingMarkInsOn)
	r = strings.ReplaceAll(r, markInsOff, settingMarkInsOff)
	r = strings.ReplaceAll(r, markChgOn, settingMarkChgOn)
	r = strings.ReplaceAll(r, markChgOff, settingMarkChgOff)
	r = strings.ReplaceAll(r, markSepOn, settingMarkSepOn)
	r = strings.ReplaceAll(r, markSepOff, settingMarkSepOff)
	r = strings.ReplaceAll(r, markWntOn, settingMarkWntOn)
	r = strings.ReplaceAll(r, markWntOff, settingMarkWntOff)
	r = strings.ReplaceAll(r, markGotOn, settingMarkGotOn)
	r = strings.ReplaceAll(r, markGotOff, settingMarkGotOff)
	r = strings.ReplaceAll(r, markMsgOn, settingMarkMsgOn)
	r = strings.ReplaceAll(r, markMsgOff, settingMarkMsgOff)
	return r
}

func gotWnt(got, wnt string) string {
	h := ""
	if strings.Count(got, "\n")+strings.Count(wnt, "\n") > 0 {
		h = "\n"
	}
	return fmt.Sprint(
		g(h+got),
		"\n",
		w(h+wnt),
	)
}

func gotWntDiff(got, wnt string, minRun int) string {
	return gotWnt(
		DiffString(got, wnt, DiffGot, minRun),
		DiffString(got, wnt, DiffWant, minRun),
	)
}
