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

// Intf compares the wanted int against the gotten int invoking an
// error should they not match.
func (chk *Chk) Intf(got, want int, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "int", msgFmt, msgArgs...)
}

// Int compares the wanted int against the gotten int invoking an
// error should they not match.
func (chk *Chk) Int(got, want int, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "int", msg...)
}

// IntSlicef checks two int slices for equality.
func (chk *Chk) IntSlicef(
	got, want []int, msgFmt string, msgArgs ...any,
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
		got, want, "int", defaultCmpFunc[int], msgFmt, msgArgs...,
	)
}

// IntSlice checks two int slices for equality.
func (chk *Chk) IntSlice(got, want []int, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk, got, want, "int", defaultCmpFunc[int], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// IntBoundedf checks value is within specified bounded range.
func (chk *Chk) IntBoundedf(
	got int, option BoundedOption, min, max int, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "int"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// IntBounded checks value is within specified bounded range.
func (chk *Chk) IntBounded(
	got int, option BoundedOption, min, max int, msg ...any,
) bool {
	const typeName = "int"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// IntUnboundedf checks value is within specified unbounded range.
func (chk *Chk) IntUnboundedf(
	got int, option UnboundedOption, bound int, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "int"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// IntUnbounded checks value is within specified unbounded range.
func (chk *Chk) IntUnbounded(
	got int, option UnboundedOption, bound int, msg ...any,
) bool {
	const typeName = "int"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
