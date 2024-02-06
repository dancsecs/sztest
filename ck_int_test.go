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

func tstChkInt(t *testing.T) {
	t.Run("Good", chkIntTest_Good)

	t.Run("Bad", chkIntTest_Bad)
	t.Run("BadMsg1", chkIntTest_Bad1)
	t.Run("BadMsg2", chkIntTest_Bad2)
	t.Run("BadMsg3", chkIntTest_Bad3)

	t.Run("Slice-Good", chkIntSliceTest_Good)
	t.Run("Slice-BadMsg1", chkIntSliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkIntSliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkIntSliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkIntSliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkIntSliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkIntSliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkIntSliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkIntSliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkIntSliceTest_BadMsg9)

	t.Run("Bounded", chkIntBoundedTest_All)
	t.Run("Unbounded", chkIntUnboundedTest_All)
}

func chkIntTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int(0, 0)
	chk.Int(0, 0, "not ", "displayed")
	chk.Intf(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkIntTest_Bad(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int",
			chkOutCommonMsg("", "int"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTest_Bad1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Intf(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Intf",
			chkOutCommonMsg("This message will be displayed first", "int"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTest_Bad2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Int(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int",
			chkOutCommonMsg("This message will be displayed second", "int"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTest_Bad3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Intf(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Intf",
			chkOutCommonMsg("This message will be displayed third", "int"),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_Good(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice(
		[]int{}, []int{},
		"This message will NOT be displayed",
	)
	chk.IntSlice(
		[]int{0}, []int{0},
		"This message will NOT be displayed",
	)
	chk.IntSlice(
		[]int{2, 6, -7}, []int{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.IntSlicef(
		[]int{2, 6, -7, 9}, []int{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg1(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlicef(
		[]int{2}, []int{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]int", "IntSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg2(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice(
		[]int{}, []int{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]int", "IntSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg3(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlicef(
		[]int{1, 2}, []int{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]int", "IntSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg4(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice(
		[]int{1}, []int{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]int", "IntSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg5(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlicef(
		[]int{1, 2}, []int{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int", "IntSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg6(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice(
		[]int{2, 2}, []int{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int", "IntSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg7(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlicef(
		[]int{1, 2}, []int{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int", "IntSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTest_BadMsg8(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice(
		[]int{1, 3}, []int{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int", "IntSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}
func chkIntSliceTest_BadMsg9(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.IntSlice([]int{1, 3}, []int{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]int", "IntSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkIntBoundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	min := int(33)
	max := int(35)

	// Bad: Error displayed.
	chk.IntBounded(30, BoundedClosed, min, max)
	chk.IntBounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.IntBoundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.IntBounded(33, BoundedClosed, min, max)
	chk.IntBounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.IntBoundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.IntBounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "IntBounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "int", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "int", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "int", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "int", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkIntUnboundedTest_All(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	bound := int(128)

	// Bad: Error displayed.
	chk.IntUnbounded(125, UnboundedMinClosed, bound)
	chk.IntUnbounded(126, UnboundedMinClosed, bound, "msg:", "126")
	chk.IntUnboundedf(127, UnboundedMinClosed, bound, "msg:%d", 127)

	// Good:  No error displayed.
	chk.IntUnbounded(128, UnboundedMinClosed, bound)
	chk.IntUnbounded(129, UnboundedMinClosed, bound, "not ", "displayed")
	chk.IntUnboundedf(130, UnboundedMinClosed, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: [128,MAX) - { want | want >= 128 }"
	const fName = "IntUnbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "125", fName, "int", ""),
		chkOutNumericUnbounded_(wntMsg, "126", fName, "int", "msg:126"),
		chkOutNumericUnboundedf(wntMsg, "127", fName, "int", "msg:127"),

		chkOutRelease(),
	)
}
