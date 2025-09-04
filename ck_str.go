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

const stringTypeName = "string"

func (chk *Chk) strPrepareSlice(lines []string) []string {
	if len(lines) == 0 {
		return nil
	}

	result := make([]string, len(lines))

	for i, v := range lines {
		result[i] = chk.isStringify(v)
	}

	return result
}

// Strf compares the got string against want.
//
// If they differ, the failure is reported with a formatted message built from
// msgFmt and msgArgs. Returns true if got == want.
func (chk *Chk) Strf(got, want string, msgFmt string, msgArgs ...any) bool {
	if chk.isStringify(got) == chk.isStringify(want) {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, stringTypeName, msgFmt, msgArgs...)
}

// Str compares the got string against want.
//
// If they differ, the failure is reported via the underlying testingT and the
// optional msg values are formatted and appended to the report. Returns true
// if got == want.
func (chk *Chk) Str(got, want string, msg ...any) bool {
	if chk.isStringify(got) == chk.isStringify(want) {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, stringTypeName, msg...)
}

// StrSlicef compares two string slices for equality.
//
// A mismatch is reported to the underlying test with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are exactly equal.
func (chk *Chk) StrSlicef(
	got, want []string, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = chk.isStringify(got[i]) == chk.isStringify(want[i])
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlicef(chk,
		got, want, stringTypeName, defaultCmpFunc[string],
		msgFmt, msgArgs...,
	)
}

// StrSlice compares two string slices for equality.
//
// A mismatch in length or element values is reported to the underlying test.
// Optional msg values are included in the failure output. Returns true if
// slices are exactly equal.
func (chk *Chk) StrSlice(got, want []string, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = chk.isStringify(got[i]) == chk.isStringify(want[i])
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(
		chk, chk.strPrepareSlice(got), chk.strPrepareSlice(want),
		stringTypeName, defaultCmpFunc[string], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// StrBoundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) StrBoundedf(
	got string, option BoundedOption, minV, maxV string,
	msgFmt string, msgArgs ...any,
) bool {
	got = chk.isStringify(got)
	minV = chk.isStringify(minV)
	maxV = chk.isStringify(maxV)

	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(stringTypeName, got, want, msgFmt, msgArgs...)
}

// StrBounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) StrBounded(
	got string, option BoundedOption, minV, maxV string, msg ...any,
) bool {
	got = chk.isStringify(got)
	minV = chk.isStringify(minV)
	maxV = chk.isStringify(maxV)

	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(stringTypeName, got, want, msg...)
}

// StrUnboundedf checks that got lies within the unbounded interval defined by
// bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) StrUnboundedf(
	got string, option UnboundedOption, bound string,
	msgFmt string, msgArgs ...any,
) bool {
	got = chk.isStringify(got)
	bound = chk.isStringify(bound)

	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(stringTypeName, got, want, msgFmt, msgArgs...)
}

// StrUnbounded checks that got lies within the unbounded interval defined by
// bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) StrUnbounded(
	got string, option UnboundedOption, bound string, msg ...any,
) bool {
	got = chk.isStringify(got)
	bound = chk.isStringify(bound)

	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(stringTypeName, got, want, msg...)
}
