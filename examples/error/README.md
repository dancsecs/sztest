<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Error Examples

- [Examples: No Error](#examples-no-error)
- [Examples: No Error Helper](#examples-no-error-helper)
- [Examples: Error](#examples-error)
- [Examples: Blank Error](#examples-blank-error)

[Contents](../../README.md#contents)

## Examples: No Error

<!--- gotomd::Bgn::file::./no_error/example_test.go -->
```bash
cat ./no_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := error(nil)

    chk.Err(
        err,
        "", // Empty string represents nil error.  Can be calculated.
    )
}

// Failing test.
func Test_FAIL_NoError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("unexpected error")

    chk.Err(
        err,
        "", // Empty string represents nil error.  Can be calculated.
    )
}
```
<!--- gotomd::End::file::./no_error/example_test.go -->

<!--- gotomd::Bgn::tst::./no_error/package -->
```bash
go test -v -cover ./no_error
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;NoError}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;NoError&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;NoError}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:30:&#xa0;unexpected&#xa0;err:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}unexpected&#xa0;error\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;NoError&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/error/no&#x332;error&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./no_error/package -->

[Contents](../../README.md#contents)

## Examples: No Error Helper

<!--- gotomd::Bgn::file::./no_error_helper/example_test.go -->
```bash
cat ./no_error_helper/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoErrorHelper(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := error(nil)

    chk.NoErr(err)
}

// Failing test.
func Test_FAIL_NoErrorHelper(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("unexpected error")

    chk.NoErr(err)
}
```
<!--- gotomd::End::file::./no_error_helper/example_test.go -->

<!--- gotomd::Bgn::tst::./no_error_helper/package -->
```bash
go test -v -cover ./no_error_helper
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;NoErrorHelper}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;NoErrorHelper&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;NoErrorHelper}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:27:&#xa0;unexpected&#xa0;err:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}unexpected&#xa0;error\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;NoErrorHelper&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/error/no&#x332;error&#x332;helper&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./no_error_helper/package -->

[Contents](../../README.md#contents)

## Examples: Error

<!--- gotomd::Bgn::file::./Error/example_test.go -->
```bash
cat ./Error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_Error(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("error condition")

    chk.Err(
        err,
        "error condition",
    )
}

// Failing test.
func Test_FAIL_Error(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("error condition generated")
    chk.Err(
        err,
        "error condition wanted",
    )
}
```
<!--- gotomd::End::file::./Error/example_test.go -->

<!--- gotomd::Bgn::tst::./Error/package -->
```bash
go test -v -cover ./Error
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;Error}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;Error&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;Error}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:29:&#xa0;unexpected&#xa0;err:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}error&#xa0;condition&#xa0;\color{darkturquoise}genera\color{default}ted}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}error&#xa0;condition&#xa0;\color{darkturquoise}wan\color{default}ted}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;Error&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/error/Error&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./Error/package -->

[Contents](../../README.md#contents)

## Examples: Blank Error

<!--- gotomd::Bgn::file::./blank_error/example_test.go -->
```bash
cat ./blank_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_BlankError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("")

    chk.Err(
        // blank "" error will be replaced with "sztest.EmptyErrorMessage"
        err,
        sztest.BlankErrorMessage,
    )
}

// Failing test.
func Test_FAIL_BlankError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    err := errors.New("")
    chk.Err(
        // blank "" error will be replaced with "sztest.EmptyErrorMessage"
        err,
        "Error message wanted",
    )
}
```
<!--- gotomd::End::file::./blank_error/example_test.go -->

<!--- gotomd::Bgn::tst::./blank_error/package -->
```bash
go test -v -cover ./blank_error
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;BlankError}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;BlankError&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;BlankError}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:30:&#xa0;unexpected&#xa0;err:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}sztest.Blank\color{default}Error\color{darkturquoise}M\color{default}essage}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}Error\color{darkturquoise}&#xa0;m\color{default}essage\color{red}&#xa0;wanted\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;BlankError&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/error/blank&#x332;error&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./blank_error/package -->

[Contents](../../README.md#contents)
