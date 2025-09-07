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

func tstChkString(t *testing.T) {
	t.Run("Good", chkStringTestGood)

	t.Run("Bad", chkStringTestBad)
	t.Run("BadMsg1", chkStringTestBad1)
	t.Run("BadMsg2", chkStringTestBad2)
	t.Run("BadMsg3", chkStringTestBad3)

	t.Run("Slice-Good", chkStringSliceTestGood)
	t.Run("Slice-BadMsg1", chkStringSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkStringSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkStringSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkStringSliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkStringSliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkStringSliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkStringSliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkStringSliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkStringSliceTestBadMsg9)

	t.Run("Bounded", chkStringBoundedTestAll)
	t.Run("Unbounded", chkStringUnboundedTestAll)
}

func chkStringTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Str("", "")
	chk.Str("", "", "not ", "displayed")
	chk.Strf("", "", "not %s", "displayed")

	chk.Str("same", "same")
	chk.Str("same", "same", "not ", "displayed")
	chk.Strf("same", "same", "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkStringTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Str("Blank want", "")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Str",
			chkOutCommonMsg("", stringTypeName),
			g(markAsIns("Blank want")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkStringTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Strf("", "Blank got", "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Strf",
			chkOutCommonMsg(
				"This message will be displayed first",
				stringTypeName,
			),
			g(""),
			w(markAsDel("Blank got")),
		),
		chkOutRelease(),
	)
}

func chkStringTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Str("Blank want", "", "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Str",
			chkOutCommonMsg(
				"This message will be displayed second",
				stringTypeName,
			),
			g(markAsIns("Blank want")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkStringTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Strf("got", "want", "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Strf",
			chkOutCommonMsg(
				"This message will be displayed third",
				stringTypeName,
			),
			g(markAsChg("got", "want", diffGot)),
			w(markAsChg("got", "want", diffWant)),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice(
		[]string{}, []string{},
		"This message will NOT be displayed",
	)
	chk.StrSlice(
		[]string{"0"}, []string{"0"},
		"This message will NOT be displayed",
	)
	chk.StrSlice(
		[]string{"2", "6", "-7"}, []string{"2", "6", "-7"},
		"This message will NOT be displayed",
	)

	chk.StrSlicef(
		[]string{"2", "6", "-7", "9"}, []string{"2", "6", "-7", "9"},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlicef(
		[]string{"2"}, []string{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]string", "StrSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice(
		[]string{}, []string{"1"},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]string", "StrSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlicef(
		[]string{"1", "2"}, []string{"1"},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]string", "StrSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice(
		[]string{"1"}, []string{"1", "2"},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]string", "StrSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlicef(
		[]string{"1", "2"}, []string{"2", "2"},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]string", "StrSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice(
		[]string{"2", "2"}, []string{"1", "2"},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]string", "StrSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlicef(
		[]string{"1", "2"}, []string{"1", "3"},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]string", "StrSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice(
		[]string{"1", "3"}, []string{"1", "2"},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]string", "StrSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkStringSliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.StrSlice([]string{"1", "3"}, []string{"1", "2"})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]string", "StrSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkStringBoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := "33"
	maxV := "35"

	// Bad: Error displayed.
	chk.StrBounded("30", BoundedClosed, minV, maxV)
	chk.StrBounded("31", BoundedClosed, minV, maxV, "msg:", "31")
	chk.StrBoundedf("32", BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.StrBounded("33", BoundedClosed, minV, maxV)
	chk.StrBounded("34", BoundedClosed, minV, maxV, "not ", "displayed")
	chk.StrBoundedf("35", BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.StrBounded("36", BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [\"33\",\"35\"] - " +
			"{ want | \"33\" <= want <= \"35\" }"
		fName = "StrBounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutStringBounded(wntMsg, "30", fName, stringTypeName, ""),
		chkOutStringBounded(wntMsg, "31", fName, stringTypeName, "msg:31"),
		chkOutStringBoundedf(wntMsg, "32", fName, stringTypeName, "msg:32"),

		chkOutStringBounded(wntMsg, "36", fName, stringTypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkStringUnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := "62"

	// Bad: Error displayed.
	chk.StrUnbounded("60", UnboundedMinOpen, bound)
	chk.StrUnbounded("61", UnboundedMinOpen, bound, "msg:", "61")
	chk.StrUnboundedf("62", UnboundedMinOpen, bound, "msg:%d", 62)

	// Good:  No error displayed.
	chk.StrUnbounded("63", UnboundedMinOpen, bound)
	chk.StrUnbounded("64", UnboundedMinOpen, bound, "not ", "displayed")
	chk.StrUnboundedf("65", UnboundedMinOpen, bound, "not %s", "displayed")

	const (
		wntMsg = "out of bounds: (\"62\",MAX) - { want | want > \"62\" }"
		fName  = "StrUnbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutStrUnbounded(wntMsg, "60", fName, stringTypeName, ""),
		chkOutStrUnbounded(wntMsg, "61", fName, stringTypeName, "msg:61"),
		chkOutStrUnboundedf(wntMsg, "62", fName, stringTypeName, "msg:62"),

		chkOutRelease(),
	)
}
