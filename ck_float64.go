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

import (
	"math"
	"strconv"
)

const float64TypeName = "float64"

func float64TypeString(tolerance float64) string {
	if tolerance == 0.0 {
		return float64TypeName
	}

	return float64TypeName + "(+/- " +
		strconv.FormatFloat(tolerance, 'g', -1, 64) +
		")"
}

// Float64f compares got and want within the given tolerance.
//
// The values are considered equal if |got - want| <= tolerance. A tolerance of
// 0.0 requires exact equality. On mismatch, the failure is reported with a
// formatted message built from msgFmt and msgArgs. Returns true if the
// comparison succeeds.
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

// Float64 compares got and want within the given tolerance.
//
// The values are considered equal if |got - want| <= tolerance. A tolerance of
// 0.0 requires exact equality. On mismatch, the failure is reported to the
// underlying testingT and the optional msg values are appended. Returns true
// if the comparison succeeds.
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

// Float64Slicef compares two float64 slices element-wise within the given
// tolerance.
//
// Each pair of elements must satisfy |got[i] - want[i]| <= tolerance. A
// tolerance of 0.0 requires exact equality. Length mismatches or element
// mismatches are reported to the underlying testingT with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are equal within
// tolerance.
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

// Float64Slice compares two float64 slices element-wise within the given
// tolerance.
//
// Each pair of elements must satisfy |got[i] - want[i]| <= tolerance. A
// tolerance of 0.0 requires exact equality. Length mismatches or element
// mismatches are reported to the underlying testingT. Optional msg values are
// included in the failure output. Returns true if slices are equal within
// tolerance.
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
func IsFloat64Similar(num1, num2, tolerance float64) bool {
	if math.IsNaN(num1) && math.IsNaN(num2) {
		return true
	}

	return math.Abs(num1-num2) <= tolerance
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// Float64Boundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Float64Boundedf(
	got float64, option BoundedOption, minV, maxV float64,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(float64TypeName, got, want, msgFmt, msgArgs...)
}

// Float64Bounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) Float64Bounded(
	got float64, option BoundedOption, minV, maxV float64, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(float64TypeName, got, want, msg...)
}

// Float64Unboundedf checks that got lies within the unbounded interval
// defined by bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Float64Unboundedf(
	got float64, option UnboundedOption, bound float64,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(float64TypeName, got, want, msgFmt, msgArgs...)
}

// Float64Unbounded checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) Float64Unbounded(
	got float64, option UnboundedOption, bound float64, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(float64TypeName, got, want, msg...)
}
