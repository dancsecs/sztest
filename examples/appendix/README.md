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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;GeneralForm}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;GeneralForm&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;GeneralForm}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:52:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}factor&#xa0;(0)&#xa0;out&#xa0;of&#xa0;bounds:&#xa0;"\color{darkturquoise}faili\color{default}ng&#xa0;message"}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}factor&#xa0;(0)&#xa0;out&#xa0;of&#xa0;bounds:&#xa0;"\color{darkturquoise}wro\color{default}ng&#xa0;message"}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:62:&#xa0;unexpected&#xa0;int:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{darkturquoise}4\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}\color{darkturquoise}3\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:63:&#xa0;unexpected&#xa0;string:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\emph{missing&#xa0;"Processed"&#xa0;prefix}:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}\color{green}Processed:&#xa0;\color{default}Hello}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}Hello}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:65:&#xa0;unexpected&#xa0;float64(+/-&#xa0;0.0005):}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\emph{off&#xa0;by&#xa0;-0.000400}:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}0.666\color{darkturquoise}6666666666666\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}0.666\color{darkturquoise}1\color{default}}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:68:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(1&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(1&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;Processing&#xa0;with&#xa0;factor:&#xa0;2&#xa0;and&#xa0;message:&#xa0;\color{red}Processed&#xa0;\color{default}Hello}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:73:&#xa0;Unexpected&#xa0;log&#xa0;Entry:&#xa0;got&#xa0;(3&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(3&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}&#xa0;Entered&#xa0;process(0,&#xa0;"\color{red}wro\color{default}\color{yellow}/\color{default}\color{green}faili\color{default}ng&#xa0;message")}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}1\color{default}:\color{darkturquoise}1\color{default}&#xa0;factor&#xa0;(0)&#xa0;out&#xa0;of&#xa0;bounds:&#xa0;"\color{red}wro\color{default}\color{yellow}/\color{default}\color{green}faili\color{default}ng&#xa0;message"}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}2\color{default}:\color{darkturquoise}2\color{default}&#xa0;Entered&#xa0;process(2,&#xa0;"\color{red}Processed&#xa0;\color{default}Hello")}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;GeneralForm&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/appendix/large&#x332;example&#x332;function&#xa0;0.0s}}$
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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;Main&#x332;No&#x332;Args}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;Main&#x332;No&#x332;Args&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;FAIL&#x332;Main&#x332;No&#x332;Args}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;main&#x332;test.go:82:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}angle&#xa0;required}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}angle\color{red}&#xa0;is\color{default}&#xa0;required}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;main&#x332;test.go:89:&#xa0;unexpected&#xa0;panic:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}invalid&#xa0;angle:&#xa0;not\color{darkturquoise}A\color{default}Number}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}invalid&#xa0;angle:&#xa0;not\color{darkturquoise}&#xa0;A&#xa0;\color{default}Number}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;main&#x332;test.go:111:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(14&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(14&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;00:00&#xa0;Testing&#xa0;angle&#xa0;0&#xa0;no&#xa0;flags}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;01:01&#xa0;0.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;02:02&#xa0;1.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;03:03&#xa0;Testing&#xa0;angle&#xa0;0&#xa0;with&#xa0;verbose&#xa0;flag}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;04:04&#xa0;Report&#xa0;on&#xa0;0.000000&#xa0;degrees&#xa0;(0.000000&#xa0;radians)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;05:05&#xa0;Sin(0.000000°)&#xa0;=&#xa0;0.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;06:06&#xa0;Cos(0.000000°)&#xa0;=&#xa0;1.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}07\color{default}:\color{darkturquoise}07\color{default}&#xa0;Testing&#xa0;angle&#xa0;2P\color{red}I\color{default}\color{yellow}/\color{default}\color{green}i\color{default}&#xa0;with&#xa0;radian&#xa0;flag}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;08:08&#xa0;-0.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;09:09&#xa0;1.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;10:10&#xa0;Testing&#xa0;angle&#xa0;2Pi&#xa0;with&#xa0;radian&#xa0;and&#xa0;verbose&#xa0;flag}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;11:11&#xa0;Report&#xa0;on&#xa0;6.283185&#xa0;radians&#xa0;(360.000000&#xa0;degrees)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;12:12&#xa0;Sin(6.283185)&#xa0;=&#xa0;-0.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;13:13&#xa0;Cos(6.283185)&#xa0;=&#xa0;1.000000}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;main&#x332;test.go:127:&#xa0;Unexpected&#xa0;logWithStderr&#xa0;Entry:&#xa0;got&#xa0;(4&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(4&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;Testing&#xa0;missing&#xa0;angle}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}1\color{default}:\color{darkturquoise}1\color{default}&#xa0;angle\color{red}&#xa0;is\color{default}&#xa0;required}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:2&#xa0;Testing&#xa0;invalid&#xa0;angle}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}&#xa0;invalid&#xa0;angle:&#xa0;not\color{red}&#xa0;A&#xa0;\color{default}\color{yellow}/\color{default}\color{green}A\color{default}Number}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;FAIL&#x332;Main&#x332;No&#x332;Args&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/appendix/large&#x332;example&#x332;main&#x332;function&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./large_example_main_function/package -->

[Contents](../../README.md#contents)
