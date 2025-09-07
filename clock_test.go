/*
   Golang test helper library: sztest.
   Copyright (C) 2023, 2024 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sztest

import (
	"testing"
	"time"
)

func tstChkClock(t *testing.T) {
	t.Run("defaultIncrement", chkClockDefaultIncrement)
	t.Run("useCurrentTime", chkClockCurrentTime)
	t.Run("useInvalidTick", chkClockInvalidTick)
	t.Run("lastFormat", chkClockCLockLastFormat)
	t.Run("nextFormat", chkClockCLockNextFormat)
	t.Run("useCase1", chkClockUseCase1)
	t.Run("useCase2", chkClockUseCase2)
	t.Run("useCase3", chkClockUseCase3)
	t.Run("subs", chkClockSubs)
	t.Run("offset", chkClockOffset)
}

func chkClockDefaultIncrement(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	c := newTstClock(time.Now(), nil)
	chk.DurSlice(c.inc, []time.Duration{time.Millisecond})
}

func chkClockCurrentTime(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	ts0 := chk.ClockNext()

	time.Sleep(time.Nanosecond)

	ts1 := chk.ClockNext()

	chk.True(ts0.Equal(chk.ClockTick(0)))
	chk.True(ts1.Equal(chk.ClockTick(1)))
	chk.False(ts0.Equal(ts1))
}

func chkClockInvalidTick(t *testing.T) {
	chk := CaptureLog(t)
	defer chk.Release()

	chk.Panic(
		func() {
			chk.ClockTick(-1)
		},
		"unknown tick index: -1",
	)

	chk.Panic(
		func() {
			chk.ClockTick(0)
		},
		"unknown tick index: 0",
	)

	chk.Panic(
		func() {
			chk.ClockTick(1)
		},
		"unknown tick index: 1",
	)

	chk.Log(
		"unknown tick index: -1",
		"unknown tick index: 0",
		"unknown tick index: 1",
	)
}

func chkClockCLockLastFormat(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	nextTS := chk.ClockNext()

	chk.Str(chk.ClockLastFmtTime(), nextTS.Format("150405"))
	chk.Str(chk.ClockLastFmtDate(), nextTS.Format("20060102"))
	chk.Str(chk.ClockLastFmtTS(), nextTS.Format("20060102150405"))
	chk.Str(chk.ClockLastFmtNano(), nextTS.Format("20060102150405.000000000"))

	chk.ClockSetCusA(time.RFC822Z)
	chk.Str(chk.ClockLastFmtCusA(), nextTS.Format(time.RFC822Z))

	chk.ClockSetCusB(time.RFC3339Nano)
	chk.Str(chk.ClockLastFmtCusB(), nextTS.Format(time.RFC3339Nano))

	chk.ClockSetCusC(time.ANSIC)
	chk.Str(chk.ClockLastFmtCusC(), nextTS.Format(time.ANSIC))
}

func chkClockCLockNextFormat(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.Str(chk.ClockNextFmtTime(), chk.ClockLastFmtTime())
	chk.Str(chk.ClockNextFmtDate(), chk.ClockLastFmtDate())
	chk.Str(chk.ClockNextFmtTS(), chk.ClockLastFmtTS())
	chk.Str(chk.ClockNextFmtNano(), chk.ClockLastFmtNano())

	chk.ClockSetCusA(time.RFC822Z)
	chk.Str(chk.ClockNextFmtCusA(), chk.ClockLastFmtCusA())

	chk.ClockSetCusB(time.RFC3339Nano)
	chk.Str(chk.ClockNextFmtCusB(), chk.ClockLastFmtCusB())

	chk.ClockSetCusC(time.ANSIC)
	chk.Str(chk.ClockNextFmtCusC(), chk.ClockLastFmtCusC())
}

func chkClockUseCase1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	tsBeforeCurrent := chk.ClockLast()
	tsBefore := chk.ClockNext()

	chk.True(tsBeforeCurrent.Before(tsBefore))

	resetAllFunc := chk.ClockSet(
		time.Date(2999, 12, 25, 13, 15, 45, 555555555, time.Local),
		time.Microsecond, time.Microsecond*10,
	)

	ts1 := chk.ClockNext()
	ts2 := chk.ClockNext()
	ts3 := chk.ClockNext()

	resetDayFunc := chk.ClockOffsetDay(-2, time.Microsecond*100)

	tsd1 := chk.ClockNext()
	tsd2 := chk.ClockNext()
	tsd3 := chk.ClockNext()
	tsd3c := chk.ClockLast()

	resetDayFunc()

	ts4 := chk.ClockNext()
	ts5 := chk.ClockNext()
	ts6 := chk.ClockNext()

	resetAllFunc()

	tsAfter := chk.ClockNext()

	chk.True(tsBefore.Before(tsAfter))

	chk.True(tsBefore.Before(ts1))
	chk.Dur(ts2.Sub(ts1), time.Microsecond)
	chk.Dur(ts3.Sub(ts2), time.Microsecond*10)

	chk.True(tsd1.Before(ts1))
	chk.Dur(tsd2.Sub(tsd1), time.Microsecond*100)
	chk.Dur(tsd3.Sub(tsd2), time.Microsecond*100)
	chk.True(tsd3.Equal(tsd3c))

	chk.Dur(ts4.Sub(ts3), time.Microsecond)
	chk.Dur(ts5.Sub(ts4), time.Microsecond*10)
	chk.Dur(ts6.Sub(ts5), time.Microsecond)

	chk.True(tsAfter.Before(ts1))
}

func chkClockUseCase2(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	_ = chk.ClockOffsetDay(-2, time.Millisecond, time.Millisecond*10)

	n := time.Now()
	t1 := chk.ClockNext()
	chk.True(t1.Before(n))

	_ = chk.ClockSet(
		time.Date(2999, 12, 25, 13, 15, 45, 555555555, time.Local),
	)

	t2 := chk.ClockNext()
	t3 := chk.ClockNext()
	t4 := chk.ClockNext()
	chk.Dur(t3.Sub(t2), time.Millisecond)
	chk.Dur(t4.Sub(t3), time.Millisecond*10)
}

func chkClockUseCase3(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	_ = chk.ClockSet(
		time.Date(2999, 12, 25, 13, 15, 45, 555555555, time.Local),
	)

	t1 := chk.ClockNext()
	t2 := chk.ClockNext()

	chk.Dur(t2.Sub(t1), time.Millisecond)
}

func chkClockSubs(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	tstTS := "2006/01/02 15:03:04"
	chk.ClockSetCusA(tstTS)
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkCusA0}}")

	tstTS = "20060102"
	chk.ClockSetCusB(tstTS)
	chk.ClockSetCusA("")
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkCusB1}}")

	chk.ClockSetSub(ClkFmtDate)
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkDate2}}")

	tstTS = "150405"
	chk.ClockSetCusC(tstTS)
	chk.ClockSetCusB("")
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkCusC3}}")

	chk.ClockSetSub(ClkFmtTime)
	chk.ClockSetCusC("")
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkTime4}}")

	tstTS = "20060102150405"

	chk.ClockAddSub(ClkFmtTS)
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkTS5}}")

	tstTS = "20060102150405.000000000"

	chk.ClockAddSub(ClkFmtNano)
	chk.Str(chk.ClockNext().Format(tstTS), "{{clkNano6}}")
}

func chkClockOffset(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	ts1 := chk.ClockNext() // Just the real time

	_ = chk.ClockOffset(time.Second) // Use internal clock + 1 second.

	ts2 := chk.ClockNext()
	chk.True(ts1.Before(ts2))
	chk.DurUnbounded(ts2.Sub(ts1), UnboundedMinClosed, time.Second)

	_ = chk.ClockOffset(time.Minute) // Use lastTS + 1 minute.

	ts3 := chk.ClockNext()
	chk.True(ts2.Before(ts3))
	chk.DurUnbounded(ts3.Sub(ts2), UnboundedMinClosed, time.Minute)
}
