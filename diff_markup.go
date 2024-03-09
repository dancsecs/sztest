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
	switch dType { //nolint:exhaustive // Default handles all other cases.
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

func resolveMarksForDisplay(line string) string {
	line = strings.ReplaceAll(line, markDelOn, settingMarkDelOn)
	line = strings.ReplaceAll(line, markDelOff, settingMarkDelOff)
	line = strings.ReplaceAll(line, markInsOn, settingMarkInsOn)
	line = strings.ReplaceAll(line, markInsOff, settingMarkInsOff)
	line = strings.ReplaceAll(line, markChgOn, settingMarkChgOn)
	line = strings.ReplaceAll(line, markChgOff, settingMarkChgOff)
	line = strings.ReplaceAll(line, markSepOn, settingMarkSepOn)
	line = strings.ReplaceAll(line, markSepOff, settingMarkSepOff)
	line = strings.ReplaceAll(line, markWntOn, settingMarkWntOn)
	line = strings.ReplaceAll(line, markWntOff, settingMarkWntOff)
	line = strings.ReplaceAll(line, markGotOn, settingMarkGotOn)
	line = strings.ReplaceAll(line, markGotOff, settingMarkGotOff)
	line = strings.ReplaceAll(line, markMsgOn, settingMarkMsgOn)
	line = strings.ReplaceAll(line, markMsgOff, settingMarkMsgOff)

	return line
}

func gotWnt(got, wnt string) string {
	prefix := ""
	if strings.Count(got, "\n")+strings.Count(wnt, "\n") > 0 {
		prefix = "\n"
	}

	return fmt.Sprint(
		g(prefix+got),
		"\n",
		w(prefix+wnt),
	)
}

func gotWntDiff(got, wnt string, minRun int) string {
	return gotWnt(
		DiffString(got, wnt, DiffGot, minRun),
		DiffString(got, wnt, DiffWant, minRun),
	)
}
