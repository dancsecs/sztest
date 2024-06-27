<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Panic examples

- [Examples: No Panic](#examples-no-panic)
- [Examples: No Panic Helper](#examples-no-panic-helper)
- [Examples: Panic](#examples-panic)
- [Examples: Blank Panic](#examples-blank-panic)

[Contents](../../README.md#contents)

## Examples: No Panic

<!--- gotomd::Bgn::file::./no_panic/example_test.go -->
```bash
cat ./no_panic/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoPanic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Panic(
        func() {
            // Exit function without panicking.
        },
        "", // Permits a NoPanic wnt to be calculated by the test.
    )
}

// Failing test.
func Test_FAIL_NoPanic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Panic(
        func() {
            panic("this is the panic generated")
        },
        // Expecting no panic to be thrown
        "",
    )
}
```
<!--- gotomd::End::file::./no_panic/example_test.go -->

<!--- gotomd::Bgn::tst::./no_panic/package -->
```bash
go test -v -cover ./no_panic
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;NoPanic}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;NoPanic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;NoPanic}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:27:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}this&#xa0;is&#xa0;the&#xa0;panic&#xa0;generated\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;NoPanic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/panic/no&#x332;panic&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./no_panic/package -->

[Contents](../../README.md#contents)

## Examples: No Panic Helper

<!--- gotomd::Bgn::file::./no_panic_helper/example_test.go -->
```bash
cat ./no_panic_helper/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoPanicHelper(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.NoPanic(
        func() {
            // Exit function without panicking.
        },
    )
}

// Failing test.
func Test_FAIL_NoPanicHelper(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.NoPanic(
        func() {
            panic("panic message") // Unexpected Panic
        },
    )
}
```
<!--- gotomd::End::file::./no_panic_helper/example_test.go -->

<!--- gotomd::Bgn::tst::./no_panic_helper/package -->
```bash
go test -v -cover ./no_panic_helper
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;NoPanicHelper}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;NoPanicHelper&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;NoPanicHelper}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:26:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}panic&#xa0;message\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;NoPanicHelper&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/panic/no&#x332;panic&#x332;helper&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./no_panic_helper/package -->

[Contents](../../README.md#contents)

## Examples: Panic

<!--- gotomd::Bgn::file::./panic/example_test.go -->
```bash
cat ./panic/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_Panic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Panic(
        func() {
            panic("expected panic message")
        },
        "expected panic message",
    )
}

// Failing test.
func Test_FAIL_Panic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    // Failure.  The invoked function's panic message
    // is not what is expected.
    chk.Panic(
        func() {
            panic("this is the panic generated")
        },
        "this is the panic wanted",
    )
}
```
<!--- gotomd::End::file::./panic/example_test.go -->

<!--- gotomd::Bgn::tst::./panic/package -->
```bash
go test -v -cover ./panic
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;Panic}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;Panic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;Panic}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:29:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}this&#xa0;is&#xa0;the&#xa0;panic&#xa0;\color{darkturquoise}genera\color{default}ted}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}this&#xa0;is&#xa0;the&#xa0;panic&#xa0;\color{darkturquoise}wan\color{default}ted}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;Panic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/panic/panic&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./panic/package -->

[Contents](../../README.md#contents)

## Examples: Blank Panic

<!--- gotomd::Bgn::file::./blank_panic/example_test.go -->
```bash
cat ./blank_panic/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_BlankPanic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Panic(
        func() {
            panic("") // Blank panic message.
        },
        sztest.BlankPanicMessage, // Panic without message is expected.
        // Message returned in place of an empty string representing
        // that an empty ("") panic was issued by the function.
    )
}

// Failing test.
func Test_FAIL_BlankPanic(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Panic(
        func() {
            panic("") // Empty panic message will be flagged.
        },
        "",
        // sztest.BlankPanicMessage will be returned.
    )
}
```
<!--- gotomd::End::file::./blank_panic/example_test.go -->

<!--- gotomd::Bgn::tst::./blank_panic/package -->
```bash
go test -v -cover ./blank_panic
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;BlankPanic}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;BlankPanic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;BlankPanic}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:29:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}sztest.BlankPanicMessage\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;BlankPanic&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/panic/blank&#x332;panic&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./blank_panic/package -->

[Contents](../../README.md#contents)
