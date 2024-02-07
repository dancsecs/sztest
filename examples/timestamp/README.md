<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Timestamp Examples

- [Examples: Logging](examples/timestamp/README.md#examples-logging)

[Contents](../../README.md#contents)

## Examples: Logging

This demonstrates using the internal clock functions to test logging with
timestamps.

<!--- gotomd::Bgn::file::./logging/example.go example_test.go -->
```bash
cat ./logging/example.go
```

```go
/*
Package example provides an example of using the sztest Clock utility to test
code that uses relative timestamps.  It simulates logging values received
on a channel and issuing warnings if the elapsed time between received
values exceeds a specified duration.
*/
//nolint:goCheckNoGlobals // Ok.
package example

import (
    "fmt"
    "strconv"
    "time"
)

// Now provides a link to get the current time that can be replaced by
// (*Chk).ClockNext to facilitate testing date related code.
var now = time.Now

func writeSample(ts time.Time, v int64, late bool) {
    const base10 = 10
    l := ts.Format("20060102150405") + " - " + strconv.FormatInt(v, base10)
    if late {
        l += " DELAYED"
    }
    fmt.Println(l)
}

// LogValues logs the values feed to it providing a warning if the
// duration between samples exceeds the provided duration.
func LogValues(warnDelay time.Duration) chan<- int64 {
    var ch = make(chan int64)
    go func() {
        v := <-ch
        last := now()
        writeSample(last, v, false)
        for {
            v := <-ch
            ts := now()
            writeSample(ts, v, ts.Sub(last) >= warnDelay)
            last = ts
        }
    }()
    return ch
}
```

```bash
cat ./logging/example_test.go
```

```go
package example

import (
    "testing"
    "time"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TimestampLogging(t *testing.T) {
    chk := sztest.CaptureStdout(t)
    defer chk.Release()

    origNow := now
    defer func() {
        now = origNow
    }()
    chk.ClockSetSub(sztest.ClockSubTS)
    chk.ClockSet(
        time.Date(1999, time.July, 7, 8, 9, 10, 0, time.Local),
        time.Second*9, time.Second*7, time.Second*11,
    )
    now = chk.ClockNext

    ch := LogValues(time.Second * 10)

    ch <- 10
    ch <- 20
    ch <- 30
    ch <- 40
    ch <- 50
    ch <- 60
    ch <- 70

    chk.Stdout(
        "" +
            "{{clkTS0}} - 10\n" +
            "{{clkTS1}} - 20\n" +
            "{{clkTS2}} - 30\n" +
            "{{clkTS3}} - 40 DELAYED\n" +
            "{{clkTS4}} - 50\n" +
            "{{clkTS5}} - 60\n" +
            "{{clkTS6}} - 70 DELAYED\n" +
            "",
    )
}

// Failing test.
func Test_Fail_BoundedFloat64WithNoMessage(t *testing.T) {
    chk := sztest.CaptureStdout(t)
    defer chk.Release()

    chk.FailFast(false)
    origNow := now
    defer func() {
        now = origNow
    }()
    chk.ClockSetSub(sztest.ClockSubTS)
    chk.ClockSet(
        time.Date(1999, time.July, 7, 8, 9, 10, 0, time.Local),
        time.Second*9, time.Second*7, time.Second*11,
    )
    now = chk.ClockNext

    ch := LogValues(time.Second * 10)

    ch <- 10
    ch <- 20
    ch <- 30
    ch <- 40
    ch <- 50
    ch <- 60
    ch <- 70

    chk.Stdout(
        "" +
            "{{clkTS0}} - 10\n" +
            "{{clkTS1}} - 20\n" +
            "{{clkTS2}} - 30\n" +
            "{{clkTS2}} - 40 DELAYED\n" + // TS incorrectly repeated.
            "{{clkTS4}} - 50\n" +
            "{{clkTS5}} - 60\n" +
            "{{clkTS6}} - 70\n" + // Missing DELAYED flag.
            "",
    )
}
```
<!--- gotomd::End::file::./logging/example.go example_test.go -->

<!--- gotomd::Bgn::tst::./logging/package -->
```bash
go test -v -cover ./logging
```

$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;PASS&#x332;TimestampLogging}}$
<br>
$\small{\texttt{---\unicode{160}PASS:\unicode{160}Test&#x332;PASS&#x332;TimestampLogging\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{===\unicode{160}RUN\unicode{160}\unicode{160}\unicode{160}Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}example&#x332;test.go:76:\unicode{160}Unexpected\unicode{160}stdout\unicode{160}Entry:\unicode{160}got\unicode{160}(7\unicode{160}lines)\unicode{160}-\unicode{160}want\unicode{160}(7\unicode{160}lines)}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}0:0\unicode{160}19990707080910\unicode{160}-\unicode{160}10}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}1:1\unicode{160}19990707080919\unicode{160}-\unicode{160}20}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}2:2\unicode{160}19990707080926\unicode{160}-\unicode{160}30}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}\unicode{160}199907070809\color{red}26\color{default}\color{yellow}/\color{default}\color{green}37\color{default}\unicode{160}-\unicode{160}40\unicode{160}DELAYED}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}4:4\unicode{160}19990707080946\unicode{160}-\unicode{160}50}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}5:5\unicode{160}19990707080953\unicode{160}-\unicode{160}60}}$
<br>
$\small{\texttt{\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\unicode{160}\color{darkturquoise}6\color{default}:\color{darkturquoise}6\color{default}\unicode{160}19990707081004\unicode{160}-\unicode{160}70\color{green}\unicode{160}DELAYED\color{default}}}$
<br>
$\small{\texttt{---\unicode{160}FAIL:\unicode{160}Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage\unicode{160}(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:\unicode{160}100.0&#xFE6A;\unicode{160}of\unicode{160}statements}}$
<br>
$\small{\texttt{FAIL\unicode{160}github.com/dancsecs/sztest/examples/timestamp/logging\unicode{160}0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./logging/package -->

[Contents](../../README.md#contents)
