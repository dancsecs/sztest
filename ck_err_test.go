/*
   Golang test helper library: sztest.
   Copyright (C) 2023-2025 Leslie Dancsecs

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
	"errors"
	"testing"
)

const (
	tstStr1 = "entry1"
	tstStr2 = "entry2"
	tstStr3 = "entry3"
)

func tstChkErr(t *testing.T) {
	t.Run("Good", chkErrTestGood)

	t.Run("Bad", chkErrTestBad)
	t.Run("BadMsg1", chkErrTestBad1)
	t.Run("BadMsg2", chkErrTestBad2)
	t.Run("BadMsg3", chkErrTestBad3)

	t.Run("NoErrBadMsg1", chkErrTestNoErrBad1)
	t.Run("NoErrBadMsg2", chkErrTestNoErrBad2)

	t.Run("Slice-Good", chkErrSliceTestGood)
	t.Run("Slice-BadMsg1", chkErrSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkErrSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkErrSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkErrSliceTestBadMsg4)

	t.Run("ErrChain", chkErrChain)
}

func chkErrTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Err(nil, "", "not ", "displayed")
	chk.Errf(nil, "", "not %s", "displayed")
	chk.Err(
		errors.New("expected error"),
		"expected error",
		"This message will NOT be displayed",
	)

	chk.NoErr(nil, "This message will NOT be displayed")
	chk.NoErrf(nil, "This message will NOT be displayed")

	chk.Err(errors.New(""), BlankErrorMessage)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkErrTestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Err(
		errors.New("error mismatch for \"expected error\""),
		"expected error",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Err",
			chkOutCommonMsg("", errTypeName),
			g(
				markAsIns(`error mismatch for "`)+
					"expected error"+
					markAsIns(`"`),
			),
			w("expected error"),
		),
		chkOutRelease(),
	)
}

func chkErrTestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Errf(
		errors.New("error mismatch for \"\""),
		"", "This err message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Errf",
			chkOutCommonMsg(
				"This err message will be displayed first",
				errTypeName,
			),
			g(markAsIns(`error mismatch for ""`)),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrTestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Errf(errors.New("error mismatch for \"<nil>\""),
		"<nil>",
		"This err message will be displayed %s", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Errf",
			chkOutCommonMsg(
				"This err message will be displayed second",
				errTypeName,
			),
			g(markAsIns("error mismatch for \"")+"<nil>"+markAsIns("\"")),
			w("<nil>"),
		),
		chkOutRelease(),
	)
}

func chkErrTestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Err(
		errors.New("error mismatch for \"expected error\""), "expected error",
		"This err message will be displayed ", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Err",
			chkOutCommonMsg(
				"This err message will be displayed third",
				errTypeName,
			),
			g(
				markAsIns(
					"error mismatch for \"")+"expected error"+markAsIns("\""),
			),
			w("expected error"),
		),
		chkOutRelease(),
	)
}

func chkErrTestNoErrBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.NoErr(
		errors.New("expected error"),
		"This err message will be displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"NoErr",
			chkOutCommonMsg("This err message will be displayed", errTypeName),
			g(markAsIns("expected error")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrTestNoErrBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.NoErrf(
		errors.New("expected error"),
		"This err message will be displayed %s",
		"noErrSecond",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"NoErrf",
			chkOutCommonMsg(
				"This err message will be displayed noErrSecond",
				errTypeName,
			),
			g(markAsIns("expected error")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstErr1 := errors.New(tstStr1)
	tstErr2 := errors.New(tstStr2)
	tstErr3 := errors.New(tstStr3)

	chk.ErrSlice(
		nil, nil, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		nil, []string{}, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{}, nil, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{}, []string{}, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{tstErr1},
		[]string{tstStr1},
		"This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{tstErr1, tstErr2},
		[]string{tstStr1, tstStr2},
		"This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{tstErr1, tstErr2, tstErr3},
		[]string{tstStr1, tstStr2, tstStr3},
		"This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{tstErr2, tstErr3},
		[]string{tstStr2, tstStr3},
		"This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{tstErr3},
		[]string{tstStr3},
		"This message will NOT be displayed",
	)

	chk.ErrSlicef(
		[]error{tstErr3, tstErr2}, []string{tstStr3, tstStr2},
		"This message will NOT be %s",
		"displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkErrSliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ErrSlicef(nil, []string{""}, "This message will be %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1, "[]err", "ErrSlicef", "This message will be displayed",
			chkOutLnWnt("0", ""),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ErrSlice(
		[]error{errors.New("x")},
		nil,
		"This message will be displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0, "[]err", "ErrSlice", "This message will be displayed",
			chkOutLnGot("0", "x"),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstErr3 := errors.New(tstStr3)

	chk.ErrSlicef(
		[]error{tstErr3},
		[]string{tstStr1},
		"This message will be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 1, "[]err", "ErrSlicef", "This message will be displayed",
			chkOutLnChanged("0", "0", "entry"+markAsChg("3", "1", diffMerge)),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstErr1 := errors.New(tstStr1)
	tstErr3 := errors.New(tstStr3)

	chk.ErrSlice(
		[]error{tstErr1, tstErr1, tstErr3},
		[]string{tstStr1, tstStr2, tstStr3},
		"This message will be displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			3, 3, "[]err", "ErrSlice", "This message will be displayed",
			chkOutLnSame("0", "0", "entry1"),
			chkOutLnChanged("1", "1", "entry"+markAsChg("1", "2", diffMerge)),
			chkOutLnSame("2", "2", "entry3"),
		),
		chkOutRelease(),
	)
}

func chkErrChain(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	tstErr1 := errors.New(tstStr1)
	tstErr3 := errors.New(tstStr3)

	chk.Str(
		chk.ErrChain(tstErr1, "and this statement", tstErr3),
		tstErr1.Error()+
			": and this statement: "+
			tstErr3.Error(),
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}
