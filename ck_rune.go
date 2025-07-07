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

const runeTypeName = "rune"

// Runef compares the wanted rune against the gotten rune invoking an
// error should they not match.
func (chk *Chk) Runef(got, want rune, msgFmt string, msgArgs ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, runeTypeName, msgFmt, msgArgs...)
}

// Rune compares the wanted rune against the gotten rune invoking an
// error should they not match.
func (chk *Chk) Rune(got, want rune, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, runeTypeName, msg...)
}

// RuneSlicef checks two rune slices for equality.
func (chk *Chk) RuneSlicef(
	got, want []rune, msgFmt string, msgArgs ...any,
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
		got, want, runeTypeName, defaultCmpFunc[rune], msgFmt, msgArgs...,
	)
}

// RuneSlice checks two rune slices for equality.
func (chk *Chk) RuneSlice(got, want []rune, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk, got, want, runeTypeName, defaultCmpFunc[rune], msg...)
}

//
// Bounded and Unbounded Ranges.
//

// RuneBoundedf checks value is within specified bounded range.
func (chk *Chk) RuneBoundedf(
	got rune,
	option BoundedOption,
	minV, maxV rune,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(runeTypeName, got, want, msgFmt, msgArgs...)
}

// RuneBounded checks value is within specified bounded range.
func (chk *Chk) RuneBounded(
	got rune, option BoundedOption, minV, maxV rune, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(runeTypeName, got, want, msg...)
}

// RuneUnboundedf checks value is within specified unbounded range.
func (chk *Chk) RuneUnboundedf(
	got rune,
	option UnboundedOption,
	bound rune,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(runeTypeName, got, want, msgFmt, msgArgs...)
}

// RuneUnbounded checks value is within specified unbounded range.
func (chk *Chk) RuneUnbounded(
	got rune, option UnboundedOption, bound rune, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(runeTypeName, got, want, msg...)
}
