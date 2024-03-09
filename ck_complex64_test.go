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

func tstChkComplex128(t *testing.T) {
	t.Run("Good", chkComplex64Test_Good)

	t.Run("Bad", chkComplex64Test_Bad)
	t.Run("BadMsg1", chkComplex64Test_Bad1)
	t.Run("BadMsg2", chkComplex64Test_Bad2)
	t.Run("BadMsg3", chkComplex64Test_Bad3)

	t.Run("Slice-Good", chkComplex64SliceTest_Good)
	t.Run("Slice-BadMsg1", chkComplex64SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkComplex64SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkComplex64SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkComplex64SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkComplex64SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkComplex64SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkComplex64SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkComplex64SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkComplex64SliceTest_BadMsg9)
}

func chkComplex64Test_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64(0, 0)
	chk.Complex64(1, 1, "not ", "displayed")
	chk.Complex64f(2, 2, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkComplex64Test_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex64",
			chkOutCommonMsg("", "complex64"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64Test_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex64f",
			chkOutCommonMsg("This message will be displayed first", "complex64"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64Test_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex64",
			chkOutCommonMsg("This message will be displayed second", "complex64"),
			g(markAsChg("(2", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(2", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64Test_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex64f",
			chkOutCommonMsg("This message will be displayed third", "complex64"),
			g(markAsChg("(0", "(1", DiffGot)+"+0i)"),
			w(markAsChg("(0", "(1", DiffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice(
		[]complex64{}, []complex64{},
		"This message will NOT be displayed",
	)
	chk.Complex64Slice(
		[]complex64{0}, []complex64{0},
		"This message will NOT be displayed",
	)
	chk.Complex64Slice(
		[]complex64{2, 6, 7}, []complex64{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.Complex64Slicef(
		[]complex64{2, 6, 7, 9}, []complex64{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slicef(
		[]complex64{2}, []complex64{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]complex64", "Complex64Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice(
		[]complex64{}, []complex64{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]complex64", "Complex64Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "(1+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slicef(
		[]complex64{1, 2}, []complex64{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]complex64", "Complex64Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnGot("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice(
		[]complex64{1}, []complex64{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]complex64", "Complex64Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnWnt("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slicef(
		[]complex64{1, 2}, []complex64{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex64", "Complex64Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "(1+0i)"),
			chkOutLnSame("1", "0", "(2+0i)"),
			chkOutLnWnt("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice(
		[]complex64{2, 2}, []complex64{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex64", "Complex64Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "(1+0i)"),
			chkOutLnSame("0", "1", "(2+0i)"),
			chkOutLnGot("1", "(2+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slicef(
		[]complex64{1, 2}, []complex64{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex64", "Complex64Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(2", "(3")+"+0i)",
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice(
		[]complex64{1, 3}, []complex64{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]complex64", "Complex64Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(3", "(2")+"+0i)",
		),
		chkOutRelease(),
	)
}

func chkComplex64SliceTest_BadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex64Slice([]complex64{1, 3}, []complex64{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]complex64", "Complex64Slice", "",
			chkOutLnSame("0", "0", "(1+0i)"),
			chkOutLnChanged("1", "1", "(3", "(2")+"+0i)",
		),
		chkOutRelease(),
	)
}
