<!--- gotomd::Auto:: See github.com/dancsecs/gotomd ** DO NOT MODIFY ** -->

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
// Package example demostraits a larger test function.
package example

import (
    "fmt"
    "log"
)

// Process returs known changes to its parameters for testing.
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;GeneralForm}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;GeneralForm\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;GeneralForm}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:52:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}factor\unicode{160}(0)\unicode{160}out\unicode{160}of\unicode{160}bounds:\unicode{160}"\color{darkturquoise}faili\color{default}ng\unicode{160}message"}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}factor\unicode{160}(0)\unicode{160}out\unicode{160}of\unicode{160}bounds:\unicode{160}"\color{darkturquoise}wro\color{default}ng\unicode{160}message"}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:62:\unicode{160}unexpected\unicode{160}int:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}\color{darkturquoise}4\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}\color{darkturquoise}3\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:63:\unicode{160}unexpected\unicode{160}string:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{missing\unicode{160}"Processed"\unicode{160}prefix}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}\color{green}Processed:\unicode{160}\color{default}Hello}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}Hello}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:65:\unicode{160}unexpected\unicode{160}float64(+/-\unicode{160}0.0005):}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{off\unicode{160}by\unicode{160}-0.000400}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}0.666\color{darkturquoise}6666666666666\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}0.666\color{darkturquoise}1\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:68:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(1\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(1\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}Processing\unicode{160}with\unicode{160}factor:\unicode{160}2\unicode{160}and\unicode{160}message:\unicode{160}\color{red}Processed\unicode{160}\color{default}Hello}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:73:\unicode{160}Unexpected\unicode{160}log\unicode{160}Entry:\unicode{160}got\unicode{160}(3\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(3\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}0\color{default}:\color{darkturquoise}0\color{default}\unicode{160}Entered\unicode{160}process(0,\unicode{160}"\color{red}wro\color{default}\color{yellow}/\color{default}\color{green}faili\color{default}ng\unicode{160}message")}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}1\color{default}:\color{darkturquoise}1\color{default}\unicode{160}factor\unicode{160}(0)\unicode{160}out\unicode{160}of\unicode{160}bounds:\unicode{160}"\color{red}wro\color{default}\color{yellow}/\color{default}\color{green}faili\color{default}ng\unicode{160}message"}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}2\color{default}:\color{darkturquoise}2\color{default}\unicode{160}Entered\unicode{160}process(2,\unicode{160}"\color{red}Processed\unicode{160}\color{default}Hello")}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;GeneralForm\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}100.0&#xFE6A;\unicode{160}of\unicode{160}statements}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/appendix/large&#x332;example&#x332;function\unicode{160}0.0s}}$
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
    chk.SetupArgsAndFlags([]string{"progname"})
    chk.Panic(
        main,
        "angle required",
    )

    log.Println("Testing invalid angle")
    chk.SetupArgsAndFlags([]string{"progname", "notANumber"})
    chk.Panic(
        main,
        "invalid angle: notANumber",
    )

    fmt.Println("Testing angle 0 no flags")
    chk.SetupArgsAndFlags([]string{"progname", "0"})
    chk.NoPanic(main)

    fmt.Println("Testing angle 0 with verbose flag")
    chk.SetupArgsAndFlags([]string{"progname", "-v", "0"})
    chk.NoPanic(main)

    twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
    fmt.Println("Testing angle 2Pi with radian flag")
    chk.SetupArgsAndFlags([]string{"progname", "-r", twoPi})
    chk.NoPanic(main)

    fmt.Println("Testing angle 2Pi with radian and verbose flag")
    chk.SetupArgsAndFlags([]string{"progname", "-v", "-r", twoPi})
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
    chk.SetupArgsAndFlags([]string{"progname"})
    chk.Panic(
        main,
        "angle is required",
    )

    log.Println("Testing invalid angle")
    chk.SetupArgsAndFlags([]string{"progname", "notANumber"})
    chk.Panic(
        main,
        "invalid angle: not A Number",
    )

    fmt.Println("Testing angle 0 no flags")
    chk.SetupArgsAndFlags([]string{"progname", "0"})
    chk.NoPanic(main)

    fmt.Println("Testing angle 0 with verbose flag")
    chk.SetupArgsAndFlags([]string{"progname", "-v", "0"})
    chk.NoPanic(main)

    twoPi := strconv.FormatFloat(math.Pi*2, 'f', -1, 64)
    fmt.Println("Testing angle 2Pi with radian flag")
    chk.SetupArgsAndFlags([]string{"progname", "-r", twoPi})
    chk.NoPanic(main)

    fmt.Println("Testing angle 2Pi with radian and verbose flag")
    chk.SetupArgsAndFlags([]string{"progname", "-v", "-r", twoPi})
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;Main&#x332;No&#x332;Args}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;Main&#x332;No&#x332;Args\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;Main&#x332;No&#x332;Args}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}main&#x332;test.go:82:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}angle\unicode{160}required}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}angle\color{red}\unicode{160}is\color{default}\unicode{160}required}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}main&#x332;test.go:89:\unicode{160}unexpected\unicode{160}panic:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}invalid\unicode{160}angle:\unicode{160}not\color{darkturquoise}A\color{default}Number}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}invalid\unicode{160}angle:\unicode{160}not\color{darkturquoise}\unicode{160}A\unicode{160}\color{default}Number}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}main&#x332;test.go:111:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(14\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(14\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}00:00\unicode{160}Testing\unicode{160}angle\unicode{160}0\unicode{160}no\unicode{160}flags}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}01:01\unicode{160}0.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}02:02\unicode{160}1.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}03:03\unicode{160}Testing\unicode{160}angle\unicode{160}0\unicode{160}with\unicode{160}verbose\unicode{160}flag}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}04:04\unicode{160}Report\unicode{160}on\unicode{160}0.000000\unicode{160}degrees\unicode{160}(0.000000\unicode{160}radians)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}05:05\unicode{160}Sin(0.000000°)\unicode{160}=\unicode{160}0.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}06:06\unicode{160}Cos(0.000000°)\unicode{160}=\unicode{160}1.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}07\color{default}:\color{darkturquoise}07\color{default}\unicode{160}Testing\unicode{160}angle\unicode{160}2P\color{red}I\color{default}\color{yellow}/\color{default}\color{green}i\color{default}\unicode{160}with\unicode{160}radian\unicode{160}flag}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}08:08\unicode{160}-0.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}09:09\unicode{160}1.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}10:10\unicode{160}Testing\unicode{160}angle\unicode{160}2Pi\unicode{160}with\unicode{160}radian\unicode{160}and\unicode{160}verbose\unicode{160}flag}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}11:11\unicode{160}Report\unicode{160}on\unicode{160}6.283185\unicode{160}radians\unicode{160}(360.000000\unicode{160}degrees)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}12:12\unicode{160}Sin(6.283185)\unicode{160}=\unicode{160}-0.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}13:13\unicode{160}Cos(6.283185)\unicode{160}=\unicode{160}1.000000}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}main&#x332;test.go:127:\unicode{160}Unexpected\unicode{160}logWithStderr\unicode{160}Entry:\unicode{160}got\unicode{160}(4\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(4\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}Testing\unicode{160}missing\unicode{160}angle}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}1\color{default}:\color{darkturquoise}1\color{default}\unicode{160}angle\color{red}\unicode{160}is\color{default}\unicode{160}required}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}Testing\unicode{160}invalid\unicode{160}angle}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}\unicode{160}invalid\unicode{160}angle:\unicode{160}not\color{red}\unicode{160}A\unicode{160}\color{default}\color{yellow}/\color{default}\color{green}A\color{default}Number}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;Main&#x332;No&#x332;Args\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}100.0&#xFE6A;\unicode{160}of\unicode{160}statements}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/appendix/large&#x332;example&#x332;main&#x332;function\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./large_example_main_function/package -->

[Contents](../../README.md#contents)
