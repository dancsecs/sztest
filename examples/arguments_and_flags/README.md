<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Arguments And Flags

- [Example: Valid Flag](#example-arguments-and-flags-valid-flag)
- [Example: Invalid Flag](#example-arguments-and-flags-invalid-flag)
- [Example: Invalid Integer](#example-arguments-and-flags-invalid-integer)

[Contents](../../README.md#contents)

## Example: Arguments And Flags: Valid Flag

<!--- gotomd::Bgn::file::./valid_flag/example_test.go -->
```bash
cat ./valid_flag/example_test.go
```

```go
package example

import (
    "flag"
    "testing"

    "github.com/dancsecs/sztest"
)

var processedArgument string

func main() {
    var strValue string
    flag.StringVar(&strValue, "s", "defaultStrValue",
        "usage of default string value",
    )
    flag.Parse()

    processedArgument = "Received: " + strValue
}

// Passing test.
func Test_ArgsAndFlags_SingleGoodFlag(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.SetArgs(
        "program/name",
        "-s",
        "str from arg",
    )

    main()

    chk.Str(processedArgument, "Received: str from arg")
}
```
<!--- gotomd::End::file::./valid_flag/example_test.go -->

<!--- gotomd::Bgn::tst::./valid_flag/package -->
```bash
go test -v -cover ./valid_flag
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;ArgsAndFlags&#x332;SingleGoodFlag}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;ArgsAndFlags&#x332;SingleGoodFlag&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/arguments&#x332;and&#x332;flags/valid&#x332;flag&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./valid_flag/package -->

[Contents](../../README.md#contents)

## Example: Arguments And Flags: Invalid Flag

<!--- gotomd::Bgn::file::./invalid_flag/example_test.go -->
```bash
cat ./invalid_flag/example_test.go
```

```go
package example

import (
    "flag"
    "testing"

    "github.com/dancsecs/sztest"
)

var processedArgument string

func main() {
    var strValue string
    flag.StringVar(&strValue, "s", "defaultStrValue",
        "usage of default string value",
    )
    flag.Parse()

    processedArgument = "Received: " + strValue
}

func Test_ArgsAndFlags_InvalidFlag(t *testing.T) {
    chk := sztest.CaptureStderr(t)
    defer chk.Release()

    chk.SetArgs(
        "program/name",
        "-x",
        "str from arg",
    )

    chk.Panic(
        main,
        "flag provided but not defined: -x",
    )

    chk.Str(processedArgument, "") // Not processed.

    chk.Stderr(
        "flag provided but not defined: -x",
        "Usage of program/name:",
        "\\s -s string", // Note: initial leading space.
        "\\s   \tusage of default string value (default \"defaultStrValue\")",
    )
}
```
<!--- gotomd::End::file::./invalid_flag/example_test.go -->

<!--- gotomd::Bgn::tst::./invalid_flag/package -->
```bash
go test -v -cover ./invalid_flag
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;ArgsAndFlags&#x332;InvalidFlag}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;ArgsAndFlags&#x332;InvalidFlag&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/arguments&#x332;and&#x332;flags/invalid&#x332;flag&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./invalid_flag/package -->

[Contents](../../README.md#contents)

## Example: Arguments And Flags: Invalid Integer

<!--- gotomd::Bgn::file::./invalid_integer/example_test.go -->
```bash
cat ./invalid_integer/example_test.go
```

```go
package example

import (
    "flag"
    "testing"

    "github.com/dancsecs/sztest"
)

func main() {
    var intValue int
    flag.IntVar(&intValue, "n", 10,
        "usage of int value",
    )
    flag.Parse()
}

func Test_ArgsAndFlags_InvalidInteger(t *testing.T) {
    chk := sztest.CaptureStderr(t)
    defer chk.Release()

    chk.SetArgs(
        "program/name",
        "-n",
        "thisIsNotAnInteger",
    )

    chk.Panic(
        main,
        "invalid value \"thisIsNotAnInteger\" for flag -n: parse error",
    )

    chk.Stderr(
        "invalid value \"thisIsNotAnInteger\" for flag -n: parse error",
        "Usage of program/name:",
        "\\s -n int", // Note: initial leading space.
        "\\s   \tusage of int value (default 10)",
    )
}
```
<!--- gotomd::End::file::./invalid_integer/example_test.go -->

<!--- gotomd::Bgn::tst::./invalid_integer/package -->
```bash
go test -v -cover ./invalid_integer
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;ArgsAndFlags&#x332;InvalidInteger}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;ArgsAndFlags&#x332;InvalidInteger&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/arguments&#x332;and&#x332;flags/invalid&#x332;integer&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./invalid_integer/package -->

[Contents](../../README.md#contents)
