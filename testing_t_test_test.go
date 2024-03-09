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
	"regexp"
	"strings"
	"testing"
)

/*
This file is named testingT_test_test and tests the helper functions defined
in testingT_test.  By separating these test the file testingT_test can be
temporarily renamed testingT_test_a permitted coverage information to be
collected.
*/

func test_SzTesting_Prerequisites(t *testing.T) {
	t.Run("Helper", testPrerequisiteHelper)
	t.Run("Logf", testPrerequisiteLogf)
	t.Run("Error", testPrerequisiteError)
	t.Run("Errorf", testPrerequisiteErrorf)
	t.Run("FailNow", testPrerequisiteFailNow)
	t.Run("SkipNow", testPrerequisiteSkipNow)
	t.Run("Name", testPrerequisiteName)
	t.Run("PrepareSlice", testPrerequisitePrepareSlice)
	t.Run("CheckGood", testPrerequisiteCheckGood)
	t.Run("CheckBad1", testPrerequisiteCheckBad1)
	t.Run("CheckBad2", testPrerequisiteCheckBad2)
	t.Run("CheckBad3", testPrerequisiteCheckBad3)

	// test chkOut Helpers

	t.Run("FreezeMarks", tstFreezeMarks)
	t.Run("chkOutHelper", tstChkOutHelper)
	t.Run("chkOutCapture", tstChkOutCapture)
	t.Run("chkOutRelease", tstChkOutRelease)
	t.Run("chkOutLogf", tstChkOutLogf)
	t.Run("chkOutErrorNoFail", tstChkOutErrorNoFail)
	t.Run("chkOutError", tstChkOutError)
	t.Run("chkOutErrorf", tstChkOutErrorf)
	t.Run("chkOutFatalf", tstChkOutFatalf)
	t.Run("chkOutPush", tstChkOutPush)
	t.Run("tstChkOutIsError", tstChkOutIsError)
	t.Run("tstChkOutCommonMsg", tstChkOutCommonMsg)
	t.Run("tstChkOutIsSliceError", tstChkOutIsSliceError)
	t.Run("tstChkOutNumericBoundedBad", tstChkOutNumericBoundedBad)
	t.Run("tstChkOutNumericUnboundedBad", tstChkOutNumericUnboundedBad)
	t.Run("tstChkOutStringBoundedBad", tstChkOutStringBoundedBad)
	t.Run("tstChkOutStringUnboundedBad", tstChkOutStringUnboundedBad)
}

func testPrerequisiteHelper(t *testing.T) {
	const tstName = "testPrerequisiteHelper"
	iT := new(iTst)

	iT.Helper()
	iT.Helper()

	// Check for iTst invocations.
	wntOut := "" +
		"Helper: " + tstName + "\n" +
		tstOutHelper(tstName) +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteName(t *testing.T) {
	const tstName = "testPrerequisiteName"
	iT := new(iTst)

	got := iT.Name()
	wnt := testName

	if got != wnt {
		t.Error(errGotWnt("it.Name()", got, wnt))
	}

	// Check for iTst invocations.
	wntOut := ""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteSkipNow(t *testing.T) {
	const tstName = "testPrerequisiteSkipNow"
	iT := new(iTst)

	iT.SkipNow()
	iT.SkipNow()

	// Check for iTst invocations.
	wntOut := "" +
		"Skip Now: " + tstName + "\n" +
		tstOutSkipNow(tstName) +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteFailNow(t *testing.T) {
	const tstName = "testPrerequisiteFailNow"
	iT := new(iTst)

	iT.FailNow()
	iT.FailNow()

	// Check for iTst invocations.
	wntOut := "" +
		"Fail Now: " + tstName + "\n" +
		tstOutFailNow(tstName) +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteError(t *testing.T) {
	const tstName = "testPrerequisiteError"
	iT := new(iTst)

	iT.Error("the error with no args")

	iT.Error("the error with ", 2, " args")

	// Check for iTst invocations.
	wntOut := "" +
		"Error: " + tstName + "\n" +
		"the error with no args\n" +
		tstOutError(tstName) +
		"the error with 2 args\n" +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteErrorf(t *testing.T) {
	const tstName = "testPrerequisiteErrorf"
	iT := new(iTst)

	iT.Errorf("the error with no args")

	iT.Errorf("the error with %d args", 1)

	// Check for iTst invocations.
	wntOut := "" +
		"Errorf: " + tstName + "\n" +
		"the error with no args\n" +
		tstOutErrorf(tstName) +
		"the error with 1 args\n" +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteLogf(t *testing.T) {
	const tstName = "testPrerequisiteLogf"
	iT := new(iTst)

	iT.Logf("the msg with no args")

	iT.Logf("the msg with %d args", 1)

	// Check for iTst invocations.
	wntOut := "" +
		"the msg with no args\n" +
		"the msg with 1 args\n" +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisitePrepareSlice(t *testing.T) {
	const tstName = "testPrerequisitePrepareSlice"
	iT := new(iTst)

	got := strings.Join(
		iT.prepareSlice(strings.TrimSpace,
			"",
			"firstLine",
			"\nsecondLine\nthirdLine",
			"fourthLine",
			"fifthLine",
			"\n\n\n",
			"\n\nsixthLine\n\n\nseventhLine\n\n\neighthLine\n\n\n\n",
			"   ",
			"ninthLine\n",
			"",
			" tenthLine   \n\n\n",
		),
		"\n",
	)
	wnt := "" +
		"firstLine\n" +
		"secondLine\n" +
		"thirdLine\n" +
		"fourthLine\n" +
		"fifthLine\n" +
		"sixthLine\n" +
		"seventhLine\n" +
		"eighthLine\n" +
		"ninthLine\n" +
		"tenthLine" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	// Check for iTst invocations.
	wntOut := ""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteCheckGood(t *testing.T) {
	const tstName = "testPrerequisiteCheckGood"
	iT := new(iTst)
	iTCheck := new(iTst)

	iT.output = "" +
		"firstLine\n" +
		"secondLine\n" +
		"thirdLine\n" +
		"fourthLine\n" +
		"fifthLine\n" +
		"sixthLine\n" +
		"seventhLine\n" +
		"eighthLine\n" +
		"ninthLine\n" +
		"tenthLine" +
		""
	iT.check(iTCheck,
		"",
		"firstLine",
		"\nsecondLine\nthirdLine",
		"fourthLine",
		"fifthLine",
		"\n\n\n",
		"\n\nsixthLine\n\n\nseventhLine\n\n\neighthLine\n\n\n\n",
		"   ",
		"ninthLine\n",
		"",
		" tenthLine   \n\n\n",
	)

	// Check for iTst invocations.
	wntOut := ""
	if iTCheck.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func testPrerequisiteCheckBad1(t *testing.T) {
	const tstName = "testPrerequisiteCheckBad1"
	iT := new(iTst)
	iTCheck := new(iTst)

	iT.output = "" +
		"firstLine\n" +
		"secondLine\n" +
		"thirdLine\n" +
		"fourthLine\n" +
		"fifthLine\n" +
		"extraLine\n" +
		"sixthLine\n" +
		"seventhLine\n" +
		"eighthLine\n" +
		"ninthLine\n" +
		"tenthLine" +
		""
	iT.check(iTCheck,
		"",
		"firstLine",
		"\nsecondLine\nthirdLine",
		"fourthLine",
		"fifthLine",
		"\n\n\n",
		"\n\nsixthLine\n\n\nseventhLine\n\n\neighthLine\n\n\n\n",
		"   ",
		"ninthLine\n",
		"",
		" tenthLine   \n\n\n",
	)

	// Check for iTst invocations.
	wntOut := "" +
		tstOutHelper("(*iTst).check") +
		tstOutErrorf("(*iTst).check") +
		"Unexpected Log Entry for Internal Testing Object:" +
		" got (11 lines) - want (10 lines)\n" +
		"00:00 firstLine\n" +
		"01:01 secondLine\n" +
		"02:02 thirdLine\n" +
		"03:03 fourthLine\n" +
		"04:04 fifthLine\n" +
		settingMarkInsOn + "05" + settingMarkInsOff + ":-- " +
		settingMarkInsOn + "extraLine" + settingMarkInsOff + "\n" +
		"06:05 sixthLine\n" +
		"07:06 seventhLine\n" +
		"08:07 eighthLine\n" +
		"09:08 ninthLine\n" +
		"10:09 tenthLine\n" +
		tstOutFailNow("(*iTst).check") +
		""
	if iTCheck.output != wntOut {
		t.Error(errGotWnt(tstName, iTCheck.output, wntOut))
	}
}

func testPrerequisiteCheckBad2(t *testing.T) {
	const tstName = "testPrerequisiteCheckBad2"
	iT := new(iTst)
	iTCheck := new(iTst)

	iT.output = "" +
		"firstLine\n" +
		"secondLine\n" +
		"thirdLine\n" +
		"fourthLine\n" +
		"fifthLine\n" +
		"sixthLine\n" +
		"seventhLine\n" +
		"eighthLine\n" +
		"ninthLine\n" +
		"tenthLine" +
		""
	iT.check(iTCheck,
		"",
		"firstLine",
		"\nsecondLine\nthirdLine",
		"fourthLine",
		"fifthLine",
		"\n\n\n",
		"extraLine",
		"\n\nsixthLine\n\n\nseventhLine\n\n\neighthLine\n\n\n\n",
		"   ",
		"ninthLine\n",
		"",
		" tenthLine   \n\n\n",
	)

	// Check for iTst invocations.
	wntOut := "" +
		tstOutHelper("(*iTst).check") +
		tstOutErrorf("(*iTst).check") +
		"Unexpected Log Entry for Internal Testing Object:" +
		" got (10 lines) - want (11 lines)\n" +
		"00:00 firstLine\n" +
		"01:01 secondLine\n" +
		"02:02 thirdLine\n" +
		"03:03 fourthLine\n" +
		"04:04 fifthLine\n" +
		"--:" + settingMarkDelOn + "05" + settingMarkDelOff + " " +
		settingMarkDelOn + "extraLine" + settingMarkDelOff + "\n" +
		"05:06 sixthLine\n" +
		"06:07 seventhLine\n" +
		"07:08 eighthLine\n" +
		"08:09 ninthLine\n" +
		"09:10 tenthLine\n" +
		tstOutFailNow("(*iTst).check") +
		""
	if iTCheck.output != wntOut {
		t.Error(errGotWnt(tstName, iTCheck.output, wntOut))
	}
}

func testPrerequisiteCheckBad3(t *testing.T) {
	const tstName = "testPrerequisiteCheckBad3"
	iT := new(iTst)
	iTCheck := new(iTst)

	// // Test for missing close mark.
	// got, err = freezeMarks(settingMarkInsOn)
	// if err == nil ||
	// 	err.Error() != `no closing mark found for "|-|InSOff|-|" in ""` {
	// 	t.Fatalf("unexpected error for: %v", err)
	// }
	// wnt = ""
	//
	// if got != wnt {
	// 	t.Error(errGotWnt(tstName, got, wnt))
	// }

	iT.output = settingMarkInsOn + "\n"
	iT.check(iTCheck,
		settingMarkInsOn,
	)

	squishTerminalCode := regexp.MustCompile(
		`(?m)\".*?\"`,
	)

	// Check for iTst invocations.
	wntOut := "" +
		tstOutHelper("(*iTst).check.func1") +
		tstOutError("(*iTst).check.func1") +
		`no closing mark found for "" in ""` + "\n" +
		tstOutFailNow("(*iTst).check.func1") +
		tstOutHelper("(*iTst).check.func2") +
		tstOutError("(*iTst).check.func2") +
		`no closing mark found for "" in ""` + "\n" +
		tstOutFailNow("(*iTst).check.func2") +
		""

	s := squishTerminalCode.ReplaceAllString(iTCheck.output, `""`)
	if s != wntOut {
		t.Error(errGotWnt(tstName, s, wntOut))
	}
}

func tstFreezeMarks(t *testing.T) {
	const tstName = "testFreezeMarks"
	iT := new(iTst)

	s := "No markups in string so will remain unchanged"
	got, err := freezeMarks(s)
	if err != nil {
		t.Fatal(err)
	}
	wnt := s

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	s = "this string has representations of all default marks used in " +
		"testing including empty strings:\n" +
		settingMarkInsOn + settingMarkInsOff + "\n" +
		settingMarkDelOn + settingMarkDelOff + "\n" +
		settingMarkChgOn + settingMarkChgOff + "\n" +
		settingMarkWntOn + settingMarkWntOff + "\n" +
		settingMarkGotOn + settingMarkGotOff + "\n" +
		settingMarkSepOn + settingMarkSepOff + "\n" +
		settingMarkMsgOn + settingMarkMsgOff + "\n" +
		settingMarkInsOn + "ins" + settingMarkInsOff + "\n" +
		settingMarkDelOn + "del" + settingMarkDelOff + "\n" +
		settingMarkChgOn + "chg" + settingMarkChgOff + "\n" +
		settingMarkWntOn + "wnt" + settingMarkWntOff + "\n" +
		settingMarkGotOn + "got" + settingMarkGotOff + "\n" +
		settingMarkSepOn + "sep" + settingMarkSepOff + "\n" +
		settingMarkMsgOn + "msg" + settingMarkMsgOff + "\n" +
		""
	got, err = freezeMarks(s)
	if err != nil {
		t.Fatal(err)
	}
	wnt = "this string has representations of all default marks used in " +
		"testing including empty strings:\n" +
		internalTestMarkInsOn + internalTestMarkInsOff + "\n" +
		internalTestMarkDelOn + internalTestMarkDelOff + "\n" +
		internalTestMarkChgOn + internalTestMarkChgOff + "\n" +
		internalTestMarkWntOn + internalTestMarkWntOff + "\n" +
		internalTestMarkGotOn + internalTestMarkGotOff + "\n" +
		internalTestMarkSepOn + internalTestMarkSepOff + "\n" +
		internalTestMarkMsgOn + internalTestMarkMsgOff + "\n" +
		internalTestMarkInsOn + "ins" + internalTestMarkInsOff + "\n" +
		internalTestMarkDelOn + "del" + internalTestMarkDelOff + "\n" +
		internalTestMarkChgOn + "chg" + internalTestMarkChgOff + "\n" +
		internalTestMarkWntOn + "wnt" + internalTestMarkWntOff + "\n" +
		internalTestMarkGotOn + "got" + internalTestMarkGotOff + "\n" +
		internalTestMarkSepOn + "sep" + internalTestMarkSepOff + "\n" +
		internalTestMarkMsgOn + "msg" + internalTestMarkMsgOff + "\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	// Test for missing close mark.
	got, err = freezeMarks(settingMarkInsOn)
	if err == nil ||
		err.Error() != `no closing mark found for "|-|InSOff|-|" in ""` {
		t.Fatalf("unexpected error for: %v", err)
	}
	wnt = ""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	// Test for wrong close mark.
	got, err = freezeMarks(settingMarkInsOn + "abc" + settingMarkGotOn)
	if err == nil ||
		err.Error() !=
			`unexpected closing mark: Got: "|-|GoTOn|-|"  Want: "|-|InSOff|-|"` {
		t.Fatalf("unexpected error for: %v", err)
	}
	wnt = ""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	// Check for iTst invocations.
	wntOut := "" +
		""
	if iT.output != wntOut {
		t.Error(errGotWnt(tstName, iT.output, wntOut))
	}
}

func tstChkOutHelper(t *testing.T) {
	const tstName = "tstChkOutHelper"

	got := chkOutHelper("FuncName")

	// Check for iTst invocations.
	wnt := "" +
		"Helper: (*Chk).FuncName\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutCapture(t *testing.T) {
	const tstName = "tstChkOutCapture"

	got := chkOutCapture("DataType")

	// Check for iTst invocations.
	wnt := "" +
		"Helper: CaptureDataType\n" +
		"Helper: newChk\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutRelease(t *testing.T) {
	const tstName = "tstChkOutRelease"

	got := chkOutRelease()

	// Check for iTst invocations.
	wnt := "" +
		"Helper: (*Chk).Release\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutLogf(t *testing.T) {
	const tstName = "tstChkOutLogf"

	got := chkOutLogf("<--MSG-->")

	// Check for iTst invocations.
	wnt := "" +
		"Helper: (*Chk).Logf\n" +
		"<--MSG-->\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutErrorNoFail(t *testing.T) {
	const tstName = "tstChkOutErrorNoFail"

	got := chkOutErrorfNoFail()

	// Check for iTst invocations.
	wnt := "" +
		"Helper: (*Chk).Errorf\n" +
		"Error: (*Chk).Errorf\n" +
		""
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutErrorfNoFail("Line 1")
	wnt += "Line 1\n"
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutErrorfNoFail("Line 1", "Line 2")
	wnt += "Line 2\n"
	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutError(t *testing.T) {
	const tstName = "tstChkOutError"

	got := chkOutError()
	wnt := "" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutError("message 1")
	wnt = "" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message 1\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutError("message 1", "message 2")
	wnt = "" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message 1\n" +
		"message 2\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutErrorf(t *testing.T) {
	const tstName = "tstChkOutErrorf"

	got := chkOutErrorf()
	wnt := "" +
		"Helper: (*Chk).Errorf\n" +
		"Error: (*Chk).Errorf\n" +
		"Fail Now: (*Chk).Errorf\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutErrorf("message 1")
	wnt = "" +
		"Helper: (*Chk).Errorf\n" +
		"Error: (*Chk).Errorf\n" +
		"message 1\n" +
		"Fail Now: (*Chk).Errorf\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutErrorf("message 1", "message 2")
	wnt = "" +
		"Helper: (*Chk).Errorf\n" +
		"Error: (*Chk).Errorf\n" +
		"message 1\n" +
		"message 2\n" +
		"Fail Now: (*Chk).Errorf\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutFatalf(t *testing.T) {
	const tstName = "tstChkOutFatalf"

	got := chkOutFatalf("the message")
	wnt := "" +
		"Helper: (*Chk).Fatalf\n" +
		"Error: (*Chk).Fatalf\n" +
		"the message\n" +
		"Fail Now: (*Chk).Fatalf\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutPush(t *testing.T) {
	const tstName = "tstChkOutPush"

	got := chkOutPush("Pre", "")
	wnt := "" +
		"Helper: (*Chk).PushPreReleaseFunc\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutPush("Post", "functionName")
	wnt = "" +
		"Helper: (*Chk).PushPostReleaseFunc.functionName\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutIsError(t *testing.T) {
	const tstName = "tstChkOutIsErrorf"

	got := chkOutIsError("", "message line 1")
	wnt := "" +
		"Helper: (*Chk).errChk\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsError("", "message line 1", "message line 2")
	wnt = "" +
		"Helper: (*Chk).errChk\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"message line 2\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsError("functionName", "message line 1")
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errChk\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsError("functionNamef", "message line 1")
	wnt = "" +
		"Helper: (*Chk).functionNamef\n" +
		"Helper: (*Chk).errChkf\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsError("functionName", "message line 1", "message line 2")
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errChk\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"message line 2\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsError("functionNamef", "message line 1", "message line 2")
	wnt = "" +
		"Helper: (*Chk).functionNamef\n" +
		"Helper: (*Chk).errChkf\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		"message line 1\n" +
		"message line 2\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutCommonMsg(t *testing.T) {
	const tstName = "tstChkOutCommonMsg"

	got := chkOutCommonMsg("", "dataType")
	wnt := "" +
		"unexpected dataType:" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutCommonMsg("message", "dataType")
	wnt = "" +
		"unexpected dataType:\n" + markMsgOn + "message" + markMsgOff + ":" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutIsSliceError(t *testing.T) {
	const tstName = "tstChkOutIsSliceError"

	got := chkOutIsSliceError(false, 1, 2, "[]dataType", "", "message1")
	wnt := "" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 1 Wnt: 2 [\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsSliceError(false, 1, 2, "[]dataType", "", "message1", "line1")
	wnt = "" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 1 Wnt: 2 [\n" +
		"line1\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsSliceError(
		false, 1, 2, "[]dataType", "", "message1", "line1", "line2",
	)
	wnt = "" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 1 Wnt: 2 [\n" +
		"line1\n" +
		"line2\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsSliceError(
		false, 2, 1, "[]dataType", "functionName", "message1",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 2 Wnt: 1 [\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsSliceError(
		false, 2, 1, "[]dataType", "functionName", "message1", "line1",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 2 Wnt: 1 [\n" +
		"line1\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutIsSliceError(
		false, 2, 1, "[]dataType", "functionName", "message1", "line1", "line2",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: errSlice[...]\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message1", "[]dataType") + "\n" +
		"Length Got: 2 Wnt: 1 [\n" +
		"line1\n" +
		"line2\n" +
		"]\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutNumericBoundedBad(t *testing.T) {
	const tstName = "tstChkOutNumericBoundedBad"

	got := chkOutNumericBounded_(
		"wantMessage", "got", "", "dataType", "message",
	)
	wnt := "" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutNumericBounded_(
		"wantMessage", "got", "functionName", "dataType", "message",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutNumericUnboundedBad(t *testing.T) {
	const tstName = "tstChkOutNumericUnboundedBad"

	got := chkOutNumericUnbounded_(
		"wantMessage", "got", "", "dataType", "message",
	)
	wnt := "" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutNumericUnbounded_(
		"wantMessage", "got", "functionName", "dataType", "message",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutStringBoundedBad(t *testing.T) {
	const tstName = "tstChkOutStringBoundedBad"

	got := chkOutStringBounded_(
		"wantMessage", "got", "", "dataType", "message",
	)
	wnt := "" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutStringBounded_(
		"wantMessage", "got", "functionName", "dataType", "message",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}

func tstChkOutStringUnboundedBad(t *testing.T) {
	const tstName = "tstChkOutStringUnboundedBad"

	got := chkOutStrUnbounded_(
		"wantMessage", "got", "", "dataType", "message",
	)
	wnt := "" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}

	got = chkOutStrUnbounded_(
		"wantMessage", "got", "functionName", "dataType", "message",
	)
	wnt = "" +
		"Helper: (*Chk).functionName\n" +
		"Helper: (*Chk).errGotWnt\n" +
		"Helper: (*Chk).Error\n" +
		"Error: (*Chk).Error\n" +
		chkOutCommonMsg("message", "dataType") + "\n" +
		g("got") + "\n" +
		w("wantMessage") + "\n" +
		"Fail Now: (*Chk).Error\n" +
		""

	if got != wnt {
		t.Error(errGotWnt(tstName, got, wnt))
	}
}
