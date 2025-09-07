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

func tstChkUint16(t *testing.T) {
	t.Run("Good", chkUint16TestGood)

	t.Run("Bad", chkUint16TestBad)
	t.Run("BadMsg1", chkUint16TestBad1)
	t.Run("BadMsg2", chkUint16TestBad2)
	t.Run("BadMsg3", chkUint16TestBad3)

	t.Run("Slice-Good", chkUint16SliceTestGood)
	t.Run("Slice-BadMsg1", chkUint16SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkUint16SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkUint16SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkUint16SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkUint16SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkUint16SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkUint16SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkUint16SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkUint16SliceTestBadMsg9)

	t.Run("Bounded", chkUint16BoundedTestAll)
	t.Run("Unbounded", chkUint16UnboundedTestAll)
}

func chkUint16TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16(0, 0)
	chk.Uint16(0, 0, "not ", "displayed")
	chk.Uint16f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint16TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16",
			chkOutCommonMsg("", uint16TypeName),
			g(markAsChg("2", "1", diffGot)),
			w(markAsChg("2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16f",
			chkOutCommonMsg(
				"This message will be displayed first",
				uint16TypeName,
			),
			g(markAsChg("2", "1", diffGot)),
			w(markAsChg("2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16",
			chkOutCommonMsg(
				"This message will be displayed second",
				uint16TypeName,
			),
			g(markAsChg("2", "1", diffGot)),
			w(markAsChg("2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16f",
			chkOutCommonMsg(
				"This message will be displayed third",
				uint16TypeName,
			),
			g(markAsChg("0", "1", diffGot)),
			w(markAsChg("0", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice(
		[]uint16{}, []uint16{},
		"This message will NOT be displayed",
	)
	chk.Uint16Slice(
		[]uint16{0}, []uint16{0},
		"This message will NOT be displayed",
	)
	chk.Uint16Slice(
		[]uint16{2, 6, 7}, []uint16{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.Uint16Slicef(
		[]uint16{2, 6, 7, 9}, []uint16{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slicef(
		[]uint16{2}, []uint16{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]uint16", "Uint16Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice(
		[]uint16{}, []uint16{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]uint16", "Uint16Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slicef(
		[]uint16{1, 2}, []uint16{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]uint16", "Uint16Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice(
		[]uint16{1}, []uint16{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]uint16", "Uint16Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slicef(
		[]uint16{1, 2}, []uint16{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint16", "Uint16Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice(
		[]uint16{2, 2}, []uint16{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint16", "Uint16Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slicef(
		[]uint16{1, 2}, []uint16{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint16", "Uint16Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice(
		[]uint16{1, 3}, []uint16{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint16", "Uint16Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16Slice([]uint16{1, 3}, []uint16{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]uint16", "Uint16Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkUint16BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := uint16(33)
	maxV := uint16(35)

	// Bad: Error displayed.
	chk.Uint16Bounded(30, BoundedClosed, minV, maxV)
	chk.Uint16Bounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.Uint16Boundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Uint16Bounded(33, BoundedClosed, minV, maxV)
	chk.Uint16Bounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.Uint16Boundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Uint16Bounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Uint16Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, uint16TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, uint16TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, uint16TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, uint16TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUint16UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := uint16(62)

	// Bad: Error displayed.
	chk.Uint16Unbounded(60, UnboundedMinOpen, bound)
	chk.Uint16Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Uint16Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Uint16Unbounded(63, UnboundedMinOpen, bound)
	chk.Uint16Unbounded(64, UnboundedMinOpen, bound, "mot ", "displayed")
	chk.Uint16Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "Uint16Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, uint16TypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, uint16TypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, uint16TypeName, "msg:62"),

		chkOutRelease(),
	)
}
