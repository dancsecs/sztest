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

const complex128TypeName = "complex128"

// Complex128f compares the wanted complex128 against the gotten complex128
// invoking an error should they not match.
func (chk *Chk) Complex128f(
	got, want complex128, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, complex128TypeName, msgFmt, msgArgs...)
}

// Complex128 compares the wanted complex128 against the gotten complex128
// invoking an error should they not match.
func (chk *Chk) Complex128(got, want complex128, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, complex128TypeName, msg...)
}

// Complex128Slicef checks two complex128 slices for equality.
func (chk *Chk) Complex128Slicef(
	got, want []complex128, msgFmt string, msgArgs ...any,
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
		got, want, complex128TypeName, defaultCmpFunc[complex128],
		msgFmt, msgArgs...,
	)
}

// Complex128Slice checks two complex128 slices for equality.
func (chk *Chk) Complex128Slice(
	got, want []complex128, msg ...any,
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
		got, want, complex128TypeName, defaultCmpFunc[complex128], msg...,
	)
}
