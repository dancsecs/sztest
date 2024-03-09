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

func tstChkUint16(t *testing.T) {
	t.Run("Good", chkUint16Test_Good)

	t.Run("Bad", chkUint16Test_Bad)
	t.Run("BadMsg1", chkUint16Test_Bad1)
	t.Run("BadMsg2", chkUint16Test_Bad2)
	t.Run("BadMsg3", chkUint16Test_Bad3)

	t.Run("Slice-Good", chkUint16SliceTest_Good)
	t.Run("Slice-BadMsg1", chkUint16SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkUint16SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkUint16SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkUint16SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkUint16SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkUint16SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkUint16SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkUint16SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkUint16SliceTest_BadMsg9)

	t.Run("Bounded", chkUint16BoundedTest_All)
	t.Run("Unbounded", chkUint16UnboundedTest_All)
}

func chkUint16Test_Good(t *testing.T) {
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

func chkUint16Test_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16",
			chkOutCommonMsg("", "uint16"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16Test_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16f(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16f",
			chkOutCommonMsg("This message will be displayed first", "uint16"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16Test_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16",
			chkOutCommonMsg("This message will be displayed second", "uint16"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16Test_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Uint16f(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Uint16f",
			chkOutCommonMsg("This message will be displayed third", "uint16"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkUint16SliceTest_Good(t *testing.T) {
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

func chkUint16SliceTest_BadMsg1(t *testing.T) {
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

func chkUint16SliceTest_BadMsg2(t *testing.T) {
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

func chkUint16SliceTest_BadMsg3(t *testing.T) {
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

func chkUint16SliceTest_BadMsg4(t *testing.T) {
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

func chkUint16SliceTest_BadMsg5(t *testing.T) {
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

func chkUint16SliceTest_BadMsg6(t *testing.T) {
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

func chkUint16SliceTest_BadMsg7(t *testing.T) {
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

func chkUint16SliceTest_BadMsg8(t *testing.T) {
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

func chkUint16SliceTest_BadMsg9(t *testing.T) {
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

func chkUint16BoundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := uint16(33)
	max := uint16(35)

	// Bad: Error displayed.
	chk.Uint16Bounded(30, BoundedClosed, min, max)
	chk.Uint16Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Uint16Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Uint16Bounded(33, BoundedClosed, min, max)
	chk.Uint16Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Uint16Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Uint16Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Uint16Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "uint16", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "uint16", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "uint16", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "uint16", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkUint16UnboundedTest_All(t *testing.T) {
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

	const wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
	const fName = "Uint16Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "60", fName, "uint16", ""),
		chkOutNumericUnbounded_(wntMsg, "61", fName, "uint16", "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, "uint16", "msg:62"),

		chkOutRelease(),
	)
}
