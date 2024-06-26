<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

<!---
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
-->

# Package sztest

## Overview

Provides a single self contained go package of test helpers.

- ```got/wnt``` helper testing
- ```'error'``` testing
- ```'panic'``` testing
- ```'os.Stdout'``` and ```'os.Stderr'``` capture and testing
- ```'package log'``` capture and testing
- ```'io.Reader``` and  ```'io.Writer'``` interface testing
- ```'os.Args'``` and ```'os.Flag'``` setup testing
- temporary directories, files and scripts
- timestamps

Package ```sztest``` only imports core go libraries minimizing your module's
dependencies.

---

## Contents

- [Usage](#usage)
  - [Examples: General Form](#example-general-form)
- [Builtin Got/Wnt](#builtin-gotwnt-checks)
  - [Examples: Integer With No Message](examples/got_wnt/README.md#examples-integer-with-no-message)
  - [Examples: Float32 With Unformatted Message](examples/got_wnt/README.md#examples-float32-with-unformatted-message)
  - [Examples: String With Formatted Message](examples/got_wnt/README.md#examples-string-with-formatted-message)
- [Builtin Slices](#array-slices)
  - [Examples: String Slice With no Message](examples/slice/README.md#examples-string-slice-with-no-message)
  - [Examples: Float64 Slice With Unformatted Message](examples/slice/README.md#examples-float64-slice-with-unformatted-message)
  - [Examples: Uint64 Slice With Formatted Message](examples/slice/README.md#examples-uint64-slice-with-formatted-message)
- [Builtin Bounded Intervals](#bounded-intervals)
  - [Examples: Float64 With no Message](examples/bounded/README.md#examples-float64-with-no-message)
- [Builtin Unbounded Intervals](#unbounded-intervals)
  - [Examples: Float64 With no Message](examples/unbounded/README.md#examples-float64-with-no-message)
- [Errors](#errors)
  - [Examples: No Error](examples/error/README.md#examples-no-error)
  - [Examples: No Error Helper](examples/error/README.md#examples-no-error-helper)
  - [Examples: Error](examples/error/README.md#examples-error)
  - [Examples: Blank Error](examples/error/README.md#examples-blank-error)
- [Panics](#panics)
  - [Examples: No Panic](examples/panic/README.md#examples-no-panic)
  - [Examples: No Panic Helper](examples/panic/README.md#examples-no-panic-helper)
  - [Examples: Panic](examples/panic/README.md#examples-panic)
  - [Examples: Blank Panic](examples/panic/README.md#examples-blank-panic)
- [Output](#output)
  - [Examples: Output Stdout](examples/output/README.md#examples-output-stdout)
  - [Examples: Output Stderr](examples/output/README.md#examples-output-stderr)
  - [Examples: Output Stderr And Stdout](examples/output/README.md#examples-output-stderr-and-stdout)
  - [Examples: Output Log](examples/output/README.md#examples-output-log)
  - [Examples: Output Log And Stdout](examples/output/README.md#examples-output-log-and-stdout)
  - [Examples: Output Log And Stderr](examples/output/README.md#examples-output-log-and-stderr)
  - [Examples: Output Log And Stderr And Stdout](examples/output/README.md#examples-output-log-and-stderr-and-stdout)
  - [Examples: Output Log With Stderr](examples/output/README.md#examples-output-log-with-stderr)
  - [Examples: Output Log With Stderr And Stdout](examples/output/README.md#examples-output-log-with-stderr-and-stdout)
- [IO Interface](#io-interface)
  - [Example: IO Read Error](examples/io_interface/README.md#example-io-read-error)
  - [Example: IO Write Error](examples/io_interface/README.md#example-io-write-error)
  - [Example: IO Read Seek Error](examples/io_interface/README.md#example-io-read-seek-error)
  - [Example: IO Write Seek Error](examples/io_interface/README.md#example-io-write-seek-error)
  - [Example: IO Close Error](examples/io_interface/README.md#example-io-close-error)
- [Arguments And Flags](#arguments-and-flags)
  - [Example: Single Good Flag](examples/arguments_and_flags/README.md#example-arguments-and-flags-single-good-flag)
  - [Example: Invalid Flag](examples/arguments_and_flags/README.md#example-arguments-and-flags-invalid-flag)
  - [Example: Invalid Integer](examples/arguments_and_flags/README.md#example-arguments-and-flags-invalid-integer)
- [Environment Variables](#environment-variables)
- [Temporary directories, files, scripts](#temporary-directories-files-scripts)
  - [Example: Temporary Directory](examples/temporary_dir_file_script/README.md#example-temporary-directory)
  - [Example: Temporary File](examples/temporary_dir_file_script/README.md#example-temporary-file)
  - [Example: Temporary Unix Script](examples/temporary_dir_file_script/README.md#example-temporary-unix-script)
  - [Example: Temporary Directory Tree](examples/temporary_dir_file_script/README.md#example-temporary-directory-tree)
- [Timestamps](#timestamps)
  - [Example: Logging](examples/timestamp/README.md#example-logging)
- [Configuration](CONFIGURE.md)
  - [Example: Default Markup](CONFIGURE.md#example-default-markup)
  - [Example: Ascii Markup](CONFIGURE.md#example-ascii-markup)
  - [Example: Unicode Markup](CONFIGURE.md#example-unicode-markup)
- [Appendices](#appendices)
  - [Appendix A: Capture* creation functions](#appendix-a-list-of-sztestcapture-create-functions)
  - [Appendix B: List of got/wnt test methods](#appendix-b-list-of-gotwnt-test-methods)
  - [Appendix C: List of got/wnt slice test methods](#appendix-c-list-of-gotwnt-slice-test-methods)
  - [Appendix D: List of Bounded and Unbounded Interval tests](#appendix-d-list-of-bounded-and-unbounded-interval-tests)
  - [Appendix E: Builtin Ansi Terminal Markup](CONFIGURE.md#appendix-e-builtin-ansi-terminal-markup)
  - [Appendix F: Large Example Function](examples/appendix/README.md#appendix-f-large-example-function)
  - [Appendix G: Large Example Main Function](examples/appendix/README.md#appendix-g-large-example-main-function)
  - [Appendix H: License](#appendix-h-license)
  - [Appendix I: To Be Documented](#appendix-i-to-be-documented)

## Usage

A``` '*sztest.Chk' ``` object is created in the test function by calling one of
the ```sztest.Capture*``` functions and then deferring its ```Release()```
method to run on the completion of the test function.  Common got/want type
testing is provided for all go builtin types as well as some common aliases and
interfaces.

### Example: General Form

<!--- gotomd::Bgn::file::./examples/general_form/example_test.go -->
```bash
cat ./examples/general_form/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_PASS_GeneralForm(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    s1 := "Value Got/Wnt"
    s2 := "Value Got/Wnt"

    chk.Str(s1, s2)
    chk.Str(s1, s2, "unformatted", " message", " not", " displayed")
    chk.Strf(s1, s2, "formatted %s %s %s", "message", "not", "displayed")
}

func Test_FAIL_GeneralForm(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.FailFast(false) // Do not stop on first problem.

    s1 := "Value Got"
    s2 := "Value Wnt"

    chk.Str(s1, s2)
    chk.Str(s1, s2, "unformatted", " message", " displayed")
    chk.Strf(s1, s2, "formatted %s %s", "message", "displayed")
}
```
<!--- gotomd::End::file::./examples/general_form/example_test.go -->

<!--- gotomd::Bgn::tst::./examples/general_form/package -->
```bash
go test -v -cover ./examples/general_form
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;GeneralForm}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;GeneralForm&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;GeneralForm}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:30:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Wnt\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:31:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\emph{unformatted&#xa0;message&#xa0;displayed}:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Wnt\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:32:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\emph{formatted&#xa0;message&#xa0;displayed}:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}Value&#xa0;\color{darkturquoise}Wnt\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;GeneralForm&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/general&#x332;form&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./examples/general_form/package -->

The ```*sztest.Chk``` object is created without capturing anything with the
Release method being deferred until the function exits (This opening pattern
will be used by all test functions). There are three string tests.  The first
two pass but the third fails producing the highlighted test differences.

[Contents](#contents)

## Builtin Got/Wnt Checks

The most basic test is to compare something ```"Got"``` from the code being
tested to the ```"Wnt"``` expected by the test.  If the values do not exactly
match an error is registered for the test using the (```testing.T.Error```)
function and the error is displayed as exampled above. Got/Want functions are
provided for all core data types as well as some aliases and interfaces. The
general forms are:

```go
func (*CHK) Type(got, wnt Type, msg ...any) bool
func (*CHK) Typef(got, wnt Type, msgFmt string, msgArgs ...any) bool

/*
Were Type is one of:

    // Basic Types
        Bool,
        Byte,
        Complex64, Complex128
        Float32, Float64,  // Includes extra tolerance parameter.
        Int, Int8, Int16, Int32, Int64,
        Rune, Str,
        Uint, Uint8, Uint16 , Uint32, Uint64,
        Uintptr

    // Aliases.
        Dur  // time.Duration
*/
```

providing for an optional message and returning ```true``` if the test passed.

- [Examples: Integer With No Message](examples/got_wnt/README.md#examples-integer-with-no-message)
- [Examples: Float32 With Unformatted Message](examples/got_wnt/README.md#examples-float32-with-unformatted-message)
- [Examples: String With Formatted Message](examples/got_wnt/README.md#examples-string-with-formatted-message)

[Contents](#contents)

## Array Slices

Array slice tests are provided for all core data types.  The arrays must match
exactly (except for float types which have a tolerance argument) otherwise a
failure will be registered. The general forms are:

```go
func (*CHK) TypeSlice(got, wnt []Type, msg ...any)
func (*CHK) TypeSlicef(got, wnt []Type, fmtMsg string, msgArgs ...any)

/*
Were Type is one of:

    // Basic Types
        Bool,
        Byte,
        Complex64, Complex128,
        Float32, Float64,  // Includes extra tolerance parameter.
        Int, Int8, Int16, Int32, Int64,
        Rune, Str,
        Uint, Uint8, Uint16 , Uint32, Uint64,
        Uintptr

    // Aliases
        Dur  // time.Duration
*/
```

with error slices tested with:

```go
// Errors NOTE:  Got/Wnt are different types.
func (*Chk) ErrSlice(got []error, wnt []string, msg ...any) bool
```

For a complete list of builtin got/wnt slice tests and their helpers see
[Appendix C: List of got/wnt slice test methods](
         #appendix-c-list-of-gotwnt-slice-test-methods).

- [Examples: String Slice With no Message](examples/slice/README.md#examples-string-slice-with-no-message)
- [Examples: Float64 Slice With Unformatted Message](examples/slice/README.md#examples-float64-slice-with-unformatted-message)
- [Examples: Uint64 Slice With Formatted Message](examples/slice/README.md#examples-uint64-slice-with-formatted-message)

[Contents](#contents)

## Bounded Intervals

These tests compare a comparable got against a range of values. The general
forms are:

```go
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

func (*CHK) TypeBounded(got Type, option BoundedOption,  min, max Type, msg ...any)
func (*CHK) TypeBoundedf(got Type, option BoundedOption,  min, max Type, fmtMsg string, msgArgs ...any)

/*
Were Type is one of:

    // Basic Types
        Byte,
        Float32, Float64,
        Int, Int8, Int16, Int32, Int64,
        Rune, Str,
        Uint, Uint8, Uint16 , Uint32, Uint64

    // Aliases
        Dur  // time.Duration
*/
```

- [Examples: Float64 With no Message](examples/bounded/README.md#examples-float64-with-no-message)

[Contents](#contents)

## Unbounded Intervals

These tests compare a comparable got against a range of values. The general
forms are:

```go
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

func (*CHK) TypeUnbounded(got Type, option UnboundedOption, bound Type, msg ...any)
func (*CHK) TypeUnboundedf(got Type, option UnboundedOption, bound Type, fmtMsg string, msgArgs ...any)

/*
Were Type is one of:

    // Basic Types
        Byte,
        Float32, Float64,
        Int, Int8, Int16, Int32, Int64,
        Rune, Str,
        Uint, Uint8, Uint16 , Uint32, Uint64

    // Aliases
        Dur  // time.Duration
*/
```

- [Examples: Float64 With no Message](examples/unbounded/README.md#examples-float64-with-no-message)

[Contents](#contents)

## Errors

Error conditions are checked using the following method:

<!--- gotomd::Bgn::dcls::./Chk.Err -->
```go
func (chk *Chk) Err(got error, want string, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Err -->

and its helper methods:

<!--- gotomd::Bgn::dcls::./Chk.Errf Chk.NoErr Chk.NoErrf -->
```go
func (chk *Chk) Errf(got error, want string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) NoErr(got error, msg ...any) bool
func (chk *Chk) NoErrf(got error, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Errf Chk.NoErr Chk.NoErrf -->

Please note with these methods the got and wnt are different data types with
the got being an error and the wnt being a string.  So what happens if the
error is not nil but empty?

```go
errors.New("")
```

then the error returned is represented by the constant

<!--- gotomd::Bgn::dcls::./BlankErrorMessage -->
```go
const BlankErrorMessage = "sztest.BlankErrorMessage"
```
<!--- gotomd::End::dcls::./BlankErrorMessage -->

- [Examples: No Error](examples/error/README.md#examples-no-error)
- [Examples: No Error Helper](examples/error/README.md#examples-no-error-helper)
- [Examples: Error](examples/error/README.md#examples-error)
- [Examples: Blank Error](examples/error/README.md#examples-blank-error)

[Contents](#contents)

## Panics

Insuring that your code properly terminates when it encounters an untenable
state is important to verify.  To facilitate this the library defines a panic
check function:

<!--- gotomd::Bgn::dcls::./Chk.Panic -->
```go
func (chk *Chk) Panic(gotF func(), want string, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Panic -->

where gotF is a function that is expected to issue a panic and wnt is the
string representation of the expected panic.  An empty ("") wnt string
represents that no panic should be thrown.  The string

<!--- gotomd::Bgn::dcls::./BlankPanicMessage -->
```go
const BlankPanicMessage = "sztest.BlankPanicMessage"
```
<!--- gotomd::End::dcls::./BlankPanicMessage -->

is returned to represent an empty ("") panic was thrown differentiating it
from no panic being thrown.

There are three helper functions:

<!--- gotomd::Bgn::dcls::./Chk.Panicf Chk.NoPanic Chk.NoPanicf -->
```go
func (chk *Chk) Panicf(gotF func(), want string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) NoPanic(gotF func(), msg ...any) bool
func (chk *Chk) NoPanicf(gotF func(), msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Panicf Chk.NoPanic Chk.NoPanicf -->

- [Examples: No Panic](examples/panic/README.md#examples-no-panic)
- [Examples: No Panic Helper](examples/panic/README.md#examples-no-panic-helper)
- [Examples: Panic](examples/panic/README.md#examples-panic)
- [Examples: Blank Panic](examples/panic/README.md#examples-blank-panic)

[Contents](#contents)

## Output

Programs writing to standard outputs (```os.Stdout```, ```os.Stderr```) and
the go log package (which may be distinct from os.Stderr) can have the
outputs captured and reviewed as part of testing.  This can be to confirm
failing conditions are properly logged and reported as part of full testing.

Each can be captured individually or the log package and os.Stderr can be
combined together into a single captured feed.  Selection of the feeds is
instantiated when the check object is initially created.  See
[Appendix A: Capture* creation functions](#appendix-a-list-of-sztestcapture-create-functions)
for a complete list.

- [Examples: Output Stdout](examples/output/README.md#examples-output-stdout)
- [Examples: Output Stderr](examples/output/README.md#examples-output-stderr)
- [Examples: Output Stderr And Stdout](examples/output/README.md#examples-output-stderr-and-stdout)
- [Examples: Output Log](examples/output/README.md#examples-output-log)
- [Examples: Output Log And Stdout](examples/output/README.md#examples-output-log-and-stdout)
- [Examples: Output Log And Stderr](examples/output/README.md#examples-output-log-and-stderr)
- [Examples: Output Log And Stderr And Stdout](examples/output/README.md#examples-output-log-and-stderr-and-stdout)
- [Examples: Output Log With Stderr](examples/output/README.md#examples-output-log-with-stderr)
- [Examples: Output Log With Stderr And Stdout](examples/output/README.md#examples-output-log-with-stderr-and-stdout)

[Contents](#contents)

## IO Interface

The check object implements some ```io interfaces``` permitting easy simulation
of hard to duplicate error and panic situations.  IO interface methods
implemented by the ```*Chk``` object are:

<!--- gotomd::Bgn::dcls::./Chk.Seek Chk.Read Chk.Write Chk.Close -->
```go
func (chk *Chk) Seek(_ int64, _ int) (int64, error)
func (chk *Chk) Read(dataBuf []byte) (int, error)
func (chk *Chk) Write(data []byte) (int, error)
func (chk *Chk) Close() error
```
<!--- gotomd::End::dcls::./Chk.Seek Chk.Read Chk.Write Chk.Close -->

Each of the above functions can have errors set to be returned on the next
call with the following methods:

<!--- gotomd::Bgn::dcls::./Chk.SetReadError Chk.SetWriteError Chk.SetSeekError Chk.SetCloseError -->
```go
func (chk *Chk) SetReadError(pos int, err error)
func (chk *Chk) SetWriteError(pos int, err error)
func (chk *Chk) SetSeekError(pos int64, err error)
func (chk *Chk) SetCloseError(err error)
```
<!--- gotomd::End::dcls::./Chk.SetReadError Chk.SetWriteError Chk.SetSeekError Chk.SetCloseError -->

Once set the next call to the corresponding io method (read, Write ,Seek,
Close or a composition function) will return the pos and error provided.  The
error is cleared so subsequent calls result in the default action.

Data for read actions (and position errors) is setup using the following
methods:

<!--- gotomd::Bgn::dcls::./Chk.SetIOReaderData Chk.SetIOReaderError -->
```go
func (chk *Chk) SetIOReaderData(d ...string)
func (chk *Chk) SetIOReaderError(byteCount int, err error)
```
<!--- gotomd::End::dcls::./Chk.SetIOReaderData Chk.SetIOReaderError -->

The data will be returned (one byte at time) until the data is exhausted
resulting in an ```io.EOF``` error or the optional byteCount is reached then
the supplied error will be returned.

Data written can be retrieved using the following method:

<!--- gotomd::Bgn::dcls::./Chk.GetIOWriterData -->
```go
func (chk *Chk) GetIOWriterData() []byte
```
<!--- gotomd::End::dcls::./Chk.GetIOWriterData -->

while a positional error condition can be setup to be returned when the nth
byte is written.  This is setup with:

<!--- gotomd::Bgn::dcls::./Chk.SetIOWriterError -->
```go
func (chk *Chk) SetIOWriterError(n int, err error)
```
<!--- gotomd::End::dcls::./Chk.SetIOWriterError -->

- [Example: IO Read Error](examples/io_interface/README.md#example-io-read-error)
- [Example: IO Write Error](examples/io_interface/README.md#example-io-write-error)
- [Example: IO Read Seek Error](examples/io_interface/README.md#example-io-read-seek-error)
- [Example: IO Write Seek Error](examples/io_interface/README.md#example-io-write-seek-error)
- [Example: IO Close Error](examples/io_interface/README.md#example-io-close-error)

[Contents](#contents)

## Arguments And Flags

In order to test main default argument processing, test args and a clean flag
environment are implemented with:

<!--- gotomd::Bgn::dcls::./Chk.SetArgs -->
```go
func (chk *Chk) SetArgs(progName string, args ...string)
```
<!--- gotomd::End::dcls::./Chk.SetArgs -->

where both ```os.Args``` and ```flags.CommandLine``` are saved and replaced
with the provided ```args``` and a ```NewFlagSet``` respectively.  Original
vales are restored when the ```chk.Release()``` method is called.  NOTE: The
new *default flag set* is set to ```panicOnError```.

- [Example: Single Good Flag](examples/arguments_and_flags/README.md#example-arguments-and-flags-single-good-flag)
- [Example: Invalid Flag](examples/arguments_and_flags/README.md#example-arguments-and-flags-invalid-flag)
- [Example: Invalid Integer](examples/arguments_and_flags/README.md#example-arguments-and-flags-invalid-integer)

[Contents](#contents)

## Environment Variables

System environment variables mat be set or deleted using the following:

<!--- gotomd::Bgn::dcls::./Chk.SetEnv Chk.DelEnv -->
```go
func (chk *Chk) SetEnv(name, value string)
func (chk *Chk) DelEnv(name string)
```
<!--- gotomd::End::dcls::./Chk.SetEnv Chk.DelEnv -->

Original values are restores when the ```chk.Release()``` method is called.

[Contents](#contents)

## Temporary directories, files, scripts

Testing underlying os file interfacing code can be somewhat automated by
using some builtin helpers.  Directories, files and scripts created through
```*chk``` methods will be automatically deleted on a successful test.  The
items are not deleted on failure or if the following helper method is invoked
from the test.

<!--- gotomd::Bgn::dcls::./Chk.KeepTmpFiles -->
```go
func (chk *Chk) KeepTmpFiles()
```
<!--- gotomd::End::dcls::./Chk.KeepTmpFiles -->

The default temporary dir for the current test function is both created
and identified with:

<!--- gotomd::Bgn::dcls::./Chk.CreateTmpDir -->
```go
func (chk *Chk) CreateTmpDir() string
```
<!--- gotomd::End::dcls::./Chk.CreateTmpDir -->

directly (or indirectly by one of these helper functions)

<!--- gotomd::Bgn::dcls::./Chk.CreateTmpFile Chk.CreateTmpFileIn Chk.CreateTmpFileAs Chk.CreateTmpUnixScript Chk.CreateTmpUnixScriptIn Chk.CreateTmpUnixScriptAs Chk.CreateTmpSubDir -->
```go
func (chk *Chk) CreateTmpFile(data []byte) string
func (chk *Chk) CreateTmpFileIn(path string, data []byte) string
func (chk *Chk) CreateTmpFileAs(path, fName string, data []byte) string
func (chk *Chk) CreateTmpUnixScript(lines []string) string
func (chk *Chk) CreateTmpUnixScriptIn(path string, lines []string) string
func (chk *Chk) CreateTmpUnixScriptAs(path, fName string, lines []string) string
func (chk *Chk) CreateTmpSubDir(subDirs ...string) string
```
<!--- gotomd::End::dcls::./Chk.CreateTmpFile Chk.CreateTmpFileIn Chk.CreateTmpFileAs Chk.CreateTmpUnixScript Chk.CreateTmpUnixScriptIn Chk.CreateTmpUnixScriptAs Chk.CreateTmpSubDir -->

which all return the path constructed by creating a new sub directory in the
default temp directory.

This can be set using the environment variable:

```bash
SZTEST_TMP_DIR="/custom/tmp"
```

or set from within the test with:

<!--- gotomd::Bgn::dcls::./Chk.SetTmpDir -->
```go
func (chk *Chk) SetTmpDir(dir string) string
```
<!--- gotomd::End::dcls::./Chk.SetTmpDir -->

otherwise it defaults to

```go
os.TempDir()
```

Permissions used when creating these objects can be defined with the following
environment variables

```bash
SZTEST_PERM_DIR="0700"
SZTEST_PERM_FILE="0600"
SZTEST_PERM_EXE="0700"
```

 or from within the test with:

<!--- gotomd::Bgn::dcls::./Chk.SetPermDir Chk.SetPermFile Chk.SetPermExe -->
```go
func (chk *Chk) SetPermDir(p os.FileMode) os.FileMode
func (chk *Chk) SetPermFile(p os.FileMode) os.FileMode
func (chk *Chk) SetPermExe(p os.FileMode) os.FileMode
```
<!--- gotomd::End::dcls::./Chk.SetPermDir Chk.SetPermFile Chk.SetPermExe -->

- [Example: Temporary Directory](examples/temporary_dir_file_script/README.md#example-temporary-directory)
- [Example: Temporary File](examples/temporary_dir_file_script/README.md#example-temporary-file)
- [Example: Temporary Unix Script](examples/temporary_dir_file_script/README.md#example-temporary-unix-script)
- [Example: Temporary Directory Tree](examples/temporary_dir_file_script/README.md#example-temporary-directory-tree)

[Contents](#contents)

## Timestamps

Predictable timestamps are provided to permit full testing of applications
using timestamps.  In order to facilitate this the application must use its own
timestamp function pointer that defaults to be the standard```time.Now```
function and can be replaced by the the testing clock function
```(*Chk).ClockNext```.

> Replacing the internal time.Now method is possible using an external monkey
patch library such as [go-mpatch](https://github.com/undefinedlabs/go-mpatch)
using something similar to:

```go
//  ...

import (
   "github.com/dancsecs/sztest"
   "github.com/undefinedlabs/go-mpatch"
)

func Test_UsesTimeStamps(t *testing) {
  chk:=sztest.CaptureStdout(t)
  defer chk.Release()

  patch,err:=mpatch.PatchMethod(time.Now, chk.ClockNext)
  chk.NoErr(err)
  defer func() {
    _ = patch.Unpatch()
  }()

  // Run tests that use golang's default time.Now function.
  // ...
}
```

Chk.ClockNext may be invoked indirectly with the formatting convenience
methods:

<!--- gotomd::Bgn::dcls::./Chk.ClockNextFmtTime Chk.ClockNextFmtDate Chk.ClockNextFmtTS Chk.ClockNextFmtNano Chk.ClockNextFmtCusA Chk.ClockNextFmtCusB Chk.ClockNextFmtCusC -->
```go
func (chk *Chk) ClockNextFmtTime() string
func (chk *Chk) ClockNextFmtDate() string
func (chk *Chk) ClockNextFmtTS() string
func (chk *Chk) ClockNextFmtNano() string
func (chk *Chk) ClockNextFmtCusA() string
func (chk *Chk) ClockNextFmtCusB() string
func (chk *Chk) ClockNextFmtCusC() string
```
<!--- gotomd::End::dcls::./Chk.ClockNextFmtTime Chk.ClockNextFmtDate Chk.ClockNextFmtTS Chk.ClockNextFmtNano Chk.ClockNextFmtCusA Chk.ClockNextFmtCusB Chk.ClockNextFmtCusC -->

As timestamps are generated they are saved and can be queried with the
function:

<!--- gotomd::Bgn::doc::./Chk.ClockTick -->
```go
func (chk *Chk) ClockTick(i int) time.Time
```

ClockTick returns i'th time returned.
<!--- gotomd::End::doc::./Chk.ClockTick -->

Further chk substitutions can be generated for each timestamp produced
including up to three custom date formats with the following constants and
methods:

<!--- gotomd::Bgn::doc::./ClockSubNone -->
```go

```

Clock substitutions.
<!--- gotomd::End::doc::./ClockSubNone -->

<!--- gotomd::Bgn::doc::./Chk.ClockSetSub Chk.ClockAddSub Chk.ClockRemoveSub Chk.ClockSetCusA Chk.ClockSetCusB Chk.ClockSetCusC -->
```go
func (chk *Chk) ClockSetSub(i int)
```

ClockSetSub sets the fields to set substitutions for.

```go
func (chk *Chk) ClockAddSub(i int)
```

ClockAddSub sets the fields to set substitutions for.

```go
func (chk *Chk) ClockRemoveSub(i int)
```

ClockRemoveSub resets the fields to set substitutions for.

```go
func (chk *Chk) ClockSetCusA(f string)
```

ClockSetCusA sets the custom date format to set tick substitution values.

```go
func (chk *Chk) ClockSetCusB(f string)
```

ClockSetCusB sets the custom date format to set tick substitution values.

```go
func (chk *Chk) ClockSetCusC(f string)
```

ClockSetCusC sets the custom date format to set tick substitution values.
<!--- gotomd::End::doc::./Chk.ClockSetSub Chk.ClockAddSub Chk.ClockRemoveSub Chk.ClockSetCusA Chk.ClockSetCusB Chk.ClockSetCusC -->

The time (and increments used between successive timestamps can be set with:

<!--- gotomd::Bgn::doc::./Chk.ClockSet -->
```go
func (chk *Chk) ClockSet(setTime time.Time, inc ...time.Duration) func()
```

ClockSet set the current test time and optionally sets the increments if
provided.  It returns a func to reset the clk back to its state when
this function was called.
<!--- gotomd::End::doc::./Chk.ClockSet -->

or

<!--- gotomd::Bgn::doc::./Chk.ClockOffsetDay -->
```go
func (chk *Chk) ClockOffsetDay(dayOffset int, inc ...time.Duration) func()
```

ClockOffsetDay adjusts the current clock by the number of specified days
with negative numbers representing the past.  It returns a func to reset
the clk back to its state when this function was called.
<!--- gotomd::End::doc::./Chk.ClockOffsetDay -->

or the clock can be adjusted with:

<!--- gotomd::Bgn::doc::./Chk.ClockOffset -->
```go
func (chk *Chk) ClockOffset(d time.Duration) func()
```

ClockOffset moves the current clock by the specified amount.  No
defined increments are applied and if a clock has not yet been set the
current time advanced by the specified amount will be used. Nothing is
returned.
<!--- gotomd::End::doc::./Chk.ClockOffset -->

while the last time returned can e retrieved with:

<!--- gotomd::Bgn::doc::./Chk.ClockLast -->
```go
func (chk *Chk) ClockLast() time.Time
```

ClockLast returns the last timestamp generated.
<!--- gotomd::End::doc::./Chk.ClockLast -->

or the formatting convenience methods:

<!--- gotomd::Bgn::dcls::./Chk.ClockLastFmtTime Chk.ClockLastFmtDate Chk.ClockLastFmtTS Chk.ClockLastFmtNano Chk.ClockLastFmtCusA Chk.ClockLastFmtCusB Chk.ClockLastFmtCusC -->
```go
func (chk *Chk) ClockLastFmtTime() string
func (chk *Chk) ClockLastFmtDate() string
func (chk *Chk) ClockLastFmtTS() string
func (chk *Chk) ClockLastFmtNano() string
func (chk *Chk) ClockLastFmtCusA() string
func (chk *Chk) ClockLastFmtCusB() string
func (chk *Chk) ClockLastFmtCusC() string
```
<!--- gotomd::End::dcls::./Chk.ClockLastFmtTime Chk.ClockLastFmtDate Chk.ClockLastFmtTS Chk.ClockLastFmtNano Chk.ClockLastFmtCusA Chk.ClockLastFmtCusB Chk.ClockLastFmtCusC -->

- [Example: Logging](examples/timestamp/README.md#example-logging)

[Contents](#contents)

## Appendices

### Appendix A: List of ```sztest.Capture*``` Create Functions

<!--- gotomd::Bgn::doc::./CaptureNothing -->
```go
func CaptureNothing(t testingT) *Chk
```

CaptureNothing returns a new sztest object without any logs or
standard io being captured.
<!--- gotomd::End::doc::./CaptureNothing -->

<!--- gotomd::Bgn::doc::./CaptureStdout -->
```go
func CaptureStdout(t testingT) *Chk
```

CaptureStdout returns a new *sztest.Chk reference
capturing:

- os.Stdout

which must be tested by calling the methods:

- (*Chk).Stdout(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureStdout -->

<!--- gotomd::Bgn::doc::./CaptureLog -->
```go
func CaptureLog(t testingT) *Chk
```

CaptureLog returns a new *sztest.Chk reference
capturing:

- log.Writer() io.Writer

which must be tested by calling the methods:

- (*Chk).Log(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLog -->

<!--- gotomd::Bgn::doc::./CaptureLogAndStdout -->
```go
func CaptureLogAndStdout(t testingT) *Chk
```

CaptureLogAndStdout returns a new *sztest.Chk reference
capturing:

- log.Writer() io.Writer
- os.Stdout

which must be tested by calling the methods:

- (*Chk).Log(wantLines ...string) bool
- (*Chk).Stdout(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLogAndStdout -->

<!--- gotomd::Bgn::doc::./CaptureLogAndStderr -->
```go
func CaptureLogAndStderr(t testingT) *Chk
```

CaptureLogAndStderr returns a new *sztest.Chk reference
capturing:

- log.Writer() io.Writer
- os.Stderr

which must be tested by calling the methods:

- (*Chk).Log(wantLines ...string) bool
- (*Chk).Stderr(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLogAndStderr -->

<!--- gotomd::Bgn::doc::./CaptureLogAndStderrAndStdout -->
```go
func CaptureLogAndStderrAndStdout(t testingT) *Chk
```

CaptureLogAndStderrAndStdout returns a new *sztest.Chk reference
capturing:

- log.Writer() io.Writer
- os.Stderr
- os.Stdout

which must be tested by calling the methods:

- (*Chk).Log(wantLines ...string) bool
- (*Chk).Stderr(wantLines ...string) bool
- (*Chk).Stdout(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLogAndStderrAndStdout -->

<!--- gotomd::Bgn::doc::./CaptureLogWithStderr -->
```go
func CaptureLogWithStderr(t testingT) *Chk
```

CaptureLogWithStderr returns a new *sztest.Chk reference
combining and capturing:

- (log.Writer() io.Writer) + os.Stderr

which must be tested by calling ONE the methods:

- (*Chk).Log(wantLines ...string) bool
- OR
- (*Chk).Stderr(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLogWithStderr -->

<!--- gotomd::Bgn::doc::./CaptureLogWithStderrAndStdout -->
```go
func CaptureLogWithStderrAndStdout(t testingT) *Chk
```

CaptureLogWithStderrAndStdout returns a new *sztest.Chk reference
capturing:

- (log.Writer() io.Writer) + os.Stderr
- os.Stdout

which must be tested by calling ONE the methods:

- (*Chk).Log(wantLines ...string) bool
- OR
- (*Chk).Stderr(wantLines ...string) bool

and the method:

- (*Chk).Stdout(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureLogWithStderrAndStdout -->

<!--- gotomd::Bgn::doc::./CaptureStderr -->
```go
func CaptureStderr(t testingT) *Chk
```

CaptureStderr returns a new *sztest.Chk reference
capturing:

- os.Stderr

which must be tested by calling the method:

- (*Chk).Stderr(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureStderr -->

<!--- gotomd::Bgn::doc::./CaptureStderrAndStdout -->
```go
func CaptureStderrAndStdout(t testingT) *Chk
```

CaptureStderrAndStdout returns a new *sztest.Chk reference
capturing:

- os.Stderr
- os.Stdout

which must be tested by calling the methods:

- (*Chk).Stderr(wantLines ...string) bool
- (*Chk).Stdout(wantLines ...string) bool

before (*Chk).Release() is invoked.
<!--- gotomd::End::doc::./CaptureStderrAndStdout -->

[Contents](#contents)

### Appendix B: List of got/wnt test methods

#### Unformatted

<!--- gotomd::Bgn::dcls::./Chk.Bool Chk.False Chk.True Chk.Byte Chk.Complex64 Chk.Complex128 Chk.Float32 Chk.Float64 Chk.Int Chk.Int8 Chk.Int16 Chk.Int32 Chk.Int64 Chk.Rune Chk.Str Chk.Uint Chk.Uint8 Chk.Uint16 Chk.Uint32 Chk.Uint64 Chk.Uintptr Chk.Dur -->
```go
func (chk *Chk) Bool(got, want bool, msg ...any) bool
func (chk *Chk) False(got bool, msg ...any) bool
func (chk *Chk) True(got bool, msg ...any) bool
func (chk *Chk) Byte(got, want byte, msg ...any) bool
func (chk *Chk) Complex64(got, want complex64, msg ...any) bool
func (chk *Chk) Complex128(got, want complex128, msg ...any) bool
func (chk *Chk) Float32(got, want, tolerance float32, msg ...any) bool
func (chk *Chk) Float64(got, want, tolerance float64, msg ...any) bool
func (chk *Chk) Int(got, want int, msg ...any) bool
func (chk *Chk) Int8(got, want int8, msg ...any) bool
func (chk *Chk) Int16(got, want int16, msg ...any) bool
func (chk *Chk) Int32(got, want int32, msg ...any) bool
func (chk *Chk) Int64(got, want int64, msg ...any) bool
func (chk *Chk) Rune(got, want rune, msg ...any) bool
func (chk *Chk) Str(got, want string, msg ...any) bool
func (chk *Chk) Uint(got, want uint, msg ...any) bool
func (chk *Chk) Uint8(got, want uint8, msg ...any) bool
func (chk *Chk) Uint16(got, want uint16, msg ...any) bool
func (chk *Chk) Uint32(got, want uint32, msg ...any) bool
func (chk *Chk) Uint64(got, want uint64, msg ...any) bool
func (chk *Chk) Uintptr(got, want uintptr, msg ...any) bool
func (chk *Chk) Dur(got, want time.Duration, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Bool Chk.False Chk.True Chk.Byte Chk.Complex64 Chk.Complex128 Chk.Float32 Chk.Float64 Chk.Int Chk.Int8 Chk.Int16 Chk.Int32 Chk.Int64 Chk.Rune Chk.Str Chk.Uint Chk.Uint8 Chk.Uint16 Chk.Uint32 Chk.Uint64 Chk.Uintptr Chk.Dur -->

#### Formatted

<!--- gotomd::Bgn::dcls::./Chk.Boolf Chk.Falsef Chk.Truef Chk.Bytef Chk.Complex64f Chk.Complex128f Chk.Float32f Chk.Float64f Chk.Intf Chk.Int8f Chk.Int16f Chk.Int32f Chk.Int64f Chk.Runef Chk.Strf Chk.Uintf Chk.Uint8f Chk.Uint16f Chk.Uint32f Chk.Uint64f Chk.Uintptrf Chk.Durf -->
```go
func (chk *Chk) Boolf(got, want bool, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Falsef(got bool, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Truef(got bool, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Bytef(got, want byte, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Complex64f(got, want complex64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Complex128f(got, want complex128, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float32f(got, want, tolerance float32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float64f(got, want, tolerance float64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Intf(got, want int, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int8f(got, want int8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int16f(got, want int16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int32f(got, want int32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int64f(got, want int64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Runef(got, want rune, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Strf(got, want string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uintf(got, want uint, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint8f(got, want uint8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint16f(got, want uint16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint32f(got, want uint32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint64f(got, want uint64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uintptrf(got, want uintptr, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Durf(got, want time.Duration, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Boolf Chk.Falsef Chk.Truef Chk.Bytef Chk.Complex64f Chk.Complex128f Chk.Float32f Chk.Float64f Chk.Intf Chk.Int8f Chk.Int16f Chk.Int32f Chk.Int64f Chk.Runef Chk.Strf Chk.Uintf Chk.Uint8f Chk.Uint16f Chk.Uint32f Chk.Uint64f Chk.Uintptrf Chk.Durf -->

#### Pointers and References Unformatted

<!--- gotomd::Bgn::dcls::./Chk.Nil Chk.NotNil -->
```go
func (chk *Chk) Nil(got any, msg ...any) bool
func (chk *Chk) NotNil(got any, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Nil Chk.NotNil -->

#### Pointers and References Formatted

<!--- gotomd::Bgn::dcls::./Chk.Nilf Chk.NotNilf -->
```go
func (chk *Chk) Nilf(got any, msgFmt string, msgArgs ...any) bool
func (chk *Chk) NotNilf(got any, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Nilf Chk.NotNilf -->

[Contents](#contents)

### Appendix C: List of got/wnt slice test methods

#### Unformatted Slice

<!--- gotomd::Bgn::dcls::./Chk.BoolSlice Chk.ByteSlice Chk.Complex64Slice Chk.Complex128Slice Chk.Float32Slice Chk.Float64Slice Chk.IntSlice Chk.Int8Slice Chk.Int16Slice Chk.Int32Slice Chk.Int64Slice Chk.RuneSlice Chk.StrSlice Chk.UintSlice Chk.Uint8Slice Chk.Uint16Slice Chk.Uint32Slice Chk.Uint64Slice Chk.UintptrSlice Chk.DurSlice Chk.ErrSlice -->
```go
func (chk *Chk) BoolSlice(got, want []bool, msg ...any) bool
func (chk *Chk) ByteSlice(got, want []byte, msg ...any) bool
func (chk *Chk) Complex64Slice(got, want []complex64, msg ...any) bool
func (chk *Chk) Complex128Slice(got, want []complex128, msg ...any) bool
func (chk *Chk) Float32Slice(got, want []float32, tolerance float32, msg ...any) bool
func (chk *Chk) Float64Slice(got, want []float64, tolerance float64, msg ...any) bool
func (chk *Chk) IntSlice(got, want []int, msg ...any) bool
func (chk *Chk) Int8Slice(got, want []int8, msg ...any) bool
func (chk *Chk) Int16Slice(got, want []int16, msg ...any) bool
func (chk *Chk) Int32Slice(got, want []int32, msg ...any) bool
func (chk *Chk) Int64Slice(got, want []int64, msg ...any) bool
func (chk *Chk) RuneSlice(got, want []rune, msg ...any) bool
func (chk *Chk) StrSlice(got, want []string, msg ...any) bool
func (chk *Chk) UintSlice(got, want []uint, msg ...any) bool
func (chk *Chk) Uint8Slice(got, want []uint8, msg ...any) bool
func (chk *Chk) Uint16Slice(got, want []uint16, msg ...any) bool
func (chk *Chk) Uint32Slice(got, want []uint32, msg ...any) bool
func (chk *Chk) Uint64Slice(got, want []uint64, msg ...any) bool
func (chk *Chk) UintptrSlice(got, want []uintptr, msg ...any) bool
func (chk *Chk) DurSlice(got, want []time.Duration, msg ...any) bool
func (chk *Chk) ErrSlice(got []error, want []string, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.BoolSlice Chk.ByteSlice Chk.Complex64Slice Chk.Complex128Slice Chk.Float32Slice Chk.Float64Slice Chk.IntSlice Chk.Int8Slice Chk.Int16Slice Chk.Int32Slice Chk.Int64Slice Chk.RuneSlice Chk.StrSlice Chk.UintSlice Chk.Uint8Slice Chk.Uint16Slice Chk.Uint32Slice Chk.Uint64Slice Chk.UintptrSlice Chk.DurSlice Chk.ErrSlice -->

#### Formatted Slice

<!--- gotomd::Bgn::dcls::./Chk.BoolSlicef Chk.ByteSlicef Chk.Complex64Slicef Chk.Complex128Slicef Chk.Float32Slicef Chk.Float64Slicef Chk.IntSlicef Chk.Int8Slicef Chk.Int16Slicef Chk.Int32Slicef Chk.Int64Slicef Chk.RuneSlicef Chk.StrSlicef Chk.UintSlicef Chk.Uint8Slicef Chk.Uint16Slicef Chk.Uint32Slicef Chk.Uint64Slicef Chk.UintptrSlicef Chk.DurSlicef Chk.ErrSlicef -->
```go
func (chk *Chk) BoolSlicef(got, want []bool, msgFmt string, msgArgs ...any) bool
func (chk *Chk) ByteSlicef(got, want []byte, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Complex64Slicef(got, want []complex64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Complex128Slicef(got, want []complex128, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float32Slicef(got, want []float32, tolerance float32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float64Slicef(got, want []float64, tolerance float64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) IntSlicef(got, want []int, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int8Slicef(got, want []int8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int16Slicef(got, want []int16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int32Slicef(got, want []int32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int64Slicef(got, want []int64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) RuneSlicef(got, want []rune, msgFmt string, msgArgs ...any) bool
func (chk *Chk) StrSlicef(got, want []string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) UintSlicef(got, want []uint, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint8Slicef(got, want []uint8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint16Slicef(got, want []uint16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint32Slicef(got, want []uint32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint64Slicef(got, want []uint64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) UintptrSlicef(got, want []uintptr, msgFmt string, msgArgs ...any) bool
func (chk *Chk) DurSlicef(got, want []time.Duration, msgFmt string, msgArgs ...any) bool
func (chk *Chk) ErrSlicef(got []error, want []string, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.BoolSlicef Chk.ByteSlicef Chk.Complex64Slicef Chk.Complex128Slicef Chk.Float32Slicef Chk.Float64Slicef Chk.IntSlicef Chk.Int8Slicef Chk.Int16Slicef Chk.Int32Slicef Chk.Int64Slicef Chk.RuneSlicef Chk.StrSlicef Chk.UintSlicef Chk.Uint8Slicef Chk.Uint16Slicef Chk.Uint32Slicef Chk.Uint64Slicef Chk.UintptrSlicef Chk.DurSlicef Chk.ErrSlicef -->

[Contents](#contents)

### Appendix D: List of Bounded and Unbounded Interval tests

#### Bounded Unformatted

<!--- gotomd::Bgn::dcls::./Chk.ByteBounded Chk.Float32Bounded Chk.Float64Bounded Chk.IntBounded Chk.Int8Bounded Chk.Int16Bounded Chk.Int32Bounded Chk.Int64Bounded Chk.RuneBounded Chk.StrBounded Chk.UintBounded Chk.Uint8Bounded Chk.Uint16Bounded Chk.Uint32Bounded Chk.Uint64Bounded Chk.DurBounded -->
```go
func (chk *Chk) ByteBounded(got byte, option BoundedOption, min, max byte, msg ...any) bool
func (chk *Chk) Float32Bounded(got float32, option BoundedOption, min, max float32, msg ...any) bool
func (chk *Chk) Float64Bounded(got float64, option BoundedOption, min, max float64, msg ...any) bool
func (chk *Chk) IntBounded(got int, option BoundedOption, min, max int, msg ...any) bool
func (chk *Chk) Int8Bounded(got int8, option BoundedOption, min, max int8, msg ...any) bool
func (chk *Chk) Int16Bounded(got int16, option BoundedOption, min, max int16, msg ...any) bool
func (chk *Chk) Int32Bounded(got int32, option BoundedOption, min, max int32, msg ...any) bool
func (chk *Chk) Int64Bounded(got int64, option BoundedOption, min, max int64, msg ...any) bool
func (chk *Chk) RuneBounded(got rune, option BoundedOption, min, max rune, msg ...any) bool
func (chk *Chk) StrBounded(got string, option BoundedOption, min, max string, msg ...any) bool
func (chk *Chk) UintBounded(got uint, option BoundedOption, min, max uint, msg ...any) bool
func (chk *Chk) Uint8Bounded(got uint8, option BoundedOption, min, max uint8, msg ...any) bool
func (chk *Chk) Uint16Bounded(got uint16, option BoundedOption, min, max uint16, msg ...any) bool
func (chk *Chk) Uint32Bounded(got uint32, option BoundedOption, min, max uint32, msg ...any) bool
func (chk *Chk) Uint64Bounded(got uint64, option BoundedOption, min, max uint64, msg ...any) bool
func (chk *Chk) DurBounded(got time.Duration, option BoundedOption, min, max time.Duration, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.ByteBounded Chk.Float32Bounded Chk.Float64Bounded Chk.IntBounded Chk.Int8Bounded Chk.Int16Bounded Chk.Int32Bounded Chk.Int64Bounded Chk.RuneBounded Chk.StrBounded Chk.UintBounded Chk.Uint8Bounded Chk.Uint16Bounded Chk.Uint32Bounded Chk.Uint64Bounded Chk.DurBounded -->

#### Bounded Formatted

<!--- gotomd::Bgn::dcls::./Chk.ByteBoundedf Chk.Float32Boundedf Chk.Float64Boundedf Chk.IntBoundedf Chk.Int8Boundedf Chk.Int16Boundedf Chk.Int32Boundedf Chk.Int64Boundedf Chk.RuneBoundedf Chk.StrBoundedf Chk.UintBoundedf Chk.Uint8Boundedf Chk.Uint16Boundedf Chk.Uint32Boundedf Chk.Uint64Boundedf Chk.DurBoundedf -->
```go
func (chk *Chk) ByteBoundedf(got byte, option BoundedOption, min, max byte, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float32Boundedf(got float32, option BoundedOption, min, max float32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float64Boundedf(got float64, option BoundedOption, min, max float64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) IntBoundedf(got int, option BoundedOption, min, max int, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int8Boundedf(got int8, option BoundedOption, min, max int8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int16Boundedf(got int16, option BoundedOption, min, max int16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int32Boundedf(got int32, option BoundedOption, min, max int32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int64Boundedf(got int64, option BoundedOption, min, max int64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) RuneBoundedf(got rune, option BoundedOption, min, max rune, msgFmt string, msgArgs ...any) bool
func (chk *Chk) StrBoundedf(got string, option BoundedOption, min, max string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) UintBoundedf(got uint, option BoundedOption, min, max uint, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint8Boundedf(got uint8, option BoundedOption, min, max uint8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint16Boundedf(got uint16, option BoundedOption, min, max uint16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint32Boundedf(got uint32, option BoundedOption, min, max uint32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint64Boundedf(got uint64, option BoundedOption, min, max uint64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) DurBoundedf(got time.Duration, option BoundedOption, min, max time.Duration, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.ByteBoundedf Chk.Float32Boundedf Chk.Float64Boundedf Chk.IntBoundedf Chk.Int8Boundedf Chk.Int16Boundedf Chk.Int32Boundedf Chk.Int64Boundedf Chk.RuneBoundedf Chk.StrBoundedf Chk.UintBoundedf Chk.Uint8Boundedf Chk.Uint16Boundedf Chk.Uint32Boundedf Chk.Uint64Boundedf Chk.DurBoundedf -->

#### Unbounded Unformatted

<!--- gotomd::Bgn::dcls::./Chk.ByteUnbounded Chk.Float32Unbounded Chk.Float64Unbounded Chk.IntUnbounded Chk.Int8Unbounded Chk.Int16Unbounded Chk.Int32Unbounded Chk.Int64Unbounded Chk.RuneUnbounded Chk.StrUnbounded Chk.UintUnbounded Chk.Uint8Unbounded Chk.Uint16Unbounded Chk.Uint32Unbounded Chk.Uint64Unbounded Chk.DurUnbounded -->
```go
func (chk *Chk) ByteUnbounded(got byte, option UnboundedOption, bound byte, msg ...any) bool
func (chk *Chk) Float32Unbounded(got float32, option UnboundedOption, bound float32, msg ...any) bool
func (chk *Chk) Float64Unbounded(got float64, option UnboundedOption, bound float64, msg ...any) bool
func (chk *Chk) IntUnbounded(got int, option UnboundedOption, bound int, msg ...any) bool
func (chk *Chk) Int8Unbounded(got int8, option UnboundedOption, bound int8, msg ...any) bool
func (chk *Chk) Int16Unbounded(got int16, option UnboundedOption, bound int16, msg ...any) bool
func (chk *Chk) Int32Unbounded(got int32, option UnboundedOption, bound int32, msg ...any) bool
func (chk *Chk) Int64Unbounded(got int64, option UnboundedOption, bound int64, msg ...any) bool
func (chk *Chk) RuneUnbounded(got rune, option UnboundedOption, bound rune, msg ...any) bool
func (chk *Chk) StrUnbounded(got string, option UnboundedOption, bound string, msg ...any) bool
func (chk *Chk) UintUnbounded(got uint, option UnboundedOption, bound uint, msg ...any) bool
func (chk *Chk) Uint8Unbounded(got uint8, option UnboundedOption, bound uint8, msg ...any) bool
func (chk *Chk) Uint16Unbounded(got uint16, option UnboundedOption, bound uint16, msg ...any) bool
func (chk *Chk) Uint32Unbounded(got uint32, option UnboundedOption, bound uint32, msg ...any) bool
func (chk *Chk) Uint64Unbounded(got uint64, option UnboundedOption, bound uint64, msg ...any) bool
func (chk *Chk) DurUnbounded(got time.Duration, option UnboundedOption, bound time.Duration, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.ByteUnbounded Chk.Float32Unbounded Chk.Float64Unbounded Chk.IntUnbounded Chk.Int8Unbounded Chk.Int16Unbounded Chk.Int32Unbounded Chk.Int64Unbounded Chk.RuneUnbounded Chk.StrUnbounded Chk.UintUnbounded Chk.Uint8Unbounded Chk.Uint16Unbounded Chk.Uint32Unbounded Chk.Uint64Unbounded Chk.DurUnbounded -->

#### Unbounded Formatted

<!--- gotomd::Bgn::dcls::./Chk.ByteUnboundedf Chk.Float32Unboundedf Chk.Float64Unboundedf Chk.IntUnboundedf Chk.Int8Unboundedf Chk.Int16Unboundedf Chk.Int32Unboundedf Chk.Int64Unboundedf Chk.RuneUnboundedf Chk.StrUnboundedf Chk.UintUnboundedf Chk.Uint8Unboundedf Chk.Uint16Unboundedf Chk.Uint32Unboundedf Chk.Uint64Unboundedf Chk.DurUnboundedf -->
```go
func (chk *Chk) ByteUnboundedf(got byte, option UnboundedOption, bound byte, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float32Unboundedf(got float32, option UnboundedOption, bound float32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Float64Unboundedf(got float64, option UnboundedOption, bound float64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) IntUnboundedf(got int, option UnboundedOption, bound int, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int8Unboundedf(got int8, option UnboundedOption, bound int8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int16Unboundedf(got int16, option UnboundedOption, bound int16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int32Unboundedf(got int32, option UnboundedOption, bound int32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Int64Unboundedf(got int64, option UnboundedOption, bound int64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) RuneUnboundedf(got rune, option UnboundedOption, bound rune, msgFmt string, msgArgs ...any) bool
func (chk *Chk) StrUnboundedf(got string, option UnboundedOption, bound string, msgFmt string, msgArgs ...any) bool
func (chk *Chk) UintUnboundedf(got uint, option UnboundedOption, bound uint, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint8Unboundedf(got uint8, option UnboundedOption, bound uint8, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint16Unboundedf(got uint16, option UnboundedOption, bound uint16, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint32Unboundedf(got uint32, option UnboundedOption, bound uint32, msgFmt string, msgArgs ...any) bool
func (chk *Chk) Uint64Unboundedf(got uint64, option UnboundedOption, bound uint64, msgFmt string, msgArgs ...any) bool
func (chk *Chk) DurUnboundedf(got time.Duration, option UnboundedOption, bound time.Duration, msgFmt string, msgArgs ...any) bool
```
<!--- gotomd::End::dcls::./Chk.ByteUnboundedf Chk.Float32Unboundedf Chk.Float64Unboundedf Chk.IntUnboundedf Chk.Int8Unboundedf Chk.Int16Unboundedf Chk.Int32Unboundedf Chk.Int64Unboundedf Chk.RuneUnboundedf Chk.StrUnboundedf Chk.UintUnboundedf Chk.Uint8Unboundedf Chk.Uint16Unboundedf Chk.Uint32Unboundedf Chk.Uint64Unboundedf Chk.DurUnboundedf -->

[Contents](#contents)

### Appendix E: Builtin Ansi Terminal Markup

[See CONFIGURE.md Appendix E: Builtin Ansi Terminal Markup](CONFIGURE.md#appendix-e-builtin-ansi-terminal-markup)

[Contents](#contents)

### Appendix F: Large Example Function

[See Appendix F Example: Large Example Function](examples/appendix/README.md#appendix-f-large-example-function)

[Contents](#contents)

### Appendix G: Large Example Main Function

[See Appendix G: Large Example Main Function](examples/appendix/README.md#appendix-g-large-example-main-function)

[Contents](#contents)

### Appendix H: License

/*
   Golang testing utility.
   Copyright (C) 2023  Leslie Dancsecs

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

[Contents](#contents)

### Appendix I: To Be Documented

<!--- gotomd::Bgn::dcls::./Chk.FailFast Chk.Logf Chk.Errorf Chk.Error Chk.Fatalf Chk.Name Chk.T Chk.PushPreReleaseFunc Chk.PushPostReleaseFunc Chk.Release CompareArrays -->
```go
func (chk *Chk) FailFast(failFast bool) bool
func (chk *Chk) Logf(msgFmt string, msgArgs ...any)
func (chk *Chk) Errorf(msgFmt string, msgArgs ...any)
func (chk *Chk) Error(args ...any)
func (chk *Chk) Fatalf(msgFmt string, msgArgs ...any)
func (chk *Chk) Name() string
func (chk *Chk) T() testingT
func (chk *Chk) PushPreReleaseFunc(newFunc func() error)
func (chk *Chk) PushPostReleaseFunc(newFunc func() error)
func (chk *Chk) Release()
func CompareArrays[T chkType](got, wnt []T) string
```
<!--- gotomd::End::dcls::./Chk.FailFast Chk.Logf Chk.Errorf Chk.Error Chk.Fatalf Chk.Name Chk.T Chk.PushPreReleaseFunc Chk.PushPostReleaseFunc Chk.Release CompareArrays -->

<!--- gotomd::Bgn::dcls::./Chk.LastErr Chk.Log Chk.Stderr Chk.Stdout Chk.SetStdinData -->
```go
func (*Chk) LastErr(args ...any) error
func (chk *Chk) Log(wantLines ...string) bool
func (chk *Chk) Stderr(wantLines ...string) bool
func (chk *Chk) Stdout(wantLines ...string) bool
func (chk *Chk) SetStdinData(lines ...string)
```
<!--- gotomd::End::dcls::./Chk.LastErr Chk.Log Chk.Stderr Chk.Stdout Chk.SetStdinData -->

Errors can be tested using:

<!--- gotomd::Bgn::dcls::./Chk.Err -->
```go
func (chk *Chk) Err(got error, want string, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Err -->

noting that the got and wnt are different types.  Finally, checking for nil
pointers or nil references can be tested with:

<!--- gotomd::Bgn::dcls::./Chk.Nil -->
```go
func (chk *Chk) Nil(got any, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.Nil -->

Some checks have helper functions provided for convenience.

<!--- gotomd::Bgn::dcls::./Chk.False Chk.True Chk.NoErr Chk.NotNil -->
```go
func (chk *Chk) False(got bool, msg ...any) bool
func (chk *Chk) True(got bool, msg ...any) bool
func (chk *Chk) NoErr(got error, msg ...any) bool
func (chk *Chk) NotNil(got any, msg ...any) bool
```
<!--- gotomd::End::dcls::./Chk.False Chk.True Chk.NoErr Chk.NotNil -->

For a complete list of builtin Got/Wnt tests and their helpers see
[Appendix B: List of got/wnt Test Methods](#appendix-b-list-of-gotwnt-test-methods)

[Contents](#contents)
