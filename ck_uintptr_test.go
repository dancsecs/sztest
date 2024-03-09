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

func tstChkUintptr(t *testing.T) {
	t.Run("Good", chkUintptrTestGood)

	t.Run("Bad", chkUintptrTestBad)
	t.Run("BadMsg1", chkUintptrTestBad1)
	t.Run("BadMsg2", chkUintptrTestBad2)
	t.Run("BadMsg3", chkUintptrTestBad3)

	t.Run("Slice-Good", chkUintptrSliceTestGood)
	t.Run("Slice-BadMsg1", chkUintptrSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkUintptrSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkUintptrSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkUintptrSliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkUintptrSliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkUintptrSliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkUintptrSliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkUintptrSliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkUintptrSliceTestBadMsg9)
}

func chkUintptrTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintptr(0, 0)
	chk.Uintptr(0, 0, "not ", "displayed")
	chk.Uintptrf(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUintptrTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintptr(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptr",
			chkOutCommonMsg("", uintptrTypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintptrf(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptrf",
			chkOutCommonMsg("This message will be displayed first", uintptrTypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintptr(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptr",
			chkOutCommonMsg("This message will be displayed second", uintptrTypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintptrf(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptrf",
			chkOutCommonMsg("This message will be displayed third", uintptrTypeName),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice(
		[]uintptr{}, []uintptr{},
		"This message will NOT be displayed",
	)
	chk.UintptrSlice(
		[]uintptr{0}, []uintptr{0},
		"This message will NOT be displayed",
	)
	chk.UintptrSlice(
		[]uintptr{2, 6, 7}, []uintptr{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.UintptrSlicef(
		[]uintptr{2, 6, 7, 9}, []uintptr{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlicef(
		[]uintptr{2}, []uintptr{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]uintptr", "UintptrSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice(
		[]uintptr{}, []uintptr{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]uintptr", "UintptrSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlicef(
		[]uintptr{1, 2}, []uintptr{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]uintptr", "UintptrSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice(
		[]uintptr{1}, []uintptr{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]uintptr", "UintptrSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlicef(
		[]uintptr{1, 2}, []uintptr{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uintptr", "UintptrSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice(
		[]uintptr{2, 2}, []uintptr{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uintptr", "UintptrSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlicef(
		[]uintptr{1, 2}, []uintptr{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uintptr", "UintptrSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice(
		[]uintptr{1, 3}, []uintptr{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uintptr", "UintptrSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintptrSlice([]uintptr{1, 3}, []uintptr{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]uintptr", "UintptrSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}
