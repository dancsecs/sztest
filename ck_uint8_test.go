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

func tstChkUint8(t *testing.T) {
	t.Run("Good", chkUint8Test_Good)

	t.Run("Bad", chkUint8Test_Bad)
	t.Run("BadMsg1", chkUint8Test_Bad1)
	t.Run("BadMsg2", chkUint8Test_Bad2)
	t.Run("BadMsg3", chkUint8Test_Bad3)

	t.Run("Slice-Good", chkUint8SliceTest_Good)
	t.Run("Slice-BadMsg1", chkUint8SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkUint8SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkUint8SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkUint8SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkUint8SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkUint8SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkUint8SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkUint8SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkUint8SliceTest_BadMsg9)

	t.Run("Bounded", chkUint8BoundedTest_All)
	t.Run("Unbounded", chkUint8UnboundedTest_All)
}

func chkUint8Test_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8(0, 0)
	chk.Uint8(0, 0, "not ", "displayed")
	chk.Uint8f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint8Test_Bad(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8",
			chkOutCommonMsg("", "uint8"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8Test_Bad1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8f",
			chkOutCommonMsg("This message will be displayed first", "uint8"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8Test_Bad2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8",
			chkOutCommonMsg("This message will be displayed second", "uint8"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8Test_Bad3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8f",
			chkOutCommonMsg("This message will be displayed third", "uint8"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice(
		[]uint8{}, []uint8{},
		"This message will NOT be displayed",
	)
	chk.Uint8Slice(
		[]uint8{0}, []uint8{0},
		"This message will NOT be displayed",
	)
	chk.Uint8Slice(
		[]uint8{2, 6, 7}, []uint8{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.Uint8Slicef(
		[]uint8{2, 6, 7, 9}, []uint8{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slicef(
		[]uint8{2}, []uint8{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]uint8", "Uint8Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice(
		[]uint8{}, []uint8{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]uint8", "Uint8Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slicef(
		[]uint8{1, 2}, []uint8{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]uint8", "Uint8Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg4(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice(
		[]uint8{1}, []uint8{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]uint8", "Uint8Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg5(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slicef(
		[]uint8{1, 2}, []uint8{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint8", "Uint8Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg6(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice(
		[]uint8{2, 2}, []uint8{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint8", "Uint8Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg7(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slicef(
		[]uint8{1, 2}, []uint8{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint8", "Uint8Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg8(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice(
		[]uint8{1, 3}, []uint8{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint8", "Uint8Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTest_BadMsg9(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Uint8Slice([]uint8{1, 3}, []uint8{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]uint8", "Uint8Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkUint8BoundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	min := uint8(33)
	max := uint8(35)

	// Bad: Error displayed.
	chk.Uint8Bounded(30, BoundedClosed, min, max)
	chk.Uint8Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Uint8Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Uint8Bounded(33, BoundedClosed, min, max)
	chk.Uint8Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Uint8Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Uint8Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Uint8Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "uint8", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "uint8", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "uint8", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "uint8", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUint8UnboundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	bound := uint8(62)

	// Bad: Error displayed.
	chk.Uint8Unbounded(60, UnboundedMinOpen, bound)
	chk.Uint8Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Uint8Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Uint8Unbounded(63, UnboundedMinOpen, bound)
	chk.Uint8Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Uint8Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Uint8Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "60", fName, "uint8", ""),
		chkOutNumericUnbounded_(wntMsg, "61", fName, "uint8", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "uint8", "msg:62"),

		chkOutRelease(),
	)
}
