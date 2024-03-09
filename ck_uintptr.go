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

// Uintptrf compares the wanted uintptr against the gotten uintptr invoking an
// error should they not match.
func (chk *Chk) Uintptrf(
	got, want uintptr, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, want, "uintptr", msgFmt, msgArgs...)
}

// Uintptr compares the wanted uintptr against the gotten uintptr invoking an
// error should they not match.
func (chk *Chk) Uintptr(got, want uintptr, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, want, "uintptr", msg...)
}

// UintptrSlicef checks two uintptr slices for equality.
func (chk *Chk) UintptrSlicef(
	got, want []uintptr, msgFmt string, msgArgs ...any,
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
		got, want, "uintptr", defaultCmpFunc[uintptr],
		msgFmt, msgArgs...,
	)
}

// UintptrSlice checks two uintptr slices for equality.
func (chk *Chk) UintptrSlice(got, want []uintptr, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlice(chk, got, want, "uintptr", defaultCmpFunc[uintptr], msg...)
}
