<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Slice Examples

- [Examples: String Slice With no Message](#examples-string-slice-with-no-message)
- [Examples: Float64 Slice With Unformatted Message](#examples-float64-slice-with-unformatted-message)
- [Examples: Uint64 Slice With Formatted Message](#examples-uint64-slice-with-formatted-message)

[Contents](../../README.md#contents)

## Examples: String Slice With No Message

This demonstrates that each element in a got slice of string values matches the
wnt slice.

<!--- gotomd::Bgn::file::./string_slice_with_no_message/example_test.go -->
```bash
cat ./string_slice_with_no_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func addPrefix(list []string, prefix string) []string {
    prefixedList := make([]string, len(list))
    for i, entry := range list {
        prefixedList[i] = prefix + entry
    }
    return prefixedList
}

// Passing test.
func Test_PASS_StringSliceWithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testList := []string{"Alpha", "Bravo", "Charlie", "Delta"}

    chk.StrSlice(
        addPrefix(testList, "-->"), // Got.
        []string{
            "-->Alpha",
            "-->Bravo",
            "-->Charlie",
            "-->Delta",
        }, // Want.
    )
}

// Failing test.
func Test_Fail_StringSliceWithNoMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testList := []string{"Alpha", "Bravo", "Charlie", "Delta"}

    chk.StrSlice(
        addPrefix(testList, "-->"), // Got.
        []string{
            "-->Alpha",
            "-->Sheen",
            "-->Delta",
            "-->Echo",
        }, // Want.
    )
}
```
<!--- gotomd::End::file::./string_slice_with_no_message/example_test.go -->

<!--- gotomd::Bgn::tst::./string_slice_with_no_message/package -->
```bash
go test -v -cover ./string_slice_with_no_message
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;StringSliceWithNoMessage}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;StringSliceWithNoMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;StringSliceWithNoMessage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:43:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;[]string:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Length&#xA0;&#x34F;&#xA0;&#x34F;Got:&#xA0;&#x34F;&#xA0;&#x34F;4&#xA0;&#x34F;&#xA0;&#x34F;Wnt:&#xA0;&#x34F;&#xA0;&#x34F;4&#xA0;&#x34F;&#xA0;&#x34F;[}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;-->Alpha}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{1}}:{\color{darkturquoise}{1}}&#xA0;&#x34F;&#xA0;&#x34F;-->{\color{red}{Sheen}}{\color{yellow}{/}}{\color{green}{Bravo}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{2}}:-&#xA0;&#x34F;&#xA0;&#x34F;{\color{green}{-->Charlie}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;3:2&#xA0;&#x34F;&#xA0;&#x34F;-->Delta}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{-->Echo}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;]}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;StringSliceWithNoMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/slice/string&#xA0;&#x332;&#xA0;&#x332;slice&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;no&#xA0;&#x332;&#xA0;&#x332;message&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./string_slice_with_no_message/package -->

[Contents](../../README.md#contents)

## Examples: Float64 Slice With Unformatted Message

This demonstrates that each element in a slice of float64 values is within
the specified tolerance between the got and wnt slice elements.  It uses an
additional message to display the function as it was called should a mismatch
occur.

<!--- gotomd::Bgn::file::./float64_slice_with_unformatted_message/example_test.go -->
```bash
cat ./float64_slice_with_unformatted_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func scale(vector []float64, factor float64) []float64 {
    scaledVector := make([]float64, len(vector))
    for i, v := range vector {
        scaledVector[i] = v * factor
    }
    return scaledVector
}

// Passing test.
func Test_PASS_Float64SliceWithUnformattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testVector := []float64{1, 2, 3, 4, 5, 6}

    chk.Float64Slice(
        []float64{2, 4, 6, 8, 10, 12},               // Got.
        scale(testVector, 2),                        // Want.
        1.0,                                         // Tolerance.
        "function scale(", testVector, ", ", 0, ")", // Additional message.
    )
}

// Failing test.
func Test_Fail_Float64SliceWithUnformattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testVector := []float64{1, 2, 3, 4, 5, 6}

    chk.Float64Slice(
        scale(testVector, 2),                        // Got.
        []float64{2, 4, 6, 8.1, 10, 12},             // Want.
        0.01,                                        // Tolerance.
        "function scale(", testVector, ", ", 0, ")", // Additional message.
    )
}
```
<!--- gotomd::End::file::./float64_slice_with_unformatted_message/example_test.go -->

<!--- gotomd::Bgn::tst::./float64_slice_with_unformatted_message/package -->
```bash
go test -v -cover ./float64_slice_with_unformatted_message
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Float64SliceWithUnformattedMessage}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Float64SliceWithUnformattedMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;Float64SliceWithUnformattedMessage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:40:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;[]float64(+/-&#xA0;&#x34F;&#xA0;&#x34F;0.01):}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\emph{function&#xA0;&#x34F;&#xA0;&#x34F;scale([1&#xA0;&#x34F;&#xA0;&#x34F;2&#xA0;&#x34F;&#xA0;&#x34F;3&#xA0;&#x34F;&#xA0;&#x34F;4&#xA0;&#x34F;&#xA0;&#x34F;5&#xA0;&#x34F;&#xA0;&#x34F;6],&#xA0;&#x34F;&#xA0;&#x34F;0)}}:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Length&#xA0;&#x34F;&#xA0;&#x34F;Got:&#xA0;&#x34F;&#xA0;&#x34F;6&#xA0;&#x34F;&#xA0;&#x34F;Wnt:&#xA0;&#x34F;&#xA0;&#x34F;6&#xA0;&#x34F;&#xA0;&#x34F;[}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:1&#xA0;&#x34F;&#xA0;&#x34F;4}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;6}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{3}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{8.1}}{\color{yellow}{/}}{\color{green}{8}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;4:4&#xA0;&#x34F;&#xA0;&#x34F;10}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;5:5&#xA0;&#x34F;&#xA0;&#x34F;12}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;]}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;Float64SliceWithUnformattedMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/slice/float64&#xA0;&#x332;&#xA0;&#x332;slice&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;unformatted&#xA0;&#x332;&#xA0;&#x332;message&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./float64_slice_with_unformatted_message/package -->

> The failing element of the slice can easily be identified and can be read as
the Wnt needs to have the 8.1 deleted and the 8 inserted as it is beyond the
specified tolerance.

[Contents](../../README.md#contents)

## Examples: Uint64 Slice With Formatted Message

<!--- gotomd::Bgn::file::./uint64_slice_with_formatted_message/example_test.go -->
```bash
cat ./uint64_slice_with_formatted_message/example_test.go
```

```go
package example

import (
    "testing"

    "github.com/dancsecs/sztest"
)

// Example function being tested.
func scale(vector []uint64, factor uint64) []uint64 {
    scaledVector := make([]uint64, len(vector))
    for i, v := range vector {
        scaledVector[i] = v * factor
    }
    return scaledVector
}

// Passing test.
func Test_PASS_Uint64SliceWithFormattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testVector := []uint64{1, 2, 3, 4, 5, 6}

    chk.Uint64Slicef(
        scale(testVector, 2),                   // Got.
        []uint64{2, 4, 6, 8, 10, 12},           // Wnt.
        "function scale(%v,%d)", testVector, 0, // Additional message.
    )
}

// Failing test.
func Test_Fail_Uint64SliceWithFormattedMessage(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    testVector := []uint64{1, 2, 3, 4, 5, 6}

    chk.Uint64Slicef(
        scale(testVector, 2),                   // Got.
        []uint64{2, 4, 6, 9, 10, 12, 14},       // Wnt.
        "function scale(%v,%d)", testVector, 0, // Additional message.
    )
}
```
<!--- gotomd::End::file::./uint64_slice_with_formatted_message/example_test.go -->

<!--- gotomd::Bgn::tst::./uint64_slice_with_formatted_message/package -->
```bash
go test -v -cover ./uint64_slice_with_formatted_message
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Uint64SliceWithFormattedMessage}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;Uint64SliceWithFormattedMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;Uint64SliceWithFormattedMessage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:39:&#xA0;&#x34F;&#xA0;&#x34F;unexpected&#xA0;&#x34F;&#xA0;&#x34F;[]uint64:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\emph{function&#xA0;&#x34F;&#xA0;&#x34F;scale([1&#xA0;&#x34F;&#xA0;&#x34F;2&#xA0;&#x34F;&#xA0;&#x34F;3&#xA0;&#x34F;&#xA0;&#x34F;4&#xA0;&#x34F;&#xA0;&#x34F;5&#xA0;&#x34F;&#xA0;&#x34F;6],0)}}:}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Length&#xA0;&#x34F;&#xA0;&#x34F;Got:&#xA0;&#x34F;&#xA0;&#x34F;6&#xA0;&#x34F;&#xA0;&#x34F;Wnt:&#xA0;&#x34F;&#xA0;&#x34F;7&#xA0;&#x34F;&#xA0;&#x34F;[}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;2}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:1&#xA0;&#x34F;&#xA0;&#x34F;4}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;6}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{3}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{9}}{\color{yellow}{/}}{\color{green}{8}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;4:4&#xA0;&#x34F;&#xA0;&#x34F;10}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;5:5&#xA0;&#x34F;&#xA0;&#x34F;12}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{6}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{14}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;]}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;Uint64SliceWithFormattedMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;[no&#xA0;&#x34F;&#xA0;&#x34F;statements]}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/slice/uint64&#xA0;&#x332;&#xA0;&#x332;slice&#xA0;&#x332;&#xA0;&#x332;with&#xA0;&#x332;&#xA0;&#x332;formatted&#xA0;&#x332;&#xA0;&#x332;message&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./uint64_slice_with_formatted_message/package -->

[Contents](../../README.md#contents)
