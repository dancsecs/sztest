<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Bounded Examples

- [Examples: Float64 With no Message](examples/bounded/README.md#examples-float64-with-no-message)

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
    chk.Float64Bounded(fromFloat(valueToTest), sztest.BoundedOpen, -20.0, 20.0)
}

// Failing test.
func Test_Fail_BoundedFloat64WithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    const valueToTest = 1.234567
    chk.Float64Bounded(fromFloat(valueToTest), sztest.BoundedOpen, 0.7, 20.0)
}
```
<!--- gotomd::End::file::./float64_with_no_message/example_test.go -->

<!--- gotomd::Bgn::tst::./float64_with_no_message/package -->
```bash
go test -v -cover ./float64_with_no_message
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;BoundedFloat64WithNoMessage&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:29:&#xa0;unexpected&#xa0;float64:}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{magenta}GOT:&#xa0;\color{default}0.6172835}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{cyan}WNT:&#xa0;\color{default}out&#xa0;of&#xa0;bounds:&#xa0;(0.7,20)&#xa0;-&#xa0;{&#xa0;want&#xa0;|&#xa0;0.7&#xa0;<&#xa0;want&#xa0;<&#xa0;20&#xa0;}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/bounded/float64&#x332;with&#x332;no&#x332;message&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./float64_with_no_message/package -->

[Contents](../../README.md#contents)
