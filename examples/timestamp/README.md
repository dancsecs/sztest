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
    chk.ClockSetSub(sztest.ClkFmtTS)
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
            "{{clkTS6}} - 70 DELAYED",
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
    chk.ClockSetSub(sztest.ClkFmtTS)
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
            "", // Extra Line.
    )
}
```
<!--- gotomd::End::file::./logging/example.go example_test.go -->

<!--- gotomd::Bgn::tst::./logging/package -->
```bash
go test -v -cover ./logging
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;TimestampLogging}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;PASS&#xA0;&#x332;&#xA0;&#x332;TimestampLogging&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;example&#xA0;&#x332;&#xA0;&#x332;test.go:75:&#xA0;&#x34F;&#xA0;&#x34F;Unexpected&#xA0;&#x34F;&#xA0;&#x34F;stdout&#xA0;&#x34F;&#xA0;&#x34F;Entry:&#xA0;&#x34F;&#xA0;&#x34F;got&#xA0;&#x34F;&#xA0;&#x34F;(7&#xA0;&#x34F;&#xA0;&#x34F;lines)&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;want&#xA0;&#x34F;&#xA0;&#x34F;(8&#xA0;&#x34F;&#xA0;&#x34F;lines)}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;0:0&#xA0;&#x34F;&#xA0;&#x34F;19990707080910&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;10}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;1:1&#xA0;&#x34F;&#xA0;&#x34F;19990707080919&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;20}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;2:2&#xA0;&#x34F;&#xA0;&#x34F;19990707080926&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;30}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{3}}:{\color{darkturquoise}{3}}&#xA0;&#x34F;&#xA0;&#x34F;199907070809{\color{red}{26}}{\color{yellow}{/}}{\color{green}{37}}&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;40&#xA0;&#x34F;&#xA0;&#x34F;DELAYED}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;4:4&#xA0;&#x34F;&#xA0;&#x34F;19990707080946&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;50}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;5:5&#xA0;&#x34F;&#xA0;&#x34F;19990707080953&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;60}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;{\color{darkturquoise}{6}}:{\color{darkturquoise}{6}}&#xA0;&#x34F;&#xA0;&#x34F;19990707081004&#xA0;&#x34F;&#xA0;&#x34F;-&#xA0;&#x34F;&#xA0;&#x34F;70{\color{green}{&#xA0;&#x34F;&#xA0;&#x34F;DELAYED}}}}$
<br>
$\small{\texttt{&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;-:{\color{red}{7}}&#xA0;&#x34F;&#xA0;&#x34F;{\color{red}{}}}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;FAIL:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;Fail&#xA0;&#x332;&#xA0;&#x332;BoundedFloat64WithNoMessage&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{FAIL}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{FAIL&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/timestamp/logging&#xA0;&#x34F;&#xA0;&#x34F;0.0s}}$
<br>
$\small{\texttt{FAIL}}$
<br>
<!--- gotomd::End::tst::./logging/package -->

[Contents](../../README.md#contents)
