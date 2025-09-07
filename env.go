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
	"os"
)

// SetEnv sets or updates the named environment variable to the given value.
// Changes are reverted automatically when chk.Release() is called. Any error
// encountered is reported to the underlying *testingT.
func (chk *Chk) SetEnv(name, value string) {
	var reverseFunc func() error

	chk.T().Helper()

	currentValue, found := os.LookupEnv(name)
	if found {
		reverseFunc = func() error {
			return os.Setenv(name, currentValue)
		}
	} else {
		reverseFunc = func() error {
			return os.Unsetenv(name)
		}
	}

	chk.PushPreReleaseFunc(reverseFunc)
	chk.NoErr(os.Setenv(name, value))
}

// DelEnv removes the named environment variable if it exists. The removal is
// reverted automatically when chk.Release() is called. Any error encountered
// is reported to the underlying *testingT.
func (chk *Chk) DelEnv(name string) {
	var reverseFunc func() error

	chk.T().Helper()

	currentValue, found := os.LookupEnv(name)
	if found {
		reverseFunc = func() error {
			return os.Setenv(name, currentValue)
		}
	} else {
		reverseFunc = func() error {
			return os.Unsetenv(name)
		}
	}

	chk.PushPreReleaseFunc(reverseFunc)
	chk.NoErr(os.Unsetenv(name))
}
