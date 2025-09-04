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

const byteTypeName = "byte"

// Bytef compares the got byte against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Bytef(got, want byte, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, byteTypeName, msgFmt, msgArgs...)
}

// Byte compares the got byte against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Byte(got, want byte, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, byteTypeName, msg...)
}

// ByteSlicef compares two byte slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
func (chk *Chk) ByteSlicef(
	got, want []byte, msgFmt string, msgArgs ...any,
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
		got, want, byteTypeName, defaultCmpFunc[byte], msgFmt, msgArgs...,
	)
}

// ByteSlice compares two byte slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
func (chk *Chk) ByteSlice(got, want []byte, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk, got, want, byteTypeName, defaultCmpFunc[byte], msg...)
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// ByteBoundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) ByteBoundedf(
	got byte,
	option BoundedOption,
	minV, maxV byte,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)

	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(byteTypeName, got, want, msgFmt, msgArgs...)
}

// ByteBounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) ByteBounded(
	got byte, option BoundedOption, minV, maxV byte, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)

	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(byteTypeName, got, want, msg...)
}

// ByteUnboundedf checks that got lies within the unbounded interval defined by
// bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) ByteUnboundedf(
	got byte,
	option UnboundedOption,
	bound byte,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)

	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(byteTypeName, got, want, msgFmt, msgArgs...)
}

// ByteUnbounded checks that got lies within the unbounded interval defined by
// bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) ByteUnbounded(
	got byte, option UnboundedOption, bound byte, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)

	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(byteTypeName, got, want, msg...)
}
