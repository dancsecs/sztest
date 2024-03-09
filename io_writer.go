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

// GetIOWriterData returns the bytes received on the ioWriter interface.
func (chk *Chk) GetIOWriterData() []byte {
	return chk.wData
}

// SetIOWriterError provides to limit the amount of bytes that will be read
// before an error (or the supplied error will be returned.)
func (chk *Chk) SetIOWriterError(n int, err error) {
	chk.wErrPos = n
	chk.wErr = err
	chk.wData = make([]byte, 0)
}

// SetWriteError primes the chk object to return the provided error on the
// next Write operation.
func (chk *Chk) SetWriteError(pos int, err error) {
	chk.ioWriteErrPos = pos
	chk.ioWriteErr = err
	chk.ioWriteErrSet = true
}

// Write implements the ioReader interface.
func (chk *Chk) Write(b []byte) (int, error) {
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
	for _, nb := range b {
		if chk.wErrPos == 0 {
			if chk.wErr == nil {
				return count, ErrForcedOutOfSpace
			}
			return count, chk.wErr
		}
		chk.wData = append(chk.wData, nb)
		chk.wErrPos--
		count++
	}
	return count, nil
}
