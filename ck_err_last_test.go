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

var (
	err1 = errors.New("test error 1")
	err2 = errors.New("test error 2")
)

func tstChkErrLast(t *testing.T) {
	t.Run("ErrLast1", chkErrLastTest1)
	t.Run("ErrLast2", chkErrLastTest2)

	t.Run("Slice-Good", chkErrSliceTestGood)
	t.Run("Slice-BadMsg1", chkErrSliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkErrSliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkErrSliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkErrSliceTestBadMsg4)
	t.Run("NilError", chkErrLastTestNilError)
}

func chkErrLastTest1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.Err(
		chk.LastErr(),
		ErrInvalidLastArg.Error(),
	)

	chk.Err(
		chk.LastErr(1),
		ErrInvalidLastArg.Error(),
	)

	chk.Err(
		chk.LastErr(1, 2),
		ErrInvalidLastArg.Error(),
	)

	chk.Err(
		chk.LastErr(err1),
		err1.Error(),
	)

	chk.Err(
		chk.LastErr(1, err1),
		err1.Error(),
	)

	chk.Err(
		chk.LastErr(1, err1, err2),
		err2.Error(),
	)
	chk.Err(
		chk.LastErr(1, err1, err2, 2),
		ErrInvalidLastArg.Error(),
	)
}

func chkErrLastTest2(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	fBad1 := func() int {
		return 0
	}

	fBad2 := func() (int, int) {
		return 0, 1
	}

	//nolint:stylecheck // OK testing last arg is not error.
	fBad3 := func() (int, error, int) {
		return 0, err1, 1
	}

	fGood1 := func() error {
		return err1
	}

	fGood2 := func() (int, error) {
		return 0, err1
	}

	fGood3 := func() (int, error, error) {
		return 0, err1, err2
	}

	chk.Err(
		chk.LastErr(fBad1()),
		ErrInvalidLastArg.Error(),
	)
	chk.Err(
		chk.LastErr(fBad2()),
		ErrInvalidLastArg.Error(),
	)
	chk.Err(
		chk.LastErr(fBad3()),
		ErrInvalidLastArg.Error(),
	)

	chk.Err(
		chk.LastErr(fGood1()),
		err1.Error(),
	)
	chk.Err(
		chk.LastErr(fGood2()),
		err1.Error(),
	)
	chk.Err(
		chk.LastErr(fGood3()),
		err2.Error(),
	)
}

func chkErrLastTestNilError(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.NoErr(chk.LastErr(nil))
}
