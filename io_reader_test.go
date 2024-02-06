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
	"bufio"
	"errors"
	"fmt"
	"io"
	"testing"
)

func tstChkIoReader(t *testing.T) {
	t.Run("IOReaderNoError1", chkIoReaderTest_IOReaderNoError1)
	t.Run("IOReaderNoError2", chkIoReaderTest_IOReaderNoError2)
	t.Run("IOReaderError1", chkIoReader_IOReaderError1)
	t.Run("IOReaderError2", chkIoReader_IOReaderError2)
	t.Run("IOReaderError3", chkIoReader_IOReaderError3)
	t.Run("SetReadError", chkIOReader_SetReadError)
	t.Run("SetStdinData", chkIOReader_SetStdinData)
}

func chkIoReaderTest_IOReaderNoError1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOReaderData(
		"This is the first line.\n",
		"This is the second line.\n",
		"This is the third line.\n",
		"This is the fourth line without linefeed.",
	)
	got := ""
	s := bufio.NewScanner(chk)
	for s.Scan() {
		got += s.Text() + "\n"
	}
	chk.NoErr(s.Err())

	chk.Str(
		got,
		""+
			"This is the first line.\n"+
			"This is the second line.\n"+
			"This is the third line.\n"+
			"This is the fourth line without linefeed.\n"+
			"",
	)
}

func chkIoReaderTest_IOReaderNoError2(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOReaderData(
		"This is the first line.\n",
		"This is the second line.\n",
		"This is the third line.\n",
		"This is the fourth line with linefeed.\n",
	)
	got := ""
	s := bufio.NewScanner(chk)
	for s.Scan() {
		got += s.Text() + "\n"
	}
	chk.NoErr(s.Err())

	chk.Str(
		got,
		""+
			"This is the first line.\n"+
			"This is the second line.\n"+
			"This is the third line.\n"+
			"This is the fourth line with linefeed.\n"+
			"",
	)
}

func chkIoReader_IOReaderError1(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOReaderData(
		"This is the first line.\n",
		"This is the second line.\n",
		"This is the third line.\n",
		"This is the fourth line with linefeed.\n",
	)

	got := ""
	s := bufio.NewScanner(chk)
	for s.Scan() {
		got += s.Text() + "\n"
	}
	chk.NoErr(s.Err())

	_, err := chk.Read(make([]byte, 10))
	chk.Err(
		err,
		"Read past end of data",
	)

	chk.Str(
		got,
		""+
			"This is the first line.\n"+
			"This is the second line.\n"+
			"This is the third line.\n"+
			"This is the fourth line with linefeed.\n"+
			"",
	)
}

func chkIoReader_IOReaderError2(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOReaderData(
		"This is the first line.\n",
		"This is the second line.\n",
		"This is the third line.\n",
		"This is the fourth line with linefeed.\n",
	)
	chk.SetIOReaderError(0, errors.New("We should get this error"))
	_, err := chk.Read(make([]byte, 10))
	chk.Err(
		err,
		"We should get this error",
	)
}

func chkIoReader_IOReaderError3(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetIOReaderData(
		"This is the first line.\n",
		"This is the second line.\n",
		"This is the third line.\n",
		"This is the fourth line with linefeed.\n",
	)
	chk.SetIOReaderError(55, errors.New("This error after 55 characters"))

	got := ""
	s := bufio.NewScanner(chk)
	for s.Scan() {
		got += s.Text() + "\n"
	}
	chk.Err(
		s.Err(),
		"This error after 55 characters",
	)

	chk.Str(
		got,
		""+
			"This is the first line.\n"+
			"This is the second line.\n"+
			"This i\n"+
			"",
	)
}

func chkIOReader_SetReadError(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	buf := make([]byte, 1)
	p, err := chk.Read(buf)
	chk.Int(p, 0)
	chk.Err(err, io.EOF.Error())

	chk.SetReadError(24, errors.New("the read error"))
	p, err = chk.Read(buf)
	chk.Int(p, 24)
	chk.Err(err, "the read error")
}

func chkIOReader_SetStdinData(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetStdinData("hello\n")

	var str string
	n, err := fmt.Scan(&str)

	chk.NoErr(err)
	chk.Int(n, 1)
	chk.Str(str, "hello")
}
