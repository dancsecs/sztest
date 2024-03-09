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

const int32TypeName = "int32"

// Int32f compares the wanted int32 against the gotten int32 invoking an
// error should they not match.
func (chk *Chk) Int32f(got, want int32, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, int32TypeName, msgFmt, msgArgs...)
}

// Int32 compares the wanted int32 against the gotten int32 invoking an
// error should they not match.
func (chk *Chk) Int32(got, want int32, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, int32TypeName, msg...)
}

// Int32Slicef checks two int32 slices for equality.
func (chk *Chk) Int32Slicef(
	got, want []int32, msgFmt string, msgArgs ...any,
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
		got, want, int32TypeName, defaultCmpFunc[int32], msgFmt, msgArgs...,
	)
}

// Int32Slice checks two int32 slices for equality.
func (chk *Chk) Int32Slice(got, want []int32, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(
		chk, got, want, int32TypeName, defaultCmpFunc[int32], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// Int32Boundedf checks value is within specified bounded range.
func (chk *Chk) Int32Boundedf(
	got int32, option BoundedOption, min, max int32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int32TypeName, got, want, msgFmt, msgArgs...)
}

// Int32Bounded checks value is within specified bounded range.
func (chk *Chk) Int32Bounded(
	got int32, option BoundedOption, min, max int32, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int32TypeName, got, want, msg...)
}

// Int32Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Int32Unboundedf(
	got int32, option UnboundedOption, bound int32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int32TypeName, got, want, msgFmt, msgArgs...)
}

// Int32Unbounded checks value is within specified unbounded range.
func (chk *Chk) Int32Unbounded(
	got int32, option UnboundedOption, bound int32, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int32TypeName, got, want, msg...)
}
