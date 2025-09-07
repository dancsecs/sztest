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

func tstChkRune(t *testing.T) {
	t.Run("Good", chkRuneTestGood)

	t.Run("Bad", chkRuneTestBad)
	t.Run("BadMsg1", chkRuneTestBad1)
	t.Run("BadMsg2", chkRuneTestBad2)
	t.Run("BadMsg3", chkRuneTestBad3)

	t.Run("Slice-Good", chkRuneSliceTestGood)
	t.Run("Slice-BadMsg1", chkRuneSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkRuneSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkRuneSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkRuneSliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkRuneSliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkRuneSliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkRuneSliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkRuneSliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkRuneSliceTestBadMsg9)

	t.Run("Bounded", chkRuneBoundedTestAll)
	t.Run("Unbounded", chkRuneUnboundedTestAll)
}

func chkRuneTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Rune(0, 0)
	chk.Rune(0, 0, "not ", "displayed")
	chk.Runef(0, 0, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkRuneTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Rune(-2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Rune",
			chkOutCommonMsg("", runeTypeName),
			g(markAsChg("-2", "1", diffGot)),
			w(markAsChg("-2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkRuneTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Runef(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Runef",
			chkOutCommonMsg(
				"This message will be displayed first",
				runeTypeName,
			),
			g(markAsChg("2", "1", diffGot)),
			w(markAsChg("2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkRuneTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Rune(-2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Rune",
			chkOutCommonMsg(
				"This message will be displayed second",
				runeTypeName,
			),
			g(markAsChg("-2", "1", diffGot)),
			w(markAsChg("-2", "1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkRuneTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Runef(0, -1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Runef",
			chkOutCommonMsg(
				"This message will be displayed third",
				runeTypeName,
			),
			g(markAsChg("0", "-1", diffGot)),
			w(markAsChg("0", "-1", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice(
		[]rune{}, []rune{},
		"This message will NOT be displayed",
	)
	chk.RuneSlice(
		[]rune{0}, []rune{0},
		"This message will NOT be displayed",
	)
	chk.RuneSlice(
		[]rune{2, 6, -7}, []rune{2, 6, -7},
		"This message will NOT be displayed",
	)

	chk.RuneSlicef(
		[]rune{2, 6, -7, 9}, []rune{2, 6, -7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlicef(
		[]rune{2}, []rune{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]rune", "RuneSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice(
		[]rune{}, []rune{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]rune", "RuneSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlicef(
		[]rune{1, 2}, []rune{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]rune", "RuneSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice(
		[]rune{1}, []rune{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]rune", "RuneSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlicef(
		[]rune{1, 2}, []rune{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]rune", "RuneSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice(
		[]rune{2, 2}, []rune{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]rune", "RuneSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlicef(
		[]rune{1, 2}, []rune{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]rune", "RuneSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice(
		[]rune{1, 3}, []rune{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]rune", "RuneSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkRuneSliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.RuneSlice([]rune{1, 3}, []rune{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]rune", "RuneSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkRuneBoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := rune(33)
	maxV := rune(35)

	// Bad: Error displayed.
	chk.RuneBounded(30, BoundedClosed, minV, maxV)
	chk.RuneBounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.RuneBoundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.RuneBounded(33, BoundedClosed, minV, maxV)
	chk.RuneBounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.RuneBoundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.RuneBounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "RuneBounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, runeTypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, runeTypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, runeTypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, runeTypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkRuneUnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := rune(62)

	// Bad: Error displayed.
	chk.RuneUnbounded(60, UnboundedMinOpen, bound)
	chk.RuneUnbounded(61, UnboundedMinOpen, bound, "msg:", "61")
	chk.RuneUnboundedf(62, UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.RuneUnbounded(63, UnboundedMinOpen, bound)
	chk.RuneUnbounded(64, UnboundedMinOpen, bound, "not ", "displayed")
	chk.RuneUnboundedf(65, UnboundedMinOpen, bound, "not %s", "displayed")

	const (
		wntMsg = "out of bounds: (62,MAX) - { want | want > 62 }"
		fName  = "RuneUnbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "60", fName, runeTypeName, ""),
		chkOutNumericUnbounded(wntMsg, "61", fName, runeTypeName, "msg:61"),
		chkOutNumericUnboundedf(wntMsg, "62", fName, runeTypeName, "msg:62"),

		chkOutRelease(),
	)
}
