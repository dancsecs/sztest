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

import "testing"

func tstChkComplex64(t *testing.T) {
	t.Run("Good", chkComplex128TestGood)

	t.Run("Bad", chkComplex128TestBad)
	t.Run("BadMsg1", chkComplex128TestBad1)
	t.Run("BadMsg2", chkComplex128TestBad2)
	t.Run("BadMsg3", chkComplex128TestBad3)

	t.Run("Slice-Good", chkComplex128SliceTestGood)
	t.Run("Slice-BadMsg1", chkComplex128SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkComplex128SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkComplex128SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkComplex128SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkComplex128SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkComplex128SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkComplex128SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkComplex128SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkComplex128SliceTestBadMsg9)
}

func chkComplex128TestGood(t *testing.T) {
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

func chkComplex128TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128",
			chkOutCommonMsg("", complex128TypeName),
			g(markAsChg("(2", "(1", diffGot)+"+0i)"),
			w(markAsChg("(2", "(1", diffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128f",
			chkOutCommonMsg(
				"This message will be displayed first",
				complex128TypeName,
			),
			g(markAsChg("(2", "(1", diffGot)+"+0i)"),
			w(markAsChg("(2", "(1", diffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128",
			chkOutCommonMsg(
				"This message will be displayed second",
				complex128TypeName,
			),
			g(markAsChg("(2", "(1", diffGot)+"+0i)"),
			w(markAsChg("(2", "(1", diffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Complex128f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Complex128f",
			chkOutCommonMsg(
				"This message will be displayed third",
				complex128TypeName,
			),
			g(markAsChg("(0", "(1", diffGot)+"+0i)"),
			w(markAsChg("(0", "(1", diffWant)+"+0i)"),
		),
		chkOutRelease(),
	)
}

func chkComplex128SliceTestGood(t *testing.T) {
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

func chkComplex128SliceTestBadMsg1(t *testing.T) {
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

func chkComplex128SliceTestBadMsg2(t *testing.T) {
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

func chkComplex128SliceTestBadMsg3(t *testing.T) {
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

func chkComplex128SliceTestBadMsg4(t *testing.T) {
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

func chkComplex128SliceTestBadMsg5(t *testing.T) {
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

func chkComplex128SliceTestBadMsg6(t *testing.T) {
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

func chkComplex128SliceTestBadMsg7(t *testing.T) {
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

func chkComplex128SliceTestBadMsg8(t *testing.T) {
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

func chkComplex128SliceTestBadMsg9(t *testing.T) {
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
