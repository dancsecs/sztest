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

// LastErr returns the last argument in the list as an error.  If there are
// no arguments or the last parameter is not an Error interface then
// ErrInvalidLastErrArg is returned.
func (*Chk) LastErr(args ...any) error {
	var (
		err     error
		lastErr error
		ok      bool
	)

	if len(args) == 0 {
		err = ErrInvalidLastArg
	}

	if err == nil {
		lastErr, ok = args[len(args)-1].(error)
		if !ok {
			if args[len(args)-1] != nil {
				err = ErrInvalidLastArg
			}
		}
	}

	if err == nil {
		return lastErr
	}

	return err
}
