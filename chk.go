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
	"os"
	"reflect"
	"strings"
	"time"
)

type chkBoundedNumericType interface {
	~float32 | ~float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type chkBoundedStringType interface {
	~string
}

type chkBoundedType interface {
	chkBoundedNumericType |
		chkBoundedStringType
}

type chkType interface {
	bool |
		~complex64 | ~complex128 |
		chkBoundedType
}

type captureOption int

const (
	// Nothing is captured.
	captureNothing captureOption = iota

	// Writes to the standard log package is captured and must be tested with
	// the chk.Log method.
	captureLog

	// Writes to both the standard log package and os.Stderr are captured and
	// must be tested with the chk.Log method.
	captureLogWithStderr

	// Writes to both the standard log package and to os.Stderr are individually
	// captured and must be tested with the chk.Log and chk.Stderr methods.
	captureLogAndStderr

	// Writes both the standard log package and os.Stderr are captured and must
	// be tested with the chk.Log method.  Further Writes to os.Stdout are also
	// captured and must be tested with the chk.Stdout method.
	captureLogWithStderrAndStdout

	// Writes to the standard log package, to os.Stderr and os.Stdout are each
	// individually captured and must be tested with the chk.Log, chk.Stderr and
	// chk.Stdout methods.
	captureLogAndStderrAndStdout

	// Writes to the standard log package and os.Stdout are each
	// individually captured and must be tested with the chk.Log and
	// chk.Stdout methods.
	captureLogAndStdout

	// Writes to os.Stderr are captured and must be tested with the chk.Stderr
	// method.
	captureStderr

	// Writes to os.Stdout are captured and must be tested with the chk.Stdout
	// method.
	captureStdout

	// Writes to os.Stderr and os.Stdout are captured and must be tested with
	// the chk.Stderr and chk.Stdout methods.
	captureStderrAndStdout
)

// Chk structure provides a selector and data to perform testing functions.
type Chk struct {
	t testingT

	subs             []substitution
	markupForDisplay markupFunction

	// Invoked by the method Release which must be called by the owner.
	releaseFunc func() error

	// Logging info.
	//

	errBuf          *bytes.Buffer
	errOrig         *os.File
	logBuf          *bytes.Buffer
	logOrig         io.Writer
	logOrigLogFlags int
	outBuf          *bytes.Buffer
	outOrig         *os.File

	faultCount uint
	nextTmpID  int

	// io.Reader and io.Writer field
	rData   []byte
	rPos    int
	rLeft   int
	rErrPos int
	rErr    error

	wData   []byte
	wErrPos int
	wErr    error

	ioSeekErrPos int64
	ioSeekErr    error

	ioReadErrPos int
	ioReadErr    error

	ioWriteErrPos int
	ioWriteErr    error

	ioCloseErr error

	ioSeekErrSet  bool
	ioReadErrSet  bool
	ioWriteErrSet bool
	ioCloseErrSet bool

	runningPanicFunction bool

	// Group smaller width fields at end to minimize structure space.
	errOn      bool
	errChecked bool
	errIncLog  bool
	outOn      bool
	outChecked bool
	logOn      bool
	logChecked bool

	keepTmpFiles  bool
	tmpDirCreated bool

	clk      *tstClk
	clkSub   int
	clkCusA  string
	clkCusB  string
	clkCusC  string
	clkTicks []time.Time
}

func newChk(t testingT, option captureOption) *Chk {
	t.Helper()
	chk := &Chk{
		t:       t,
		rData:   make([]byte, 0),
		rErrPos: -1,
		wData:   make([]byte, 0),
		wErrPos: -1,
		clk:     newTstClock(time.Now(), []time.Duration{time.Millisecond}),
	}
	chk.markupForDisplay = resolveMarksForDisplay

	chk.setupLoggers(option)
	return chk
}

// CaptureNothing returns a new sztest object without any logs or
// standard io being captured.
func CaptureNothing(t testingT) *Chk {
	t.Helper()
	return newChk(t, captureNothing)
}

// FailFast sets the action takin after an error is discovered.
func (chk *Chk) FailFast(failFast bool) bool {
	oldSetting := settingFailFast
	settingFailFast = failFast
	return oldSetting
}

// Logf passthrough to t.
func (chk *Chk) Logf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()
	chk.t.Logf(msgFmt, msgArgs...)
}

// Error passthrough to t.
func (chk *Chk) Error(args ...any) {
	chk.t.Helper()
	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprint(args...)))
	if settingFailFast {
		chk.t.FailNow()
	}
}

// Errorf passthrough to t.
func (chk *Chk) Errorf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()
	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprintf(msgFmt, msgArgs...)))
	if settingFailFast {
		chk.t.FailNow()
	}
}

// Fatalf passthrough to t.
func (chk *Chk) Fatalf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()
	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprintf(msgFmt, msgArgs...)))
	chk.t.FailNow()
}

// Name returns the name of the saved test object.
func (chk *Chk) Name() string {
	return strings.ReplaceAll(chk.t.Name(), string(os.PathSeparator), "-")
}

// T returns an interface to sztest object provided on creation.
func (chk *Chk) T() testingT {
	return chk.t
}

// KeepTmpFiles stops the removal of tmp files when the check is
// fault free.
func (chk *Chk) KeepTmpFiles() {
	chk.keepTmpFiles = true
}

// PushPreReleaseFunc adds a new release function to the front of the
// queue.
func (chk *Chk) PushPreReleaseFunc(newFunc func() error) {
	chk.t.Helper()
	if chk.releaseFunc == nil {
		chk.releaseFunc = func() error {
			chk.t.Helper()
			return newFunc()
		}
		return
	}
	prevFunc := chk.releaseFunc
	chk.releaseFunc = func() error {
		chk.t.Helper()
		err := newFunc()
		if err == nil {
			err = prevFunc()
		}
		return err
	}
}

// PushPostReleaseFunc adds a new release function to the end of the
// queue.
func (chk *Chk) PushPostReleaseFunc(newFunc func() error) {
	chk.t.Helper()
	if chk.releaseFunc == nil {
		chk.releaseFunc = func() error {
			chk.t.Helper()
			return newFunc()
		}
		return
	}
	prevFunc := chk.releaseFunc
	chk.releaseFunc = func() error {
		chk.t.Helper()
		err := prevFunc()
		if err == nil {
			err = newFunc()
		}
		return err
	}
}

// Release invokes all pushed release functions.
func (chk *Chk) Release() {
	chk.t.Helper()
	if !chk.runningPanicFunction {
		if r := recover(); r != nil {
			panic(r)
		} else if chk.releaseFunc != nil {
			err := chk.releaseFunc()
			chk.releaseFunc = nil
			if err != nil {
				chk.Fatalf("release caused error: %v", err)
			}
		}
	}
}

// Generic tests.

func errMsgHeader(typeName string, msgArgs ...any) string {
	msg := ""
	if len(msgArgs) > 0 {
		msg = fmt.Sprint(msgArgs...)
		if msg != "" {
			msg = ":\n" + markMsgOn + msg + markMsgOff
		}
	}
	return "unexpected " + typeName + msg + ":\n"
}

func errMsgHeaderf(typeName, msgFmt string, msgArgs ...any) string {
	msg := fmt.Sprintf(msgFmt, msgArgs...)
	if msg != "" {
		msg = ":\n" + markMsgOn + msg + markMsgOff
	}
	return "unexpected " + typeName + msg + ":\n"
}

func (chk *Chk) errGotWnt(typeName string, got, wnt any, msg ...any) bool {
	gotf := fmt.Sprintf("%v", got)
	wntf := fmt.Sprintf("%v", wnt)
	chk.t.Helper()
	chk.Error(
		errMsgHeader(typeName, msg...) + gotWnt(gotf, wntf),
	)
	return false
}

func (chk *Chk) errGotWntf(
	typeName string, got, wnt any, msgFmt string, msgArgs ...any,
) bool {
	gotf := fmt.Sprintf("%v", got)
	wntf := fmt.Sprintf("%v", wnt)
	chk.t.Helper()
	chk.Error(
		errMsgHeaderf(typeName, msgFmt, msgArgs...) + gotWnt(gotf, wntf),
	)
	return false
}

func (chk *Chk) isStringify(value any) string {
	s := fmt.Sprintf("%v", value)
	s = chk.subStr(s)
	s = fmt.Sprintf("%q", s)
	s = s[1 : len(s)-1]                  // trim "%q" added quotemarks
	s = strings.ReplaceAll(s, `\"`, `"`) // reverse "%q" action on quotes
	s = strings.ReplaceAll(s, `\\`, `\`) // reverse "%q" action on backslashes
	return s
}

func (chk *Chk) errChk(
	got, want any,
	typeName string,
	msg ...any,
) bool {
	chk.t.Helper()
	wStr := chk.isStringify(want)
	gStr := chk.isStringify(got)
	chk.Error(
		errMsgHeader(typeName, msg...) +
			gotWntDiff(gStr, wStr, settingDiffChars),
	)
	return false
}

func (chk *Chk) errChkf(
	got, want any,
	typeName string,
	msgFmt string,
	msgArgs ...any,
) bool {
	chk.t.Helper()
	wStr := chk.isStringify(want)
	gStr := chk.isStringify(got)
	chk.Error(
		errMsgHeaderf(typeName, msgFmt, msgArgs...) +
			gotWntDiff(gStr, wStr, settingDiffChars),
	)
	return false
}

func errSlice[V chkType](
	chk *Chk,
	got, want []V, typeName string, cmp func(a, b V) bool, msg ...any,
) bool {
	chk.t.Helper()
	var errMsg string

	d1 := false
	gDiff := DiffSlice(
		got,
		want,
		newDiffLnFmt(len(got), len(want)),
		&d1,
		settingDiffSlice,
		settingDiffChars,
		cmp,
	)
	if d1 {
		errMsg = fmt.Sprint("Length Got: ", len(got), " Wnt: ", len(want))
	} else {
		errMsg = "invalid invocation: arrays are identical"
	}
	chk.Error(
		errMsgHeader("[]"+typeName, msg...) +
			errMsg +
			" [\n" + strings.Join(gDiff, "\n") + "\n]",
	)
	return false
}

func errSlicef[V chkType](
	chk *Chk,
	got, want []V, typeName string, cmp func(a, b V) bool,
	msgFmt string, msgArgs ...any,
) bool {
	chk.t.Helper()
	var errMsg string

	d1 := false
	gDiff := DiffSlice(
		got,
		want,
		newDiffLnFmt(len(got), len(want)),
		&d1,
		settingDiffSlice,
		settingDiffChars,
		cmp,
	)
	if d1 {
		errMsg = fmt.Sprint("Length Got: ", len(got), " Wnt: ", len(want))
	} else {
		errMsg = "invalid invocation: arrays are identical"
	}
	chk.Error(
		errMsgHeaderf("[]"+typeName, msgFmt, msgArgs...) +
			errMsg +
			" [\n" + strings.Join(gDiff, "\n") + "\n]",
	)
	return false
}

// BoundedOption constant type.
type BoundedOption int

const (
	// BoundedOpen (a,b) = { x | a < x < b }.
	BoundedOpen BoundedOption = iota
	// BoundedClosed [a,b] = { x | a ≦ x ≦ b }.
	BoundedClosed
	// BoundedMinOpen (a,b] = { x | a < x ≦ b }.
	BoundedMinOpen
	// BoundedMaxClosed (a,b] = { x | a < x ≦ b }.
	BoundedMaxClosed
	// BoundedMaxOpen [a,b) = { x | a ≦ x < b }.
	BoundedMaxOpen
	// BoundedMinClosed [a,b) = { x | a ≦ x < b }.
	BoundedMinClosed
)

// UnboundedOption constant type.
type UnboundedOption int

const (
	// UnboundedMinOpen (a,+∞) = { x | x > a }.
	UnboundedMinOpen UnboundedOption = iota
	// UnboundedMinClosed [a,+∞) = { x | x ≧ a }.
	UnboundedMinClosed
	// UnboundedMaxOpen (-∞, b) = { x | x < b }.
	UnboundedMaxOpen
	// UnboundedMaxClosed (-∞, b] = { x | x ≦ b }.
	UnboundedMaxClosed
)

func inBoundedRange[V chkBoundedType](
	got V, option BoundedOption, min, max V,
) (bool, string) {
	var inRange bool
	var want string
	v := "%v"
	if reflect.TypeOf(got).Name() == "string" {
		v = "\"%v\""
	}
	switch option {
	case BoundedOpen:
		// IntervalBoundedOpen (a,b) = { x | a < x < b }
		inRange = min < got && got < max
		want = "(" + v + "," + v + ") - { want | " + v + " < want < " + v + " }"
	case BoundedClosed:
		// IntervalBoundedClosed [a,b] = { x | a <= x <= b }
		inRange = min <= got && got <= max
		want = "[" + v + "," + v + "] - { want | " + v + " <= want <= " + v + " }"
	case BoundedMinOpen, BoundedMaxClosed:
		// IntervalBoundedLeftOpen (a,b] = { x | a < x ≦ b }
		inRange = min < got && got <= max
		want = "(" + v + "," + v + "] - { want | " + v + " < want <= " + v + " }"
	case BoundedMaxOpen, BoundedMinClosed:
		// IntervalBoundedRightOpen [a,b) = { x | a ≦ x < b }
		inRange = min <= got && got < max
		want = "[" + v + "," + v + ") - { want | " + v + " <= want < " + v + " }"
	default:
		return false, fmt.Sprint("unknown bounded option ", option)
	}
	if inRange {
		return true, ""
	}
	return false, fmt.Sprintf("out of bounds: "+want, min, max, min, max)
}

func inUnboundedRange[V chkBoundedType](
	got V, option UnboundedOption, bound V,
) (bool, string) {
	var inRange bool
	var want string
	v := "%v"
	if reflect.TypeOf(got).Name() == "string" {
		v = "\"%v\""
	}
	switch option {
	case UnboundedMinOpen:
		// (a,+∞) = { x | x > a }
		inRange = got > bound
		want = "(" + v + ",MAX) - { want | want > " + v + " }"
	case UnboundedMinClosed:
		// [a,+∞) = { x | x ≧ a }
		inRange = got >= bound
		want = "[" + v + ",MAX) - { want | want >= " + v + " }"
	case UnboundedMaxOpen:
		// (-∞, b) = { x | x < b }
		inRange = got < bound
		want = "(MIN," + v + ") - { want | want < " + v + " }"
	case UnboundedMaxClosed:
		// (-∞, b] = { x | x ≦ b }
		inRange = got <= bound
		want = "(MIN," + v + "] - { want | want <= " + v + " }"
	default:
		return false, fmt.Sprint("unknown unbounded option ", option)
	}
	if inRange {
		return true, ""
	}
	return false, fmt.Sprintf("out of bounds: "+want, bound, bound)
}
