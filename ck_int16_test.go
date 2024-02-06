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

func tstChkInt16(t *testing.T) {
	t.Run("Good", chkInt16Test_Good)

	t.Run("Bad", chkInt16Test_Bad)
	t.Run("BadMsg1", chkInt16Test_Bad1)
	t.Run("BadMsg2", chkInt16Test_Bad2)
	t.Run("BadMsg3", chkInt16Test_Bad3)

	t.Run("Slice-Good", chkInt16SliceTest_Good)
	t.Run("Slice-BadMsg1", chkInt16SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkInt16SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkInt16SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkInt16SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkInt16SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkInt16SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkInt16SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkInt16SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkInt16SliceTest_BadMsg9)

	t.Run("Bounded", chkInt16BoundedTest_All)
	t.Run("Unbounded", chkInt16UnboundedTest_All)
}

func chkInt16Test_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16(0, 0)
	chk.Int16(0, 0, "not ", "displayed")
	chk.Int16f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt16Test_Bad(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16",
			chkOutCommonMsg("", "int16"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16Test_Bad1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16f",
			chkOutCommonMsg("This message will be displayed first", "int16"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16Test_Bad2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16",
			chkOutCommonMsg("This message will be displayed second", "int16"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16Test_Bad3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16f",
			chkOutCommonMsg("This message will be displayed third", "int16"),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice(
		[]int16{}, []int16{},
		"This message will NOT be displayed",
	)
	chk.Int16Slice(
		[]int16{0}, []int16{0},
		"This message will NOT be displayed",
	)
	chk.Int16Slice(
		[]int16{2, 6, -7}, []int16{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.Int16Slicef(
		[]int16{2, 6, -7, 9}, []int16{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slicef(
		[]int16{2}, []int16{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]int16", "Int16Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice(
		[]int16{}, []int16{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]int16", "Int16Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slicef(
		[]int16{1, 2}, []int16{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]int16", "Int16Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg4(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice(
		[]int16{1}, []int16{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]int16", "Int16Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg5(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slicef(
		[]int16{1, 2}, []int16{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int16", "Int16Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg6(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice(
		[]int16{2, 2}, []int16{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int16", "Int16Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg7(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slicef(
		[]int16{1, 2}, []int16{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int16", "Int16Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg8(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice(
		[]int16{1, 3}, []int16{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int16", "Int16Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTest_BadMsg9(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int16Slice([]int16{1, 3}, []int16{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]int16", "Int16Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkInt16BoundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	min := int16(33)
	max := int16(35)

	// Bad: Error displayed.
	chk.Int16Bounded(30, BoundedClosed, min, max)
	chk.Int16Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Int16Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Int16Bounded(33, BoundedClosed, min, max)
	chk.Int16Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Int16Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Int16Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Int16Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "int16", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "int16", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "int16", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "int16", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt16UnboundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	bound := int16(62)

	// Bad: Error displayed.
	chk.Int16Unbounded(60, UnboundedMinOpen, bound)
	chk.Int16Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Int16Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Int16Unbounded(63, UnboundedMinOpen, bound)
	chk.Int16Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Int16Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Int16Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "60", fName, "int16", ""),
		chkOutNumericUnbounded_(wntMsg, "61", fName, "int16", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "int16", "msg:62"),

		chkOutRelease(),
	)
}
