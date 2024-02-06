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
	t.Run("Good", chkUintptrTest_Good)

	t.Run("Bad", chkUintptrTest_Bad)
	t.Run("BadMsg1", chkUintptrTest_Bad1)
	t.Run("BadMsg2", chkUintptrTest_Bad2)
	t.Run("BadMsg3", chkUintptrTest_Bad3)

	t.Run("Slice-Good", chkUintptrSliceTest_Good)
	t.Run("Slice-BadMsg1", chkUintptrSliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkUintptrSliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkUintptrSliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkUintptrSliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkUintptrSliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkUintptrSliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkUintptrSliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkUintptrSliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkUintptrSliceTest_BadMsg9)
}

func chkUintptrTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrTest_Bad(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uintptr(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptr",
			chkOutCommonMsg("", "uintptr"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTest_Bad1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uintptrf(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptrf",
			chkOutCommonMsg("This message will be displayed first", "uintptr"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTest_Bad2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uintptr(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptr",
			chkOutCommonMsg("This message will be displayed second", "uintptr"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrTest_Bad3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uintptrf(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintptrf",
			chkOutCommonMsg("This message will be displayed third", "uintptr"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintptrSliceTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg4(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg5(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg6(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg7(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg8(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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

func chkUintptrSliceTest_BadMsg9(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
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
