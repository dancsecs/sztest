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
)

func tstChkPanic(t *testing.T) {
	t.Run("Good", chkPanicTest_Good)

	t.Run("Bad", chkPanicTest_Bad)
	t.Run("BadMsg1", chkPanicTest_Bad1)
	t.Run("BadMsg2", chkPanicTest_Bad2)
	t.Run("BadMsg3", chkPanicTest_Bad3)
	t.Run("BadMsg3", chkPanicTest_Bad4)
	t.Run("BadMsg3", chkPanicTest_Bad5)
}

func chkPanicTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Panic(
		func() {
		},
		"",
		"This message will NOT be displayed",
	)
	chk.NoPanic(
		func() {},
		"This message will NOT be displayed",
	)
	chk.NoPanicf(
		func() {},
		"This message will NOT be displayed",
	)
	chk.Panic(
		func() {
			panic("same")
		},
		"same",
		"This message will NOT be displayed",
	)

	chk.Panicf(
		func() {
			panic("same")
		},
		"same",
		"not %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Panic(
		func() {
			panic("Blank want")
		},
		"",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Panic",
			chkOutCommonMsg("", "panic"),
			g(markAsIns("Blank want")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Panicf(
		func() {
			panic("")
		},
		"",
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Panicf",
			chkOutCommonMsg("This message will be displayed first", "panic"),
			g(markAsIns(BlankPanicMessage)),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Panic(
		func() {
			panic("Blank want")
		},
		"",
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Panic",
			chkOutCommonMsg("This message will be displayed second", "panic"),
			g(markAsIns("Blank want")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Panicf(
		func() {
			panic("got")
		},
		"want",
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Panicf",
			chkOutCommonMsg("This message will be displayed third", "panic"),
			g(markAsChg("got", "want", DiffGot)),
			w(markAsChg("got", "want", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.NoPanic(
		func() {
			panic("Unexpected Panic")
		},
		"This message will be displayed ", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"NoPanic",
			chkOutCommonMsg("This message will be displayed fifth", "panic"),
			g(markAsIns("Unexpected Panic")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkPanicTest_Bad5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.NoPanicf(
		func() {
			panic("Unexpected Panicf")
		},
		"This message will be displayed %s", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"NoPanicf",
			chkOutCommonMsg("This message will be displayed sixth", "panic"),
			g(markAsIns("Unexpected Panicf")),
			w(""),
		),
		chkOutRelease(),
	)
}
