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
	t.Run("Good", chkUint8TestGood)

	t.Run("Bad", chkUint8TestBad)
	t.Run("BadMsg1", chkUint8TestBad1)
	t.Run("BadMsg2", chkUint8TestBad2)
	t.Run("BadMsg3", chkUint8TestBad3)

	t.Run("Slice-Good", chkUint8SliceTestGood)
	t.Run("Slice-BadMsg1", chkUint8SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkUint8SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkUint8SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkUint8SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkUint8SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkUint8SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkUint8SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkUint8SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkUint8SliceTestBadMsg9)

	t.Run("Bounded", chkUint8BoundedTestAll)
	t.Run("Unbounded", chkUint8UnboundedTestAll)
}

func chkUint8TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint8(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8",
			chkOutCommonMsg("", uint8TypeName),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint8f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8f",
			chkOutCommonMsg(
				"This message will be displayed first",
				uint8TypeName,
			),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint8(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8",
			chkOutCommonMsg(
				"This message will be displayed second",
				uint8TypeName,
			),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint8f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint8f",
			chkOutCommonMsg(
				"This message will be displayed third",
				uint8TypeName,
			),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint8SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

func chkUint8BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := uint8(33)
	maxV := uint8(35)

	// Bad: Error displayed.
	chk.Uint8Bounded(30, BoundedClosed, minV, maxV)
	chk.Uint8Bounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.Uint8Boundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Uint8Bounded(33, BoundedClosed, minV, maxV)
	chk.Uint8Bounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.Uint8Boundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Uint8Bounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Uint8Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, uint8TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, uint8TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, uint8TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, uint8TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUint8UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "Uint8Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, uint8TypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, uint8TypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, uint8TypeName, "msg:62"),

		chkOutRelease(),
	)
}
