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

const uint32TypeName = "uint32"

// Uint32f compares the got uint32 against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Uint32f(
	got, want uint32, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, uint32TypeName, msgFmt, msgArgs...)
}

// Uint32 compares the got uint32 against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Uint32(got, want uint32, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, uint32TypeName, msg...)
}

// Uint32Slicef compares two uint32 slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
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
		got, want, uint32TypeName, defaultCmpFunc[uint32],
		msgFmt, msgArgs...,
	)
}

// Uint32Slice compares two uint32 slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
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

	return errSlice(
		chk, got, want, uint32TypeName, defaultCmpFunc[uint32], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// Uint32Boundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Uint32Boundedf(
	got uint32, option BoundedOption, minV, maxV uint32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint32TypeName, got, want, msgFmt, msgArgs...)
}

// Uint32Bounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) Uint32Bounded(
	got uint32, option BoundedOption, minV, maxV uint32, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint32TypeName, got, want, msg...)
}

// Uint32Unboundedf checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Uint32Unboundedf(
	got uint32, option UnboundedOption, bound uint32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint32TypeName, got, want, msgFmt, msgArgs...)
}

// Uint32Unbounded checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) Uint32Unbounded(
	got uint32, option UnboundedOption, bound uint32, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint32TypeName, got, want, msg...)
}
