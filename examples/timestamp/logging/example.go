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
