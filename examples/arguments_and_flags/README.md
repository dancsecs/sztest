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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;SingleGoodFlag}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;SingleGoodFlag&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/arguments&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;flags/valid&#xA0;&#x332;&#xA0;&#x332;flag&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;InvalidFlag}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:39:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stderr&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(4&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(4&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;flag&#xA0;&#x34F;&#xA0;&#x34F;provided&#xA0;&#x34F;&#xA0;&#x34F;but&#xA0;&#x34F;&#xA0;&#x34F;not&#xA0;&#x34F;&#xA0;&#x34F;defined:&#xA0;&#x34F;&#xA0;&#x34F;-x}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:1&#xA0;&#x34F;&#xA0;&#x34F;Usage&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;program/name:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{2}}:{\color{darkturquoise}{2}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{\s}}{\color{yellow}{/}}{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;}}&#xA0;&#x34F;&#xA0;&#x34F;-s&#xA0;&#x34F;&#xA0;&#x34F;string}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{3}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{\s}}{\color{yellow}{/}}{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;}}&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;\tusage&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;default&#xA0;&#x34F;&#xA0;&#x34F;string&#xA0;&#x34F;&#xA0;&#x34F;value&#xA0;&#x34F;&#xA0;&#x34F;(default&#xA0;&#x34F;&#xA0;&#x34F;"defaultStrValue")}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;InvalidFlag&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/arguments&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;flags/invalid&#xA0;&#x332;&#xA0;&#x332;flag&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
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
        "  -n int",
        "    \tusage of int value (default 10)",
    )
}
```
<!--- gotomd::End::file::./invalid_integer/example_test.go -->

<!--- gotomd::Bgn::tst::./invalid_integer/package -->
```bash
go test -v -cover ./invalid_integer
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;InvalidInteger}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;ArgsAndFlags&#xA0;&#x332;&#xA0;&#x332;InvalidInteger&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/arguments&#xA0;&#x332;&#xA0;&#x332;and&#xA0;&#x332;&#xA0;&#x332;flags/invalid&#xA0;&#x332;&#xA0;&#x332;integer&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
<!--- gotomd::End::tst::./invalid_integer/package -->

[Contents](../../README.md#contents)
