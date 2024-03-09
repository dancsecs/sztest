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

import "testing"

func tstChkBool(t *testing.T) {
	t.Run("Good", chkBoolTestGood)

	t.Run("Bad", chkBoolTestBad)
	t.Run("BadWithMsg1", chkBoolTestBadWithMsg1)
	t.Run("BadWithMsg2", chkBoolTestBadWithMsg2)
	t.Run("BadWithMsg3", chkBoolTestBadWithMsg3)
	t.Run("BadWithMsgFmt1", chkBoolTestBadWithMsgFmt1)
	t.Run("BadWithMsgFmt2", chkBoolTestBadWithMsgFmt2)
	t.Run("BadWithMsgFmt3", chkBoolTestBadWithMsgFmt3)

	t.Run("Slice-Good", chkBoolSliceTestGood)
	t.Run("Slice-BadWithMsg1", chkBoolSliceTestBadWithMessage1)
	t.Run("Slice-BadWithMsg2", chkBoolSliceTestBadWithMessage2)
	t.Run("Slice-BadWithMsg3", chkBoolSliceTestBadWithMessage3)
	t.Run("Slice-BadWithMsg4", chkBoolSliceTestBadWithMessage4)
	t.Run("Slice-BadWithMsg5", chkBoolSliceTestBadWithMessage5)
	t.Run("Slice-BadWithMsg6", chkBoolSliceTestBadWithMessage6)
	t.Run("Slice-BadWithMsg7", chkBoolSliceTestBadWithMessage7)
	t.Run("Slice-BadWithMsg8", chkBoolSliceTestBadWithMessage8)
	t.Run("Slice-BadWithMsg9", chkBoolSliceTestBadWithMessage9)

	t.Run("Helper-Good", chkBoolHelperTestGood)
	t.Run("Helper-Bad1", chkBoolHelperTestBad1)
	t.Run("Helper-Bad2", chkBoolHelperTestBad2)
	t.Run("Helper-Bad3", chkBoolHelperTestBad3)
	t.Run("Helper-Bad4", chkBoolHelperTestBad4)
}

func chkBoolTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bool(false, false)
	chk.Bool(true, true, "not ", "displayed")
	chk.Boolf(true, true, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkBoolTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bool(false, true)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Bool",
			chkOutCommonMsg("", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bool(false, true, "This message will be displayed first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Bool",
			chkOutCommonMsg("This message will be displayed first", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bool(true, false, "This message will ", "be displayed second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Bool",
			chkOutCommonMsg("This message will be displayed second", "bool"),
			g(markAsChg("true", "false", DiffGot)),
			w(markAsChg("true", "false", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bool(true, false, "This message will", " ", "be displayed third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Bool",
			chkOutCommonMsg("This message will be displayed third", "bool"),
			g(markAsChg("true", "false", DiffGot)),
			w(markAsChg("true", "false", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsgFmt1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Boolf(false, true, "This message will be displayed first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Boolf",
			chkOutCommonMsg("This message will be displayed first", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsgFmt2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Boolf(true, false, "This message will be displayed %s", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Boolf",
			chkOutCommonMsg("This message will be displayed second", "bool"),
			g(markAsChg("true", "false", DiffGot)),
			w(markAsChg("true", "false", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolTestBadWithMsgFmt3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Boolf(false, true, "This %s will be displayed %s", "message", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Boolf",
			chkOutCommonMsg("This message will be displayed third", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice(
		[]bool{}, []bool{},
		"This message will NOT be displayed",
	)
	chk.BoolSlice(
		[]bool{false}, []bool{false},
		"This message will NOT be displayed",
	)
	chk.BoolSlice(
		[]bool{true, true, false}, []bool{true, true, false},
		"This message will NOT be displayed",
	)

	chk.BoolSlicef(
		[]bool{true, true, false, true}, []bool{true, true, false, true},
		"This message will NOT be displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice(
		[]bool{true}, []bool{},
		"This message will be displayed ", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]bool", "BoolSlice",
			"This message will be displayed first",
			chkOutLnGot("0", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlicef(
		[]bool{}, []bool{false},
		"This message will be displayed %s", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]bool", "BoolSlicef",
			"This message will be displayed second",
			chkOutLnWnt("0", "false"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice(
		[]bool{false, true}, []bool{false},
		"This message will be displayed ", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]bool", "BoolSlice",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "false"),
			chkOutLnGot("1", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlicef(
		[]bool{false}, []bool{false, true},
		"This message will be displayed %s", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]bool", "BoolSlicef",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "false"),
			chkOutLnWnt("1", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlicef(
		[]bool{false, true}, []bool{true, true},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]bool", "BoolSlicef",
			"This message will be displayed fifth",
			markAsIns("0")+":- "+markAsIns("false"),
			chkOutLnSame("1", "0", "true"),
			chkOutLnWnt("1", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice(
		[]bool{true, true}, []bool{false, true},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]bool", "BoolSlice",
			"This message will be displayed sixth",
			"-:"+markAsDel("0")+" "+markAsDel("false"),
			chkOutLnSame("0", "1", "true"),
			chkOutLnGot("1", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlicef(
		[]bool{false, true}, []bool{false, false},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]bool", "BoolSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "false"),
			chkOutLnChanged("1", "1", "true", "false"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice(
		[]bool{false, false}, []bool{false, true},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]bool", "BoolSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "false"),
			chkOutLnChanged("1", "1", "false", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolSliceTestBadWithMessage9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.BoolSlice([]bool{false, false}, []bool{false, true})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]bool", "BoolSlice", "",
			chkOutLnSame("0", "0", "false"),
			chkOutLnChanged("1", "1", "false", "true"),
		),
		chkOutRelease(),
	)
}

func chkBoolHelperTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.False(false)
	chk.False(false, "This message will NOT be displayed")

	chk.Falsef(false, "This message will NOT be displayed")
	chk.Falsef(false, "This message will NOT be %s", "displayed")

	chk.True(true)
	chk.True(true, "This message will NOT be displayed")

	chk.Truef(true, "This message will NOT be displayed")
	chk.Truef(true, "This message will NOT be %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkBoolHelperTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Truef(false, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Truef",
			chkOutCommonMsg("This message will be displayed first", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolHelperTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Falsef(true, "This message will be displayed %s", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("Falsef",
			chkOutCommonMsg("This message will be displayed second", "bool"),
			g(markAsChg("true", "false", DiffGot)),
			w(markAsChg("true", "false", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolHelperTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.True(false, "This message will", " ", "be displayed third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("True",
			chkOutCommonMsg("This message will be displayed third", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkBoolHelperTestBad4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.False(true)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("False",
			chkOutCommonMsg("", "bool"),
			g(markAsChg("true", "false", DiffGot)),
			w(markAsChg("true", "false", DiffWant)),
		),
		chkOutRelease(),
	)
}
