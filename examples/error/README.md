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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;NoError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;NoError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;NoError}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:30:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;err:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{green}{unexpected&#xA0;&#x34F;&#xA0;&#x34F;error}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;NoError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/error/no&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;NoErrorHelper}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;NoErrorHelper&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;NoErrorHelper}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:27:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;err:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{green}{unexpected&#xA0;&#x34F;&#xA0;&#x34F;error}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;NoErrorHelper&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/error/no&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x332;&#xA0;&#x332;helper&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Error}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Error&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Error}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:29:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;err:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}error&#xA0;&#x34F;&#xA0;&#x34F;condition&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{genera}}ted}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}error&#xA0;&#x34F;&#xA0;&#x34F;condition&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{wan}}ted}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Error&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/error/Error&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;BlankError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;BlankError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;BlankError}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:30:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;err:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{green}{sztest.Blank}}Error{\color{darkturquoise}{M}}essage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}Error{\color{darkturquoise}{&#xA0;&#x34F;&#xA0;&#x34F;m}}essage{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;wanted}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;BlankError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/error/blank&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./blank_error/package -->

[Contents](../../README.md#contents)
