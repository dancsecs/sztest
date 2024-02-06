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

func (chk *Chk) strPrepareSlice(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = chk.isStringify(v)
	}
	return r
}

// Strf compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Strf(got, want string, msgFmt string, msgArgs ...any) bool {
	if chk.isStringify(got) == chk.isStringify(want) {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(got, want, "string", msgFmt, msgArgs...)
}

// Str compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Str(got, want string, msg ...any) bool {
	if chk.isStringify(got) == chk.isStringify(want) {
		return true
	}
	chk.t.Helper()
	return chk.errChk(got, want, "string", msg...)
}

// StrSlicef checks two string slices for equality.
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
		got, want, "string", defaultCmpFunc[string],
		msgFmt, msgArgs...,
	)
}

// StrSlice checks two string slices for equality.
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
		"string", defaultCmpFunc[string], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// StrBoundedf checks value is within specified bounded range.
func (chk *Chk) StrBoundedf(
	got string, option BoundedOption, min, max string,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "string"
	got = chk.isStringify(got)
	min = chk.isStringify(min)
	max = chk.isStringify(max)

	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// StrBounded checks value is within specified bounded range.
func (chk *Chk) StrBounded(
	got string, option BoundedOption, min, max string, msg ...any,
) bool {
	const typeName = "string"
	got = chk.isStringify(got)
	min = chk.isStringify(min)
	max = chk.isStringify(max)

	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// StrUnboundedf checks value is within specified unbounded range.
func (chk *Chk) StrUnboundedf(
	got string, option UnboundedOption, bound string,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "string"
	got = chk.isStringify(got)
	bound = chk.isStringify(bound)

	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// StrUnbounded checks value is within specified unbounded range.
func (chk *Chk) StrUnbounded(
	got string, option UnboundedOption, bound string, msg ...any,
) bool {
	const typeName = "string"
	got = chk.isStringify(got)
	bound = chk.isStringify(bound)

	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
