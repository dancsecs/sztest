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

// GetIOWriterData returns all bytes written to the io.Writer interface
// so far, and clears the internal buffer. This is useful for verifying
// output in tests.
func (chk *Chk) GetIOWriterData() []byte {
	data := chk.wData

	chk.wData = []byte{}

	return data
}

// SetIOWriterError configures the io.Writer interface to return the
// specified error after writing n bytes. Once triggered, the error is
// cleared and subsequent writes proceed as normal unless another error
// is set.
func (chk *Chk) SetIOWriterError(n int, err error) {
	chk.wErrPos = n
	chk.wErr = err
	chk.wData = make([]byte, 0)
}

// SetWriteError primes the chk object to return the given position and
// error on the very next Write call. The error is returned once, and then
// cleared automatically.
func (chk *Chk) SetWriteError(pos int, err error) {
	chk.ioWriteErrPos = pos
	chk.ioWriteErr = err
	chk.ioWriteErrSet = true
}

// Write implements the io.Writer interface. Data is recorded internally
// and can be retrieved via GetIOWriterData. Pending errors set with
// SetIOWriterError or SetWriteError take precedence and are returned
// according to their rules.
func (chk *Chk) Write(data []byte) (int, error) {
	if chk.ioWriteErrSet {
		writePos := chk.ioWriteErrPos
		writeErr := chk.ioWriteErr
		chk.ioWriteErrPos = 0
		chk.ioWriteErr = nil
		chk.ioWriteErrSet = false

		return writePos, writeErr
	}

	count := 0

	if chk.wErr != nil && chk.wErrPos <= 0 {
		return 0, chk.wErr
	}

	for _, nextByte := range data {
		if chk.wErrPos == 0 {
			if chk.wErr == nil {
				return count, ErrForcedOutOfSpace
			}

			return count, chk.wErr
		}

		chk.wData = append(chk.wData, nextByte)
		chk.wErrPos--
		count++
	}

	return count, nil
}
