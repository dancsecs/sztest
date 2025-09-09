<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Output Examples

- [Examples: Output Stdout](#examples-output-stdout)
- [Examples: Output Stderr](#examples-output-stderr)
- [Examples: Output Stderr And Stdout](#examples-output-stderr-and-stdout)
- [Examples: Output Log](#examples-output-log)
- [Examples: Output Log And Stdout](#examples-output-log-and-stdout)
- [Examples: Output Log And Stderr](#examples-output-log-and-stderr)
- [Examples: Output Log And Stderr And Stdout](#examples-output-log-and-stderr-and-stdout)
- [Examples: Output Log With Stderr](#examples-output-log-with-stderr)
- [Examples: Output Log With Stderr And Stdout](#examples-output-log-with-stderr-and-stdout)

[Contents](../../README.md#contents)

## Examples Output Stdout

<!--- gotomd::Bgn::file::./stdout/example_test.go -->
```bash
cat ./stdout/example_test.go
```

```go
package example

import (
    "fmt"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureStdout(t *testing.T) {
    chk := sztest.CaptureStdout(t)
    defer chk.Release()

    arg1 := "1"
    arg2 := "2"

    fmt.Printf("Line %s\n", arg1)
    fmt.Printf("Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)
    // Stdout output as expected?
    chk.Stdout(
        "Line {{arg1}}",
        "Line {{arg2}}",
    )
}

// Failing test.
func Test_FAIL_CaptureStdout(t *testing.T) {
    chk := sztest.CaptureStdout(t)
    defer chk.Release()

    arg1 := "1"
    arg2 := "2"

    fmt.Printf("Line %s\n", arg1)
    fmt.Print("Missing in want\n")
    fmt.Printf("Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)
    // Stdout output as expected?
    chk.Stdout(
        "Line {{arg1}}",
        "Line {{arg2}}",
    )
}
```
<!--- gotomd::End::file::./stdout/example_test.go -->

<!--- gotomd::Bgn::tst::./stdout/package -->
```bash
go test -v -cover ./stdout
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStdout}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStdout}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:45:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(2&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{1}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:1&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/stdout&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./stdout/package -->

> The failure highlights that there was a line in the output (the got) that is
missing from the ```chk.Stdout()``` list (the want).  IE the line Missing in
want indicates that this line needs to be added to the want.

[Contents](../../README.md#contents)

### Examples: Output Stderr

<!--- gotomd::Bgn::file::./stderr/example_test.go -->
```bash
cat ./stderr/example_test.go
```

```go
package example

import (
    "fmt"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureStderr(t *testing.T) {
    chk := sztest.CaptureStderr(t)
    defer chk.Release()

    arg1 := "1"
    arg2 := "2"

    fmt.Fprintf(os.Stderr, "Line %s\n", arg1)
    fmt.Fprintf(os.Stderr, "Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)
    // Stderr output as expected?
    chk.Stderr(
        "Line {{arg1}}",
        "Line {{arg2}}",
    )
}

// Failing test.
func Test_FAIL_CaptureStderr(t *testing.T) {
    chk := sztest.CaptureStderr(t)
    defer chk.Release()

    arg1 := "1"
    arg2 := "2"

    fmt.Fprintf(os.Stderr, "Line %s\n", arg1)
    fmt.Fprintf(os.Stderr, "Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)
    // Stderr output as expected?
    chk.Stderr(
        "Line {{arg1}}",
        "Missing in got",
        "Line {{arg2}}",
    )
}
```
<!--- gotomd::End::file::./stderr/example_test.go -->

<!--- gotomd::Bgn::tst::./stderr/package -->
```bash
go test -v -cover ./stderr
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStderr}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStderr}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:45:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(2&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:2&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/stderr&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./stderr/package -->

> The failure highlights that there was a line in the ```chk.Stderr()``` list
(the want) that is missing from the got.  IE the line Missing in got
indicates that this line needs to be removed from the want.

[Contents](../../README.md#contents)

### Examples: Output Stderr And Stdout

<!--- gotomd::Bgn::file::./stderr_and_stdout/example_test.go -->
```bash
cat ./stderr_and_stdout/example_test.go
```

```go
package example

import (
    "fmt"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// / Passing test.
func Test_PASS_CaptureStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureStderrAndStdout(t)
    defer chk.Release()

    arg1 := "1"
    arg2 := "2"

    fmt.Printf("Stdout: Line %s\n", arg1)
    fmt.Printf("Stdout: Line %s\n", arg2)

    fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg1)
    fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)

    // Stdout output as expected?
    chk.Stdout(
        "Stdout: Line {{arg1}}",
        "Stdout: Line {{arg2}}",
    )

    // Stderr output as expected?
    chk.Stderr(
        "Stderr: Line {{arg1}}",
        "Stderr: Line {{arg2}}",
    )
}

// Failing test.
func Test_FAIL_CaptureStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureStderrAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    arg1 := "1"
    arg2 := "2"

    fmt.Printf("Stdout: Line %s\n", arg1)
    fmt.Println("Stdout: Missing In Want")
    fmt.Printf("Stdout: Line %s\n", arg2)

    fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg1)
    fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg2)

    chk.AddSub("{{arg1}}", arg1)
    chk.AddSub("{{arg2}}", arg2)

    // Stdout output as expected?
    chk.Stdout(
        "Stdout: Line {{arg1}}",
        "Stdout: Line {{arg2}}",
    )

    // Stderr output as expected?
    chk.Stderr(
        "Stderr: Line {{arg1}}",
        "StdErr: Missing in got",
        "Stderr: Line {{arg2}}",
    )
}
```
<!--- gotomd::End::file::./stderr_and_stdout/example_test.go -->

<!--- gotomd::Bgn::tst::./stderr_and_stdout/package -->
```bash
go test -v -cover ./stderr_and_stdout
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:62:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(2&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;Stdout:&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{1}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{Stdout:&#xA0;&#x34F;&#xA0;&#x34F;Missing&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:1&#xA0;&#x34F;&#xA0;&#x34F;Stdout:&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:68:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(2&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(3&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;Stderr:&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{StdErr:&#xA0;&#x34F;&#xA0;&#x34F;Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:2&#xA0;&#x34F;&#xA0;&#x34F;Stderr:&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/stderr&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stdout&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./stderr_and_stdout/package -->

> Note we first set ```chk.FailFast(false)``` to have all checks within this
test run before terminating the test function.

[Contents](../../README.md#contents)

### Examples: Output Log

<!--- gotomd::Bgn::file::./log/example_test.go -->
```bash
cat ./log/example_test.go
```

```go
package example

import (
    "log"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLog(t *testing.T) {
    chk := sztest.CaptureLog(t)
    defer chk.Release()

    log.Print("Line 1")
    log.Print("Line 2")

    // Log output as expected?
    chk.Log(
        "Line 1",
        "Line 2",
    )
}

// Failing test.
func Test_FAIL_CaptureLog(t *testing.T) {
    chk := sztest.CaptureLog(t)
    defer chk.Release()

    log.Print("ONLY In Got")
    log.Print("SAME Line 1")
    log.Printf("CHANGED: This is first.")
    log.Print("SAME Line 2")
    log.Printf("CHANGED: This will be second. (Missing in want)")

    // Log output as expected?
    chk.Log(
        "SAME Line 1",
        "CHANGED: (Missing in got) This will be first.",
        "SAME Line 2",
        "CHANGED: This is second.",
        "ONLY in want",
    )
}
```
<!--- gotomd::End::file::./log/example_test.go -->

<!--- gotomd::Bgn::tst::./log/package -->
```bash
go test -v -cover ./log
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLog}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLog&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLog}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:37:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;}}This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}{\color{yellow}{/}}{\color{green}{is}}&#xA0;&#x34F;&#xA0;&#x34F;first.}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;3:2&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{4}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLog&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log/package -->

[Contents](../../README.md#contents)

### Examples: Output Log And Stdout

<!--- gotomd::Bgn::file::./log_and_stdout/example_test.go -->
```bash
cat ./log_and_stdout/example_test.go
```

```go
package example

import (
    "fmt"
    "log"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogAndStderr(t *testing.T) {
    chk := sztest.CaptureLogAndStdout(t)
    defer chk.Release()

    log.Print("logged line 1")
    log.Print("logged line 2")

    fmt.Println("stdout line 1")
    fmt.Println("stdout line 2")

    // Log output as expected?
    chk.Log(
        "logged line 1",
        "logged line 2",
    )

    chk.Stdout(
        "stdout line 1",
        "stdout line 2",
    )
}

// Failing test.
func Test_FAIL_CaptureLogAndStderr(t *testing.T) {
    chk := sztest.CaptureLogAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    log.Print("logged ONLY In Got")
    log.Print("logged SAME Line 1")
    log.Printf("logged CHANGED: This is first.")
    log.Print("logged SAME Line 2")
    log.Printf("logged CHANGED: This will be second. (Missing in want)")

    fmt.Println("this line will be different")

    // Log output as expected?
    chk.Log(
        "logged SAME Line 1",
        "logged CHANGED: (Missing in got) This will be first.",
        "logged SAME Line 2",
        "logged CHANGED: This is second.",
        "logged ONLY in want",
    )

    chk.Stdout(
        "this line will not be the same",
    )
}
```
<!--- gotomd::End::file::./log_and_stdout/example_test.go -->

<!--- gotomd::Bgn::tst::./log_and_stdout/package -->
```bash
go test -v -cover ./log_and_stdout
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:50:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;}}This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}{\color{yellow}{/}}{\color{green}{is}}&#xA0;&#x34F;&#xA0;&#x34F;first.}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;3:2&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{4}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:58:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{not&#xA0;&#x34F;&#xA0;&#x34F;}}be&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{the&#xA0;&#x34F;&#xA0;&#x34F;same}}{\color{yellow}{/}}{\color{green}{different}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stdout&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_and_stdout/package -->

[Contents](../../README.md#contents)

### Examples: Output Log And Stderr

<!--- gotomd::Bgn::file::./log_and_stderr/example_test.go -->
```bash
cat ./log_and_stderr/example_test.go
```

```go
package example

import (
    "fmt"
    "log"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogAndStderr(t *testing.T) {
    chk := sztest.CaptureLogAndStderr(t)
    defer chk.Release()

    log.Print("logged line 1")
    log.Print("logged line 2")

    fmt.Fprintln(os.Stderr, "stderr line 1")
    fmt.Fprintln(os.Stderr, "stderr line 2")

    // Log output as expected?
    chk.Log(
        "logged line 1",
        "logged line 2",
    )

    chk.Stderr(
        "stderr line 1",
        "stderr line 2",
    )
}

// Failing test.
func Test_FAIL_CaptureLogAndStderr(t *testing.T) {
    chk := sztest.CaptureLogAndStderr(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    log.Print("logged ONLY In Got")
    log.Print("logged SAME Line 1")
    log.Printf("logged CHANGED: This is first.")
    log.Print("logged SAME Line 2")
    log.Printf("logged CHANGED: This will be second. (Missing in want)")

    fmt.Fprintln(os.Stderr, "this line will be different")

    // Log output as expected?
    chk.Log(
        "logged SAME Line 1",
        "logged CHANGED: (Missing in got) This will be first.",
        "logged SAME Line 2",
        "logged CHANGED: This is second.",
        "logged ONLY in want",
    )

    chk.Stderr(
        "this line will not be the same",
    )
}
```
<!--- gotomd::End::file::./log_and_stderr/example_test.go -->

<!--- gotomd::Bgn::tst::./log_and_stderr/package -->
```bash
go test -v -cover ./log_and_stderr
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:51:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;}}This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}{\color{yellow}{/}}{\color{green}{is}}&#xA0;&#x34F;&#xA0;&#x34F;first.}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;3:2&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{4}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:59:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{not&#xA0;&#x34F;&#xA0;&#x34F;}}be&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{the&#xA0;&#x34F;&#xA0;&#x34F;same}}{\color{yellow}{/}}{\color{green}{different}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stderr&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_and_stderr/package -->

[Contents](../../README.md#contents)

### Examples: Output Log And Stderr And Stdout

<!--- gotomd::Bgn::file::./log_and_stderr_and_stdout/example_test.go -->
```bash
cat ./log_and_stderr_and_stdout/example_test.go
```

```go
package example

import (
    "fmt"
    "log"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogAndStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureLogAndStderrAndStdout(t)
    defer chk.Release()

    log.Print("logged line 1")
    log.Print("logged line 2")

    fmt.Println("stdout line 1")
    fmt.Println("stdout line 2")

    fmt.Fprintln(os.Stderr, "stderr line 1")
    fmt.Fprintln(os.Stderr, "stderr line 2")

    // Log output as expected?
    chk.Log(
        "logged line 1",
        "logged line 2",
    )

    chk.Stdout(
        "stdout line 1",
        "stdout line 2",
    )

    chk.Stderr(
        "stderr line 1",
        "stderr line 2",
    )
}

// Failing test.
func Test_FAIL_CaptureLogAndStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureLogAndStderrAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    log.Print("logged ONLY In Got")
    log.Print("logged SAME Line 1")
    log.Printf("logged CHANGED: This is first.")
    log.Print("logged SAME Line 2")
    log.Printf("logged CHANGED: This will be second. (Missing in want)")

    fmt.Println("this stdout line will be different")

    fmt.Fprintln(os.Stderr, "this stderr line will be different")

    // Log output as expected?
    chk.Log(
        "logged SAME Line 1",
        "logged CHANGED: (Missing in got) This will be first.",
        "logged SAME Line 2",
        "logged CHANGED: This is second.",
        "logged ONLY in want",
    )

    chk.Stdout(
        "this stdout line will not be the same",
    )

    chk.Stderr(
        "this stderr line will not be the same",
    )
}
```
<!--- gotomd::End::file::./log_and_stderr_and_stdout/example_test.go -->

<!--- gotomd::Bgn::tst::./log_and_stderr_and_stdout/package -->
```bash
go test -v -cover ./log_and_stderr_and_stdout
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:61:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;log&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(5&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;}}This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}{\color{yellow}{/}}{\color{green}{is}}&#xA0;&#x34F;&#xA0;&#x34F;first.}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;3:2&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{4}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{4}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:69:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{not&#xA0;&#x34F;&#xA0;&#x34F;}}be&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{the&#xA0;&#x34F;&#xA0;&#x34F;same}}{\color{yellow}{/}}{\color{green}{different}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:73:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{0}}:{\color{darkturquoise}{0}}&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{not&#xA0;&#x34F;&#xA0;&#x34F;}}be&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{the&#xA0;&#x34F;&#xA0;&#x34F;same}}{\color{yellow}{/}}{\color{green}{different}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogAndStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stderr&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stdout&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_and_stderr_and_stdout/package -->

[Contents](../../README.md#contents)

### Examples: Output Log With Stderr

<!--- gotomd::Bgn::file::./log_with_stderr/example_test.go -->
```bash
cat ./log_with_stderr/example_test.go
```

```go
package example

import (
    "fmt"
    "log"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogWithStderr(t *testing.T) {
    chk := sztest.CaptureLogWithStderr(t)
    defer chk.Release()

    log.Print("logged line 1")

    fmt.Fprintln(os.Stderr, "stderr line 1")

    log.Print("logged line 2")

    fmt.Fprintln(os.Stderr, "stderr line 2")

    // Log output as expected? (either chk.Log or chk.Stderr)
    chk.Log(
        "logged line 1",
        "stderr line 1",
        "logged line 2",
        "stderr line 2",
    )
}

// Failing test.
func Test_FAIL_CaptureLogWithStderr(t *testing.T) {
    chk := sztest.CaptureLogWithStderr(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    log.Print("logged ONLY In Got")
    log.Print("logged SAME Line 1")

    fmt.Fprintln(os.Stderr, "this stderr line will be the same")

    log.Printf("logged CHANGED: This is first.")
    log.Print("logged SAME Line 2")
    log.Printf("logged CHANGED: This will be second. (Missing in want)")

    // Log output as expected? (either chk.Log or chk.Stderr)
    chk.Stderr(
        "logged SAME Line 1",
        "logged CHANGED: (Missing in got) This will be first.",
        "this stderr line will be the same",
        "logged SAME Line 2",
        "logged CHANGED: This is second.",
        "logged ONLY in want",
    )
}
```
<!--- gotomd::End::file::./log_with_stderr/example_test.go -->

<!--- gotomd::Bgn::tst::./log_with_stderr/package -->
```bash
go test -v -cover ./log_with_stderr
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:51:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;logWithStderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(6&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(6&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;be&#xA0;&#x34F;&#xA0;&#x34F;first.}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;be&#xA0;&#x34F;&#xA0;&#x34F;the&#xA0;&#x34F;&#xA0;&#x34F;same}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{3}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;first.}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;4:3&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{5}}:{\color{darkturquoise}{4}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{5}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderr&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;stderr&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_with_stderr/package -->

[Contents](../../README.md#contents)

### Examples: Output Log With Stderr And Stdout

<!--- gotomd::Bgn::file::./log_with_stderr_and_stdout/example_test.go -->
```bash
cat ./log_with_stderr_and_stdout/example_test.go
```

```go
package example

import (
    "fmt"
    "log"
    "os"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogWithStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureLogWithStderrAndStdout(t)
    defer chk.Release()

    log.Print("logged line 1")

    fmt.Fprintln(os.Stderr, "stderr line 1")

    log.Print("logged line 2")

    fmt.Fprintln(os.Stderr, "stderr line 2")

    fmt.Println("stdout line 1")

    // Log output as expected? (either chk.Log or chk.Stderr)
    chk.Log(
        "logged line 1",
        "stderr line 1",
        "logged line 2",
        "stderr line 2",
    )

    chk.Stdout(
        "stdout line 1",
    )
}

// Failing test.
func Test_FAIL_CaptureLogWithStderrAndStdout(t *testing.T) {
    chk := sztest.CaptureLogWithStderrAndStdout(t)
    defer chk.Release()

    chk.FailFast(false) // Process all checks in this test.

    log.Print("logged ONLY In Got")
    log.Print("logged SAME Line 1")

    fmt.Fprintln(os.Stderr, "this stderr line will be the same")

    log.Printf("logged CHANGED: This is first.")
    log.Print("logged SAME Line 2")
    log.Printf("logged CHANGED: This will be second. (Missing in want)")

    // Log output as expected? (either chk.Log or chk.Stderr)
    chk.Stderr(
        "logged SAME Line 1",
        "logged CHANGED: (Missing in got) This will be first.",
        "this stderr line will be the same",
        "logged SAME Line 2",
        "logged CHANGED: This is second.",
        "logged ONLY in want",
    )

    chk.Stdout(
        "stdout line 2",
    )
}
```
<!--- gotomd::End::file::./log_with_stderr_and_stdout/example_test.go -->

<!--- gotomd::Bgn::tst::./log_with_stderr_and_stdout/package -->
```bash
go test -v -cover ./log_with_stderr_and_stdout
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:57:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;logWithStderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(6&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(6&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{0}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;In&#xA0;&#x34F;&#xA0;&#x34F;Got}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:0&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;1}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{1}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;got)&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;be&#xA0;&#x34F;&#xA0;&#x34F;first.}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;this&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;will&#xA0;&#x34F;&#xA0;&#x34F;be&#xA0;&#x34F;&#xA0;&#x34F;the&#xA0;&#x34F;&#xA0;&#x34F;same}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{3}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;is&#xA0;&#x34F;&#xA0;&#x34F;first.}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;4:3&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;SAME&#xA0;&#x34F;&#xA0;&#x34F;Line&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{5}}:{\color{darkturquoise}{4}}&#xA0;&#x34F;&#xA0;&#x34F;logged&#xA0;&#x34F;&#xA0;&#x34F;CHANGED:&#xA0;&#x34F;&#xA0;&#x34F;This&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{is}}{\color{yellow}{/}}{\color{green}{will&#xA0;&#x34F;&#xA0;&#x34F;be}}&#xA0;&#x34F;&#xA0;&#x34F;second.{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;(Missing&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want)}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{5}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{logged&#xA0;&#x34F;&#xA0;&#x34F;ONLY&#xA0;&#x34F;&#xA0;&#x34F;in&#xA0;&#x34F;&#xA0;&#x34F;want}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:66:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(0&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(1&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{0}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{stdout&#xA0;&#x34F;&#xA0;&#x34F;line&#xA0;&#x34F;&#xA0;&#x34F;2}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;FAIL&#xA0;&#x332;&#xA0;&#x332;CaptureLogWithStderrAndStdout&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/output/log&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;stderr&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;stdout&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_with_stderr_and_stdout/package -->

[Contents](../../README.md#contents)
