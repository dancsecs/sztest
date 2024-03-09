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
	"fmt"
	"io"
	"os"
	"strings"
)

// SetIOReaderData initializes the IO reader interface for testing.
func (chk *Chk) SetIOReaderData(d ...string) {
	lines := ""
	for _, e := range d {
		lines += e
	}
	chk.rData = []byte(lines)
	chk.rLeft = len(chk.rData)
	chk.rPos = 0
	chk.rErrPos = -1
	chk.rErr = nil
}

// SetIOReaderError initializes the IO reader to return en error after the
// specified number of bytes have been read.
func (chk *Chk) SetIOReaderError(byteCount int, err error) {
	chk.rErrPos = byteCount
	chk.rErr = err
}

// SetReadError primes the chk object to return the provided error on the
// next read operation.
func (chk *Chk) SetReadError(pos int, err error) {
	chk.ioReadErrPos = pos
	chk.ioReadErr = err
	chk.ioReadErrSet = true
}

// Read implements the ioReader interface.
func (chk *Chk) Read(b []byte) (n int, err error) {
	if chk.ioReadErrSet {
		readPos := chk.ioReadErrPos
		readErr := chk.ioReadErr
		chk.ioReadErrPos = 0
		chk.ioReadErr = nil
		chk.ioReadErrSet = false
		return readPos, readErr
	}
	if chk.rErr != nil && (chk.rErrPos <= 0 || chk.rLeft <= 0) {
		return 0, chk.rErr
	}
	if chk.rLeft <= 0 {
		chk.rErr = ErrReadPastEndOfData
		return 0, io.EOF
	}

	i := 0
	mi := len(b)
	for i < mi && chk.rLeft > 0 {
		b[i] = chk.rData[chk.rPos]
		chk.rPos++
		chk.rLeft--
		chk.rErrPos--
		i++
		if chk.rErr != nil && (chk.rErrPos <= 0 || chk.rLeft <= 0) {
			break
		}
	}
	return i, nil
}

// SetStdinData sets the os.Stdin to stream the provided data.
func (chk *Chk) SetStdinData(d ...string) {
	chk.t.Helper()
	origStdin := os.Stdin
	r, w, err := os.Pipe()

	if chk.NoErr(err) {
		chk.PushPostReleaseFunc(func() error {
			os.Stdin = origStdin
			err = w.Close()
			if err == nil {
				err = r.Close()
			}
			return err //nolint:wrapcheck // Ok.
		})

		chk.NoErr(err)
		os.Stdin = r

		fmt.Fprint(w, strings.Join(d, ""))
	}
}
