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
	t.Run("RemoveLogPrefixes", chkLogTestRemoveLogPrefixes)
	t.Run("WriteNoneCaptureNothing", chkLogTestWriteNoneCaptureNothing)
	t.Run("LogLoggerNoCheckLogger", chkLogTestLogLoggerNoCheckLogger)
	t.Run("LogStdoutNoCheckStdout", chkLogTestLogStdoutNoCheckStdout)
	t.Run("LogStdoutNoCheckStderr", chkLogTestLogStdoutNoCheckStderr)
	t.Run("WriteNoneLogLogger", chkLogTestWriteNoneLogLogger)
	t.Run("WriteNoneLogStderr", chkLogTestWriteNoneLogStderr)
	t.Run("WriteNoneLogStdout", chkLogTestWriteNoneLogStdout)
	t.Run("WriteNoneLogLoggerAndStderr", chkLogTestWriteNoneLogLoggerAndStderr)
	t.Run("WriteNoneLogLoggerAndStdout", chkLogTestWriteNoneLogLoggerAndStdout)
	t.Run("WriteNoneLogStderrAndStdout", chkLogTestWriteNoneLogStderrAndStdout)
	t.Run("WriteNoneLogLoggerWithStderrChkStderr",
		chkLogTestWriteNoneLogLoggerWithStderrChkStderr)
	t.Run("WriteNoneLogLoggerWithStderrChkLog",
		chkLogTestWriteNoneLogLoggerWithStderrChkLog)
	t.Run("WriteNoneLogLoggerWithStderrChkBoth",
		chkLogTestWriteNoneLogLoggerWithStderrChkBoth)
	t.Run("WriteNoneLogLoggerAndStderrAndStdout",
		chkLogTestWriteNoneLogLoggerAndStderrAndStdout)
	t.Run("WriteNoneLogLoggerWithStderrAndStdoutChkStderr",
		chkLogTestWriteNoneLogLoggerWithStderrAndStdoutChkStderr)
	t.Run("WriteNoneLogLoggerWithStderrAndStdOutChkLog",
		chkLogTestWriteNoneLogLoggerWithStderrAndStdOutChkLog)
	t.Run("WriteNoneLogLoggerWithStderrAndStdOutChkBoth",
		chkLogTestWriteNoneLogLoggerWithStderrAndStdOutChkBoth)
	t.Run("CheckLogging_NotStdoutWithExpected",
		chkLogTestCheckLoggingNotStdoutWithExpected)
	t.Run("CheckLogging_NoLoggingWithExpected",
		chkLogTestCheckLoggingNoLoggingWithExpected)
	t.Run("CheckLogging_NotStderrWithExpected",
		chkLogTestCheckLoggingNotStderrWithExpected)
	t.Run("CheckLogging_StdoutMissing",
		chkLogTestCheckLoggingStdoutMissing)
	t.Run("CheckLogging_LoggerMissing",
		chkLogTestCheckLoggingLoggerMissing)
	t.Run("CheckLogging_StderrMissing",
		chkLogTestCheckLoggingStderrMissing)
	t.Run("WriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck",
		chkLogTestWriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck,
	)
	t.Run("ReleasePanicInteractionPanicInternal",
		chkLogTestReleasePanicInteractionPanicInternal)
	t.Run("LeadingAndTrainingSpaces",
		chkLogTestLeadingAndTrainingSpaces)
	t.Run("Slog", chkLogTestSlog)
}

//nolint:cyclop,funlen,gocognit // Ok.
func chkLogTestRemoveLogPrefixes(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
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

	var (
		flags  int
		prefix string
	)

	explain := ""

	addExplanation := func(s string) {
		if explain != "" {
			explain += " | "
		}

		explain += s
	}

	for bPrefix := range 2 {
		for bMsgPrefix := range 2 {
			for bShortFile := range 2 {
				for bLongFile := range 2 {
					for bms := range 2 {
						for btm := range 2 {
							for bdt := range 2 {
								flags = 0
								explain = ""

								if bPrefix == 0 {
									prefix = ""
								} else {
									prefix = "PREFIX"

									addExplanation("PREFIX")
								}

								if bMsgPrefix != 0 {
									flags |= log.Lmsgprefix

									addExplanation("Lmsgprefix")
								}

								if bShortFile != 0 {
									flags |= log.Lshortfile

									addExplanation("Lshortfile")
								}

								if bLongFile != 0 {
									flags |= log.Llongfile

									addExplanation("Llongfile")
								}

								if bms != 0 {
									flags |= log.Lmicroseconds

									addExplanation("Lmicroseconds")
								}

								if btm != 0 {
									flags |= log.Ltime

									addExplanation("Ltime")
								}

								if bdt != 0 {
									flags |= log.Ldate

									addExplanation("Ldate")
								}

								chk.StrSlice(
									strings.Split(
										removeLogPrefixes(
											createLogString(prefix, flags),
										),
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

func chkLogTestWriteNoneCaptureNothing(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("Stdout"),
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutError(
			"invalid log.Writer check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
	)
}

// Nothing Checked.

func chkLogTestLogLoggerNoCheckLogger(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
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

func chkLogTestLogStdoutNoCheckStdout(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStdout(iT)
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

func chkLogTestLogStdoutNoCheckStderr(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStderr(iT)
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

func chkLogTestWriteNoneLogLogger(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogStderr(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
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

func chkLogTestWriteNoneLogStdout(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStdout(iT)
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
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogLoggerAndStderr(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogAndStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogLoggerAndStdout(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogAndStdout(iT)
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
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogStderrAndStdout(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStderrAndStdout(iT)
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

func chkLogTestWriteNoneLogLoggerWithStderrChkStderr(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderr(iT)
	iT.chk = chk

	chk.Stdout()

	chk.Stderr()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogLoggerWithStderrChkLog(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderr(iT)
	iT.chk = chk

	chk.Stdout()

	chk.Log()

	chk.Release()
	iT.check(t,
		chkOutCapture("LogWithStderr"),
		chkOutHelper("setupStderrLogger"),
		chkOutPush("Pre", ""),
		chkOutHelper("Stdout"),
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogLoggerWithStderrChkBoth(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Stderr"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestWriteNoneLogLoggerAndStderrAndStdout(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogAndStderrAndStdout(iT)
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

func chkLogTestWriteNoneLogLoggerWithStderrAndStdoutChkStderr(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderrAndStdout(iT)
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

func chkLogTestWriteNoneLogLoggerWithStderrAndStdOutChkLog(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderrAndStdout(iT)
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

func chkLogTestWriteNoneLogLoggerWithStderrAndStdOutChkBoth(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderrAndStdout(iT)
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

func chkLogTestCheckLoggingNotStdoutWithExpected(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStdout(iT)
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
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestCheckLoggingNoLoggingWithExpected(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected log Entry: got (0 lines) - want (1 lines)",
		chkOutLnWnt("0", "This line should be flagged as not in log"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Stderr"),
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestCheckLoggingNotStderrWithExpected(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
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

func chkLogTestCheckLoggingStdoutMissing(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStdout(iT)
	iT.chk = chk

	//nolint:forbidigo // Ok testing print capture.
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
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestCheckLoggingLoggerMissing(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutHelper("Error"),
		"Error: (*Chk).Error",
		"Unexpected log Entry: got (1 lines) - want (0 lines)",
		chkOutLnGot("0", "Forced error message"),
		"Fail Now: (*Chk).Error",
		chkOutHelper("Stderr"),
		chkOutError(
			"invalid os.Stderr check without information being captured",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestCheckLoggingStderrMissing(t *testing.T) {
	iT := new(iTst)
	chk := CaptureStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
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

func chkLogTestWriteLogLoggerWithStderrAndStdoutWithCheckErrorBlockingCheck(
	t *testing.T,
) {
	iT := new(iTst)
	chk := CaptureLogWithStderr(iT)
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
		chkOutError(
			"invalid os.Stdout check without information being captured",
		),
		chkOutHelper("Log"),
		chkOutHelper("compareLog"),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkLogTestReleasePanicInteractionPanicInternal(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogWithStderr(iT)
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

func chkLogTestLeadingAndTrainingSpaces(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLogAndStderrAndStdout(iT)
	iT.chk = chk

	//nolint:forbidigo // Ok testing print capture.
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

func chkLogTestSlog(t *testing.T) {
	iT := new(iTst)
	chk := CaptureLog(iT)
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
