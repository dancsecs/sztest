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

func tstChkComplex64(t *testing.T) {
	t.Run("Good", chkComplex128Test_Good)

	t.Run("Bad", chkComplex128Test_Bad)
	t.Run("BadMsg1", chkComplex128Test_Bad1)
	t.Run("BadMsg2", chkComplex128Test_Bad2)
	t.Run("BadMsg3", chkComplex128Test_Bad3)

	t.Run("Slice-Good", chkComplex128SliceTest_Good)
	t.Run("Slice-BadMsg1", chkComplex128SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkComplex128SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkComplex128SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkComplex128SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkComplex128SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkComplex128SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkComplex128SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkComplex128SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkComplex128SliceTest_BadMsg9)
}

func chkComplex128Test_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128(0, 0)
	chk.Complex128(0, 0, "not ", "displayed")
	chk.Complex128f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkComplex128Test_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128",
			chkOutCommonMsg("", "complex128"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128Test_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128f",
			chkOutCommonMsg("This message will be displayed first", "complex128"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128Test_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128",
			chkOutCommonMsg("This message will be displayed second", "complex128"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128Test_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128f",
			chkOutCommonMsg("This message will be displayed third", "complex128"),
			g(markAsChg("(0", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(0", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice(
		[]complex128{}, []complex128{},
		"This message will NOT be displayed",
	)
	chk.Complex128Slice(
		[]complex128{0}, []complex128{0},
		"This message will NOT be displayed",
	)
	chk.Complex128Slice(
		[]complex128{2, 6, 7}, []complex128{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.Complex128Slicef(
		[]complex128{2, 6, 7, 9}, []complex128{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slicef(
		[]complex128{2}, []complex128{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]complex128", "Complex128Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice(
		[]complex128{}, []complex128{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]complex128", "Complex128Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "(1+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slicef(
		[]complex128{1, 2}, []complex128{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]complex128", "Complex128Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnGot("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice(
		[]complex128{1}, []complex128{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]complex128", "Complex128Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnWnt("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slicef(
		[]complex128{1, 2}, []complex128{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex128", "Complex128Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "(1+0i)"),
			chkOutLnSame("1", "0", "(2+0i)"),
			chkOutLnWnt("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice(
		[]complex128{2, 2}, []complex128{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex128", "Complex128Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "(1+0i)"),
			chkOutLnSame("0", "1", "(2+0i)"),
			chkOutLnGot("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slicef(
		[]complex128{1, 2}, []complex128{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex128", "Complex128Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(2", "(3")+"+0i)",
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice(
		[]complex128{1, 3}, []complex128{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex128", "Complex128Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(3", "(2")+"+0i)",
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTest_BadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128Slice([]complex128{1, 3}, []complex128{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]complex128", "Complex128Slice", "",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(3", "(2")+"+0i)",
		),
		chkOutRelease(),
	)
}
