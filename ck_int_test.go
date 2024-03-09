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
	t.Run("Good", chkIntTestGood)

	t.Run("Bad", chkIntTestBad)
	t.Run("BadMsg1", chkIntTestBad1)
	t.Run("BadMsg2", chkIntTestBad2)
	t.Run("BadMsg3", chkIntTestBad3)

	t.Run("Slice-Good", chkIntSliceTestGood)
	t.Run("Slice-BadMsg1", chkIntSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkIntSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkIntSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkIntSliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkIntSliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkIntSliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkIntSliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkIntSliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkIntSliceTestBadMsg9)

	t.Run("Bounded", chkIntBoundedTestAll)
	t.Run("Unbounded", chkIntUnboundedTestAll)
}

func chkIntTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int",
			chkOutCommonMsg("", intTypeName),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Intf(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Intf",
			chkOutCommonMsg("This message will be displayed first", intTypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int",
			chkOutCommonMsg("This message will be displayed second", intTypeName),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Intf(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Intf",
			chkOutCommonMsg("This message will be displayed third", intTypeName),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkIntSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntSliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkIntBoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "IntBounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, intTypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, intTypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, intTypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, intTypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkIntUnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: [128,MAX) - { want | want >= 128 }"
		fName  = "IntUnbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "125", fName, intTypeName, ""),
		chkOutNumericUnbounded(wntMsg, "126", fName, intTypeName, "msg:126"),
		chkOutNumericUnboundedf(wntMsg, "127", fName, intTypeName, "msg:127"),

		chkOutRelease(),
	)
}
