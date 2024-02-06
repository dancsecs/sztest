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

// Complex64f compares the wanted complex64 against the gotten complex64
// invoking an error should they not match.
func (chk *Chk) Complex64f(
	got, want complex64, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "complex64", msgFmt, msgArgs...)
}

// Complex64 compares the wanted complex64 against the gotten complex64
// invoking an error should they not match.
func (chk *Chk) Complex64(got, want complex64, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "complex64", msg...)
}

// Complex64Slicef checks two complex64 slices for equality.
func (chk *Chk) Complex64Slicef(
	got, want []complex64, msgFmt string, msgArgs ...any,
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
		got, want, "complex64", defaultCmpFunc[complex64],
		msgFmt, msgArgs...,
	)
}

// Complex64Slice checks two complex64 slices for equality.
func (chk *Chk) Complex64Slice(
	got, want []complex64, msg ...any,
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
	return errSlice(chk,
		got, want, "complex64", defaultCmpFunc[complex64], msg...,
	)
}
