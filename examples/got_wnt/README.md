<!--- gotomd::Auto:: See github.com/dancsecs/gotomd ** DO NOT MODIFY ** -->

# Got / Want Examples

- [Examples: Integer With No Message](#examples-integer-with-no-message)
- [Examples: Float32 With Unformatted Message](#examples-float32-with-unformatted-message)
- [Examples: String With Formatted Message](#examples-string-with-formatted-message)

[Contents](../../README.md#contents)

## Examples: Integer With No Message

This is the simplest form of a builtin Got/Wnt test.  Just comparing the wanted
value with the gotten value and an error registered if they are not equal.
No additional context information is provided.

<!--- gotomd::Bgn::file::./integer_with_no_message/example_test.go -->
```bash
cat ./integer_with_no_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func addIntegers(int1, int2 int) int {
    return int1 + int2
}

// Passing test.
func Test_PASS_IntegerWithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Int(
        addIntegers(2, 3), // Got.
        5,                 // Want.
    )
}

// Failing test.
func Test_FAIL_IntegerWithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Int(
        addIntegers(1237456, 1000), // Got.
        1237456,                    // Want.
    )
}
```
<!--- gotomd::End::file::./integer_with_no_message/example_test.go -->

<!--- gotomd::Bgn::tst::./integer_with_no_message/package -->
```bash
go test -v -cover ./integer_with_no_message
```

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;IntegerWithNoMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;IntegerWithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;IntegerWithNoMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:30:\unicode{160}unexpected\unicode{160}int:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}123\color{darkturquoise}8\color{default}456}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}123\color{darkturquoise}7\color{default}456}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;IntegerWithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/got&#x332;wnt/integer&#x332;with&#x332;no&#x332;message\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./integer_with_no_message/package -->

> Here the failing test has its fourth number highlighted as changed while
the passing test produced no output.

[Contents](../../README.md#contents)

## Examples: Float32 With Unformatted Message

> This example shows the float Got/Wnt test with an unformatted message.  Due to
the nature of floats it is the only builtin type check that includes a
tolerance factor.  If the absolute value of the difference between the got and
want values is less than the tolerance then the two floats will be considered
equivalent.

<!--- gotomd::Bgn::file::./float32_with_unformatted_message/example_test.go -->
```bash
cat ./float32_with_unformatted_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func fromFloat(f float32) float32 {
    return f / 2.0
}

// Passing test.
func Test_PASS_Float32WithUnformattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    const valueToTest = 1.234567
    chk.Float32(
        fromFloat(valueToTest),                  // Got.
        0.617,                                   // Want.
        0.001,                                   // Tolerance.
        "function fromFloat(", valueToTest, ")", // Additional message.
    )
}

// Failing test.
func Test_Fail_Float32WithUnformattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    const valueToTest = 2.468024
    chk.Float32(
        fromFloat(valueToTest),                  // Got.
        1.2356,                                  // Want.
        0.0005,                                  // Tolerance.
        "function fromFloat(", valueToTest, ")", // Additional message.
    )
}
```
<!--- gotomd::End::file::./float32_with_unformatted_message/example_test.go -->

<!--- gotomd::Bgn::tst::./float32_with_unformatted_message/package -->
```bash
go test -v -cover ./float32_with_unformatted_message
```

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;Float32WithUnformattedMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;Float32WithUnformattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;Float32WithUnformattedMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:34:\unicode{160}unexpected\unicode{160}float32(+/-\unicode{160}0.0005000000237487257):}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{function\unicode{160}fromFloat(2.468024)}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}1.23\color{darkturquoise}4012\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}1.23\color{darkturquoise}56\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;Float32WithUnformattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/got&#x332;wnt/float32&#x332;with&#x332;unformatted&#x332;message\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./float32_with_unformatted_message/package -->

> Here the thousandth fractional position is flagged as being out of tolerance
and is highlighted as changed/different while the specific
tolerance value used has been added to the type name.  Finally the
additional unformatted message is displayed just before the GOT: line.

[Contents](../../README.md#contents)

## Examples: String With Formatted Message

This example shows a string Got/Wnt test with a formatted message.

<!--- gotomd::Bgn::file::./string_with_formatted_message/example_test.go -->
```bash
cat ./string_with_formatted_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_StringWithFormattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Strf(
        "This string being tested.",
        "This string being tested.",
        // Optional formatted message.
        "%s message with %s information", "Formatted", "additional",
    )
}

// Failing test.
func Test_FAIL_StringWithFormattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.Strf(
        "This (extra) is the Got string being tested.",
        "This is the Wnt string (missing) being tested.",
        // Optional formatted message.
        "%s message with %s information", "Formatted", "additional",
    )
}
```
<!--- gotomd::End::file::./string_with_formatted_message/example_test.go -->

<!--- gotomd::Bgn::tst::./string_with_formatted_message/package -->
```bash
go test -v -cover ./string_with_formatted_message
```

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;StringWithFormattedMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;StringWithFormattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;FAIL&#x332;StringWithFormattedMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:27:\unicode{160}unexpected\unicode{160}string:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{Formatted\unicode{160}message\unicode{160}with\unicode{160}additional\unicode{160}information}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}This\color{green}\unicode{160}(extra)\color{default}\unicode{160}is\unicode{160}the\unicode{160}\color{darkturquoise}Go\color{default}t\unicode{160}string\unicode{160}being\unicode{160}tested.}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}This\unicode{160}is\unicode{160}the\unicode{160}\color{darkturquoise}Wn\color{default}t\unicode{160}string\color{red}\unicode{160}(missing)\color{default}\unicode{160}being\unicode{160}tested.}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;FAIL&#x332;StringWithFormattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/got&#x332;wnt/string&#x332;with&#x332;formatted&#x332;message\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./string_with_formatted_message/package -->

> Here the got string has extra information (extra) not found in the want
string while the want string has missing information (missing) not found
in the got string.  Then there is a changed area area between the got Go and
the want Wnt.  Finally the additional formatted message is displayed just
before the GOT: line.

[Contents](../../README.md#contents)
