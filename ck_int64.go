/*
   Golang test helper library: sztest.
   Copyright (C) 2023-2025 Leslie Dancsecs

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

// Int64f compares the got int64 against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Int64f(got, want int64, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, int64TypeName, msgFmt, msgArgs...)
}

// Int64 compares the got int64 against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Int64(got, want int64, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, int64TypeName, msg...)
}

// Int64Slicef compares two int64 slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
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

// Int64Slice compares two int64 slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
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

// Int64Boundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Int64Boundedf(
	got int64, option BoundedOption, minV, maxV int64,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(int64TypeName, got, want, msgFmt, msgArgs...)
}

// Int64Bounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) Int64Bounded(
	got int64, option BoundedOption, minV, maxV int64, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(int64TypeName, got, want, msg...)
}

// Int64Unboundedf checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
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

// Int64Unbounded checks that got lies within the unbounded interval defined by
// bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
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
