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

// SetSeekError primes the chk object to return the provided error.
func (chk *Chk) SetSeekError(pos int64, err error) {
	chk.ioSeekErrPos = pos
	chk.ioSeekErr = err
	chk.ioSeekErrSet = true
}

// Seek implements the interface to simulate a Seek operation returning an
// error if provided.
func (chk *Chk) Seek(_ int64, _ int) (int64, error) {
	seekPos := chk.ioSeekErrPos
	seekErr := chk.ioSeekErr

	if chk.ioSeekErrSet {
		chk.ioSeekErrPos = 0
		chk.ioSeekErr = nil
		chk.ioSeekErrSet = false
	}

	return seekPos, seekErr
}
