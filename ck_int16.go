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

const int16TypeName = "int16"

// Int16f compares the wanted int16 against the gotten int16 invoking an
// error should they not match.
func (chk *Chk) Int16f(got, want int16, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, int16TypeName, msgFmt, msgArgs...,
	)
}

// Int16 compares the wanted int16 against the gotten int16 invoking an
// error should they not match.
func (chk *Chk) Int16(got, want int16, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, int16TypeName, msg...)
}

// Int16Slicef checks two int16 slices for equality.
func (chk *Chk) Int16Slicef(
	got, want []int16, msgFmt string, msgArgs ...any,
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
		got, want, int16TypeName, defaultCmpFunc[int16], msgFmt, msgArgs...,
	)
}

// Int16Slice checks two int16 slices for equality.
func (chk *Chk) Int16Slice(got, want []int16, msg ...any) bool {
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
		chk, got, want, int16TypeName, defaultCmpFunc[int16], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// Int16Boundedf checks value is within specified bounded range.
func (chk *Chk) Int16Boundedf(
	got int16, option BoundedOption, min, max int16,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int16TypeName, got, want, msgFmt, msgArgs...)
}

// Int16Bounded checks value is within specified bounded range.
func (chk *Chk) Int16Bounded(
	got int16, option BoundedOption, min, max int16, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int16TypeName, got, want, msg...)
}

// Int16Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Int16Unboundedf(
	got int16, option UnboundedOption, bound int16,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int16TypeName, got, want, msgFmt, msgArgs...)
}

// Int16Unbounded checks value is within specified unbounded range.
func (chk *Chk) Int16Unbounded(
	got int16, option UnboundedOption, bound int16, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int16TypeName, got, want, msg...)
}
