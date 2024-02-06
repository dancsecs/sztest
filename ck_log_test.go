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
	"log"
	"log/slog"
	"os"
	"strings"
	"testing"
)

func tstChkLogging(t *testing.T) {
	t.Run("RemoveLogPrefixes", chkLogTest__RemoveLogPrefixes)
	t.Run("WriteNoneCaptureNothing", chkLogTest_WriteNoneCaptureNothing)
	t.Run("LogLoggerNoCheckLogger", chkLogTest_LogLoggerNoCheckLogger)
	t.Run("LogStdoutNoCheckStdout", chkLogTest_LogStdoutNoCheckStdout)
	t.Run("LogStdoutNoCheckStderr", chkLogTest_LogStdoutNoCheckStderr)
	t.Run("WriteNoneLogLogger", chkLogTest_WriteNoneLogLogger)
	t.Run("WriteNoneLogStderr", chkLogTest_WriteNoneLogStderr)
	t.Run("WriteNoneLogStdout", chkLogTest_WriteNoneLogStdout)
	t.Run("WriteNoneLogLoggerAndStderr", chkLogTest_WriteNoneLogLoggerAndStderr)
	t.Run("WriteNoneLogLoggerAndStdout", chkLogTest_WriteNoneLogLoggerAndStdout)
	t.Run("WriteNoneLogStderrAndStdout", chkLogTest_WriteNoneLogStderrAndStdout)
	t.Run("WriteNoneLogLoggerWithStderrChkStderr",
		chkLogTest_WriteNoneLogLoggerWithStderrChkStderr)
	t.Run("WriteNoneLogLoggerWithStderrChkLog",
		chkLogTest_WriteNoneLogLoggerWithStderrChkLog)
	t.Run("WriteNoneLogLoggerWithStderrChkBoth",
		chkLogTest_WriteNoneLogLoggerWithStderrChkBoth)
	t.Run("WriteNoneLogLoggerAndStderrAndStdout",
		chkLogTest_WriteNoneLogLoggerAndStderrAndStdout)
	t.Run("WriteNoneLogLoggerWithStderrAndStdoutChkStderr",
		chkLogTest_WriteNoneLogLoggerWithStderrAndStdoutChkStderr)
	t.Run("WriteNoneLogLoggerWithStderrAndStdOutChkLog",
		chkLogTest_WriteNoneLogLoggerWithStderrAndStdOutChkLog)
	t.Run("WriteNoneLogLoggerWithStderrAndStdOutChkBoth",
		chkLogTest_WriteNoneLogLoggerWithStderrAndStdOutChkBoth)
	t.Run("CheckLogging_NotStdoutWithExpected",
		chkLogTest_CheckLogging_NotStdoutWithExpected)
	t.Run("CheckLogging_NoLoggingWithExpected",
		chkLogTest_CheckLogging_NoLoggingWithExpected)
	t.Run("CheckLogging_NotStderrWithExpected",
		chkLogTest_CheckLogging_NotStderrWithExpected)
	t.Run("CheckLogging_StdoutMissing",
		chkLogTest_CheckLogging_StdoutMissing)
	t.Run("CheckLogging_LoggerMissing",
		chkLogTest_CheckLogging_LoggerMissing)
	t.Run("CheckLogging_StderrMissing",
		chkLogTest_CheckLogging_StderrMissing)
	t.Run("WriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck",
		chkLogTest_WriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck)
	t.Run("ReleasePanicInteractionPanicInternal",
		chkLogTest_ReleasePanicInteractionPanicInternal)
	t.Run("LeadingAndTrainingSpaces",
		chkLogTest_LeadingAndTrainingSpaces)
	t.Run("Slog", chkLogTest_Slog)
}

func chkLogTest__RemoveLogPrefixes(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	lastPrefix := log.Prefix()
	lastFlags := log.Flags()
	lastOutput := log.Writer()
	defer func() {
		log.SetPrefix(lastPrefix)
		log.SetFlags(lastFlags)
		log.SetOutput(lastOutput)
	}()

	createLogString := func(prefix string, flags int) string {
		rawLog := &strings.Builder{}

		log.SetPrefix(prefix)
		log.SetFlags(flags)

		log.SetOutput(rawLog)

		log.Print("This is line 1")
		log.Print("  This is line 2 prefixed by two spaces.")
		log.Print("Line number 3.")

		return rawLog.String()
	}

	finalResult := strings.Split(createLogString("", 0), "\n")
	f := 0
	p := ""

	explain := ""

	addExplanation := func(s string) {
		if explain != "" {
			explain += " | "
		}
		explain += s
	}

	for bPrefix := 0; bPrefix < 2; bPrefix++ {
		for bMsgPrefix := 0; bMsgPrefix < 2; bMsgPrefix++ {
			for bShortFile := 0; bShortFile < 2; bShortFile++ {
				for bLongFile := 0; bLongFile < 2; bLongFile++ {
					for bms := 0; bms < 2; bms++ {
						for btm := 0; btm < 2; btm++ {
							for bdt := 0; bdt < 2; bdt++ {
								f = 0
								explain = ""

								if bPrefix == 0 {
									p = ""
								} else {
									p = "PREFIX"
									addExplanation("PREFIX")
								}

								if bMsgPrefix != 0 {
									f |= log.Lmsgprefix
									addExplanation("Lmsgprefix")
								}

								if bShortFile != 0 {
									f |= log.Lshortfile
									addExplanation("Lshortfile")
								}

								if bLongFile != 0 {
									f |= log.Llongfile
									addExplanation("Llongfile")
								}

								if bms != 0 {
									f |= log.Lmicroseconds
									addExplanation("Lmicroseconds")
								}

								if btm != 0 {
									f |= log.Ltime
									addExplanation("Ltime")
								}

								if bdt != 0 {
									f |= log.Ldate
									addExplanation("Ldate")
								}

								chk.StrSlice(
									strings.Split(
										removeLogPrefixes(createLogString(p, f)),
										"\n",
									),
									finalResult,
									explain,
								)
							}
						}
					}
				}
			}
		}
	}
	iT.check(t,
		chkOutCapture("Nothing"),
	)
}

// Nothing Captured Nothing written.

func chkLogTest_WriteNoneCaptureNothing(t *testing.T) {
	iT := iTst{}
	chk := CaptureNothing(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
	)
}

// Nothing Checked.

func chkLogTest_LogLoggerNoCheckLogger(t *testing.T) {
	iT := iTst{}
	chk := CaptureLog(&iT)
	iT.chk = chk

	chk.Release()
	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
		chkOutHelper("setupLogLogger.func1"),
		chkOutError("log.Writer data was collected but never checked"),
	)
}

func chkLogTest_LogStdoutNoCheckStdout(t *testing.T) {
	iT := iTst{}
	chk := CaptureStdout(&iT)
	iT.chk = chk

	chk.Release()
	iT.check(t,
		chkOutCapture("Stdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
		chkOutHelper("setupStdoutLogger.func1"),
		chkOutError("os.Stdout data was collected but never checked"),
	)
}

func chkLogTest_LogStdoutNoCheckStderr(t *testing.T) {
	iT := iTst{}
	chk := CaptureStderr(&iT)
	iT.chk = chk

	chk.Release()
	iT.check(t,
		chkOutCapture("Stderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
		chkOutHelper("setupStderrLogger.func1"),
		chkOutError("os.Stderr data was collected but never checked"),
	)
}

// Nothing written.

func chkLogTest_WriteNoneLogLogger(t *testing.T) {
	iT := iTst{}
	chk := CaptureLog(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogStderr(t *testing.T) {
	iT := iTst{}
	chk := CaptureStderr(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Stderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogStdout(t *testing.T) {
	iT := iTst{}
	chk := CaptureStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Stdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerAndStderr(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogAndStderr(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogAndStderr"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerAndStdout(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogStderrAndStdout(t *testing.T) {
	iT := iTst{}
	chk := CaptureStderrAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("StderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrChkStderr(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderr(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrChkLog(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderr(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrChkBoth(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderr(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerAndStderrAndStdout(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogAndStderrAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogAndStderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrAndStdoutChkStderr(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderrAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrAndStdOutChkLog(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderrAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteNoneLogLoggerWithStderrAndStdOutChkBoth(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderrAndStdout(&iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_NotStdoutWithExpected(t *testing.T) {
	iT := iTst{}
	chk := CaptureStdout(&iT)
	iT.chk = chk

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout(`
		This line should be flagged as not in log
	`,
	)

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Stdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected stdout Entry: got (0 lines) - want (1 lines)",
		chkOutLnWnt("0", "This line should be flagged as not in log"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_NoLoggingWithExpected(t *testing.T) {
	iT := iTst{}
	chk := CaptureLog(&iT)
	iT.chk = chk

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout()

	chk.Log(`
		This line should be flagged as not in log
	`,
	)

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected log Entry: got (0 lines) - want (1 lines)",
		chkOutLnWnt("0", "This line should be flagged as not in log"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_NotStderrWithExpected(t *testing.T) {
	iT := iTst{}
	chk := CaptureStderr(&iT)
	iT.chk = chk

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout()

	chk.Log()

	chk.Stderr(`
		This line should be flagged as not in log
	`,
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Stderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected stderr Entry: got (0 lines) - want (1 lines)",
		chkOutLnWnt("0", "This line should be flagged as not in log"),
		"Fail Now: (*Chk).Error",
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_StdoutMissing(t *testing.T) {
	iT := iTst{}
	chk := CaptureStdout(&iT)
	iT.chk = chk

	fmt.Print("Forced error message\n")

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Stdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected stdout Entry: got (1 lines) - want (0 lines)",
		chkOutLnGot("0", "Forced error message"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_LoggerMissing(t *testing.T) {
	iT := iTst{}
	chk := CaptureLog(&iT)
	iT.chk = chk

	log.Print("Forced error message\n")

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected log Entry: got (1 lines) - want (0 lines)",
		chkOutLnGot("0", "Forced error message"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Stderr"),
		chkOutError("invalid os.Stderr check without information being captured"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_CheckLogging_StderrMissing(t *testing.T) {
	iT := iTst{}
	chk := CaptureStderr(&iT)
	iT.chk = chk

	fmt.Fprint(os.Stderr, "Forced error message\n")

	chk.markupForDisplay = func(s string) string {
		return s
	}

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Stderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected stderr Entry: got (1 lines) - want (0 lines)",
		chkOutLnGot("0", "Forced error message"),
		"Fail Now: (*Chk).Error",
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_WriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck(
	t *testing.T,
) {
	iT := iTst{}
	chk := CaptureLogWithStderr(&iT)
	iT.chk = chk

	chk.Str("this", "that")

	chk.Stdout()

	chk.Log()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutIsError(
			"Str",
			chkOutCommonMsg("", "string"),
			g(markAsChg("this", "that", DiffGot)),
			w(markAsChg("this", "that", DiffWant)),
		),
		chkOutHelper("Stdout"),
		chkOutError("invalid os.Stdout check without information being captured"),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_ReleasePanicInteractionPanicInternal(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogWithStderr(&iT)
	iT.chk = chk

	chk.Panic(
		func() {
			panic("panic message")
		},
		"panic message",
	)

	chk.Panic(
		func() {
			panic("panic \tmessage")
		},
		"panic \tmessage difference",
	)

	chk.Log(`
		`)

	chk.Release()

	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutIsError(
			"Panic",
			chkOutCommonMsg("", "panic"),
			g("panic \\tmessage"),
			w("panic \\tmessage"+markAsDel(" difference")),
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_LeadingAndTrainingSpaces(t *testing.T) {
	iT := iTst{}
	chk := CaptureLogAndStderrAndStdout(&iT)
	iT.chk = chk

	fmt.Println("   stdout   ")
	fmt.Fprintln(os.Stderr, "   stderr   ")
	log.Print("   logger   ")

	chk.Log(`
    \s  logger  \s
		`)

	chk.Stderr(`
    \s  stderr  \s
		`)

	chk.Stdout(`
    \s  stdout  \s
		`)

	chk.Release()

	iT.check(t,
		chkOutCapture("LogAndStderrAndStdout"),
		chkOutHelper("setupStdoutLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stdout"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTest_Slog(t *testing.T) {
	iT := iTst{}
	chk := CaptureLog(&iT)
	iT.chk = chk

	slog.Info("    logger     ")

	chk.Log(`INFO     logger    \s`)

	chk.Release()

	iT.check(t,
		chkOutCapture("Log"),
		chkOutHelper("setupLogLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}
