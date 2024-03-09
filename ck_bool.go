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

// Helpers.

// Truef simply invokes Bool with want set to true and msg formatted.
func (chk *Chk) Truef(got bool, msgFmt string, msgArgs ...any) bool {
	if got {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, true, "bool", msgFmt, msgArgs...)
}

// True simply invokes Bool with want set to true.
func (chk *Chk) True(got bool, msg ...any) bool {
	if got {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, true, "bool", msg...)
}

// Falsef simply invokes Bool with want set to true and msg formatted.
func (chk *Chk) Falsef(got bool, msgFmt string, msgArgs ...any) bool {
	if !got {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, false, "bool", msgFmt, msgArgs...)
}

// False simply invokes Bool with want set to true.
func (chk *Chk) False(got bool, msg ...any) bool {
	if !got {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, false, "bool", msg...)
}

// Boolf compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Boolf(got, want bool, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, want, "bool", msgFmt, msgArgs...)
}

// Bool compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Bool(got, want bool, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, want, "bool", msg...)
}

// BoolSlicef checks two boolean slices for equality.
func (chk *Chk) BoolSlicef(
	got, want []bool, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlicef(chk,
		got, want, "bool", defaultCmpFunc[bool], msgFmt, msgArgs...,
	)
}

// BoolSlice checks two boolean slices for equality.
func (chk *Chk) BoolSlice(got, want []bool, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlice(chk, got, want, "bool", defaultCmpFunc[bool], msg...)
}
