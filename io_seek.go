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

// SetSeekError primes the chk object to return the provided error on a
// future Seek call. The error is returned once, after which normal seek
// behavior resumes.
func (chk *Chk) SetSeekError(pos int64, err error) {
	chk.ioSeekErrPos = pos
	chk.ioSeekErr = err
	chk.ioSeekErrSet = true
}

// Seek implements the io.Seeker interface. It updates the current seek
// position and returns any pending error set via SetSeekError. If no
// error is pending, it behaves as a successful seek.
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
