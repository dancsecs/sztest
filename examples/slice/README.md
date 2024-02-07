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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;StringSliceWithNoMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;StringSliceWithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;StringSliceWithNoMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:43:\unicode{160}unexpected\unicode{160}[]string:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}Length\unicode{160}Got:\unicode{160}4\unicode{160}Wnt:\unicode{160}4\unicode{160}[}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}-->Alpha}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}1\color{default}:\color{darkturquoise}1\color{default}\unicode{160}-->\color{red}Sheen\color{default}\color{yellow}/\color{default}\color{green}Bravo\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{green}2\color{default}:-\unicode{160}\color{green}-->Charlie\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}3:2\unicode{160}-->Delta}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}3\color{default}\unicode{160}\color{red}-->Echo\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}]}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;StringSliceWithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/slice/string&#x332;slice&#x332;with&#x332;no&#x332;message\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;Float64SliceWithUnformattedMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;Float64SliceWithUnformattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;Float64SliceWithUnformattedMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:40:\unicode{160}unexpected\unicode{160}[]float64(+/-\unicode{160}0.01):}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{function\unicode{160}scale([1\unicode{160}2\unicode{160}3\unicode{160}4\unicode{160}5\unicode{160}6],\unicode{160}0)}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}Length\unicode{160}Got:\unicode{160}6\unicode{160}Wnt:\unicode{160}6\unicode{160}[}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:1\unicode{160}4}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}6}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}\unicode{160}\color{red}8.1\color{default}\color{yellow}/\color{default}\color{green}8\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}4:4\unicode{160}10}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}5:5\unicode{160}12}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}]}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;Float64SliceWithUnformattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/slice/float64&#x332;slice&#x332;with&#x332;unformatted&#x332;message\unicode{160}0.0s}}$
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

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;Uint64SliceWithFormattedMessage}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;Uint64SliceWithFormattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;Uint64SliceWithFormattedMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:39:\unicode{160}unexpected\unicode{160}[]uint64:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\emph{function\unicode{160}scale([1\unicode{160}2\unicode{160}3\unicode{160}4\unicode{160}5\unicode{160}6],0)}:}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}Length\unicode{160}Got:\unicode{160}6\unicode{160}Wnt:\unicode{160}7\unicode{160}[}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}2}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:1\unicode{160}4}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}6}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}\unicode{160}\color{red}9\color{default}\color{yellow}/\color{default}\color{green}8\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}4:4\unicode{160}10}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}5:5\unicode{160}12}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}-:\color{red}6\color{default}\unicode{160}\color{red}14\color{default}}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}]}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;Uint64SliceWithFormattedMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}[no\unicode{160}statements]}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/slice/uint64&#x332;slice&#x332;with&#x332;formatted&#x332;message\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./uint64_slice_with_formatted_message/package -->

[Contents](../../README.md#contents)
