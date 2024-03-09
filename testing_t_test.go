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
	"fmt"
	"math"
	"runtime"
	"strings"
)

// Testing tracer implementation of the testingT interface.

const (
	testName = "Internal Testing Object"
)

type iTst struct {
	output string
	chk    *Chk
}

func (t *iTst) Helper() {
	t.output += "Helper: " + t.getCallerName() + "\n"
}

func (t *iTst) Logf(msgFmt string, msgArgs ...any) {
	t.output += fmt.Sprintf(msgFmt, msgArgs...) + "\n"
}

func (t *iTst) Errorf(msgFmt string, msgArgs ...any) {
	t.output += "Errorf: " + t.getCallerName() + "\n" +
		fmt.Sprintf(msgFmt, msgArgs...) + "\n"
}

func (t *iTst) Error(msgArgs ...any) {
	t.output += "Error: " + t.getCallerName() + "\n" +
		fmt.Sprint(msgArgs...) + "\n"
}

func (t *iTst) FailNow() {
	t.output += "Fail Now: " + t.getCallerName() + "\n"
}

func (t *iTst) SkipNow() {
	t.output += "Skip Now: " + t.getCallerName() + "\n"
}

func (t *iTst) Name() string {
	return testName
}

//

func (t *iTst) getCallerName() string {
	const (
		pkgPrefix  = `github.com/dancsecs/sztest.`
		stackDepth = 2
	)

	calledFrom := "<unknown>"

	pc, _, _, ok := runtime.Caller(stackDepth)
	if ok {
		details := runtime.FuncForPC(pc)
		if details != nil {
			calledFrom = details.Name()
		}
	}

	return strings.TrimPrefix(calledFrom, pkgPrefix)
}

func (*iTst) prepareSlice(
	processFunc func(string) string,
	rawLines ...string,
) []string {
	var lines []string

	for _, rl := range rawLines {
		for _, l := range strings.Split(rl, "\n") {
			l = strings.TrimSpace(l)
			if l != "" {
				l = processFunc(l)
				lines = append(lines, l)
			}
		}
	}

	return lines
}

func (t *iTst) check(testT testingT, rawLines ...string) {
	wantLines := t.prepareSlice(
		func(line string) string {
			line, err := freezeMarks(line)
			if err != nil {
				testT.Helper()
				testT.Error(err.Error())
				testT.FailNow()
			}

			return prepareWantString(line)
		},
		rawLines...,
	)

	gotLines := t.prepareSlice(
		func(line string) string {
			line, err := freezeMarks(line)
			if err != nil {
				testT.Helper()
				testT.Error(err.Error())
				testT.FailNow()
			}

			return line
		},
		t.output,
	)

	ret := CompareSlices(fmt.Sprint("Unexpected Log Entry for ", t.Name()),
		gotLines,
		wantLines,
		settingDiffSlice,
		settingDiffChars,
		defaultCmpFunc[string],
		stringify,
	)

	if ret != "" {
		testT.Helper()
		testT.Errorf("%s", resolveMarksForDisplay(ret))
		testT.FailNow()
	}
}

func tstOut(fName, caller string) string {
	return fName + ": " + caller + "\n"
}

func tstOutHelper(caller string) string {
	return tstOut("Helper", caller)
}

func tstOutSkipNow(caller string) string {
	return tstOut("Skip Now", caller)
}

func tstOutFailNow(caller string) string {
	return tstOut("Fail Now", caller)
}

func tstOutErrorf(caller string) string {
	return tstOut("Errorf", caller)
}

func tstOutError(caller string) string {
	return tstOut("Error", caller)
}

// CHK specific helpers.
func chkOutHelper(funcName string) string {
	return tstOutHelper("(*Chk)." + funcName)
}

func chkOutCapture(area string) string {
	return "" +
		tstOutHelper("Capture"+area) +
		tstOutHelper("newChk")
}

func chkOutRelease() string {
	return "" +
		chkOutHelper("Release")
}

func chkOutLogf(msg string) string {
	return "" +
		chkOutHelper("Logf") +
		msg + "\n"
}

func chkOutErrorNoFail(msg ...string) string {
	const caller = "(*Chk).Error"

	fullMsg := strings.Join(msg, "\n")
	if fullMsg != "" {
		fullMsg += "\n"
	}

	return "" +
		tstOutHelper(caller) +
		tstOutError(caller) +
		fullMsg +
		""
}

func chkOutErrorfNoFail(msg ...string) string {
	const caller = "(*Chk).Errorf"

	fullMsg := strings.Join(msg, "\n")
	if fullMsg != "" {
		fullMsg += "\n"
	}

	return "" +
		tstOutHelper(caller) +
		tstOutError(caller) +
		fullMsg +
		""
}

func chkOutError(msg ...string) string {
	const caller = "(*Chk).Error"

	return "" +
		chkOutErrorNoFail(msg...) +
		tstOutFailNow(caller) +
		""
}

func chkOutErrorf(msg ...string) string {
	const caller = "(*Chk).Errorf"

	return "" +
		chkOutErrorfNoFail(msg...) +
		tstOutFailNow(caller) +
		""
}

func chkOutFatalf(msg string) string {
	const caller = "(*Chk).Fatalf"

	return "" +
		tstOutHelper(caller) +
		tstOutError(caller) +
		msg + "\n" +
		tstOutFailNow(caller) +
		""
}

func chkOutPush(pos, subFunc string) string {
	if subFunc != "" {
		subFunc = "." + subFunc
	}

	return tstOutHelper("(*Chk).Push" + pos + "ReleaseFunc" + subFunc)
}

// Generic functions.

func chkOutIsError(caller, msg string, additionLines ...string) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		if caller != "f" {
			result += chkOutHelper(caller)
		}

		isFormatted = strings.HasSuffix(caller, "f")
	}

	if isFormatted {
		result += tstOutHelper("(*Chk).errChkf")
	} else {
		result += tstOutHelper("(*Chk).errChk")
	}

	result += chkOutError(append([]string{msg}, additionLines...)...)

	return result
}

func chkOutCommonMsg(msg, dataType string) string {
	if msg == "" {
		msg = commonMsgPrefix + dataType
	} else {
		msg = commonMsgPrefix + dataType + ":\n" + markMsgOn + msg + markMsgOff
	}

	return msg + ":"
}

func chkOutIsSliceError(
	identical bool,
	gNum, wNum int,
	dataType, caller, msg string,
	additionalLines ...string,
) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		if caller != "f" {
			result += tstOutHelper("(*Chk)." + caller)
		}

		isFormatted = strings.HasSuffix(caller, "f")
	}

	lines := make([]string, 0, len(additionalLines)+3)

	lines = append(lines, chkOutCommonMsg(msg, dataType))
	if identical {
		lines = append(lines, "invalid invocation: arrays are identical [")
	} else {
		lines = append(
			lines,
			fmt.Sprint("Length Got: ", gNum, " Wnt: ", wNum, " ["),
		)
	}

	lines = append(lines, additionalLines...)
	lines = append(lines, "]")

	if isFormatted {
		result += tstOutHelper("errSlicef[...]")
	} else {
		result += tstOutHelper("errSlice[...]")
	}

	result += chkOutError(lines...)

	return result
}

func chkOutLnSame(g, w, s string) string {
	return g + ":" + w + " " + s
}

func chkOutLnChanged(gLn, wLn, gStr string, wStr ...string) string {
	if len(wStr) == 1 {
		return "" +
			markAsChg(gLn, "", DiffGot) +
			":" +
			markAsChg("", wLn, DiffWant) +
			" " +
			markAsChg(gStr, wStr[0], DiffMerge)
	}

	return "" +
		markAsChg(gLn, "", DiffGot) +
		":" +
		markAsChg("", wLn, DiffWant) +
		" " +
		gStr
}

func chkOutLnGot(gLn, gStr string) string {
	return "" +
		markAsIns(gLn) +
		":" + strings.Repeat("-", len(gLn)) +
		" " +
		markAsIns(gStr)
}

func chkOutLnWnt(wLn, wStr string) string {
	return "" +
		strings.Repeat("-", len(wLn)) +
		":" +
		markAsDel(wLn) +
		" " +
		markAsDel(wStr)
}

func chkOutNumericBoundedf(
	wantMsg, got, caller, dataType, msg string,
) string {
	return chkOutNumericBounded(wantMsg, got, caller+"f", dataType, msg)
}

func chkOutNumericBounded(
	wantMsg, got, caller, dataType, msg string,
) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		result += tstOutHelper("(*Chk)." + caller)
		isFormatted = strings.HasSuffix(caller, "f")
	}

	if isFormatted {
		result += tstOutHelper("(*Chk).errGotWntf")
	} else {
		result += tstOutHelper("(*Chk).errGotWnt")
	}

	return result +
		chkOutError(
			chkOutCommonMsg(msg, dataType),
			g(got),
			w(wantMsg),
		)
}

func chkOutNumericUnboundedf(
	wantMsg, got, caller, dataType, msg string,
) string {
	return chkOutNumericUnbounded(
		wantMsg, got, caller+"f", dataType, msg,
	)
}

func chkOutNumericUnbounded(
	wantMsg, got, caller, dataType, msg string,
) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		result += tstOutHelper("(*Chk)." + caller)
		isFormatted = strings.HasSuffix(caller, "f")
	}

	if isFormatted {
		result += tstOutHelper("(*Chk).errGotWntf")
	} else {
		result += tstOutHelper("(*Chk).errGotWnt")
	}

	return result +
		chkOutError(
			chkOutCommonMsg(msg, dataType),
			g(got),
			w(wantMsg),
		)
}

func chkOutStringBoundedf(wantMsg, got, caller, dataType, msg string) string {
	return chkOutStringBounded(wantMsg, got, caller+"f", dataType, msg)
}

func chkOutStringBounded(wantMsg, got, caller, dataType, msg string) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		result += tstOutHelper("(*Chk)." + caller)
		isFormatted = strings.HasSuffix(caller, "f")
	}

	if isFormatted {
		result += tstOutHelper("(*Chk).errGotWntf")
	} else {
		result += tstOutHelper("(*Chk).errGotWnt")
	}

	return result +
		chkOutError(
			chkOutCommonMsg(msg, dataType),
			g(got),
			w(wantMsg),
		)
}

func chkOutStrUnboundedf(wantMsg, got, caller, dataType, msg string) string {
	return chkOutStrUnbounded(wantMsg, got, caller+"f", dataType, msg)
}

func chkOutStrUnbounded(wantMsg, got, caller, dataType, msg string) string {
	var (
		result      string
		isFormatted bool
	)

	if caller != "" {
		result += tstOutHelper("(*Chk)." + caller)
		isFormatted = strings.HasSuffix(caller, "f")
	}

	if isFormatted {
		result += tstOutHelper("(*Chk).errGotWntf")
	} else {
		result += tstOutHelper("(*Chk).errGotWnt")
	}

	return result +
		chkOutError(
			chkOutCommonMsg(msg, dataType),
			g(got),
			w(wantMsg),
		)
}

// Freezing test marks enabling regular error highlighting to be employed
// when using a test version of testing.T.

const (
	internalTestMarkDelOn  = "⨴"
	internalTestMarkDelOff = "⨵"
	internalTestMarkInsOn  = "⨭"
	internalTestMarkInsOff = "⨮"
	internalTestMarkChgOn  = "«"
	internalTestMarkChgOff = "»"
	internalTestMarkSepOn  = "⧚"
	internalTestMarkSepOff = "⧛"
	internalTestMarkGotOn  = "{"
	internalTestMarkGotOff = "}"
	internalTestMarkWntOn  = "["
	internalTestMarkWntOff = "]"
	internalTestMarkMsgOn  = "<"
	internalTestMarkMsgOff = ">"
)

// findNextMark searches the string for all known marks.
//
//nolint:cyclop // Ok.
func findNextMark(data, expectedClose string,
) (int, string, string, string) {
	if data == "" {
		return -1, "", "", ""
	}

	markOpenIndex := math.MaxInt
	markOpen := ""
	markOpenInternal := ""
	markOpenExpectedInternal := ""

	findOnMark := func(eOpenMark, iOpenMark, iCloseMark string) {
		tmpIndex := strings.Index(data, eOpenMark)
		if tmpIndex >= 0 && tmpIndex < markOpenIndex {
			markOpenIndex = tmpIndex
			markOpen = eOpenMark
			markOpenInternal = iOpenMark
			markOpenExpectedInternal = iCloseMark
		}
	}

	findOnMark(settingMarkInsOn, markInsOn, markInsOff)
	findOnMark(settingMarkDelOn, markDelOn, markDelOff)
	findOnMark(settingMarkChgOn, markChgOn, markChgOff)
	findOnMark(settingMarkWntOn, markWntOn, markWntOff)
	findOnMark(settingMarkGotOn, markGotOn, markGotOff)
	findOnMark(settingMarkSepOn, markSepOn, markSepOff)
	findOnMark(settingMarkMsgOn, markMsgOn, markMsgOff)

	markCloseIndex := math.MaxInt
	markClose := ""
	markCloseInternal := ""

	findOffMark := func(mark, internalMark string) {
		tmpIndex := strings.Index(data, mark)
		if tmpIndex >= 0 &&
			tmpIndex < markOpenIndex &&
			tmpIndex <= markCloseIndex {
			if tmpIndex == markCloseIndex && markCloseInternal == expectedClose {
				return
			}

			markCloseIndex = tmpIndex
			markClose = mark
			markCloseInternal = internalMark
		}
	}

	findOffMark(settingMarkInsOff, markInsOff)
	findOffMark(settingMarkDelOff, markDelOff)
	findOffMark(settingMarkChgOff, markChgOff)
	findOffMark(settingMarkWntOff, markWntOff)
	findOffMark(settingMarkGotOff, markGotOff)
	findOffMark(settingMarkSepOff, markSepOff)
	findOffMark(settingMarkMsgOff, markMsgOff)

	if markOpenIndex < math.MaxInt || markCloseIndex < math.MaxInt {
		if markOpenIndex < markCloseIndex {
			return markOpenIndex,
				markOpen,
				markOpenInternal,
				markOpenExpectedInternal
		}

		return markCloseIndex, markClose, markCloseInternal, ""
	}

	return -1, "", "", ""
}

func translateToTestSymbols(line string) string {
	line = strings.ReplaceAll(line, markDelOn, internalTestMarkDelOn)
	line = strings.ReplaceAll(line, markDelOff, internalTestMarkDelOff)
	line = strings.ReplaceAll(line, markInsOn, internalTestMarkInsOn)
	line = strings.ReplaceAll(line, markInsOff, internalTestMarkInsOff)
	line = strings.ReplaceAll(line, markChgOn, internalTestMarkChgOn)
	line = strings.ReplaceAll(line, markChgOff, internalTestMarkChgOff)
	line = strings.ReplaceAll(line, markSepOn, internalTestMarkSepOn)
	line = strings.ReplaceAll(line, markSepOff, internalTestMarkSepOff)
	line = strings.ReplaceAll(line, markWntOn, internalTestMarkWntOn)
	line = strings.ReplaceAll(line, markWntOff, internalTestMarkWntOff)
	line = strings.ReplaceAll(line, markGotOn, internalTestMarkGotOn)
	line = strings.ReplaceAll(line, markGotOff, internalTestMarkGotOff)
	line = strings.ReplaceAll(line, markMsgOn, internalTestMarkMsgOn)
	line = strings.ReplaceAll(line, markMsgOff, internalTestMarkMsgOff)

	return line
}

func freezeMarks(source string) (string, error) {
	var (
		iCloseMarkExpected string
		newS               string
	)

	for {
		i, eNextMark, iNextMark, iNextCloseMark := findNextMark(
			source, iCloseMarkExpected,
		)

		// If no more marks are present then we are done.  Either return the
		// translated string with the all marks reversed or an error if we are
		// expecting a close mark.
		if i < 0 {
			if iCloseMarkExpected != "" {
				return "", fmt.Errorf(
					"no closing mark found for %q in %q",
					iCloseMarkExpected,
					source,
				)
			}

			return translateToTestSymbols(newS + source), nil
		}

		// Otherwise we found a Mark.  Move all text up to the next mark from
		// the string to the translated string.
		if i > 0 {
			newS += source[:i]
			source = source[i:]
		}

		// Add the internal representation, replacing the resolved marks.
		newS += iNextMark

		// Remove the resolved Mark from the source string
		source = source[len(eNextMark):]

		if iCloseMarkExpected != "" {
			// There is an open mark that needs to be closed.
			if iNextMark != iCloseMarkExpected {
				return "", fmt.Errorf(
					"unexpected closing mark: Got: %q  Want: %q",
					iNextMark,
					iCloseMarkExpected,
				)
			}

			iCloseMarkExpected = ""
		} else {
			iCloseMarkExpected = iNextCloseMark
		}
	}
}
