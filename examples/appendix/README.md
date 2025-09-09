<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Package sztest

## Example Contents

- [Appendix F: Large Example Function](#appendix-f-large-example-function)
- [Appendix G: Large Example Main Function](#appendix-g-large-example-main-function)

[Contents](../../README.md#contents)

## Appendix F: Large Example Function

This example shows several of the types of checks this library is designed to
assist with.  Suppose we have the following function:

<!--- gotomd::Bgn::file::./large_example_function/example.go -->
```bash
cat ./large_example_function/example.go
```

```go
// Package example demonstrates a larger test function.
package example

import (
    "fmt"
    "log"
)

// Process returns known changes to its parameters for testing.
func Process(factor int, msg string) (int, string, float64) {
    const mulFactor = 2
    const aThird = 1.0 / 3.0
    // Function being tested.
    log.Printf("Entered process(%d, %q)", factor, msg)

    if factor < 1 || factor > 10 {
        log.Panicf("factor (%d) out of bounds: %q", factor, msg)
    }

    fmt.Println("Processing with factor:", factor, "and message:", msg)

    return factor * mulFactor, "Processed: " + msg, float64(factor) * aThird
}
```
<!--- gotomd::End::file::./large_example_function/example.go -->

we could test all of the expected outputs with the following test (with
passing and failing version)s

<!--- gotomd::Bgn::file::./large_example_function/example_test.go -->
```bash
cat ./large_example_function/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_GeneralForm(t *testing.T) {
    chk := sztest.CaptureLogAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Don't exit test on first failure.

    // Test panic condition.
    chk.Panic(
        func() {
            Process(0, "failing message")
        },
        "factor (0) out of bounds: \"failing message\"",
    )

    // Test valid operation.
    gotInt, gotStr, gotFloat := Process(2, "Hello")

    chk.Int(gotInt, 4)
    chk.Str(gotStr, "Processed: Hello")
    chk.Float64(gotFloat, 0.6666, 0.0005) // Tolerance (± 0.0005)

    // Check output.
    chk.Stdout(
        "Processing with factor: 2 and message: Hello",
    )

    // Check logging.
    chk.Log(`
      Entered process(0, "failing message")
      factor (0) out of bounds: "failing message"
      Entered process(2, "Hello")
  `)
}

// Failing test.
func Test_FAIL_GeneralForm(t *testing.T) {
    chk := sztest.CaptureLogAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Don't exit test on first failure.

    // Test panic condition.
    chk.Panic(
        func() {
            Process(0, "failing message")
        },
        "factor (0) out of bounds: \"wrong message\"",
    )

    // Test valid operation.
    gotInt, gotStr, gotFloat := Process(2, "Hello")

    chk.Int(gotInt, 3)
    chk.Str(gotStr, "Hello", "missing", " ", `"Processed"`, " ", "prefix")
    // Tolerance (± 0.0005)
    chk.Float64f(gotFloat, 0.6661, 0.0005, "off by %f", 0.6661-0.6665)

    // Check output.
    chk.Stdout(
        "Processing with factor: 2 and message: Processed Hello",
    )

    // Check logging.
    chk.Log(`
      Entered process(0, "wrong message")
      factor (0) out of bounds: "wrong message"
      Entered process(2, "Processed Hello")
  `)
}
```
<!--- gotomd::End::file::./large_example_function/example_test.go -->

causing the following output when these tests are run:

<!--- gotomd::Bgn::tst::./large_example_function/package -->
```bash
go test -v -cover ./large_example_function
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;GeneralForm}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:37:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{Entered&#xA0;&#x34F;&#xA0;&#x34F;process(0,&#xA0;&#x34F;&#xA0;&#x34F;"failing&#xA0;&#x34F;&#xA0;&#x34F;message")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{1}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Entered&#xA0;&#x34F;&#xA0;&#x34F;process(0,}}{\color{yellow}{/}}{\color{green}{factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:}}&#xA0;&#x34F;&#xA0;&#x34F;"failing&#xA0;&#x34F;&#xA0;&#x34F;message"{\color{red}{)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{2}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;"failing&#xA0;&#x34F;&#xA0;&#x34F;m}}{\color{yellow}{/}}{\color{green}{Entered&#xA0;&#x34F;&#xA0;&#x34F;proc}}ess{\color{red}{age"}}{\color{yellow}{/}}{\color{green}{(2,&#xA0;&#x34F;&#xA0;&#x34F;"Hello")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Entered&#xA0;&#x34F;&#xA0;&#x34F;process(2,&#xA0;&#x34F;&#xA0;&#x34F;"Hello")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;GeneralForm&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;GeneralForm}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:52:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;panic:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;"{\color{darkturquoise}{faili}}ng&#xA0;&#x34F;&#xA0;&#x34F;message"}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;"{\color{darkturquoise}{wro}}ng&#xA0;&#x34F;&#xA0;&#x34F;message"}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:62:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;int:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{4}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{darkturquoise}{3}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:63:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;string:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\emph{missing&#xA0;&#x34F;&#xA0;&#x34F;"Processed"&#xA0;&#x34F;&#xA0;&#x34F;prefix}}:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{green}{Processed:&#xA0;&#x34F;&#xA0;&#x34F;}}Hello}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}Hello}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:65:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;float64(+/-&#xA0;&#x34F;&#xA0;&#x34F;0.0005):}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\emph{off&#xA0;&#x34F;&#xA0;&#x34F;by&#xA0;&#x34F;&#xA0;&#x34F;-0.000400}}:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}0.666{\color{darkturquoise}{6666666666666}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}0.666{\color{darkturquoise}{1}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:68:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;Processing&#xA0;&#x34F;&#xA0;&#x34F;with&#xA0;&#x34F;&#xA0;&#x34F;factor:&#xA0;&#x34F;&#xA0;&#x34F;2&#xA0;&#x34F;&#xA0;&#x34F;and&#xA0;&#x34F;&#xA0;&#x34F;message:&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{Processed&#xA0;&#x34F;&#xA0;&#x34F;}}Hello}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:73:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{Entered&#xA0;&#x34F;&#xA0;&#x34F;process(0,&#xA0;&#x34F;&#xA0;&#x34F;"failing&#xA0;&#x34F;&#xA0;&#x34F;message")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{1}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Entered&#xA0;&#x34F;&#xA0;&#x34F;process(0,&#xA0;&#x34F;&#xA0;&#x34F;"wro}}{\color{yellow}{/}}{\color{green}{factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;"faili}}ng&#xA0;&#x34F;&#xA0;&#x34F;message"{\color{red}{)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{2}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;factor&#xA0;&#x34F;&#xA0;&#x34F;(0)&#xA0;&#x34F;&#xA0;&#x34F;out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;"wrong&#xA0;&#x34F;&#xA0;&#x34F;m}}{\color{yellow}{/}}{\color{green}{Entered&#xA0;&#x34F;&#xA0;&#x34F;proc}}ess{\color{red}{age"}}{\color{yellow}{/}}{\color{green}{(2,&#xA0;&#x34F;&#xA0;&#x34F;"Hello")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Entered&#xA0;&#x34F;&#xA0;&#x34F;process(2,&#xA0;&#x34F;&#xA0;&#x34F;"Processed&#xA0;&#x34F;&#xA0;&#x34F;Hello")}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;GeneralForm&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/appendix/large&#xA0;&#x332;&#xA0;&#x332;example&#xA0;&#x332;&#xA0;&#x332;function&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./large_example_function/package -->

[Contents](../../README.md#contents)

## Appendix G: Large Example Main Function

Given the following main.go file:

<!--- gotomd::Bgn::file::./large_example_main_function/main.go -->
```bash
cat ./large_example_main_function/main.go
```

```go
package main

import (
    "flag"
    "fmt"
    "log"
    "math"
    "strconv"
    "strings"
)

const radiansPerDegree = math.Pi / 180.0
const degreesPerRadian = 180.0 / math.Pi
const bits64 = 64
const numDecimals = 6

func makeReport(degrees, radians float64, useRadians, verbose bool) []string {
    var rep []string
    addToReport := func(v float64, name string) {
        line := strconv.FormatFloat(v, 'f', numDecimals, bits64)
        if verbose {
            arg := ""
            if !useRadians {
                arg = strconv.FormatFloat(degrees, 'f', numDecimals, bits64) + "°"
            } else {
                arg = strconv.FormatFloat(radians, 'f', numDecimals, bits64)
            }
            line = name + "(" + arg + ") = " + line
        }
        rep = append(rep, line)
    }
    addToReport(math.Sin(radians), "Sin")
    addToReport(math.Cos(radians), "Cos")
    return rep
}

// Program takes an angle and reports its Sin and Cos values.
// -v causes a more detailed response.
// -r cause the angle input to be interrupted as radians.
func main() {
    var degrees, radians float64

    verbose := flag.Bool("v", false, "More detailed information.")
    useRadians := flag.Bool("r", false, "Value is in radians.")

    flag.Parse()

    if flag.NArg() != 1 {
        log.Panic("angle required")
    }
    v, err := strconv.ParseFloat(flag.Args()[0], bits64)

    if err != nil {
        log.Panicf("invalid angle: %s", flag.Args()[0])
    }

    if *useRadians {
        degrees = degreesPerRadian * v
        radians = v
    } else {
        degrees = v
        radians = radiansPerDegree * v
    }

    if *verbose {
        if *useRadians {
            fmt.Printf("Report on %f radians (%f degrees)\n", radians, degrees)
        } else {
            fmt.Printf("Report on %f degrees (%f radians)\n", degrees, radians)
        }
    }

    fmt.Print(
        strings.Join(makeReport(degrees, radians, *useRadians, *verbose), "\n"),
        "\n",
    )
}
```
<!--- gotomd::End::file::./large_example_main_function/main.go -->

we can provide 100% test coverage with the following test file showing both
successful and failing responses:

<!--- gotomd::Bgn::file::./large_example_main_function/main_test.go -->
```bash
cat ./large_example_main_function/main_test.go
```

```go
package main

import (
    "fmt"
    "log"
    "math"
    "strconv"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_PASS_Main_No_Args(t *testing.T) {
    chk := sztest.CaptureLogWithStderrAndStdout(t)
    defer chk.Release()

    log.Println("Testing missing angle")
    chk.SetArgs("progname")
    chk.Panic(
        main,
        "angle required",
    )

    log.Println("Testing invalid angle")
    chk.SetArgs("progname", "notANumber")
    chk.Panic(
        main,
        "invalid angle: notANumber",
    )

    fmt.Println("Testing angle 0 no flags")
    chk.SetArgs("progname", "0")
    chk.NoPanic(main)

    fmt.Println("Testing angle 0 with verbose flag")
    chk.SetArgs("progname", "-v", "0")
    chk.NoPanic(main)

    twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
    fmt.Println("Testing angle 2Pi with radian flag")
    chk.SetArgs("progname", "-r", twoPi)
    chk.NoPanic(main)

    fmt.Println("Testing angle 2Pi with radian and verbose flag")
    chk.SetArgs("progname", "-v", "-r", twoPi)
    chk.NoPanic(main)

    chk.Stdout(
        "Testing angle 0 no flags",
        "0.000000",
        "1.000000",
        "Testing angle 0 with verbose flag",
        "Report on 0.000000 degrees (0.000000 radians)",
        "Sin(0.000000°) = 0.000000",
        "Cos(0.000000°) = 1.000000",
        "Testing angle 2Pi with radian flag",
        "-0.000000",
        "1.000000",
        "Testing angle 2Pi with radian and verbose flag",
        "Report on 6.283185 radians (360.000000 degrees)",
        "Sin(6.283185) = -0.000000",
        "Cos(6.283185) = 1.000000",
    )
    chk.Log(
        "Testing missing angle",
        "angle required",
        //
        "Testing invalid angle",
        "invalid angle: notANumber",
        //
    )
}

func Test_FAIL_Main_No_Args(t *testing.T) {
    chk := sztest.CaptureLogWithStderrAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Do not terminate function on first error.

    log.Println("Testing missing angle")
    chk.SetArgs("progname")
    chk.Panic(
        main,
        "angle is required",
    )

    log.Println("Testing invalid angle")
    chk.SetArgs("progname", "notANumber")
    chk.Panic(
        main,
        "invalid angle: not A Number",
    )

    fmt.Println("Testing angle 0 no flags")
    chk.SetArgs("progname", "0")
    chk.NoPanic(main)

    fmt.Println("Testing angle 0 with verbose flag")
    chk.SetArgs("progname", "-v", "0")
    chk.NoPanic(main)

    twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
    fmt.Println("Testing angle 2Pi with radian flag")
    chk.SetArgs("progname", "-r", twoPi)
    chk.NoPanic(main)

    fmt.Println("Testing angle 2Pi with radian and verbose flag")
    chk.SetArgs("progname", "-v", "-r", twoPi)
    chk.NoPanic(main)

    chk.Stdout(
        "Testing angle 0 no flags",
        "0.000000",
        "1.000000",
        "Testing angle 0 with verbose flag",
        "Report on 0.000000 degrees (0.000000 radians)",
        "Sin(0.000000°) = 0.000000",
        "Cos(0.000000°) = 1.000000",
        "Testing angle 2PI with radian flag",
        "-0.000000",
        "1.000000",
        "Testing angle 2Pi with radian and verbose flag",
        "Report on 6.283185 radians (360.000000 degrees)",
        "Sin(6.283185) = -0.000000",
        "Cos(6.283185) = 1.000000",
    )
    chk.Log(
        "Testing missing angle",
        "angle is required",
        //
        "Testing invalid angle",
        "invalid angle: not A Number",
        //
    )
}
```
<!--- gotomd::End::file::./large_example_main_function/main_test.go -->

<!--- gotomd::Bgn::tst::./large_example_main_function/package -->
```bash
go test -v -cover ./large_example_main_function
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Main&#xA0;&#x332;&#xA0;&#x332;No&#xA0;&#x332;&#xA0;&#x332;Args}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Main&#xA0;&#x332;&#xA0;&#x332;No&#xA0;&#x332;&#xA0;&#x332;Args&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Main&#xA0;&#x332;&#xA0;&#x332;No&#xA0;&#x332;&#xA0;&#x332;Args}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;main&#xA0;&#x332;&#xA0;&#x332;test.go:82:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;panic:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}angle&#xA0;&#x34F;&#xA0;&#x34F;required}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}angle{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;is}}&#xA0;&#x34F;&#xA0;&#x34F;required}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;main&#xA0;&#x332;&#xA0;&#x332;test.go:89:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;panic:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}invalid&#xA0;&#x34F;&#xA0;&#x34F;angle:&#xA0;&#x34F;&#xA0;&#x34F;not{\color{darkturquoise}{A}}Number}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}invalid&#xA0;&#x34F;&#xA0;&#x34F;angle:&#xA0;&#x34F;&#xA0;&#x34F;not{\color{darkturquoise}{&#xA0;&#x34F;&#xA0;&#x34F;A&#xA0;&#x34F;&#xA0;&#x34F;}}Number}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;main&#xA0;&#x332;&#xA0;&#x332;test.go:111:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(14&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(14&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;00:00&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;angle&#xA0;&#x34F;&#xA0;&#x34F;0&#xA0;&#x34F;&#xA0;&#x34F;no&#xA0;&#x34F;&#xA0;&#x34F;flags}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;01:01&#xA0;&#x34F;&#xA0;&#x34F;0.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;02:02&#xA0;&#x34F;&#xA0;&#x34F;1.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;03:03&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;angle&#xA0;&#x34F;&#xA0;&#x34F;0&#xA0;&#x34F;&#xA0;&#x34F;with&#xA0;&#x34F;&#xA0;&#x34F;verbose&#xA0;&#x34F;&#xA0;&#x34F;flag}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;04:04&#xA0;&#x34F;&#xA0;&#x34F;Report&#xA0;&#x34F;&#xA0;&#x34F;on&#xA0;&#x34F;&#xA0;&#x34F;0.000000&#xA0;&#x34F;&#xA0;&#x34F;degrees&#xA0;&#x34F;&#xA0;&#x34F;(0.000000&#xA0;&#x34F;&#xA0;&#x34F;radians)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;05:05&#xA0;&#x34F;&#xA0;&#x34F;Sin(0.000000°)&#xA0;&#x34F;&#xA0;&#x34F;=&#xA0;&#x34F;&#xA0;&#x34F;0.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;06:06&#xA0;&#x34F;&#xA0;&#x34F;Cos(0.000000°)&#xA0;&#x34F;&#xA0;&#x34F;=&#xA0;&#x34F;&#xA0;&#x34F;1.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{07}}:{\color{darkturquoise}{07}}&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;angle&#xA0;&#x34F;&#xA0;&#x34F;2P{\color{red}{I}}{\color{yellow}{/}}{\color{green}{i}}&#xA0;&#x34F;&#xA0;&#x34F;with&#xA0;&#x34F;&#xA0;&#x34F;radian&#xA0;&#x34F;&#xA0;&#x34F;flag}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;08:08&#xA0;&#x34F;&#xA0;&#x34F;-0.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;09:09&#xA0;&#x34F;&#xA0;&#x34F;1.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;10:10&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;angle&#xA0;&#x34F;&#xA0;&#x34F;2Pi&#xA0;&#x34F;&#xA0;&#x34F;with&#xA0;&#x34F;&#xA0;&#x34F;radian&#xA0;&#x34F;&#xA0;&#x34F;and&#xA0;&#x34F;&#xA0;&#x34F;verbose&#xA0;&#x34F;&#xA0;&#x34F;flag}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;11:11&#xA0;&#x34F;&#xA0;&#x34F;Report&#xA0;&#x34F;&#xA0;&#x34F;on&#xA0;&#x34F;&#xA0;&#x34F;6.283185&#xA0;&#x34F;&#xA0;&#x34F;radians&#xA0;&#x34F;&#xA0;&#x34F;(360.000000&#xA0;&#x34F;&#xA0;&#x34F;degrees)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;12:12&#xA0;&#x34F;&#xA0;&#x34F;Sin(6.283185)&#xA0;&#x34F;&#xA0;&#x34F;=&#xA0;&#x34F;&#xA0;&#x34F;-0.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;13:13&#xA0;&#x34F;&#xA0;&#x34F;Cos(6.283185)&#xA0;&#x34F;&#xA0;&#x34F;=&#xA0;&#x34F;&#xA0;&#x34F;1.000000}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;main&#xA0;&#x332;&#xA0;&#x332;test.go:127:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;logWithStderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(4&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(4&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;missing&#xA0;&#x34F;&#xA0;&#x34F;angle}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{1}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;angle{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;is}}&#xA0;&#x34F;&#xA0;&#x34F;required}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;Testing&#xA0;&#x34F;&#xA0;&#x34F;invalid&#xA0;&#x34F;&#xA0;&#x34F;angle}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{3}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;invalid&#xA0;&#x34F;&#xA0;&#x34F;angle:&#xA0;&#x34F;&#xA0;&#x34F;not{\color{red}{&#xA0;&#x34F;&#xA0;&#x34F;A&#xA0;&#x34F;&#xA0;&#x34F;}}{\color{yellow}{/}}{\color{green}{A}}Number}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;Main&#xA0;&#x332;&#xA0;&#x332;No&#xA0;&#x332;&#xA0;&#x332;Args&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/appendix/large&#xA0;&#x332;&#xA0;&#x332;example&#xA0;&#x332;&#xA0;&#x332;main&#xA0;&#x332;&#xA0;&#x332;function&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./large_example_main_function/package -->

[Contents](../../README.md#contents)
