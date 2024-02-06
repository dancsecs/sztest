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

// SetCloseError primes the chk object to return the provided error.
func (chk *Chk) SetCloseError(err error) {
	chk.ioCloseErr = err
	chk.ioCloseErrSet = true
}

// Close implements the interface to simulate a close operation returning an
// error if provided.
func (chk *Chk) Close() error {
	closeErr := chk.ioCloseErr
	if chk.ioCloseErrSet {
		chk.ioCloseErr = nil
		chk.ioCloseErrSet = false
	}
	return closeErr
}
