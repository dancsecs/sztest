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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureStdout}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureStdout}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:45:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(3&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(2&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}1\color{default}:-&#xa0;\color{green}Missing&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:1&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/stdout&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureStderr}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureStderr}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:45:&#xa0;Unexpected&#xa0;stderr&#xa0;Entry:&#xa0;got&#xa0;(2&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(3&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}1\color{default}&#xa0;\color{red}Missing&#xa0;in&#xa0;got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:2&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/stderr&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:62:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(3&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(2&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;Stdout:&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}1\color{default}:-&#xa0;\color{green}Stdout:&#xa0;Missing&#xa0;In&#xa0;Want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:1&#xa0;Stdout:&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:68:&#xa0;Unexpected&#xa0;stderr&#xa0;Entry:&#xa0;got&#xa0;(2&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(3&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;Stderr:&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}1\color{default}&#xa0;\color{red}StdErr:&#xa0;Missing&#xa0;in&#xa0;got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:2&#xa0;Stderr:&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/stderr&#x332;and&#x332;stdout&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLog}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLog&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLog}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:37:&#xa0;Unexpected&#xa0;log&#xa0;Entry:&#xa0;got&#xa0;(5&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(5&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}&#xa0;CHANGED:&#xa0;\color{red}(Missing&#xa0;in&#xa0;got)&#xa0;\color{default}This&#xa0;\color{red}will&#xa0;be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}&#xa0;first.}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;3:2&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}4\color{default}&#xa0;\color{red}ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLog&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:50:&#xa0;Unexpected&#xa0;log&#xa0;Entry:&#xa0;got&#xa0;(5&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(5&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}logged&#xa0;ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;\color{red}(Missing&#xa0;in&#xa0;got)&#xa0;\color{default}This&#xa0;\color{red}will&#xa0;be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}&#xa0;first.}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;3:2&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}4\color{default}&#xa0;\color{red}logged&#xa0;ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:58:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;this&#xa0;line&#xa0;will&#xa0;\color{red}not&#xa0;\color{default}be&#xa0;\color{red}the&#xa0;same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stdout&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:51:&#xa0;Unexpected&#xa0;log&#xa0;Entry:&#xa0;got&#xa0;(5&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(5&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}logged&#xa0;ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;\color{red}(Missing&#xa0;in&#xa0;got)&#xa0;\color{default}This&#xa0;\color{red}will&#xa0;be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}&#xa0;first.}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;3:2&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}4\color{default}&#xa0;\color{red}logged&#xa0;ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:59:&#xa0;Unexpected&#xa0;stderr&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;this&#xa0;line&#xa0;will&#xa0;\color{red}not&#xa0;\color{default}be&#xa0;\color{red}the&#xa0;same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stderr&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLogAndStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:61:&#xa0;Unexpected&#xa0;log&#xa0;Entry:&#xa0;got&#xa0;(5&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(5&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}logged&#xa0;ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;\color{red}(Missing&#xa0;in&#xa0;got)&#xa0;\color{default}This&#xa0;\color{red}will&#xa0;be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}&#xa0;first.}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;3:2&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}4\color{default}&#xa0;\color{red}logged&#xa0;ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:69:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;this&#xa0;stdout&#xa0;line&#xa0;will&#xa0;\color{red}not&#xa0;\color{default}be&#xa0;\color{red}the&#xa0;same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:73:&#xa0;Unexpected&#xa0;stderr&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;this&#xa0;stderr&#xa0;line&#xa0;will&#xa0;\color{red}not&#xa0;\color{default}be&#xa0;\color{red}the&#xa0;same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLogAndStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stderr&#x332;and&#x332;stdout&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLogWithStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:51:&#xa0;Unexpected&#xa0;logWithStderr&#xa0;Entry:&#xa0;got&#xa0;(6&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(6&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}logged&#xa0;ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}1\color{default}&#xa0;\color{red}logged&#xa0;CHANGED:&#xa0;(Missing&#xa0;in&#xa0;got)&#xa0;This&#xa0;will&#xa0;be&#xa0;first.\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:2&#xa0;this&#xa0;stderr&#xa0;line&#xa0;will&#xa0;be&#xa0;the&#xa0;same}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}3\color{default}:-&#xa0;\color{green}logged&#xa0;CHANGED:&#xa0;This&#xa0;is&#xa0;first.\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;4:3&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}5\color{default}:\color{darkturquoise}4\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}5\color{default}&#xa0;\color{red}logged&#xa0;ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLogWithStderr&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#x332;with&#x332;stderr&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;CaptureLogWithStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:57:&#xa0;Unexpected&#xa0;logWithStderr&#xa0;Entry:&#xa0;got&#xa0;(6&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(6&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}0\color{default}:-&#xa0;\color{green}logged&#xa0;ONLY&#xa0;In&#xa0;Got\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:0&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;1}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}1\color{default}&#xa0;\color{red}logged&#xa0;CHANGED:&#xa0;(Missing&#xa0;in&#xa0;got)&#xa0;This&#xa0;will&#xa0;be&#xa0;first.\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:2&#xa0;this&#xa0;stderr&#xa0;line&#xa0;will&#xa0;be&#xa0;the&#xa0;same}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{green}3\color{default}:-&#xa0;\color{green}logged&#xa0;CHANGED:&#xa0;This&#xa0;is&#xa0;first.\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;4:3&#xa0;logged&#xa0;SAME&#xa0;Line&#xa0;2}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}5\color{default}:\color{darkturquoise}4\color{default}&#xa0;logged&#xa0;CHANGED:&#xa0;This&#xa0;\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will&#xa0;be\color{default}&#xa0;second.\color{green}&#xa0;(Missing&#xa0;in&#xa0;want)\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}5\color{default}&#xa0;\color{red}logged&#xa0;ONLY&#xa0;in&#xa0;want\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:66:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(0&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;-:\color{red}0\color{default}&#xa0;\color{red}stdout&#xa0;line&#xa0;2\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;CaptureLogWithStderrAndStdout&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/output/log&#x332;with&#x332;stderr&#x332;and&#x332;stdout&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_with_stderr_and_stdout/package -->

[Contents](../../README.md#contents)
