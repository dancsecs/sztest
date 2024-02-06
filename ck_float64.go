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
	"strconv"
)

func float64TypeString(t float64) string {
	if t == 0.0 {
		return "float64"
	}
	return "float64(+/- " +
		strconv.FormatFloat(t, 'g', -1, 64) +
		")"
}

// Float64f compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Float64f(
	got, want, tolerance float64, msgFmt string, msgArgs ...any,
) bool {
	if IsFloat64Similar(got, want, tolerance) {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(
		got, want, float64TypeString(tolerance), msgFmt, msgArgs...,
	)
}

// Float64 compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Float64(
	got, want, tolerance float64, msg ...any,
) bool {
	if IsFloat64Similar(got, want, tolerance) {
		return true
	}
	chk.t.Helper()
	return chk.errChk(
		got, want, float64TypeString(tolerance), msg...,
	)
}

// Float64Slicef compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Float64Slicef(
	got, want []float64, tolerance float64, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = IsFloat64Similar(got[i], want[i], tolerance)
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlicef(chk,
		got, want, float64TypeString(tolerance),
		func(a, b float64) bool {
			return IsFloat64Similar(a, b, tolerance)
		},
		msgFmt, msgArgs...,
	)
}

// Float64Slice compare the wanted boolean against the gotten bool invoking an
// error should they not match.
func (chk *Chk) Float64Slice(
	got, want []float64, tolerance float64, msg ...any,
) bool {
	l := len(got)
	equal := l == len(want)
	for i := 0; equal && i < l; i++ {
		equal = IsFloat64Similar(got[i], want[i], tolerance)
	}
	if equal {
		return true
	}
	chk.t.Helper()
	return errSlice(chk,
		got, want, float64TypeString(tolerance),
		func(a, b float64) bool {
			return IsFloat64Similar(a, b, tolerance)
		},
		msg...,
	)
}

// IsFloat64Similar compares two floats to see if they match within the
// specified tolerance.
func IsFloat64Similar(a, b, t float64) bool {
	if t == 0.0 {
		return a == b
	}
	// Are a and b within tolerance t
	switch {
	case a < b:
		return (b - a) <= t
	case a > b:
		return (a - b) <= t
	}
	return true // a == b
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// Float64Boundedf checks value is within specified bounded range.
func (chk *Chk) Float64Boundedf(
	got float64, option BoundedOption, min, max float64,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "float64"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Float64Bounded checks value is within specified bounded range.
func (chk *Chk) Float64Bounded(
	got float64, option BoundedOption, min, max float64, msg ...any,
) bool {
	const typeName = "float64"
	inRange, want := inBoundedRange(got, option, min, max)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}

// Float64Unboundedf checks value is within specified unbounded range.
func (chk *Chk) Float64Unboundedf(
	got float64, option UnboundedOption, bound float64,
	msgFmt string, msgArgs ...any,
) bool {
	const typeName = "float64"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWntf(typeName, got, want, msgFmt, msgArgs...)
}

// Float64Unbounded checks value is within specified unbounded range.
func (chk *Chk) Float64Unbounded(
	got float64, option UnboundedOption, bound float64, msg ...any,
) bool {
	const typeName = "float64"
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}
	chk.t.Helper()
	return chk.errGotWnt(typeName, got, want, msg...)
}
