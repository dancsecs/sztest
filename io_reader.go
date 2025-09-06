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
	"fmt"
	"io"
	"os"
	"strings"
)

// SetIOReaderData initializes chk’s io.Reader with the supplied strings.
// The strings are concatenated and served sequentially through Read. Once
// exhausted, Read returns io.EOF.
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

// SetIOReaderError configures chk’s io.Reader to return err once the given
// byteCount has been read. After returning err, chk clears it and continues
// serving any remaining data until EOF.
func (chk *Chk) SetIOReaderError(byteCount int, err error) {
	chk.rErrPos = byteCount
	chk.rErr = err
}

// SetReadError primes chk’s io.Reader to return (pos, err) on the next
// call to Read. After returning, chk clears the error and resumes normal
// reading for subsequent calls.
func (chk *Chk) SetReadError(pos int, err error) {
	chk.ioReadErrPos = pos
	chk.ioReadErr = err
	chk.ioReadErrSet = true
}

// Read implements io.Reader for chk. It serves data provided by
// SetIOReaderData, returns injected errors as configured by
// SetIOReaderError or SetReadError, and yields io.EOF once all data is
// consumed.
//
//nolint:cyclop // Ok.
func (chk *Chk) Read(dataBuf []byte) (int, error) {
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
	mi := len(dataBuf)

	for i < mi && chk.rLeft > 0 {
		dataBuf[i] = chk.rData[chk.rPos]
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

// SetStdinData replaces os.Stdin with a stream that sequentially provides
// the supplied lines. Once exhausted, reads return io.EOF. This is not part
// of io.Reader itself but enables testing of code that directly consumes
// os.Stdin.
func (chk *Chk) SetStdinData(lines ...string) {
	chk.t.Helper()

	wPipeClosed := false

	origStdin := os.Stdin
	rPipe, wPipe, err := os.Pipe()

	if chk.NoErr(err) {
		chk.PushPostReleaseFunc(func() error {
			os.Stdin = origStdin

			err = wPipe.Close()
			if err == nil || wPipeClosed {
				err = rPipe.Close()
			}

			return err //nolint:wrapcheck // Ok.
		})

		chk.NoErr(err)

		os.Stdin = rPipe

		_, err = fmt.Fprint(wPipe, strings.Join(lines, ""))
		chk.NoErr(err)

		chk.NoErr(wPipe.Close())

		wPipeClosed = true
	}
}
