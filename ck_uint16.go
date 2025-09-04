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

const uint16TypeName = "uint16"

// Uint16f compares the got uint16 against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Uint16f(
	got, want uint16, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, uint16TypeName, msgFmt, msgArgs...)
}

// Uint16 compares the got uint16 against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Uint16(got, want uint16, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, uint16TypeName, msg...)
}

// Uint16Slicef compares two uint16 slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
func (chk *Chk) Uint16Slicef(
	got, want []uint16, msgFmt string, msgArgs ...any,
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
		got, want, uint16TypeName, defaultCmpFunc[uint16],
		msgFmt, msgArgs...,
	)
}

// Uint16Slice compares two uint16 slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
func (chk *Chk) Uint16Slice(got, want []uint16, msg ...any) bool {
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
		chk, got, want, uint16TypeName, defaultCmpFunc[uint16], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// Uint16Boundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Uint16Boundedf(
	got uint16, option BoundedOption, minV, maxV uint16,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint16TypeName, got, want, msgFmt, msgArgs...)
}

// Uint16Bounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) Uint16Bounded(
	got uint16, option BoundedOption, minV, maxV uint16, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint16TypeName, got, want, msg...)
}

// Uint16Unboundedf checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Uint16Unboundedf(
	got uint16, option UnboundedOption, bound uint16,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(uint16TypeName, got, want, msgFmt, msgArgs...)
}

// Uint16Unbounded checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) Uint16Unbounded(
	got uint16, option UnboundedOption, bound uint16, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(uint16TypeName, got, want, msg...)
}
