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

func tstChkInt64(t *testing.T) {
	t.Run("Good", chkInt64Test_Good)

	t.Run("Bad", chkInt64Test_Bad)
	t.Run("BadMsg1", chkInt64Test_Bad1)
	t.Run("BadMsg2", chkInt64Test_Bad2)
	t.Run("BadMsg3", chkInt64Test_Bad3)

	t.Run("Slice-Good", chkInt64SliceTest_Good)
	t.Run("Slice-BadMsg1", chkInt64SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkInt64SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkInt64SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkInt64SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkInt64SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkInt64SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkInt64SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkInt64SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkInt64SliceTest_BadMsg9)

	t.Run("Bounded", chkInt64BoundedTest_All)
	t.Run("Unbounded", chkInt64UnboundedTest_All)
}

func chkInt64Test_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64(0, 0)
	chk.Int64(0, 0, "not ", "displayed")
	chk.Int64f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt64Test_Bad(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64",
			chkOutCommonMsg("", "int64"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64Test_Bad1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64f",
			chkOutCommonMsg("This message will be displayed first", "int64"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64Test_Bad2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64",
			chkOutCommonMsg("This message will be displayed second", "int64"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64Test_Bad3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64f",
			chkOutCommonMsg("This message will be displayed third", "int64"),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice(
		[]int64{}, []int64{},
		"This message will NOT be displayed",
	)
	chk.Int64Slice(
		[]int64{0}, []int64{0},
		"This message will NOT be displayed",
	)
	chk.Int64Slice(
		[]int64{2, 6, -7}, []int64{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.Int64Slicef(
		[]int64{2, 6, -7, 9}, []int64{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slicef(
		[]int64{2}, []int64{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]int64", "Int64Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice(
		[]int64{}, []int64{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]int64", "Int64Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slicef(
		[]int64{1, 2}, []int64{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]int64", "Int64Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg4(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice(
		[]int64{1}, []int64{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]int64", "Int64Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg5(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slicef(
		[]int64{1, 2}, []int64{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int64", "Int64Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg6(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice(
		[]int64{2, 2}, []int64{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int64", "Int64Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg7(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slicef(
		[]int64{1, 2}, []int64{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int64", "Int64Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg8(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice(
		[]int64{1, 3}, []int64{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int64", "Int64Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTest_BadMsg9(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int64Slice([]int64{1, 3}, []int64{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]int64", "Int64Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkInt64BoundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	min := int64(33)
	max := int64(35)

	// Bad: Error displayed.
	chk.Int64Bounded(30, BoundedClosed, min, max)
	chk.Int64Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Int64Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Int64Bounded(33, BoundedClosed, min, max)
	chk.Int64Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Int64Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Int64Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Int64Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "int64", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "int64", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "int64", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "int64", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt64UnboundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	bound := int64(62)

	// Bad: Error displayed.
	chk.Int64Unbounded(60, UnboundedMinOpen, bound)
	chk.Int64Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Int64Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Int64Unbounded(63, UnboundedMinOpen, bound)
	chk.Int64Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Int64Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Int64Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "60", fName, "int64", ""),
		chkOutNumericUnbounded_(wntMsg, "61", fName, "int64", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "int64", "msg:62"),

		chkOutRelease(),
	)
}
