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

		return t.Close() //nolint:wrapcheck // Ok.
	})
}

func (chk *Chk) copyStdout() error {
	var err error
	var rFile *os.File
	var wFile *os.File

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
	var err error
	var rFile *os.File
	var wFile *os.File

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

// CaptureStdout returns a new *sztest.Chk reference
// capturing:
//
// - os.Stdout
//
// which must be tested by calling the methods:
//
// - (*Chk).Stdout(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStdout)
}

// CaptureLog returns a new *sztest.Chk reference
// capturing:
//
// - log.Writer() io.Writer
//
// which must be tested by calling the methods:
//
// - (*Chk).Log(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLog(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLog)
}

// CaptureLogAndStdout returns a new *sztest.Chk reference
// capturing:
//
// - log.Writer() io.Writer
// - os.Stdout
//
// which must be tested by calling the methods:
//
// - (*Chk).Log(wantLines ...string) bool
// - (*Chk).Stdout(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLogAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStdout)
}

// CaptureLogAndStderr returns a new *sztest.Chk reference
// capturing:
//
// - log.Writer() io.Writer
// - os.Stderr
//
// which must be tested by calling the methods:
//
// - (*Chk).Log(wantLines ...string) bool
// - (*Chk).Stderr(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLogAndStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStderr)
}

// CaptureLogAndStderrAndStdout returns a new *sztest.Chk reference
// capturing:
//
// - log.Writer() io.Writer
// - os.Stderr
// - os.Stdout
//
// which must be tested by calling the methods:
//
// - (*Chk).Log(wantLines ...string) bool
// - (*Chk).Stderr(wantLines ...string) bool
// - (*Chk).Stdout(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLogAndStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogAndStderrAndStdout)
}

// CaptureLogWithStderr returns a new *sztest.Chk reference
// combining and capturing:
//
// - (log.Writer() io.Writer) + os.Stderr
//
// which must be tested by calling ONE the methods:
//
// - (*Chk).Log(wantLines ...string) bool
// - OR
// - (*Chk).Stderr(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLogWithStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogWithStderr)
}

// CaptureLogWithStderrAndStdout returns a new *sztest.Chk reference
// capturing:
//
// - (log.Writer() io.Writer) + os.Stderr
// - os.Stdout
//
// which must be tested by calling ONE the methods:
//
// - (*Chk).Log(wantLines ...string) bool
// - OR
// - (*Chk).Stderr(wantLines ...string) bool
//
// and the method:
//
// - (*Chk).Stdout(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureLogWithStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureLogWithStderrAndStdout)
}

// CaptureStderr returns a new *sztest.Chk reference
// capturing:
//
// - os.Stderr
//
// which must be tested by calling the method:
//
// - (*Chk).Stderr(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureStderr(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStderr)
}

// CaptureStderrAndStdout returns a new *sztest.Chk reference
// capturing:
//
// - os.Stderr
// - os.Stdout
//
// which must be tested by calling the methods:
//
// - (*Chk).Stderr(wantLines ...string) bool
// - (*Chk).Stdout(wantLines ...string) bool
//
// before (*Chk).Release() is invoked.
func CaptureStderrAndStdout(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureStderrAndStdout)
}

func (chk *Chk) prepareSlice(
	processFunc func(string) string,
	rawLines ...string,
) []string {
	var lines []string
	for _, rl := range rawLines {
		for _, l := range strings.Split(rl, "\n") {
			l = processFunc(l)
			lines = append(lines, chk.isStringify(l))
		}
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

	return lines[firstPos:lastPos]
}

func prepareWantString(line string) string {
	line = strings.TrimSpace(line)
	// Replace first har space.
	if strings.HasPrefix(line, `\s`) {
		line = " " + line[2:]
	}
	if strings.HasSuffix(line, `\s`) {
		line = line[:len(line)-2] + " "
	}

	return line
}

func (chk *Chk) compareLog(
	name, got string,
	gotFilter, wantFilter func(string) string,
	wantLines ...string,
) bool {
	chk.t.Helper()
	ret := CompareSlices(
		fmt.Sprint("Unexpected ", name, " Entry"),
		chk.prepareSlice(
			gotFilter,
			got,
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

		return true
	}

	return false
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
	var clearLogPrefix *regexp.Regexp
	var ok bool

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

// Log checks the internally captured log data with the supplied list.
func (chk *Chk) Log(wantLines ...string) bool {
	chk.t.Helper()

	if !chk.logOn && !(chk.errOn && chk.errIncLog) {
		chk.Error("invalid log.Writer check without information being captured")

		return true
	}

	time.Sleep(pauseTimeForLogToCatchup)

	var gotString string
	var name string

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
		prepareWantString,
		wantLines...,
	)
}

// Stderr checks the internally captured log data with the supplied list.
func (chk *Chk) Stderr(wantLines ...string) bool {
	chk.t.Helper()

	if !(chk.errOn) {
		chk.Error("invalid os.Stderr check without information being captured")

		return true
	}

	time.Sleep(pauseTimeForLogToCatchup)

	chk.errChecked = true

	var getFilterFunc func(string) string
	var name string
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
		prepareWantString,
		wantLines...,
	)
}

// Stdout checks the internally captured log data with the supplied list.
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
		prepareWantString,
		wantLines...,
	)
}
