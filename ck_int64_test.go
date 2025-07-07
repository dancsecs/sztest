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
	t.Run("Good", chkInt64TestGood)

	t.Run("Bad", chkInt64TestBad)
	t.Run("BadMsg1", chkInt64TestBad1)
	t.Run("BadMsg2", chkInt64TestBad2)
	t.Run("BadMsg3", chkInt64TestBad3)

	t.Run("Slice-Good", chkInt64SliceTestGood)
	t.Run("Slice-BadMsg1", chkInt64SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkInt64SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkInt64SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkInt64SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkInt64SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkInt64SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkInt64SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkInt64SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkInt64SliceTestBadMsg9)

	t.Run("Bounded", chkInt64BoundedTestAll)
	t.Run("Unbounded", chkInt64UnboundedTestAll)
}

func chkInt64TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int64(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64",
			chkOutCommonMsg("", int64TypeName),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int64f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64f",
			chkOutCommonMsg(
				"This message will be displayed first",
				int64TypeName,
			),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int64(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64",
			chkOutCommonMsg(
				"This message will be displayed second",
				int64TypeName,
			),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int64f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int64f",
			chkOutCommonMsg(
				"This message will be displayed third",
				int64TypeName,
			),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt64SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt64BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := int64(33)
	maxV := int64(35)

	// Bad: Error displayed.
	chk.Int64Bounded(30, BoundedClosed, minV, maxV)
	chk.Int64Bounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.Int64Boundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Int64Bounded(33, BoundedClosed, minV, maxV)
	chk.Int64Bounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.Int64Boundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Int64Bounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Int64Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, int64TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, int64TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, int64TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, int64TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt64UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "Int64Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, int64TypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, int64TypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, int64TypeName, "msg:62"),

		chkOutRelease(),
	)
}
