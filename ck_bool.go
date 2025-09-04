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

const boolTypeName = "bool"

// Helpers.

// Truef compares the got bool against true.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Truef(got bool, msgFmt string, msgArgs ...any) bool {
	if got {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, true, boolTypeName, msgFmt, msgArgs...)
}

// True compares the got bool against true.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) True(got bool, msg ...any) bool {
	if got {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, true, boolTypeName, msg...)
}

// Falsef compares the got bool against false.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Falsef(got bool, msgFmt string, msgArgs ...any) bool {
	if !got {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, false, boolTypeName, msgFmt, msgArgs...)
}

// False compares the got bool against false.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) False(got bool, msg ...any) bool {
	if !got {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, false, boolTypeName, msg...)
}

// Boolf compares the got bool against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Boolf(got, want bool, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, boolTypeName, msgFmt, msgArgs...)
}

// Bool compares the got bool against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Bool(got, want bool, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, boolTypeName, msg...)
}

// BoolSlicef compares two bool slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
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
		got, want, boolTypeName, defaultCmpFunc[bool], msgFmt, msgArgs...,
	)
}

// BoolSlice compares two bool slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
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

	return errSlice(chk, got, want, boolTypeName, defaultCmpFunc[bool], msg...)
}
