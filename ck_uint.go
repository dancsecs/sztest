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

// Uintf compares the wanted uint against the gotten uint invoking an
// error should they not match.
func (chk *Chk) Uintf(got, want uint, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "uint", msgFmt, msgArgs...)
}

// Uint compares the wanted uint against the gotten uint invoking an
// error should they not match.
func (chk *Chk) Uint(got, want uint, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "uint", msg...)
}

// UintSlicef checks two uint slices for equality.
func (chk *Chk) UintSlicef(
	got, want []uint, msgFmt string, msgArgs ...any,
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
		got, want, "uint", defaultCmpFunc[uint], msgFmt, msgArgs...,
	)
}

// UintSlice checks two uint slices for equality.
func (chk *Chk) UintSlice(got, want []uint, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk, got, want, "uint", defaultCmpFunc[uint], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// UintBoundedf checks value is within specified bounded range.
func (chk *Chk) UintBoundedf(
	got uint, option BoundedOption, min, max uint, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// UintBounded checks value is within specified bounded range.
func (chk *Chk) UintBounded(
	got uint, option BoundedOption, min, max uint, msg ...any,
) bool {
	const typeName = "uint"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// UintUnboundedf checks value is within specified unbounded range.
func (chk *Chk) UintUnboundedf(
	got uint, option UnboundedOption, bound uint, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// UintUnbounded checks value is within specified unbounded range.
func (chk *Chk) UintUnbounded(
	got uint, option UnboundedOption, bound uint, msg ...any,
) bool {
	const typeName = "uint"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
