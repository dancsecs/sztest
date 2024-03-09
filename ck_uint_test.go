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

func tstChkUint(t *testing.T) {
	t.Run("Good", chkUintTest_Good)

	t.Run("Bad", chkUintTest_Bad)
	t.Run("BadMsg1", chkUintTest_Bad1)
	t.Run("BadMsg2", chkUintTest_Bad2)
	t.Run("BadMsg3", chkUintTest_Bad3)

	t.Run("Slice-Good", chkUintSliceTest_Good)
	t.Run("Slice-BadMsg1", chkUintSliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkUintSliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkUintSliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkUintSliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkUintSliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkUintSliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkUintSliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkUintSliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkUintSliceTest_BadMsg9)

	t.Run("Bounded", chkUintBoundedTest_All)
	t.Run("Unbounded", chkUintUnboundedTest_All)
}

func chkUintTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint(0, 0)
	chk.Uint(0, 0, "not ", "displayed")
	chk.Uintf(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUintTest_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint",
			chkOutCommonMsg("", "uint"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintTest_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintf(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintf",
			chkOutCommonMsg("This message will be displayed first", "uint"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintTest_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint",
			chkOutCommonMsg("This message will be displayed second", "uint"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintTest_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uintf(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uintf",
			chkOutCommonMsg("This message will be displayed third", "uint"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice(
		[]uint{}, []uint{},
		"This message will NOT be displayed",
	)
	chk.UintSlice(
		[]uint{0}, []uint{0},
		"This message will NOT be displayed",
	)
	chk.UintSlice(
		[]uint{2, 6, 7}, []uint{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.UintSlicef(
		[]uint{2, 6, 7, 9}, []uint{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlicef(
		[]uint{2}, []uint{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]uint", "UintSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice(
		[]uint{}, []uint{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]uint", "UintSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlicef(
		[]uint{1, 2}, []uint{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]uint", "UintSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice(
		[]uint{1}, []uint{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]uint", "UintSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlicef(
		[]uint{1, 2}, []uint{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint", "UintSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice(
		[]uint{2, 2}, []uint{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint", "UintSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlicef(
		[]uint{1, 2}, []uint{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint", "UintSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice(
		[]uint{1, 3}, []uint{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint", "UintSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkUintSliceTest_BadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.UintSlice([]uint{1, 3}, []uint{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]uint", "UintSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkUintBoundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := uint(33)
	max := uint(35)

	// Bad: Error displayed.
	chk.UintBounded(30, BoundedClosed, min, max)
	chk.UintBounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.UintBoundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.UintBounded(33, BoundedClosed, min, max)
	chk.UintBounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.UintBoundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.UintBounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "UintBounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "uint", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "uint", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "uint", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "uint", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUintUnboundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := uint(62)

	chk.UintUnbounded(60, UnboundedMinOpen, bound)
	chk.UintUnbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.UintUnboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.UintUnbounded(63, UnboundedMinOpen, bound)
	chk.UintUnbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.UintUnboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "UintUnbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "60", fName, "uint", ""),
		chkOutNumericUnbounded_(wntMsg, "61", fName, "uint", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "uint", "msg:62"),

		chkOutRelease(),
	)
}
