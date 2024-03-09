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

func tstChkDur(t *testing.T) {
	t.Run("Good", chkDurTestGood)

	t.Run("Bad", chkDurTestBad)
	t.Run("BadMsg1", chkDurTestBad1)
	t.Run("BadMsg2", chkDurTestBad2)

	t.Run("Slice-Good", chkDurSliceTestGood)
	t.Run("Slice-BadMsg1", chkDurSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkDurSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkDurSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkDurSliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkDurSliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkDurSliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkDurSliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkDurSliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkDurSliceTestBadMsg9)

	t.Run("Bounded", chkDurBoundedTestAll)
	t.Run("Unbounded", chkDurUnboundedTestAll)
}

func chkDurTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Dur(time.Second, time.Second)
	chk.Dur(time.Second, time.Second, "not ", "displayed")
	chk.Durf(time.Second, time.Second, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkDurTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Dur(time.Second*9, time.Second*7)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Dur",
			chkOutCommonMsg("", durTypeName),
			g(markAsChg("9s", "7s", DiffGot)),
			w(markAsChg("9s", "7s", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkDurTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Durf(
		time.Second*3, time.Second*2,
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Durf",
			chkOutCommonMsg("This message will be displayed first", durTypeName),
			g(markAsChg("3s", "2s", DiffGot)),
			w(markAsChg("3s", "2s", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkDurTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Dur(
		time.Second*6, time.Second*5,
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Dur",
			chkOutCommonMsg(
				"This message will be displayed second", durTypeName,
			),
			g(markAsChg("6s", "5s", DiffGot)),
			w(markAsChg("6s", "5s", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlice(
		[]time.Duration{}, []time.Duration{},
		"This message will NOT be displayed",
	)
	chk.DurSlice(
		[]time.Duration{dOne}, []time.Duration{dOne},
		"This message will NOT be displayed",
	)
	chk.DurSlice(
		[]time.Duration{dTwo, dTwo, dOne}, []time.Duration{dTwo, dTwo, dOne},
		"This message will NOT be displayed",
	)

	chk.DurSlicef(
		[]time.Duration{dTwo, dTwo, dOne, dOne},
		[]time.Duration{dTwo, dTwo, dOne, dOne},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dTwo := time.Second * 3

	chk.DurSlice(
		[]time.Duration{dTwo}, []time.Duration{},
		"This message will be displayed ", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]time.Duration", "DurSlice",
			"This message will be displayed first",
			chkOutLnGot("0", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2

	chk.DurSlicef(
		[]time.Duration{}, []time.Duration{dOne},
		"This message will be displayed %s", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]time.Duration", "DurSlicef",
			"This message will be displayed second",
			chkOutLnWnt("0", "2s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlice(
		[]time.Duration{dOne, dTwo}, []time.Duration{dOne},
		"This message will be displayed ", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]time.Duration", "DurSlice",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "2s"),
			chkOutLnGot("1", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlicef(
		[]time.Duration{dOne}, []time.Duration{dOne, dTwo},
		"This message will be displayed %s", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]time.Duration", "DurSlicef",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "2s"),
			chkOutLnWnt("1", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlicef(
		[]time.Duration{dOne, dTwo}, []time.Duration{dTwo, dTwo},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]time.Duration", "DurSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "2s"),
			chkOutLnSame("1", "0", "3s"),
			chkOutLnWnt("1", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlice(
		[]time.Duration{dTwo, dTwo}, []time.Duration{dOne, dTwo},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]time.Duration", "DurSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "2s"),
			chkOutLnSame("0", "1", "3s"),
			chkOutLnGot("1", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlicef(
		[]time.Duration{dOne, dTwo}, []time.Duration{dOne, dOne},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]time.Duration", "DurSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "2s"),
			chkOutLnChanged("1", "1", "3s", "2s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlice(
		[]time.Duration{dOne, dOne}, []time.Duration{dOne, dTwo},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]time.Duration", "DurSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "2s"),
			chkOutLnChanged("1", "1", "2s", "3s"),
		),
		chkOutRelease(),
	)
}

func chkDurSliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dOne := time.Second * 2
	dTwo := time.Second * 3

	chk.DurSlice([]time.Duration{dOne, dOne}, []time.Duration{dOne, dTwo})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]time.Duration", "DurSlice", "",
			chkOutLnSame("0", "0", "2s"),
			chkOutLnChanged("1", "1", "2s", "3s"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkDurBoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := time.Duration(33)
	max := time.Duration(35)

	// Bad: Error displayed.
	chk.DurBounded(30, BoundedClosed, min, max)
	chk.DurBounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.DurBoundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.DurBounded(33, BoundedClosed, min, max)
	chk.DurBounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.DurBoundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.DurBounded(36, BoundedClosed, min, max)

	const (
		wntMsg = "out of bounds: [33ns,35ns] - { want | 33ns <= want <= 35ns }"
		fName  = "DurBounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30ns", fName, durTypeName, ""),
		chkOutNumericBounded(wntMsg, "31ns", fName, durTypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32ns", fName, durTypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36ns", fName, durTypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkDurUnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := time.Duration(128)

	// Bad: Error displayed.
	chk.DurUnbounded(125, UnboundedMinClosed, bound)
	chk.DurUnbounded(126, UnboundedMinClosed, bound, "msg:", "126")
	chk.DurUnboundedf(127, UnboundedMinClosed, bound, "msg:%d", 127)

	// Good:  No error displayed.
	chk.DurUnbounded(128, UnboundedMinClosed, bound)
	chk.DurUnbounded(129, UnboundedMinClosed, bound, "not ", "displayed")
	chk.DurUnboundedf(130, UnboundedMinClosed, bound, "not %s", "displayed")

	const (
		wntMsg = "out of bounds: [128ns,MAX) - { want | want >= 128ns }"
		fName  = "DurUnbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "125ns", fName, durTypeName, ""),
		chkOutNumericUnbounded(wntMsg, "126ns", fName, durTypeName, "msg:126"),
		chkOutNumericUnboundedf(wntMsg, "127ns", fName, durTypeName, "msg:127"),

		chkOutRelease(),
	)
}
