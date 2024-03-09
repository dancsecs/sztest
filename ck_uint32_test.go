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

func tstChkUint32(t *testing.T) {
	t.Run("Good", chkUint32TestGood)

	t.Run("Bad", chkUint32TestBad)
	t.Run("BadMsg1", chkUint32TestBad1)
	t.Run("BadMsg2", chkUint32TestBad2)
	t.Run("BadMsg3", chkUint32TestBad3)

	t.Run("Slice-Good", chkUint32SliceTestGood)
	t.Run("Slice-BadMsg1", chkUint32SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkUint32SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkUint32SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkUint32SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkUint32SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkUint32SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkUint32SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkUint32SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkUint32SliceTestBadMsg9)

	t.Run("Bounded", chkUint32BoundedTestAll)
	t.Run("Unbounded", chkUint32UnboundedTestAll)
}

func chkUint32TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32(0, 0)
	chk.Uint32(0, 0, "not ", "displayed")
	chk.Uint32f(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint32TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint32",
			chkOutCommonMsg("", "uint32"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint32TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint32f",
			chkOutCommonMsg("This message will be displayed first", "uint32"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint32TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint32",
			chkOutCommonMsg("This message will be displayed second", "uint32"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint32TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint32f",
			chkOutCommonMsg("This message will be displayed third", "uint32"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice(
		[]uint32{}, []uint32{},
		"This message will NOT be displayed",
	)
	chk.Uint32Slice(
		[]uint32{0}, []uint32{0},
		"This message will NOT be displayed",
	)
	chk.Uint32Slice(
		[]uint32{2, 6, 7}, []uint32{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.Uint32Slicef(
		[]uint32{2, 6, 7, 9}, []uint32{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slicef(
		[]uint32{2}, []uint32{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]uint32", "Uint32Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice(
		[]uint32{}, []uint32{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]uint32", "Uint32Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slicef(
		[]uint32{1, 2}, []uint32{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]uint32", "Uint32Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice(
		[]uint32{1}, []uint32{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]uint32", "Uint32Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slicef(
		[]uint32{1, 2}, []uint32{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint32", "Uint32Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice(
		[]uint32{2, 2}, []uint32{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint32", "Uint32Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slicef(
		[]uint32{1, 2}, []uint32{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint32", "Uint32Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice(
		[]uint32{1, 3}, []uint32{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]uint32", "Uint32Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkUint32SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint32Slice([]uint32{1, 3}, []uint32{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]uint32", "Uint32Slice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkUint32BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := uint32(33)
	max := uint32(35)

	// Bad: Error displayed.
	chk.Uint32Bounded(30, BoundedClosed, min, max)
	chk.Uint32Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Uint32Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Uint32Bounded(33, BoundedClosed, min, max)
	chk.Uint32Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Uint32Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Uint32Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Uint32Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, "uint32", ""),
		chkOutNumericBounded(wntMsg, "31", fName, "uint32", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "uint32", "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, "uint32", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUint32UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := uint32(62)

	// Bad: Error displayed.
	chk.Uint32Unbounded(60, UnboundedMinOpen, bound)
	chk.Uint32Unbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.Uint32Unboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.Uint32Unbounded(63, UnboundedMinOpen, bound)
	chk.Uint32Unbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.Uint32Unboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Uint32Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, "uint32", ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, "uint32", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "uint32", "msg:62"),

		chkOutRelease(),
	)
}
