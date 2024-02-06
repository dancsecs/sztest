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

// Uint64f compares the wanted uint64 against the gotten uint64 invoking an
// error should they not match.
func (chk *Chk) Uint64f(
	got, want uint64, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "uint64", msgFmt, msgArgs...)
}

// Uint64 compares the wanted uint64 against the gotten uint64 invoking an
// error should they not match.
func (chk *Chk) Uint64(got, want uint64, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "uint64", msg...)
}

// Uint64Slicef checks two uint64 slices for equality.
func (chk *Chk) Uint64Slicef(
	got, want []uint64, msgFmt string, msgArgs ...any,
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
		got, want, "uint64", defaultCmpFunc[uint64],
		msgFmt, msgArgs...,
	)
}

// Uint64Slice checks two uint64 slices for equality.
func (chk *Chk) Uint64Slice(got, want []uint64, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk, got, want, "uint64", defaultCmpFunc[uint64], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// Uint64Boundedf checks value is within specified bounded range.
func (chk *Chk) Uint64Boundedf(
	got uint64, option BoundedOption, min, max uint64,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint64"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint64Bounded checks value is within specified bounded range.
func (chk *Chk) Uint64Bounded(
	got uint64, option BoundedOption, min, max uint64, msg ...any,
) bool {
	const typeName = "uint64"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// Uint64Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Uint64Unboundedf(
	got uint64, option UnboundedOption, bound uint64,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "uint64"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Uint64Unbounded checks value is within specified unbounded range.
func (chk *Chk) Uint64Unbounded(
	got uint64, option UnboundedOption, bound uint64, msg ...any,
) bool {
	const typeName = "uint64"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
