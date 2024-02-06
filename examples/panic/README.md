<!--- gotomd::Auto:: See github.com/dancsecs/gotomd ** DO NOT MODIFY ** -->

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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;NoPanic}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;NoPanic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;NoPanic}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:27:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}\color{green}this\unicode{160}is\unicode{160}the\unicode{160}panic\unicode{160}generated\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;NoPanic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/panic/no&#x332;panic\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;NoPanicHelper}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;NoPanicHelper\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;NoPanicHelper}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:26:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}\color{green}panic\unicode{160}message\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;NoPanicHelper\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/panic/no&#x332;panic&#x332;helper\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;Panic}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;Panic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;Panic}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:29:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}this\unicode{160}is\unicode{160}the\unicode{160}panic\unicode{160}\color{darkturquoise}genera\color{default}ted}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}this\unicode{160}is\unicode{160}the\unicode{160}panic\unicode{160}\color{darkturquoise}wan\color{default}ted}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;Panic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/panic/panic\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;BlankPanic}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;BlankPanic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;BlankPanic}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:29:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}\color{green}sztest.BlankPanicMessage\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;BlankPanic\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/panic/blank&#x332;panic\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./blank_panic/package -->

[Contents](../../README.md#contents)
