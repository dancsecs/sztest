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

func tstChkInt32(t *testing.T) {
	t.Run("Good", chkInt32TestGood)

	t.Run("Bad", chkInt32TestBad)
	t.Run("BadMsg1", chkInt32TestBad1)
	t.Run("BadMsg2", chkInt32TestBad2)
	t.Run("BadMsg3", chkInt32TestBad3)

	t.Run("Slice-Good", chkInt32SliceTestGood)
	t.Run("Slice-BadMsg1", chkInt32SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkInt32SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkInt32SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkInt32SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkInt32SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkInt32SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkInt32SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkInt32SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkInt32SliceTestBadMsg9)

	t.Run("Bounded", chkInt32BoundedTestAll)
	t.Run("Unbounded", chkInt32UnboundedTestAll)
}

func chkInt32TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32(0, 0)
	chk.Int32(0, 0, "not ", "displayed")
	chk.Int32f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt32TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int32",
			chkOutCommonMsg("", "int32"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt32TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int32f",
			chkOutCommonMsg("This message will be displayed first", "int32"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt32TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int32",
			chkOutCommonMsg("This message will be displayed second", "int32"),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt32TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int32f",
			chkOutCommonMsg("This message will be displayed third", "int32"),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice(
		[]int32{}, []int32{},
		"This message will NOT be displayed",
	)
	chk.Int32Slice(
		[]int32{0}, []int32{0},
		"This message will NOT be displayed",
	)
	chk.Int32Slice(
		[]int32{2, 6, -7}, []int32{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.Int32Slicef(
		[]int32{2, 6, -7, 9}, []int32{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slicef(
		[]int32{2}, []int32{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]int32", "Int32Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice(
		[]int32{}, []int32{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]int32", "Int32Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slicef(
		[]int32{1, 2}, []int32{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]int32", "Int32Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice(
		[]int32{1}, []int32{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]int32", "Int32Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slicef(
		[]int32{1, 2}, []int32{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int32", "Int32Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice(
		[]int32{2, 2}, []int32{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int32", "Int32Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slicef(
		[]int32{1, 2}, []int32{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int32", "Int32Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice(
		[]int32{1, 3}, []int32{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int32", "Int32Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt32SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int32Slice([]int32{1, 3}, []int32{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]int32", "Int32Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkInt32BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := int32(33)
	max := int32(35)

	// Bad: Error displayed.
	chk.Int32Bounded(30, BoundedClosed, min, max)
	chk.Int32Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Int32Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Int32Bounded(33, BoundedClosed, min, max)
	chk.Int32Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Int32Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Int32Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Int32Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, "int32", ""),
		chkOutNumericBounded(wntMsg, "31", fName, "int32", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "int32", "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, "int32", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt32UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := int32(62)

	// Bad: Error displayed.
	chk.Int32Unbounded(60, UnboundedMinOpen, bound)
	chk.Int32Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Int32Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Int32Unbounded(63, UnboundedMinOpen, bound)
	chk.Int32Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Int32Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Int32Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, "int32", ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, "int32", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "int32", "msg:62"),

		chkOutRelease(),
	)
}
