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
	"testing"
)

func tstChkSubstitution(t *testing.T) {
	t.Run("All", chkSubstitutionTest_All)
}

func chkSubstitutionTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
	iT.chk = chk

	chk.AddSub(`[a-`, "gf")

	chk.AddSub(`{{f}}`, "flagged")

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Log(
		"This line should be {{f}} as not in log",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("AddSub"),
		chkOutFatalf("error parsing regexp: missing closing ]: `[a-`"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected log Entry: got (0 lines) - want (1 lines)",
		chkOutLnWnt("0", "This line should be flagged as not in log"),
		"Fail Now: (*Chk).Error",
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}
