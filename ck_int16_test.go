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
	t.Run("Good", chkInt16TestGood)

	t.Run("Bad", chkInt16TestBad)
	t.Run("BadMsg1", chkInt16TestBad1)
	t.Run("BadMsg2", chkInt16TestBad2)
	t.Run("BadMsg3", chkInt16TestBad3)

	t.Run("Slice-Good", chkInt16SliceTestGood)
	t.Run("Slice-BadMsg1", chkInt16SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkInt16SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkInt16SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkInt16SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkInt16SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkInt16SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkInt16SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkInt16SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkInt16SliceTestBadMsg9)

	t.Run("Bounded", chkInt16BoundedTestAll)
	t.Run("Unbounded", chkInt16UnboundedTestAll)
}

func chkInt16TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int16(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16",
			chkOutCommonMsg("", int16TypeName),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int16f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16f",
			chkOutCommonMsg("This message will be displayed first", int16TypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int16(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16",
			chkOutCommonMsg("This message will be displayed second", int16TypeName),
			g(markAsChg("-2", "1", DiffGot)),
			w(markAsChg("-2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int16f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int16f",
			chkOutCommonMsg("This message will be displayed third", int16TypeName),
			g(markAsChg("0", "-1", DiffGot)),
			w(markAsChg("0", "-1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt16SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkInt16BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Int16Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, int16TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, int16TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, int16TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, int16TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt16UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "Int16Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, int16TypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, int16TypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, int16TypeName, "msg:62"),

		chkOutRelease(),
	)
}
