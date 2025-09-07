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

import (
	"testing"
)

func tstChkInt8(t *testing.T) {
	t.Run("Good", chkInt8TestGood)

	t.Run("Bad", chkInt8TestBad)
	t.Run("BadMsg1", chkInt8TestBad1)
	t.Run("BadMsg2", chkInt8TestBad2)
	t.Run("BadMsg3", chkInt8TestBad3)

	t.Run("Slice-Good", chkInt8SliceTestGood)
	t.Run("Slice-BadMsg1", chkInt8SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkInt8SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkInt8SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkInt8SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkInt8SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkInt8SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkInt8SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkInt8SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkInt8SliceTestBadMsg9)

	t.Run("Bounded", chkInt8BoundedTestAll)
	t.Run("Unbounded", chkInt8UnboundedTestAll)
}

func chkInt8TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8(0, 0)
	chk.Int8(0, 0, "not ", "displayed")
	chk.Int8f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt8TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int8",
			chkOutCommonMsg("", int8TypeName),
			g(markAsChg("-2", "1", diffGot)),
			w(markAsChg("-2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt8TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int8f",
			chkOutCommonMsg(
				"This message will be displayed first",
				int8TypeName,
			),
			g(markAsChg("2", "1", diffGot)),
			w(markAsChg("2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt8TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int8",
			chkOutCommonMsg(
				"This message will be displayed second",
				int8TypeName,
			),
			g(markAsChg("-2", "1", diffGot)),
			w(markAsChg("-2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt8TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8f(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Int8f",
			chkOutCommonMsg(
				"This message will be displayed third",
				int8TypeName,
			),
			g(markAsChg("0", "-1", diffGot)),
			w(markAsChg("0", "-1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice(
		[]int8{}, []int8{},
		"This message will NOT be displayed",
	)
	chk.Int8Slice(
		[]int8{0}, []int8{0},
		"This message will NOT be displayed",
	)
	chk.Int8Slice(
		[]int8{2, 6, -7}, []int8{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.Int8Slicef(
		[]int8{2, 6, -7, 9}, []int8{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slicef(
		[]int8{2}, []int8{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]int8", "Int8Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice(
		[]int8{}, []int8{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]int8", "Int8Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slicef(
		[]int8{1, 2}, []int8{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]int8", "Int8Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice(
		[]int8{1}, []int8{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]int8", "Int8Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slicef(
		[]int8{1, 2}, []int8{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int8", "Int8Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice(
		[]int8{2, 2}, []int8{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int8", "Int8Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slicef(
		[]int8{1, 2}, []int8{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int8", "Int8Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice(
		[]int8{1, 3}, []int8{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]int8", "Int8Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkInt8SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Int8Slice([]int8{1, 3}, []int8{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]int8", "Int8Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkInt8BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := int8(33)
	maxV := int8(35)

	// Bad: Error displayed.
	chk.Int8Bounded(30, BoundedClosed, minV, maxV)
	chk.Int8Bounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.Int8Boundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Int8Bounded(33, BoundedClosed, minV, maxV)
	chk.Int8Bounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.Int8Boundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Int8Bounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Int8Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, int8TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, int8TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, int8TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, int8TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkInt8UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := int8(62)

	// Bad: Error displayed.
	chk.Int8Unbounded(60, UnboundedMinOpen, bound)
	chk.Int8Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Int8Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Int8Unbounded(63, UnboundedMinOpen, bound)
	chk.Int8Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Int8Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "Int8Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, int8TypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, int8TypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, int8TypeName, "msg:62"),

		chkOutRelease(),
	)
}
