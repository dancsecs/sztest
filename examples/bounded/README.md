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

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:29:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;float64:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{magenta}{GOT:&#xA0;&#x34F;&#xA0;&#x34F;}}0.6172835}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{cyan}{WNT:&#xA0;&#x34F;&#xA0;&#x34F;}}out&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;bounds:&#xA0;&#x34F;&#xA0;&#x34F;(0.7,20)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;{&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;|&#xA0;&#x34F;&#xA0;&#x34F;0.7&#xA0;&#x34F;&#xA0;&#x34F;<&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;<&#xA0;&#x34F;&#xA0;&#x34F;20&#xA0;&#x34F;&#xA0;&#x34F;}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/bounded/float64&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;no&#xA0;&#x332;&#xA0;&#x332;message&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./float64_with_no_message/package -->

[Contents](../../README.md#contents)
