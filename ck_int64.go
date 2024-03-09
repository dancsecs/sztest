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

const int64TypeName = "int64"

// Int64f compares the wanted int64 against the gotten int64 invoking an
// error should they not match.
func (chk *Chk) Int64f(got, want int64, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, int64TypeName, msgFmt, msgArgs...)
}

// Int64 compares the wanted int64 against the gotten int64 invoking an
// error should they not match.
func (chk *Chk) Int64(got, want int64, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, int64TypeName, msg...)
}

// Int64Slicef checks two int64 slices for equality.
func (chk *Chk) Int64Slicef(
	got, want []int64, msgFmt string, msgArgs ...any,
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
		got, want, int64TypeName, defaultCmpFunc[int64], msgFmt, msgArgs...,
	)
}

// Int64Slice checks two int64 slices for equality.
func (chk *Chk) Int64Slice(got, want []int64, msg ...any) bool {
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
		chk, got, want, int64TypeName, defaultCmpFunc[int64], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// Int64Boundedf checks value is within specified bounded range.
func (chk *Chk) Int64Boundedf(
	got int64, option BoundedOption, min, max int64,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int64TypeName, got, want, msgFmt, msgArgs...)
}

// Int64Bounded checks value is within specified bounded range.
func (chk *Chk) Int64Bounded(
	got int64, option BoundedOption, min, max int64, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int64TypeName, got, want, msg...)
}

// Int64Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Int64Unboundedf(
	got int64, option UnboundedOption, bound int64,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int64TypeName, got, want, msgFmt, msgArgs...)
}

// Int64Unbounded checks value is within specified unbounded range.
func (chk *Chk) Int64Unbounded(
	got int64, option UnboundedOption, bound int64, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int64TypeName, got, want, msg...)
}
