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
	"errors"
	"fmt"
	"testing"
)

func tstChkCore(t *testing.T) {
	t.Run("CaptureNothing", chkTestCaptureNothing)
	t.Run("Logf", chkTestLogf)
	t.Run("ErrorPassthrough", chkTestErrorPassthrough)
	t.Run("ErrorfPassthrough", chkTestErrorfPassthrough)
	t.Run("FatalPassthrough", chkTestFatalPassthrough)
	t.Run("NamePassthrough", chkTestNamePassthrough)
	t.Run("T", chkTestT)
	t.Run("KeepTmpFilesSet", chkTestKeepTmpFileSet)
	t.Run("FailFast", chkTestFailFast)
	t.Run("PushPreReleaseFunc", chkTestPushPreReleaseFunc)
	t.Run("PushPostReleaseFunc", chkTestPushPostReleaseFunc)
	t.Run("PushPostReleaseFuncWithError", chkTestPushPostReleaseFuncWithError)
	t.Run("ReleaseWithUnexpectedPanic", chkTestReleaseWithUnexpectedPanic)
}

func tstChkGeneric(t *testing.T) {
	t.Run("Is", chkTestChkIs)
	t.Run("IsSlice", chkTestChkIsSlice)
	t.Run(
		"InBoundedRangeUnknownBoundedOption",
		chkTestInBoundedRangeUnknownBoundedOption,
	)
	t.Run("InBoundedRange_Open", chkTestTstBoundedRangeOpen)
	t.Run("InBoundedRange_Closed", chkTestTstBoundedRangeClosed)
	t.Run(
		"InBoundedRange_OpenMinOrClosedMax",
		chkTestTstBoundedRangeOpenMinOrClosedMax,
	)
	t.Run(
		"InBoundedRange_OpenMaxOrClosedMin",
		chkTestTstBoundedRangeOpenMaxOrClosedMin,
	)
	t.Run(
		"InUnboundedRangeUnknownBoundedOption",
		chkTestInUnboundedRangeUnknownBoundedOption,
	)
	t.Run("InUnboundedRange_Open", chkTestTstUnboundedRangeOpen)
	t.Run("InUnboundedRange_Closed", chkTestTstUnboundedRangeClosed)
}

// Simply exercises the create Chk processes and the "to be deferred" Release
// function.  This pattern of using the iTst object will be used whenever
// actions taken against &iTst (or &testing.T) need to be confirmed by
// checking the log in the iT.check function.
func chkTestCaptureNothing(t *testing.T) {
	// Create a stand in object to intercept operations made against t.
	iT := new(iTst)

	// Using the assert naming convention.  NOTE:  This package is intended to
	// be used by golang testing functions and not in production code
	assert := CaptureNothing(iT)
	iT.chk = assert

	assert.Release() // Manually issue the defer.

	// Using the chk (check) naming convention used throughout all subsequent
	// tests.
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Release() // Manually issue the defer.

	// Check actions taken against the testing.T object stand in.
	iT.check(t,
		// Output from assert:= CaptureNothing(iT).
		chkOutCapture("Nothing"),
		chkOutRelease(),

		// Output from chk := CaptureNothing(iT).
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkTestLogf(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Logf("message 1")
	chk.Logf("message %d", 2)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutLogf("message 1"),
		chkOutLogf("message 2"),
		chkOutRelease(),
	)
}

func chkTestErrorPassthrough(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Error("here is error number 1")

	chk.Error("here is error number 2")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutError("here is error number 1"),
		chkOutError("here is error number 2"),
		chkOutRelease(),
	)
}

func chkTestErrorfPassthrough(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Errorf("here is error number %d", 1)

	chk.Errorf("here is error number %d", 2)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutErrorf("here is error number 1"),
		chkOutErrorf("here is error number 2"),
		chkOutRelease(),
	)
}

func chkTestFatalPassthrough(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Fatalf("here is fatal error 1")

	chk.Fatalf("here is fatal error 2")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutFatalf("here is fatal error 1"),
		chkOutFatalf("here is fatal error 2"),
		chkOutRelease(),
	)
}

func chkTestNamePassthrough(t *testing.T) {
	const area = "name passthrough"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	got := chk.Name()
	wnt := "Internal Testing Object"

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkTestT(t *testing.T) {
	const area = "T value"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	got := fmt.Sprintf("%v", chk.T())
	wnt := fmt.Sprintf("%v", iT)

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	a := chk.T()

	// Invoke helper to insure object identifiers itself properly.
	a.Helper()

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		tstOutHelper("chkTestT"), // Helper output.
		chkOutRelease(),
	)
}

func chkTestKeepTmpFileSet(t *testing.T) {
	const area = "KeepTmpFiles"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	if chk.keepTmpFiles {
		t.Error(errGotWnt(area, chk.keepTmpFiles, false))
	}

	chk.KeepTmpFiles()

	if !chk.keepTmpFiles {
		t.Error(errGotWnt(area, chk.keepTmpFiles, true))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkTestFailFast(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	orig := chk.FailFast(false)
	defer chk.FailFast(orig)

	chk.Error("failFast: false / error number ", 1)
	chk.Error("failFast: false / error number ", 2)

	chk.FailFast(true)
	chk.Error("failFast: true  / error number ", 3)
	chk.Error("failFast: true  / error number ", 4)

	chk.FailFast(false)
	chk.Error("failFast: false / error number ", 5)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutErrorNoFail("failFast: false / error number 1"),
		chkOutErrorNoFail("failFast: false / error number 2"),
		chkOutError("failFast: true  / error number 3"),
		chkOutError("failFast: true  / error number 4"),
		chkOutErrorNoFail("failFast: false / error number 5"),
		chkOutRelease(),
	)
}

func chkTestPushPreReleaseFunc(t *testing.T) {
	const area = "push pre release func"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstFlag1 := false
	tstFlag2 := false

	chk.PushPreReleaseFunc(func() error {
		iT.output += "First Pushed A\n"
		tstFlag1 = true

		return nil
	})

	chk.PushPreReleaseFunc(func() error {
		iT.output += "Second Pushed A\n"
		tstFlag2 = true

		return nil
	})

	chk.Release()

	if !tstFlag1 {
		t.Error(errGotWnt(area, tstFlag1, true))
	}

	if !tstFlag2 {
		t.Error(errGotWnt(area, tstFlag2, true))
	}

	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutPush("Pre", ""),
		chkOutPush("Pre", ""),

		chkOutRelease(),

		chkOutPush("Pre", "func2"),
		"Second Pushed A",
		chkOutPush("Pre", "func1"),
		"First Pushed A",
	)
}

func chkTestPushPostReleaseFunc(t *testing.T) {
	const area = "push post Release func"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstFlag1 := false
	tstFlag2 := false

	chk.PushPostReleaseFunc(func() error {
		iT.output += "First Pushed\n"
		tstFlag1 = true

		return nil
	})

	chk.PushPostReleaseFunc(func() error {
		iT.output += "Second Pushed\n"
		tstFlag2 = true

		return nil
	})

	chk.Release()

	if !tstFlag1 {
		t.Error(errGotWnt(area, tstFlag1, true))
	}

	if !tstFlag2 {
		t.Error(errGotWnt(area, tstFlag2, true))
	}

	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutPush("Post", ""),
		chkOutPush("Post", ""),
		chkOutRelease(),
		chkOutPush("Post", "func2"),
		chkOutPush("Post", "func1"),
		"First Pushed",
		"Second Pushed",
	)
}

func chkTestPushPostReleaseFuncWithError(t *testing.T) {
	const area = "push post release with error"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstFlag1 := false
	tstFlag2 := false

	chk.PushPostReleaseFunc(func() error {
		iT.output += "First Pushed\n"
		tstFlag1 = true

		return nil
	})

	chk.PushPostReleaseFunc(func() error {
		iT.output += "Second Pushed\n"
		tstFlag2 = true

		return errors.New("Second release forced error")
	})

	chk.Release()

	if !tstFlag1 {
		t.Error(errGotWnt(area, tstFlag1, true))
	}

	if !tstFlag2 {
		t.Error(errGotWnt(area, tstFlag2, true))
	}

	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutPush("Post", ""),
		chkOutPush("Post", ""),
		chkOutRelease(),
		chkOutPush("Post", "func2"),
		chkOutPush("Post", "func1"),
		"First Pushed",
		"Second Pushed",
		chkOutFatalf("release caused error: Second release forced error"),
	)
}

//////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////
//  TEST generic functions
//////////////////////////////////////////////////////////////////////////

func chkTestChkIs(t *testing.T) {
	const area = "generic is"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var got, wnt bool

	got = chk.errChkf(false, true, "bool", "fMessage %s", "displayed")
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = chk.errChk(false, true, "bool", "message")
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError("f",
			chkOutCommonMsg("fMessage displayed", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutIsError("",
			chkOutCommonMsg("message", "bool"),
			g(markAsChg("false", "true", DiffGot)),
			w(markAsChg("false", "true", DiffWant)),
		),
		chkOutRelease(),
	)
}

//nolint:funlen // Ok.
func chkTestChkIsSlice(t *testing.T) {
	const area = "generic isSlice"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var got, wnt bool

	got = errSlice(chk,
		[]bool{true, false},
		[]bool{true, false},
		"bool",
		defaultCmpFunc[bool],
	)

	wnt = false
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = errSlicef(chk,
		[]bool{true, false},
		[]bool{true, false},
		"bool",
		defaultCmpFunc[bool],
		"%s", "message",
	)

	wnt = false
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = errSlice(
		chk,
		[]bool{true, false},
		[]bool{true, true},
		"bool",
		defaultCmpFunc[bool],
		"message",
	)

	wnt = false
	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(true, 2, 2, "[]bool", "",
			"",
			"0:0 true",
			"1:1 false",
		),
		chkOutIsSliceError(true, 2, 2, "[]bool", "f",
			"message",
			"0:0 true",
			"1:1 false",
		),
		chkOutIsSliceError(false, 2, 2, "[]bool", "",
			"message",
			"0:0 true",
			markAsChg("1", "", DiffGot)+
				":"+
				markAsChg("", "1", DiffWant)+
				" "+
				markAsChg("false", "true", DiffMerge),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////
//  TEST inBoundedRange
//////////////////////////////////////////////////////////////////////////

func chkTestInBoundedRangeUnknownBoundedOption(t *testing.T) {
	const area = "generic inBounded unknown option"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	expMsg := "unknown bounded option 1000000"
	inRange, gotMsg := inBoundedRange(0, BoundedOption(1000000), 0, 0)

	if inRange {
		t.Error(errGotWnt(area, true, false))
	}

	if gotMsg != expMsg {
		t.Error(errGotWnt(area, gotMsg, expMsg))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func tstBounded[V chkBoundedType](
	o BoundedOption, got, min, max V, wantAccumulator *string,
) bool {
	inRange, want := inBoundedRange(got, o, min, max)
	*wantAccumulator += want

	return inRange
}

//nolint:funlen // Ok.
func chkTestTstBoundedRangeOpen(t *testing.T) {
	const area = "generic inBoundedRange open"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	min := int(-2)
	max := int(2)

	got = tstBounded(BoundedOpen, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, -2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, 2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedOpen, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: (-2,2) - { want | -2 < want < 2 }" +
		"out of bounds: (-2,2) - { want | -2 < want < 2 }" +
		"out of bounds: (-2,2) - { want | -2 < want < 2 }" +
		"out of bounds: (-2,2) - { want | -2 < want < 2 }" +
		""

	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

//nolint:funlen // Ok.
func chkTestTstBoundedRangeClosed(t *testing.T) {
	const area = "generic inBoundedRange closed"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	min := int8(-2)
	max := int8(2)

	got = tstBounded(BoundedClosed, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, -2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, 2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedClosed, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: [-2,2] - { want | -2 <= want <= 2 }" +
		"out of bounds: [-2,2] - { want | -2 <= want <= 2 }" +
		""
	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

//nolint:cyclop,funlen // Ok.
func chkTestTstBoundedRangeOpenMinOrClosedMax(t *testing.T) {
	const area = "generic inBoundedRange OpenMinOrClosedMax"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	min := int8(-2)
	max := int8(2)

	got = tstBounded(BoundedMinOpen, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, -2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, 2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinOpen, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, -2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, 2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxClosed, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		"out of bounds: (-2,2] - { want | -2 < want <= 2 }" +
		""
	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

//nolint:cyclop,funlen // Ok.
func chkTestTstBoundedRangeOpenMaxOrClosedMin(t *testing.T) {
	const area = "generic inBoundedRange OpenMaxOrClosedMin"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	min := int16(-2)
	max := int16(2)

	got = tstBounded(BoundedMaxOpen, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, -2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, 2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMaxOpen, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, -3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, -2, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, -1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, -0, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, 1, min, max, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, 2, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstBounded(BoundedMinClosed, 3, min, max, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		"out of bounds: [-2,2) - { want | -2 <= want < 2 }" +
		""
	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////
//  TEST inUnboundedRange
//////////////////////////////////////////////////////////////////////////

func chkTestInUnboundedRangeUnknownBoundedOption(t *testing.T) {
	const area = "generic inUnboundedRange unknown option"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	expMsg := "unknown unbounded option 1000000"
	inRange, gotMsg := inUnboundedRange(0, UnboundedOption(1000000), 0)

	if inRange {
		t.Error(errGotWnt(area, true, false))
	}

	if gotMsg != expMsg {
		t.Error(errGotWnt(area, gotMsg, expMsg))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func tstUnbounded[V chkBoundedType](
	o UnboundedOption, got, bound V, wantAccumulator *string,
) bool {
	inRange, want := inUnboundedRange(got, o, bound)
	if want != "" {
		*wantAccumulator += want
	}

	return inRange
}

//nolint:funlen // Ok.
func chkTestTstUnboundedRangeOpen(t *testing.T) {
	const area = "generic inUnboundedRange open"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	bound := int32(-2)

	got = tstUnbounded(UnboundedMinOpen, -3, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMinOpen, -2, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMinOpen, -1, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxOpen, -3, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxOpen, -2, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxOpen, -1, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: (-2,MAX) - { want | want > -2 }" +
		"out of bounds: (-2,MAX) - { want | want > -2 }" +
		"out of bounds: (MIN,-2) - { want | want < -2 }" +
		"out of bounds: (MIN,-2) - { want | want < -2 }" +
		""

	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

//nolint:funlen // Ok.
func chkTestTstUnboundedRangeClosed(t *testing.T) {
	const area = "generic inUnboundedRange closed"

	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	var (
		got, wnt bool
		gotBuf   string
	)

	bound := int32(-2)

	got = tstUnbounded(UnboundedMinClosed, -3, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMinClosed, -2, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMinClosed, -1, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxClosed, -3, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxClosed, -2, bound, &gotBuf)
	wnt = true

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	got = tstUnbounded(UnboundedMaxClosed, -1, bound, &gotBuf)
	wnt = false

	if got != wnt {
		t.Error(errGotWnt(area, got, wnt))
	}

	wntBuf := "" +
		"out of bounds: [-2,MAX) - { want | want >= -2 }" +
		"out of bounds: (MIN,-2] - { want | want <= -2 }" +
		""
	if gotBuf != wntBuf {
		t.Error(errGotWnt(area, gotBuf, wntBuf))
	}

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func runChkTestReleaseWithUnexpectedPanic(_ *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	defer chk.Release()
	panic("abc")
}

func chkTestReleaseWithUnexpectedPanic(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			runChkTestReleaseWithUnexpectedPanic(t)
		},
		"abc",
	)
}
