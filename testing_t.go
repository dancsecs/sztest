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

// testingT is the minimal interface sztest requires from *testing.T.
// It exists to decouple chk from the concrete *testing.T type so that
// sztest can test itself by substituting a recorder or mock. This enables
// full coverage while still integrating seamlessly with Go's testing
// framework in normal use.
type testingT interface {
	Helper()
	Logf(msgFmt string, msgArgs ...any)
	Errorf(msgFmt string, msgArgs ...any)
	Error(msgArgs ...any)
	Fatalf(msgFmt string, msgArgs ...any)
	FailNow()
	SkipNow()
	Name() string
}
