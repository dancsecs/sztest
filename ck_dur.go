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

import (
	"time"
)

const durTypeName = "time.Duration"

// Durf compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Durf(
	got, want time.Duration, msgFmt string, msgArgs ...any,
) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(got, want, durTypeName, msgFmt, msgArgs...)
}

// Dur compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Dur(got, want time.Duration, msg ...any) bool {
	if got == want {
		return true
	}

	chk.t.Helper()

	return chk.errChk(got, want, durTypeName, msg...)
}

// DurSlicef checks two time.Duration slices for equality.
func (chk *Chk) DurSlicef(
	got, want []time.Duration, msgFmt string, msgArgs ...any,
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
		got, want, durTypeName, defaultCmpFunc[time.Duration],
		msgFmt, msgArgs...,
	)
}

// DurSlice checks two time.Duration slices for equality.
func (chk *Chk) DurSlice(got, want []time.Duration, msg ...any) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = got[i] == want[i]
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk,
		got, want, durTypeName, defaultCmpFunc[time.Duration], msg...,
	)
}

//
// Bounded and Unbounded Ranges.
//

// DurBoundedf checks value is within specified bounded range.
func (chk *Chk) DurBoundedf(
	got time.Duration, option BoundedOption, min, max time.Duration,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(durTypeName, got, want, msgFmt, msgArgs...)
}

// DurBounded checks value is within specified bounded range.
func (chk *Chk) DurBounded(
	got time.Duration, option BoundedOption, min, max time.Duration, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(durTypeName, got, want, msg...)
}

// DurUnboundedf checks value is within specified unbounded range.
func (chk *Chk) DurUnboundedf(
	got time.Duration, option UnboundedOption, bound time.Duration,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(durTypeName, got, want, msgFmt, msgArgs...)
}

// DurUnbounded checks value is within specified unbounded range.
func (chk *Chk) DurUnbounded(
	got time.Duration, option UnboundedOption, bound time.Duration, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(durTypeName, got, want, msg...)
}
