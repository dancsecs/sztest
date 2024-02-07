<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Unbounded Examples

- [Examples: Float64 With no Message](examples/unbounded/README.md#examples-float64-with-no-message)

[Contents](../../README.md#contents)

## Examples: Float64 With No Message

This demonstrates that a float value is within the specified bounds.

<!--- gotomd::Bgn::file::./float64_with_no_message/example_test.go -->
```bash
cat ./float64_with_no_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func fromFloat(f float64) float64 {
    return f / 2.0
}

// Passing test.
func Test_PASS_BoundedFloat64WithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    const valueToTest = 1.234567
    chk.Float64Unbounded(fromFloat(valueToTest), sztest.UnboundedMinOpen, -2.0)
}

// Failing test.
func Test_Fail_BoundedFloat64WithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    const valueToTest = 1.234567
    chk.Float64Unbounded(fromFloat(valueToTest), sztest.UnboundedMinOpen, 2.0)
}
```
<!--- gotomd::End::file::./float64_with_no_message/example_test.go -->

<!--- gotomd::Bgn::tst::./float64_with_no_message/package -->
```bash
go test -v -cover ./float64_with_no_message
```

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;BoundedFloat64WithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:29:\unicode{160}unexpected\unicode{160}float64:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{magenta}GOT:\unicode{160}\color{default}0.6172835}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{cyan}WNT:\unicode{160}\color{default}out\unicode{160}of\unicode{160}bounds:\unicode{160}(2,MAX)\unicode{160}-\unicode{160}{\unicode{160}want\unicode{160}|\unicode{160}want\unicode{160}>\unicode{160}2\unicode{160}}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/unbounded/float64&#x332;with&#x332;no&#x332;message\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./float64_with_no_message/package -->

[Contents](../../README.md#contents)
