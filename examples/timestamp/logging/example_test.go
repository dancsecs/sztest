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
