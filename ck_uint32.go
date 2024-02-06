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

// Uint32f compares the wanted uint32 against the gotten uint32 invoking an
// error should they not match.
func (chk *Chk) Uint32f(
	got, want uint32, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "uint32", msgFmt, msgArgs...)
}

// Uint32 compares the wanted uint32 against the gotten uint32 invoking an
// error should they not match.
func (chk *Chk) Uint32(got, want uint32, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "uint32", msg...)
}

// Uint32Slicef checks two uint32 slices for equality.
func (chk *Chk) Uint32Slicef(
	got, want []uint32, msgFmt string, msgArgs ...any,
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
		got, want, "uint32", defaultCmpFunc[uint32],
		msgFmt, msgArgs...,
	)
}

// Uint32Slice checks two uint32 slices for equality.
func (chk *Chk) Uint32Slice(got, want []uint32, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk, got, want, "uint32", defaultCmpFunc[uint32], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// Uint32Boundedf checks value is within specified bounded range.
func (chk *Chk) Uint32Boundedf(
	got uint32, option BoundedOption, min, max uint32,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint32"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint32Bounded checks value is within specified bounded range.
func (chk *Chk) Uint32Bounded(
	got uint32, option BoundedOption, min, max uint32, msg ...any,
) bool {
	const typeName = "uint32"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// Uint32Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Uint32Unboundedf(
	got uint32, option UnboundedOption, bound uint32,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint32"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint32Unbounded checks value is within specified unbounded range.
func (chk *Chk) Uint32Unbounded(
	got uint32, option UnboundedOption, bound uint32, msg ...any,
) bool {
	const typeName = "uint32"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
