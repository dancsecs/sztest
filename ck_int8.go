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

// Int8f compares the wanted int8 against the gotten int8 invoking an
// error should they not match.
func (chk *Chk) Int8f(got, want int8, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChkf(got, want, "int8", msgFmt, msgArgs...)
}

// Int8 compares the wanted int8 against the gotten int8 invoking an
// error should they not match.
func (chk *Chk) Int8(got, want int8, msg ...any) bool {
	if got == want {
		return true
	}
	chk.t.Helper()

	return chk.errChk(got, want, "int8", msg...)
}

// Int8Slicef checks two int8 slices for equality.
func (chk *Chk) Int8Slicef(
	got, want []int8, msgFmt string, msgArgs ...any,
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
		got, want, "int8", defaultCmpFunc[int8], msgFmt, msgArgs...,
	)
}

// Int8Slice checks two int8 slices for equality.
func (chk *Chk) Int8Slice(got, want []int8, msg ...any) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}
	if equal {
		return true
	}
	chk.t.Helper()

	return errSlice(chk, got, want, "int8", defaultCmpFunc[int8], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// Int8Boundedf checks value is within specified bounded range.
func (chk *Chk) Int8Boundedf(
	got int8, option BoundedOption, min, max int8, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "int8"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Int8Bounded checks value is within specified bounded range.
func (chk *Chk) Int8Bounded(
	got int8, option BoundedOption, min, max int8, msg ...any,
) bool {
	const typeName = "int8"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWnt(typeName, got, want, msg...)
}

// Int8Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Int8Unboundedf(
	got int8, option UnboundedOption, bound int8, msgFmt string, msgArgs ...any,
) bool {
	const typeName = "int8"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Int8Unbounded checks value is within specified unbounded range.
func (chk *Chk) Int8Unbounded(
	got int8, option UnboundedOption, bound int8, msg ...any,
) bool {
	const typeName = "int8"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()

	return chk.errGotWnt(typeName, got, want, msg...)
}
