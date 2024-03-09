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
	"log"
	"math"
	"strconv"
	"time"
)

const (
	clkSubTime = "150405"
	clkSubDate = "20060102"
	clkSubTS   = clkSubDate + clkSubTime
	clkSubNano = clkSubTS + ".000000000"
)

// Clock substitutions.
const (
	ClockSubNone = 0           // No substitutions.
	ClockSubTime = 1 << iota   // {{clkTime#}} = HHmmSS.
	ClockSubDate               // {{clkDate#}} = YYYYMMDD.
	ClockSubTS                 // {{clkTS#}}   = YYYYMMDDHHmmSS.
	ClockSubNano               // {{clkNano#}} = YYYYMMDDHHmmSS.#########.
	ClockSubCusA               // {{clkCusA#}} = definable format string.
	ClockSubCusB               // {{clkCusB#}} = definable format string.
	ClockSubCusC               // {{clkCusC#}} = definable format string.
	ClockSubAll  = math.MaxInt // All defined substitutions.
)

// ClockSetSub sets the fields to set substitutions for.
func (chk *Chk) ClockSetSub(i int) {
	chk.clkSub = i
}

// ClockAddSub sets the fields to set substitutions for.
func (chk *Chk) ClockAddSub(i int) {
	chk.clkSub |= i
}

// ClockRemoveSub resets the fields to set substitutions for.
func (chk *Chk) ClockRemoveSub(i int) {
	chk.clkSub &= ^i
}

// ClockSetCusA sets the custom date format to set tick substitution values.
func (chk *Chk) ClockSetCusA(f string) {
	chk.clkCusA = f
	if f != "" {
		chk.ClockAddSub(ClockSubCusA)
	} else {
		chk.ClockRemoveSub(ClockSubCusA)
	}
}

// ClockSetCusB sets the custom date format to set tick substitution values.
func (chk *Chk) ClockSetCusB(f string) {
	chk.clkCusB = f
	if f != "" {
		chk.ClockAddSub(ClockSubCusB)
	} else {
		chk.ClockRemoveSub(ClockSubCusB)
	}
}

// ClockSetCusC sets the custom date format to set tick substitution values.
func (chk *Chk) ClockSetCusC(f string) {
	chk.clkCusC = f
	if f != "" {
		chk.ClockAddSub(ClockSubCusC)
	} else {
		chk.ClockRemoveSub(ClockSubCusC)
	}
}

// ClockLast returns the last timestamp generated.
func (chk *Chk) ClockLast() time.Time {
	return chk.clk.lastTS
}

// ClockLastFmtTime returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtTime() string {
	return chk.clk.lastTS.Format(clkSubTime)
}

// ClockLastFmtDate returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtDate() string {
	return chk.clk.lastTS.Format(clkSubDate)
}

// ClockLastFmtTS returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtTS() string {
	return chk.clk.lastTS.Format(clkSubTS)
}

// ClockLastFmtNano returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtNano() string {
	return chk.clk.lastTS.Format(clkSubNano)
}

// ClockLastFmtCusA returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtCusA() string {
	return chk.clk.lastTS.Format(chk.clkCusA)
}

// ClockLastFmtCusB returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtCusB() string {
	return chk.clk.lastTS.Format(chk.clkCusB)
}

// ClockLastFmtCusC returns the last time generated in the indicated format.
func (chk *Chk) ClockLastFmtCusC() string {
	return chk.clk.lastTS.Format(chk.clkCusC)
}

// ClockNext returns the current time or the next time sequence if a clock has
// been set.
func (chk *Chk) ClockNext() time.Time {
	nextTS := chk.clk.next()
	chk.clkTicks = append(chk.clkTicks, nextTS)

	if chk.clkSub != 0 { //nolint:nestif // Ok.
		idx := strconv.FormatInt(int64(len(chk.clkTicks)-1), base10)

		if chk.clkSub&ClockSubTime > 0 {
			chk.AddSub("{{clkTime"+idx+"}}", nextTS.Format(clkSubTime))
		}

		if chk.clkSub&ClockSubDate > 0 {
			chk.AddSub("{{clkDate"+idx+"}}", nextTS.Format(clkSubDate))
		}

		if chk.clkSub&ClockSubTS > 0 {
			chk.AddSub("{{clkTS"+idx+"}}", nextTS.Format(clkSubTS))
		}

		if chk.clkSub&ClockSubNano > 0 {
			chk.AddSub("{{clkNano"+idx+"}}", nextTS.Format(clkSubNano))
		}

		if chk.clkSub&ClockSubCusA > 0 {
			chk.AddSub("{{clkCusA"+idx+"}}", nextTS.Format(chk.clkCusA))
		}

		if chk.clkSub&ClockSubCusB > 0 {
			chk.AddSub("{{clkCusB"+idx+"}}", nextTS.Format(chk.clkCusB))
		}

		if chk.clkSub&ClockSubCusC > 0 {
			chk.AddSub("{{clkCusC"+idx+"}}", nextTS.Format(chk.clkCusC))
		}
	}

	return nextTS
}

// ClockNextFmtTime returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtTime() string {
	return chk.ClockNext().Format(clkSubTime)
}

// ClockNextFmtDate returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtDate() string {
	return chk.ClockNext().Format(clkSubDate)
}

// ClockNextFmtTS returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtTS() string {
	return chk.ClockNext().Format(clkSubTS)
}

// ClockNextFmtNano returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtNano() string {
	return chk.ClockNext().Format(clkSubNano)
}

// ClockNextFmtCusA returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtCusA() string {
	return chk.ClockNext().Format(chk.clkCusA)
}

// ClockNextFmtCusB returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtCusB() string {
	return chk.ClockNext().Format(chk.clkCusB)
}

// ClockNextFmtCusC returns the last time generated in the indicated format.
func (chk *Chk) ClockNextFmtCusC() string {
	return chk.ClockNext().Format(chk.clkCusC)
}

// ClockTick returns i'th time returned.
func (chk *Chk) ClockTick(i int) time.Time {
	if i < 0 || i >= len(chk.clkTicks) {
		log.Panicf("unknown tick index: %d", i)
	}

	return chk.clkTicks[i]
}

// ClockSet set the current test time and optionally sets the increments if
// provided.  It returns a func to reset the clk back to its state when
// this function was called.
func (chk *Chk) ClockSet(setTime time.Time, inc ...time.Duration) func() {
	savedClk := chk.clk

	if inc == nil {
		inc = chk.clk.inc
	}

	chk.clk = newTstClock(setTime, inc)

	return func() {
		chk.clk = savedClk
	}
}

// ClockOffsetDay adjusts the current clock by the number of specified days
// with negative numbers representing the past.  It returns a func to reset
// the clk back to its state when this function was called.
func (chk *Chk) ClockOffsetDay(dayOffset int, inc ...time.Duration) func() {
	t := chk.clk.lastTS

	return chk.ClockSet(t.AddDate(0, 0, dayOffset), inc...)
}

// ClockOffset moves the current clock by the specified amount.  No
// defined increments are applied and if a clock has not yet been set the
// current time advanced by the specified amount will be used. Nothing is
// returned.
func (chk *Chk) ClockOffset(d time.Duration) func() {
	t := chk.clk.nextTS

	return chk.ClockSet(t.Add(d))
}

type tstClk struct {
	lastTS    time.Time
	nextTS    time.Time
	inc       []time.Duration
	nextIndex int
}

func newTstClock(startAt time.Time, inc []time.Duration) *tstClk {
	if len(inc) < 1 {
		inc = []time.Duration{time.Millisecond}
	}

	return &tstClk{
		lastTS:    startAt.Add(-inc[len(inc)-1]),
		nextTS:    startAt,
		inc:       inc,
		nextIndex: 0,
	}
}

func (clk *tstClk) next() time.Time {
	clk.lastTS = clk.nextTS
	if clk.nextIndex >= len(clk.inc) {
		clk.nextIndex = 0
	}

	clk.nextTS = clk.nextTS.Add(clk.inc[clk.nextIndex])
	clk.nextIndex++

	return clk.lastTS
}
