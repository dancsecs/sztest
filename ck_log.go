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
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const pauseTimeForLogToCatchup = time.Millisecond * 5

// Log package prefix elements.
const (
	logDate   = `\d\d\d\d\/\d\d\/\d\d\s`
	logTime   = `\d\d:\d\d:\d\d\s`
	logTimeMS = `\d\d:\d\d:\d\d.\d\d\d\d\d\d\s`
	logFile   = `.+\..+\:\d+\:\s`
)

//nolint:goCheckNoGlobals // Caches log.Flag() and log.Prefix() regexps.
var logPrefixRegexpCache = make(map[string]*regexp.Regexp)

//nolint:cyclop // Ok.
func (chk *Chk) setupLoggers(option captureOption) {
	if option == captureStdout ||
		option == captureStderrAndStdout ||
		option == captureLogAndStdout ||
		option == captureLogAndStderrAndStdout ||
		option == captureLogWithStderrAndStdout {
		//
		chk.setupStdoutLogger()
	}

	if option == captureLog ||
		option == captureLogAndStderr ||
		option == captureLogAndStdout ||
		option == captureLogAndStderrAndStdout {
		//
		chk.setupLogLogger()
	}

	if option == captureStderr ||
		option == captureLogWithStderr ||
		option == captureLogAndStderr ||
		option == captureStderrAndStdout ||
		option == captureLogWithStderrAndStdout ||
		option == captureLogAndStderrAndStdout {
		//
		chk.setupStderrLogger(option == captureLogWithStderr ||
			option == captureLogWithStderrAndStdout)
	}
}

func (chk *Chk) setupStderrLogger(includeLog bool) {
	chk.t.Helper()

	var errBuf bytes.Buffer

	errBuf.Grow(settingBufferSize)

	chk.errOn = true
	chk.errBuf = &errBuf
	chk.errOrig = os.Stderr
	_ = chk.copyStderr()

	if includeLog {
		chk.errIncLog = true
		chk.logOrigLogFlags = log.Flags()
		chk.logOrig = log.Writer()
		log.SetFlags(0)
		log.SetOutput(os.Stderr)
	}

	chk.PushPreReleaseFunc(func() error {
		if !chk.errChecked {
			chk.t.Helper()

			if chk.faultCount == 0 {
				chk.Error("os.Stderr data was collected but never checked")
			}
		}

		t := os.Stderr
		os.Stderr = chk.errOrig
		err := t.Close()

		if includeLog {
			log.SetOutput(chk.logOrig)
			log.SetFlags(chk.logOrigLogFlags)
		}

		return err //nolint:wrapcheck // Ok
	})
}

func (chk *Chk) setupLogLogger() {
	chk.t.Helper()

	var logBuf bytes.Buffer

	logBuf.Grow(settingBufferSize)

	chk.logOn = true
	chk.logBuf = &logBuf
	chk.logOrigLogFlags = log.Flags()
	chk.logOrig = log.Writer()
	log.SetFlags(0)
	log.SetOutput(chk.logBuf)

	chk.PushPreReleaseFunc(func() error {
		if !chk.logChecked {
			chk.t.Helper()

			if chk.faultCount == 0 {
				chk.Error("log.Writer data was collected but never checked")
			}
		}

		log.SetOutput(chk.logOrig)
		log.SetFlags(chk.logOrigLogFlags)

		return nil
	})
}

func (chk *Chk) setupStdoutLogger() {
	chk.t.Helper()

	var outBuf bytes.Buffer

	outBuf.Grow(settingBufferSize)

	chk.outOn = true
	chk.outBuf = &outBuf
	chk.outOrig = os.Stdout
	_ = chk.copyStdout()

	chk.PushPreReleaseFunc(func() error {
		if !chk.outChecked {
			chk.t.Helper()

			if chk.faultCount == 0 {
				chk.Error("os.Stdout data was collected but never checked")
			}
		}

		t := os.Stdout
		os.Stdout = chk.outOrig

		return t.Close()
	})
}

func (chk *Chk) copyStdout() error {
	var (
		err   error
		rFile *os.File
		wFile *os.File
	)

	rFile, wFile, err = os.Pipe()
	if err == nil {
		go func() {
			defer func() {
				_ = rFile.Close()
			}()

			_, _ = io.Copy(chk.outBuf, rFile)
		}()

		os.Stdout = wFile
	}

	return err //nolint:wrapcheck // Ok.
}

func (chk *Chk) copyStderr() error {
	var (
		err   error
		rFile *os.File
		wFile *os.File
	)

	rFile, wFile, err = os.Pipe()
	if err == nil {
		go func() {
			defer func() {
				_ = rFile.Close()
			}()

			_, _ = io.Copy(chk.errBuf, rFile)
		}()

		os.Stderr = wFile
	}

	return err //nolint:wrapcheck // Ok.
}

// CaptureStdout returns a *Chk that captures os.Stdout.
//
// Call (*Chk).Stdout(wantLines...) to assert the captured stdout before
// calling chk.Release(). After Release the captured data is no longer
// available.
func CaptureStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStdout)
}

// CaptureLog returns a *Chk that captures the package logger (log.Writer()).
//
// Call (*Chk).Log(wantLines...) to assert captured log output before
// calling chk.Release().
func CaptureLog(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLog)
}

// CaptureLogAndStdout returns a *Chk that captures both log.Writer()
// and os.Stdout.
//
// Use (*Chk).Log(...) to assert the logger output and (*Chk).Stdout(...)
// to assert stdout. Perform these checks before calling chk.Release().
func CaptureLogAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStdout)
}

// CaptureLogAndStderr returns a *Chk that captures log.Writer() and os.Stderr.
//
// Use (*Chk).Log(...) to assert the logger output and (*Chk).Stderr(...)
// to assert stderr. Perform these checks before calling chk.Release().
func CaptureLogAndStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStderr)
}

// CaptureLogAndStderrAndStdout returns a *Chk that captures the package
// logger, os.Stderr and os.Stdout.
//
// Assert the captured streams with the corresponding methods
// ((*Chk).Log(...), (*Chk).Stdout(...) and (*Chk).Stderr(...)) before calling
// chk.Release().
func CaptureLogAndStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStderrAndStdout)
}

// CaptureLogWithStderr returns a *Chk that combines the package logger
// output and os.Stderr into a single capture buffer.
//
// In this combined mode the same underlying data may be inspected by either
// (*Chk).Log(...) or (*Chk).Stderr(...). Call exactly one of those two
// methods to assert the combined contents, and do so before calling
// chk.Release().
func CaptureLogWithStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogWithStderr)
}

// CaptureLogWithStderrAndStdout returns a *Chk that combines the package
// logger and os.Stderr into one capture buffer and also captures os.Stdout.
//
// Assert the combined logger/stderr with either (*Chk).Log(...) or
// (*Chk).Stderr(...), and assert stdout with (*Chk).Stdout(...). Do all
// assertions before calling chk.Release().
func CaptureLogWithStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogWithStderrAndStdout)
}

// CaptureStderr returns a *Chk that captures os.Stderr.
//
// Call (*Chk).Stderr(wantLines...) to assert the captured stderr before
// invoking chk.Release().
func CaptureStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStderr)
}

// CaptureStderrAndStdout returns a *Chk that captures both stderr and stdout.
//
// Call the corresponding assertion helpers ((*Chk).Stdout(...) and
// (*Chk).Stderr(...)) before calling chk.Release().
func CaptureStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStderrAndStdout)
}

// TrimAll normalizes a multi-line string into a compact form suitable for
// output assertions.
//
// It splits str into lines, removes common leading indentation, trims
// trailing whitespace from each line, and discards leading/trailing blank
// lines. The cleaned lines are then rejoined with a single '\n' between
// them and returned as one string.
//
// To preserve intentional leading or trailing spaces/tabs, replace the first
// or last space/tab with the escape markers `\s` or `\t`. This allows test
// data to remain both human-readable and assertion-accurate. It is especially
// useful when comparing against captured output via Log, Stdout, or Stderr.
func (chk *Chk) TrimAll(str string) string {
	lines := make([]string, 0, strings.Count(str, "\n")+1)

	for l := range strings.SplitSeq(str, "\n") {
		line := strings.TrimSpace(l)

		// Replace first hard space.
		if strings.HasPrefix(line, `\s`) {
			line = " " + line[2:]
		} else if strings.HasPrefix(line, `\t`) {
			line = "\t" + line[2:]
		}

		if strings.HasSuffix(line, `\s`) {
			line = line[:len(line)-2] + " "
		}

		lines = append(lines, line)
	}
	// remove leading blank lines
	firstPos := 0
	lastPos := len(lines)

	for firstPos < lastPos && lines[firstPos] == "" {
		firstPos++
	}
	// remove trailing blank lines
	for lastPos > firstPos && lines[lastPos-1] == "" {
		lastPos--
	}

	return strings.Join(lines[firstPos:lastPos], "\n")
}

func (chk *Chk) prepareSlice(
	processFunc func(string) string,
	rawLines ...string,
) []string {
	var lines []string

	for _, rl := range rawLines {
		for l := range strings.SplitSeq(rl, "\n") {
			l = processFunc(l)
			lines = append(lines, chk.isStringify(l))
		}
	}

	return lines
}

func (chk *Chk) compareLog(
	name, got string,
	gotFilter, wantFilter func(string) string,
	wantLines ...string,
) bool {
	chk.t.Helper()

	var gotSlice []string

	if got != "" {
		gotSlice = []string{strings.TrimSuffix(got, "\n")}
	}

	ret := CompareSlices(
		fmt.Sprint("Unexpected ", name, " Entry"),
		chk.prepareSlice(
			gotFilter,
			gotSlice...,
		),
		chk.prepareSlice(
			wantFilter,
			wantLines...,
		),
		settingDiffSlice,
		settingDiffChars,
		defaultCmpFunc[string],
		chk.isStringify,
	)

	if ret != "" {
		chk.Error(ret)

		return false
	}

	return true
}

func buildLogPrefixRegexpStr(prefix string, flags int) string {
	regExpStr := "(?m)^" // Multi-line regular expression for log header.
	if prefix != "" && (flags&log.Lmsgprefix) == 0 {
		regExpStr += prefix
	}

	if (flags & log.Ldate) != 0 {
		regExpStr += logDate
	}

	if flags&(log.Ltime|log.Lmicroseconds) != 0 {
		if (flags & log.Lmicroseconds) != 0 {
			regExpStr += logTimeMS
		} else {
			regExpStr += logTime
		}
	}

	if (flags & (log.Lshortfile | log.Llongfile)) != 0 {
		regExpStr += logFile
	}

	if prefix != "" && (flags&log.Lmsgprefix != 0) {
		regExpStr += prefix
	}

	return regExpStr
}

func removeLogPrefixes(line string) string {
	var (
		clearLogPrefix *regexp.Regexp
		ok             bool
	)

	logFlags := log.Flags()
	logPrefix := log.Prefix()

	if logFlags == 0 && logPrefix == "" {
		return line
	}

	cacheKey := fmt.Sprint(logPrefix, logFlags)
	if clearLogPrefix, ok = logPrefixRegexpCache[cacheKey]; !ok {
		re := buildLogPrefixRegexpStr(logPrefix, logFlags)
		clearLogPrefix = regexp.MustCompile(re)
		logPrefixRegexpCache[cacheKey] = clearLogPrefix
	}

	return clearLogPrefix.ReplaceAllString(line, "")
}

// Log compares the internally captured logger output against wantLines.
//
// Any decorations applied by the standard library log package (such as
// timestamps, optional flags, or prefixes) are stripped before comparison.
// This ensures that tests focus only on the actual log message content rather
// than on formatting applied by the logger itself.
//
// Returns true when the captured lines match exactly the supplied sequence.
// Failures are reported to the underlying testingT. Call this before
// chk.Release().
func (chk *Chk) Log(wantLines ...string) bool {
	chk.t.Helper()

	if !chk.logOn && !(chk.errOn && chk.errIncLog) {
		chk.Error(
			"invalid log.Writer check without information being captured",
		)

		return true
	}

	time.Sleep(pauseTimeForLogToCatchup)

	var (
		gotString string
		name      string
	)

	if chk.logOn {
		chk.logChecked = true
		name = "log"
		gotString = chk.logBuf.String()
	} else {
		chk.errChecked = true
		name = "logWithStderr"
		gotString = chk.errBuf.String()
	}

	return chk.compareLog(
		name,
		gotString,
		removeLogPrefixes,
		func(s string) string {
			return s
		},
		//		prepareWantString,
		wantLines...,
	)
}

// Stderr compares the internally captured stderr output against wantLines.
//
// It returns true on an exact match and reports test failures via the Chk's
// testingT. Call this before chk.Release().
func (chk *Chk) Stderr(wantLines ...string) bool {
	chk.t.Helper()

	if !(chk.errOn) {
		chk.Error("invalid os.Stderr check without information being captured")

		return true
	}

	time.Sleep(pauseTimeForLogToCatchup)

	chk.errChecked = true

	var (
		getFilterFunc func(string) string
		name          string
	)

	if chk.errIncLog {
		name = "logWithStderr"
		getFilterFunc = removeLogPrefixes
	} else {
		name = "stderr"
		getFilterFunc = func(s string) string {
			return s
		}
	}

	return chk.compareLog(
		name,
		chk.errBuf.String(),
		getFilterFunc,
		func(s string) string {
			return s
		},
		//		prepareWantString,
		wantLines...,
	)
}

// Stdout compares the internally captured stdout output against wantLines.
//
// It returns true on an exact match and reports test failures via the Chk's
// testingT. Call this before chk.Release().
func (chk *Chk) Stdout(wantLines ...string) bool {
	chk.t.Helper()

	if !(chk.outOn) {
		chk.Error("invalid os.Stdout check without information being captured")

		return true
	}

	time.Sleep(pauseTimeForLogToCatchup)

	chk.outChecked = true

	return chk.compareLog(
		"stdout",
		chk.outBuf.String(),
		func(s string) string {
			return s
		},
		func(s string) string {
			return s
		},
		//		prepareWantString,
		wantLines...,
	)
}
