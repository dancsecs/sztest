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

// Uint16f compares the wanted uint16 against the gotten uint16 invoking an
// error should they not match.
func (chk *Chk) Uint16f(
	got, want uint16, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "uint16", msgFmt, msgArgs...)
}

// Uint16 compares the wanted uint16 against the gotten uint16 invoking an
// error should they not match.
func (chk *Chk) Uint16(got, want uint16, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "uint16", msg...)
}

// Uint16Slicef checks two uint16 slices for equality.
func (chk *Chk) Uint16Slicef(
	got, want []uint16, msgFmt string, msgArgs ...any,
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
		got, want, "uint16", defaultCmpFunc[uint16],
		msgFmt, msgArgs...,
	)
}

// Uint16Slice checks two uint16 slices for equality.
func (chk *Chk) Uint16Slice(got, want []uint16, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk, got, want, "uint16", defaultCmpFunc[uint16], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// Uint16Boundedf checks value is within specified bounded range.
func (chk *Chk) Uint16Boundedf(
	got uint16, option BoundedOption, min, max uint16,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint16"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint16Bounded checks value is within specified bounded range.
func (chk *Chk) Uint16Bounded(
	got uint16, option BoundedOption, min, max uint16, msg ...any,
) bool {
	const typeName = "uint16"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// Uint16Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Uint16Unboundedf(
	got uint16, option UnboundedOption, bound uint16,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint16"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint16Unbounded checks value is within specified unbounded range.
func (chk *Chk) Uint16Unbounded(
	got uint16, option UnboundedOption, bound uint16, msg ...any,
) bool {
	const typeName = "uint16"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
