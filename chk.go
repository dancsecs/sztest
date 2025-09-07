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

// Internal constants used to identify the output areas to capture used by
// the various Capture* methods.
const (
	// Nothing is captured.
	captureNothing captureOption = iota

	// Writes to the standard log package is captured and must be tested with
	// the chk.Log method.
	captureLog

	// Writes to both the standard log package and os.Stderr are captured and
	// must be tested with the chk.Log method.
	captureLogWithStderr

	// Writes to both the standard log package and to os.Stderr are
	// individually captured and must be tested with the chk.Log and chk.Stderr
	// methods.
	captureLogAndStderr

	// Writes both the standard log package and os.Stderr are captured and must
	// be tested with the chk.Log method.  Further Writes to os.Stdout are also
	// captured and must be tested with the chk.Stdout method.
	captureLogWithStderrAndStdout

	// Writes to the standard log package, to os.Stderr and os.Stdout are each
	// individually captured and must be tested with the chk.Log, chk.Stderr
	// and chk.Stdout methods.
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

const commonMsgPrefix = "unexpected "

// Chk provides the core test harness used by sztest.
//
// It holds selectors and captured data (stdout, stderr, package log output,
// temp resources, environment changes, deterministic clock state, and other
// helpers) that the assertion helpers operate against.
//
// Create a *Chk with one of the Capture* constructors and always defer
// chk.Release() to restore global state and clean up resources.
//
// Example:
//
// chk := sztest.CaptureNothing(t)
// defer chk.Release()
// chk.Str(got, want)
//
// The concrete fields are intentionally unexported; use the provided
// constructors and methods to interact with a Chk instance.
type Chk struct {
	t testingT

	subs             []substitution
	markupForDisplay markupFunction

	// Invoked by the method Release for pre and post release functions.
	releaseFunc func() error

	// Output capture fields.

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
	clkSub   ClkFmt
	clkCusA  string
	clkCusB  string
	clkCusC  string
	clkTicks []time.Time
}

func newChk(t testingT, option captureOption) *Chk {
	t.Helper()

	chk := new(Chk)
	chk.t = t
	chk.rData = make([]byte, 0)
	chk.rErrPos = -1
	chk.wData = make([]byte, 0)
	chk.wErrPos = -1
	chk.clk = newTstClock(time.Now(), []time.Duration{time.Millisecond})
	chk.markupForDisplay = resolveMarksForDisplay

	chk.setupLoggers(option)

	return chk
}

// CaptureNothing returns a *Chk that performs no output capturing.
//
// Use this when a test needs the sztest helper object but does not need to
// capture stdout, stderr, or the package logger. The supplied t must be a
// testing helper (*testing.T).
//
// Always defer chk.Release() to ensure any modified global state is
// restored and temporary resources are cleaned up.
func CaptureNothing(t testingT) *Chk {
	t.Helper()

	return newChk(t, captureNothing)
}

// FailFast controls whether chk stops execution immediately after the first
// error (true) or continues accumulating further checks (false). This only
// applies to the current test and is independent of `go test -failfast`.
func (chk *Chk) FailFast(failFast bool) bool {
	oldSetting := settingFailFast
	settingFailFast = failFast

	return oldSetting
}

// Logf forwards a formatted log message to the underlying testingT.
func (chk *Chk) Logf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()
	chk.t.Logf(msgFmt, msgArgs...)
}

// Error forwards an error message to the underlying testingT.
func (chk *Chk) Error(args ...any) {
	chk.t.Helper()

	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprint(args...)))

	if settingFailFast {
		chk.t.FailNow()
	}
}

// Errorf forwards a formatted error message to the underlying testingT.
func (chk *Chk) Errorf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()

	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprintf(msgFmt, msgArgs...)))

	if settingFailFast {
		chk.t.FailNow()
	}
}

// Fatalf forwards a formatted fatal error message to the underlying testingT.
// Internally it calls Errorf before aborting the current test.
func (chk *Chk) Fatalf(msgFmt string, msgArgs ...any) {
	chk.t.Helper()

	chk.faultCount++
	chk.t.Error(chk.markupForDisplay(fmt.Sprintf(msgFmt, msgArgs...)))
	chk.t.FailNow()
}

// Name returns the name of the current test from the underlying testingT.
func (chk *Chk) Name() string {
	return strings.ReplaceAll(chk.t.Name(), string(os.PathSeparator), "-")
}

// T exposes the chk's underlying testingT object, as provided at creation.
// Useful for advanced scenarios where direct access is required.
//
//nolint:ireturn // By design.
func (chk *Chk) T() testingT {
	return chk.t
}

// KeepTmpFiles prevents automatic cleanup of the temporary
// directory tree when the test completes successfully. Applies
// only to the current test and is useful for debugging setups
// by inspecting intermediate files.
func (chk *Chk) KeepTmpFiles() {
	chk.keepTmpFiles = true
}

// PushPreReleaseFunc prepends a new cleanup function to the pre-release queue.
//
// Pre-release functions are executed before the Chk's internal cleanup. Since
// each new function is placed at the front of the queue, pre-release funcs
// run in LIFO order (most recently pushed runs first). Return a non-nil
// error from a pre-release function to signal a cleanup failure; Release will
// report such errors to the test.
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

// PushPostReleaseFunc appends a new cleanup function to the post-release
// queue.
//
// Post-release functions are executed after the Chk's internal cleanup and
// after pre-release functions. Functions are executed in the order they are
// pushed (FIFO). Each function should return a non-nil error to indicate a
// cleanup failure; Release will report such errors to the test.
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

// Release restores global state, runs any pushed pre-release functions,
// performs the Chk's internal cleanup (restoring os.Stdout/os.Stderr,
// log.Writer(), env vars, tmp files, clock state, etc.), and then runs any
// pushed post-release functions.
//
// Release must be called (typically via defer) to avoid leaking global state
// or temporary resources. Any non-nil errors returned by pushed release
// functions are reported to the underlying testingT.
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

	return commonMsgPrefix + typeName + msg + ":\n"
}

func errMsgHeaderf(typeName, msgFmt string, msgArgs ...any) string {
	msg := fmt.Sprintf(msgFmt, msgArgs...)
	if msg != "" {
		msg = ":\n" + markMsgOn + msg + markMsgOff
	}

	return commonMsgPrefix + typeName + msg + ":\n"
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

func (chk *Chk) isStringify(rawStr any) string {
	str := fmt.Sprintf("%v", rawStr)
	str = chk.subStr(str)
	str = fmt.Sprintf("%q", str)
	str = str[1 : len(str)-1]                // trim "%q" added quote marks
	str = strings.ReplaceAll(str, `\"`, `"`) // reverse "%q" on quotes
	str = strings.ReplaceAll(str, `\\`, `\`) // reverse "%q" on backslashes

	return str
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

	var (
		errMsg    string
		diffFound bool
	)

	gDiff := diffSlice(
		got,
		want,
		newDiffLnFmt(len(got), len(want)),
		&diffFound,
		settingDiffSlice,
		settingDiffChars,
		cmp,
	)

	if diffFound {
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

	var (
		errMsg    string
		diffFound bool
	)

	gDiff := diffSlice(
		got,
		want,
		newDiffLnFmt(len(got), len(want)),
		&diffFound,
		settingDiffSlice,
		settingDiffChars,
		cmp,
	)

	if diffFound {
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

// BoundedOption specifies the inclusivity of bounds in a closed interval
// check.
type BoundedOption int

const (
	// BoundedOpen checks (a,b) = { x | a < x < b }.
	BoundedOpen BoundedOption = iota

	// BoundedClosed checks [a,b] = { x | a <= x <= b }.
	BoundedClosed

	// BoundedMinOpen checks (a,b] = { x | a < x <= b }.
	// Alias of BoundedMaxClosed.
	BoundedMinOpen

	// BoundedMaxClosed checks (a,b] = { x | a < x <= b }.
	// Alias of BoundedMinOpen.
	BoundedMaxClosed

	// BoundedMaxOpen checks [a,b) = { x | a <= x < b }.
	// Alias of BoundedMinClosed.
	BoundedMaxOpen

	// BoundedMinClosed checks [a,b) = { x | a <= x < b }.
	// Alias of BoundedMaxOpen.
	BoundedMinClosed
)

// UnboundedOption specifies the inclusivity of bounds in a half-infinite
// interval check.
type UnboundedOption int

const (
	// UnboundedMinOpen checks (a,+∞) = { x | x > a }.
	UnboundedMinOpen UnboundedOption = iota

	// UnboundedMinClosed checks [a,+∞) = { x | x >= a }.
	UnboundedMinClosed

	// UnboundedMaxOpen checks (-∞, b) = { x | x < b }.
	UnboundedMaxOpen

	// UnboundedMaxClosed checks (-∞, b] = { x | x <= b }.
	UnboundedMaxClosed
)

//nolint:cyclop // Ok.
func inBoundedRange[V chkBoundedType](
	got V, option BoundedOption, minV, maxV V,
) (bool, string) {
	var (
		inRange bool
		want    string
	)

	vFmt := "%v"

	if reflect.TypeOf(got).Name() == "string" {
		vFmt = "\"%v\""
	}

	const (
		definitionSeparator = " - "
		wntLabel            = "{ want | "
	)

	switch option {
	case BoundedOpen:
		// IntervalBoundedOpen (a,b) = { x | a < x < b }
		inRange = minV < got && got < maxV
		want = "" +
			"(" + vFmt + "," + vFmt + ")" +
			definitionSeparator +
			wntLabel + vFmt + " < want < " + vFmt + " }"
	case BoundedClosed:
		// IntervalBoundedClosed [a,b] = { x | a <= x <= b }
		inRange = minV <= got && got <= maxV
		want = "" +
			"[" + vFmt + "," + vFmt + "]" +
			definitionSeparator +
			wntLabel + vFmt + " <= want <= " + vFmt + " }"
	case BoundedMinOpen, BoundedMaxClosed:
		// IntervalBoundedLeftOpen (a,b] = { x | a < x ≦ b }
		inRange = minV < got && got <= maxV
		want = "" +
			"(" + vFmt + "," + vFmt + "]" +
			definitionSeparator +
			wntLabel + vFmt + " < want <= " + vFmt + " }"
	case BoundedMaxOpen, BoundedMinClosed:
		// IntervalBoundedRightOpen [a,b) = { x | a ≦ x < b }
		inRange = minV <= got && got < maxV
		want = "" +
			"[" + vFmt + "," + vFmt + ")" +
			definitionSeparator +
			wntLabel + vFmt + " <= want < " + vFmt + " }"
	default:
		return false, fmt.Sprint("unknown bounded option ", option)
	}

	if inRange {
		return true, ""
	}

	return false, fmt.Sprintf("out of bounds: "+want, minV, maxV, minV, maxV)
}

func inUnboundedRange[V chkBoundedType](
	got V, option UnboundedOption, bound V,
) (bool, string) {
	var (
		inRange bool
		want    string
	)

	vFmt := "%v"

	if reflect.TypeOf(got).Name() == "string" {
		vFmt = "\"%v\""
	}

	switch option {
	case UnboundedMinOpen:
		// (a,+∞) = { x | x > a }
		inRange = got > bound
		want = "(" + vFmt + ",MAX) - { want | want > " + vFmt + " }"
	case UnboundedMinClosed:
		// [a,+∞) = { x | x ≧ a }
		inRange = got >= bound
		want = "[" + vFmt + ",MAX) - { want | want >= " + vFmt + " }"
	case UnboundedMaxOpen:
		// (-∞, b) = { x | x < b }
		inRange = got < bound
		want = "(MIN," + vFmt + ") - { want | want < " + vFmt + " }"
	case UnboundedMaxClosed:
		// (-∞, b] = { x | x ≦ b }
		inRange = got <= bound
		want = "(MIN," + vFmt + "] - { want | want <= " + vFmt + " }"
	default:
		return false, fmt.Sprint("unknown unbounded option ", option)
	}

	if inRange {
		return true, ""
	}

	return false, fmt.Sprintf("out of bounds: "+want, bound, bound)
}
