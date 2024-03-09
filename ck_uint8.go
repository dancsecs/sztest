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

const uint8TypeName = "uint8"

// Uint8f compares the wanted uint8 against the gotten uint8 invoking an
// error should they not match.
func (chk *Chk) Uint8f(got, want uint8, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, uint8TypeName, msgFmt, msgArgs...)
}

// Uint8 compares the wanted uint8 against the gotten uint8 invoking an
// error should they not match.
func (chk *Chk) Uint8(got, want uint8, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, uint8TypeName, msg...)
}

// Uint8Slicef checks two uint8 slices for equality.
func (chk *Chk) Uint8Slicef(
	got, want []uint8, msgFmt string, msgArgs ...any,
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
		got, want, uint8TypeName, defaultCmpFunc[uint8], msgFmt, msgArgs...,
	)
}

// Uint8Slice checks two uint8 slices for equality.
func (chk *Chk) Uint8Slice(got, want []uint8, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk, got, want, uint8TypeName, defaultCmpFunc[uint8], msg...)
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges
////////////////////////////////////////////////////////////////

// Uint8Boundedf checks value is within specified bounded range.
func (chk *Chk) Uint8Boundedf(
	got uint8, option BoundedOption, min, max uint8,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint8TypeName, got, want, msgFmt, msgArgs...)
}

// Uint8Bounded checks value is within specified bounded range.
func (chk *Chk) Uint8Bounded(
	got uint8, option BoundedOption, min, max uint8, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint8TypeName, got, want, msg...)
}

// Uint8Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Uint8Unboundedf(
	got uint8, option UnboundedOption, bound uint8,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint8TypeName, got, want, msgFmt, msgArgs...)
}

// Uint8Unbounded checks value is within specified unbounded range.
func (chk *Chk) Uint8Unbounded(
	got uint8, option UnboundedOption, bound uint8, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint8TypeName, got, want, msg...)
}
