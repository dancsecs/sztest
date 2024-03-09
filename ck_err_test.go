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
	"errors"
	"testing"
)

const (
	s1 = "entry1"
	s2 = "entry2"
	s3 = "entry3"
)

func tstChkErr(t *testing.T) {
	t.Run("Good", chkErrTest_Good)

	t.Run("Bad", chkErrTest_Bad)
	t.Run("BadMsg1", chkErrTest_Bad1)
	t.Run("BadMsg2", chkErrTest_Bad2)
	t.Run("BadMsg3", chkErrTest_Bad3)

	t.Run("NoErrBadMsg1", chkErrTest_NoErrBad1)
	t.Run("NoErrBadMsg2", chkErrTest_NoErrBad2)

	t.Run("Slice-Good", chkErrSliceTest_Good)
	t.Run("Slice-BadMsg1", chkErrSliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkErrSliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkErrSliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkErrSliceTest_BadMsg4)
}

func chkErrTest_Good(t *testing.T) {
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

func chkErrTest_Bad(t *testing.T) {
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
			chkOutCommonMsg("", "err"),
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

func chkErrTest_Bad1(t *testing.T) {
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
			chkOutCommonMsg("This err message will be displayed first", "err"),
			g(markAsIns(`error mismatch for ""`)),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrTest_Bad2(t *testing.T) {
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
			chkOutCommonMsg("This err message will be displayed second", "err"),
			g(markAsIns("error mismatch for \"")+"<nil>"+markAsIns("\"")),
			w("<nil>"),
		),
		chkOutRelease(),
	)
}

func chkErrTest_Bad3(t *testing.T) {
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
			chkOutCommonMsg("This err message will be displayed third", "err"),
			g(markAsIns("error mismatch for \"")+"expected error"+markAsIns("\"")),
			w("expected error"),
		),
		chkOutRelease(),
	)
}

func chkErrTest_NoErrBad1(t *testing.T) {
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
			chkOutCommonMsg("This err message will be displayed", "err"),
			g(markAsIns("expected error")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrTest_NoErrBad2(t *testing.T) {
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
			chkOutCommonMsg("This err message will be displayed noErrSecond", "err"),
			g(markAsIns("expected error")),
			w(""),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	e1 := errors.New(s1)
	e2 := errors.New(s2)
	e3 := errors.New(s3)

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
		[]error{e1}, []string{s1}, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{e1, e2}, []string{s1, s2}, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{e1, e2, e3}, []string{s1, s2, s3},
		"This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{e2, e3}, []string{s2, s3}, "This message will NOT be displayed",
	)
	chk.ErrSlice(
		[]error{e3}, []string{s3}, "This message will NOT be displayed",
	)

	chk.ErrSlicef(
		[]error{e3, e2}, []string{s3, s2},
		"This message will NOT be %s",
		"displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkErrSliceTest_BadMsg1(t *testing.T) {
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

func chkErrSliceTest_BadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ErrSlice([]error{errors.New("x")}, nil, "This message will be displayed")

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

func chkErrSliceTest_BadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	e3 := errors.New(s3)

	chk.ErrSlicef(
		[]error{e3}, []string{s1}, "This message will be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 1, "[]err", "ErrSlicef", "This message will be displayed",
			chkOutLnChanged("0", "0", "entry"+markAsChg("3", "1", DiffMerge)),
		),
		chkOutRelease(),
	)
}

func chkErrSliceTest_BadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	e1 := errors.New(s1)
	e3 := errors.New(s3)

	chk.ErrSlice(
		[]error{e1, e1, e3}, []string{s1, s2, s3},
		"This message will be displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			3, 3, "[]err", "ErrSlice", "This message will be displayed",
			chkOutLnSame("0", "0", "entry1"),
			chkOutLnChanged("1", "1", "entry"+markAsChg("1", "2", DiffMerge)),
			chkOutLnSame("2", "2", "entry3"),
		),
		chkOutRelease(),
	)
}
