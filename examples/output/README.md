<!--- gotomd::Auto:: See github.com/dancsecs/gotomd ** DO NOT MODIFY ** -->

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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureStdout}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureStdout}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:45:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(3\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(2\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}1\color{default}:-\unicode{160}\color{green}Missing\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:1\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/stdout\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureStderr}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureStderr}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:45:\unicode{160}Unexpected\unicode{160}stderr\unicode{160}Entry:\unicode{160}got\unicode{160}(2\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(3\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}1\color{default}\unicode{160}\color{red}Missing\unicode{160}in\unicode{160}got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:2\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/stderr\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureStderrAndStdout}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:62:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(3\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(2\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}Stdout:\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}1\color{default}:-\unicode{160}\color{green}Stdout:\unicode{160}Missing\unicode{160}In\unicode{160}Want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:1\unicode{160}Stdout:\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:68:\unicode{160}Unexpected\unicode{160}stderr\unicode{160}Entry:\unicode{160}got\unicode{160}(2\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(3\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}Stderr:\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}1\color{default}\unicode{160}\color{red}StdErr:\unicode{160}Missing\unicode{160}in\unicode{160}got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:2\unicode{160}Stderr:\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/stderr&#x332;and&#x332;stdout\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLog}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLog\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLog}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:37:\unicode{160}Unexpected\unicode{160}log\unicode{160}Entry:\unicode{160}got\unicode{160}(5\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(5\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}\unicode{160}CHANGED:\unicode{160}\color{red}(Missing\unicode{160}in\unicode{160}got)\unicode{160}\color{default}This\unicode{160}\color{red}will\unicode{160}be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}\unicode{160}first.}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}3:2\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}4\color{default}\unicode{160}\color{red}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLog\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:50:\unicode{160}Unexpected\unicode{160}log\unicode{160}Entry:\unicode{160}got\unicode{160}(5\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(5\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}logged\unicode{160}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}\color{red}(Missing\unicode{160}in\unicode{160}got)\unicode{160}\color{default}This\unicode{160}\color{red}will\unicode{160}be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}\unicode{160}first.}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}3:2\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}4\color{default}\unicode{160}\color{red}logged\unicode{160}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:58:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(1\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}this\unicode{160}line\unicode{160}will\unicode{160}\color{red}not\unicode{160}\color{default}be\unicode{160}\color{red}the\unicode{160}same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stdout\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderr}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:51:\unicode{160}Unexpected\unicode{160}log\unicode{160}Entry:\unicode{160}got\unicode{160}(5\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(5\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}logged\unicode{160}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}\color{red}(Missing\unicode{160}in\unicode{160}got)\unicode{160}\color{default}This\unicode{160}\color{red}will\unicode{160}be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}\unicode{160}first.}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}3:2\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}4\color{default}\unicode{160}\color{red}logged\unicode{160}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:59:\unicode{160}Unexpected\unicode{160}stderr\unicode{160}Entry:\unicode{160}got\unicode{160}(1\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}this\unicode{160}line\unicode{160}will\unicode{160}\color{red}not\unicode{160}\color{default}be\unicode{160}\color{red}the\unicode{160}same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stderr\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLogAndStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderrAndStdout}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:61:\unicode{160}Unexpected\unicode{160}log\unicode{160}Entry:\unicode{160}got\unicode{160}(5\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(5\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}logged\unicode{160}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}2\color{default}:\color{darkturquoise}1\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}\color{red}(Missing\unicode{160}in\unicode{160}got)\unicode{160}\color{default}This\unicode{160}\color{red}will\unicode{160}be\color{default}\color{yellow}/\color{default}\color{green}is\color{default}\unicode{160}first.}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}3:2\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}4\color{default}:\color{darkturquoise}3\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}4\color{default}\unicode{160}\color{red}logged\unicode{160}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:69:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(1\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}this\unicode{160}stdout\unicode{160}line\unicode{160}will\unicode{160}\color{red}not\unicode{160}\color{default}be\unicode{160}\color{red}the\unicode{160}same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:73:\unicode{160}Unexpected\unicode{160}stderr\unicode{160}Entry:\unicode{160}got\unicode{160}(1\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}this\unicode{160}stderr\unicode{160}line\unicode{160}will\unicode{160}\color{red}not\unicode{160}\color{default}be\unicode{160}\color{red}the\unicode{160}same\color{default}\color{yellow}/\color{default}\color{green}different\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLogAndStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log&#x332;and&#x332;stderr&#x332;and&#x332;stdout\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLogWithStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLogWithStderr}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:51:\unicode{160}Unexpected\unicode{160}logWithStderr\unicode{160}Entry:\unicode{160}got\unicode{160}(6\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(6\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}logged\unicode{160}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}1\color{default}\unicode{160}\color{red}logged\unicode{160}CHANGED:\unicode{160}(Missing\unicode{160}in\unicode{160}got)\unicode{160}This\unicode{160}will\unicode{160}be\unicode{160}first.\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}this\unicode{160}stderr\unicode{160}line\unicode{160}will\unicode{160}be\unicode{160}the\unicode{160}same}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}3\color{default}:-\unicode{160}\color{green}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}is\unicode{160}first.\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}4:3\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}5\color{default}:\color{darkturquoise}4\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}5\color{default}\unicode{160}\color{red}logged\unicode{160}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLogWithStderr\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log&#x332;with&#x332;stderr\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;CaptureLogWithStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;CaptureLogWithStderrAndStdout}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:57:\unicode{160}Unexpected\unicode{160}logWithStderr\unicode{160}Entry:\unicode{160}got\unicode{160}(6\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(6\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}0\color{default}:-\unicode{160}\color{green}logged\unicode{160}ONLY\unicode{160}In\unicode{160}Got\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:0\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}1}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}1\color{default}\unicode{160}\color{red}logged\unicode{160}CHANGED:\unicode{160}(Missing\unicode{160}in\unicode{160}got)\unicode{160}This\unicode{160}will\unicode{160}be\unicode{160}first.\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}this\unicode{160}stderr\unicode{160}line\unicode{160}will\unicode{160}be\unicode{160}the\unicode{160}same}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}3\color{default}:-\unicode{160}\color{green}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}is\unicode{160}first.\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}4:3\unicode{160}logged\unicode{160}SAME\unicode{160}Line\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}5\color{default}:\color{darkturquoise}4\color{default}\unicode{160}logged\unicode{160}CHANGED:\unicode{160}This\unicode{160}\color{red}is\color{default}\color{yellow}/\color{default}\color{green}will\unicode{160}be\color{default}\unicode{160}second.\color{green}\unicode{160}(Missing\unicode{160}in\unicode{160}want)\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}5\color{default}\unicode{160}\color{red}logged\unicode{160}ONLY\unicode{160}in\unicode{160}want\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:66:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(0\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}0\color{default}\unicode{160}\color{red}stdout\unicode{160}line\unicode{160}2\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;CaptureLogWithStderrAndStdout\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/output/log&#x332;with&#x332;stderr&#x332;and&#x332;stdout\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./log_with_stderr_and_stdout/package -->

[Contents](../../README.md#contents)
