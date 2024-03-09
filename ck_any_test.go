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

import "testing"

func tstChkAny(t *testing.T) {
	t.Run("Nil", chkAnyTest_Nil)
	t.Run("Nilf", chkAnyTest_Nilf)
	t.Run("NotNil", chkAnyTest_NotNil)
	t.Run("NotNilf", chkAnyTest_NotNilf)
}

type abcInterface interface {
	abcFunction()
}

type abcStruct struct{}

func (s *abcStruct) abcFunction() {
}

func chkAnyTest_Nil(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var nilReference *abcStruct
	nilInterface := abcInterface(nilReference)

	notNilReference := &abcStruct{}
	notNilInterface := abcInterface(notNilReference)

	chk.Nil(nil, "This message will NOT be displayed")
	chk.Nil(error(nil), "This message will NOT be displayed")
	chk.Nil(nilInterface, "This message will ", "NOT be displayed")
	chk.Nil(nilReference)

	chk.Nil(chk)
	chk.Nil(notNilInterface, "This message will be displayed")
	chk.Nil(notNilReference, "This message will ", "be displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutHelper("Nil"),
		chkOutError(chkOutCommonMsg("", "not nil interface")),

		chkOutHelper("Nil"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "not nil interface"),
		),

		chkOutHelper("Nil"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "not nil interface"),
		),

		chkOutRelease(),
	)
}

func chkAnyTest_Nilf(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var nilReference *abcStruct
	nilInterface := abcInterface(nilReference)

	notNilReference := &abcStruct{}
	notNilInterface := abcInterface(notNilReference)

	chk.Nilf(nil, "This message will NOT be displayed")
	chk.Nilf(error(nil), "This message will NOT be displayed")
	chk.Nilf(nilInterface, "This message will %s", "NOT be displayed")
	chk.Nilf(nilReference, "")

	chk.Nilf(chk, "")
	chk.Nilf(notNilInterface, "This message will be displayed")
	chk.Nilf(notNilReference, "This message will %s", "be displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutHelper("Nilf"),
		chkOutError(chkOutCommonMsg("", "not nil interface")),

		chkOutHelper("Nilf"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "not nil interface"),
		),

		chkOutHelper("Nilf"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "not nil interface"),
		),

		chkOutRelease(),
	)
}

func chkAnyTest_NotNil(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var nilReference *abcStruct
	nilInterface := abcInterface(nilReference)

	notNilReference := &abcStruct{}
	notNilInterface := abcInterface(notNilReference)

	chk.NotNil(chk)
	chk.NotNil(notNilInterface, "This message will NOT be displayed")
	chk.NotNil(notNilReference, "This message will ", "Not be displayed")

	chk.NotNil(nil, "This message will be displayed")
	chk.NotNil(error(nil), "This message will be displayed")
	chk.NotNil(nilInterface, "This message will ", "be displayed")
	chk.NotNil(nilReference)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutHelper("NotNil"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNil"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNil"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNil"),
		chkOutError(
			chkOutCommonMsg("", "nil interface"),
		),

		chkOutRelease(),
	)
}

func chkAnyTest_NotNilf(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var nilReference *abcStruct
	nilInterface := abcInterface(nilReference)

	notNilReference := &abcStruct{}
	notNilInterface := abcInterface(notNilReference)

	chk.NotNilf(chk, "")
	chk.NotNilf(notNilInterface, "This message will NOT be displayed")
	chk.NotNilf(notNilReference, "This message will %s", "Not be displayed")

	chk.NotNilf(nil, "This message will be displayed")
	chk.NotNilf(error(nil), "This message will be displayed")
	chk.NotNilf(nilInterface, "This message will %s", "be displayed")
	chk.NotNilf(nilReference, "")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutHelper("NotNilf"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNilf"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNilf"),
		chkOutError(
			chkOutCommonMsg("This message will be displayed", "nil interface"),
		),

		chkOutHelper("NotNilf"),
		chkOutError(
			chkOutCommonMsg("", "nil interface"),
		),

		chkOutRelease(),
	)
}
