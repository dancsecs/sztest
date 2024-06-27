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

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;TimestampLogging}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;TimestampLogging&#xa0;(0.0s)}}$
<br>
$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;example&#x332;test.go:76:&#xa0;Unexpected&#xa0;stdout&#xa0;Entry:&#xa0;got&#xa0;(7&#xa0;lines)&#xa0;-&#xa0;want&#xa0;(7&#xa0;lines)}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;0:0&#xa0;19990707080910&#xa0;-&#xa0;10}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;1:1&#xa0;19990707080919&#xa0;-&#xa0;20}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;2:2&#xa0;19990707080926&#xa0;-&#xa0;30}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}3\color{default}:\color{darkturquoise}3\color{default}&#xa0;199907070809\color{red}26\color{default}\color{yellow}/\color{default}\color{green}37\color{default}&#xa0;-&#xa0;40&#xa0;DELAYED}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;4:4&#xa0;19990707080946&#xa0;-&#xa0;50}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;5:5&#xa0;19990707080953&#xa0;-&#xa0;60}}$
<br>
$\small{\texttt{&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;\color{darkturquoise}6\color{default}:\color{darkturquoise}6\color{default}&#xa0;19990707081004&#xa0;-&#xa0;70\color{green}&#xa0;DELAYED\color{default}}}$
<br>
$\small{\texttt{---&#xa0;FAIL:&#xa0;Test&#x332;Fail&#x332;BoundedFloat64WithNoMessage&#xa0;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xa0;100.0&#xFE6A;&#xa0;of&#xa0;statements}}$
<br>
$\small{\texttt{FAIL&#xa0;github.com/dancsecs/sztest/examples/timestamp/logging&#xa0;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./logging/package -->

[Contents](../../README.md#contents)
