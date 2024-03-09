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

func tstChkIoWriter(t *testing.T) {
	t.Run("IOWriterNoError1", chkIoWriterTestIOWriterNoError1)
	t.Run("IOWriterError1", chkIoWriterTestIOWriterError1)
	t.Run("IOWriterError2", chkIoWriterTestIOWriterError2)
	t.Run("IOWriterError3", chkIoWriterTestIOWriterError3)
	t.Run("IOWriterError4", chkIoWriterTestIOWriterError4)
	t.Run("IOWriterError5", chkIoWriterTestIOWriterError5)
	t.Run("IOWriterError6", chkIoWriterTestIOWriterError6)
	t.Run("SetWriteError", chkIoWriterTestSetWriteError)
}

func chkIoWriterTestIOWriterNoError1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(-1, nil)

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.NoErr(err)
	_, err = chk.Write([]byte("This is the second line.\n"))
	chk.NoErr(err)
	_, err = chk.Write([]byte("This is the third line.\n"))
	chk.NoErr(err)
	_, err = chk.Write([]byte("This is the fourth line.\n"))
	chk.NoErr(err)

	chk.Str(
		string(chk.GetIOWriterData()),
		""+
			"This is the first line.\n"+
			"This is the second line.\n"+
			"This is the third line.\n"+
			"This is the fourth line.\n"+
			"",
	)
}

func chkIoWriterTestIOWriterError1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(0, nil)

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(
		err,
		ErrForcedOutOfSpace.Error(),
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		"",
	)
}

func chkIoWriterTestIOWriterError2(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(0, errors.New("this will be the custom error"))

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(
		err,
		"this will be the custom error",
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		"",
	)
}

func chkIoWriterTestIOWriterError3(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(1, nil)

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(
		err,
		ErrForcedOutOfSpace.Error(),
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		"T",
	)
}

func chkIoWriterTestIOWriterError4(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(1, errors.New("this will be the custom error"))

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(
		err,
		"this will be the custom error",
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		"T",
	)
}

func chkIoWriterTestIOWriterError5(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(30, nil)

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(err, "")
	_, err = chk.Write([]byte("This is the second line.\n"))
	chk.Err(
		err,
		ErrForcedOutOfSpace.Error(),
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		""+
			"This is the first line.\n"+
			"This i"+
			"",
	)
}

func chkIoWriterTestIOWriterError6(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOWriterError(30, errors.New("this will be the custom error"))

	_, err := chk.Write([]byte("This is the first line.\n"))
	chk.Err(err, "")
	_, err = chk.Write([]byte("This is the second line.\n"))
	chk.Err(
		err,
		"this will be the custom error",
	)

	chk.Str(
		string(chk.GetIOWriterData()),
		""+
			"This is the first line.\n"+
			"This i"+
			"",
	)
}

func chkIoWriterTestSetWriteError(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	buf := make([]byte, 1)
	p, err := chk.Write(buf)
	chk.Int(p, 1)
	chk.NoErr(err)

	chk.SetWriteError(24, errors.New("the write error"))
	p, err = chk.Write(buf)
	chk.Int(p, 24)
	chk.Err(err, "the write error")
}
