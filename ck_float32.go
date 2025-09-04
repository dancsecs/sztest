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

const float32TypeName = "float32"

func float32TypeString(tolerance float32) string {
	if tolerance == 0.0 {
		return float32TypeName
	}

	return float32TypeName + "(+/- " +
		strconv.FormatFloat(float64(tolerance), 'g', -1, 64) +
		")"
}

// Float32f compares got and want within the given tolerance.
//
// The values are considered equal if |got - want| <= tolerance. A tolerance of
// 0.0 requires exact equality. On mismatch, the failure is reported with a
// formatted message built from msgFmt and msgArgs. Returns true if the
// comparison succeeds.
func (chk *Chk) Float32f(
	got, want, tolerance float32, msgFmt string, msgArgs ...any,
) bool {
	if IsFloat32Similar(got, want, tolerance) {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(
		got, want, float32TypeString(tolerance), msgFmt, msgArgs...,
	)
}

// Float32 compares got and want within the given tolerance.
//
// The values are considered equal if |got - want| <= tolerance. A tolerance of
// 0.0 requires exact equality. On mismatch, the failure is reported to the
// underlying testingT and the optional msg values are appended. Returns true
// if the comparison succeeds.
func (chk *Chk) Float32(
	got, want, tolerance float32, msg ...any,
) bool {
	if IsFloat32Similar(got, want, tolerance) {
		return true
	}

	chk.t.Helper()

	return chk.errChk(
		got, want, float32TypeString(tolerance),
		msg...,
	)
}

// Float32Slicef compares two float64 slices element-wise within the given
// tolerance.
//
// Each pair of elements must satisfy |got[i] - want[i]| <= tolerance. A
// tolerance of 0.0 requires exact equality. Length mismatches or element
// mismatches are reported to the underlying testingT with a formatted message
// built from msgFmt and msgArgs. Returns true if slices are equal within
// tolerance.
func (chk *Chk) Float32Slicef(
	got, want []float32, tolerance float32, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = IsFloat32Similar(got[i], want[i], tolerance)
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlicef(chk,
		got, want, float32TypeString(tolerance),
		func(a, b float32) bool {
			return IsFloat32Similar(a, b, tolerance)
		},
		msgFmt, msgArgs...,
	)
}

// Float32Slice compares two float64 slices element-wise within the given
// tolerance.
//
// Each pair of elements must satisfy |got[i] - want[i]| <= tolerance. A
// tolerance of 0.0 requires exact equality. Length mismatches or element
// mismatches are reported to the underlying testingT. Optional msg values are
// included in the failure output. Returns true if slices are equal within
// tolerance.
func (chk *Chk) Float32Slice(
	got, want []float32, tolerance float32, msg ...any,
) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = IsFloat32Similar(got[i], want[i], tolerance)
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk,
		got, want, float32TypeString(tolerance),
		func(a, b float32) bool {
			return IsFloat32Similar(a, b, tolerance)
		},
		msg...,
	)
}

// IsFloat32Similar compares two floats to see if they match within the
// specified tolerance.
func IsFloat32Similar(num1, num2, tolerance float32) bool {
	if math.IsNaN(float64(num1)) && math.IsNaN(float64(num2)) {
		return true
	}

	return math.Abs(float64(num1)-float64(num2)) <= float64(tolerance)
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// Float32Boundedf checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Float32Boundedf(
	got float32, option BoundedOption, minV, maxV float32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(float32TypeName, got, want, msgFmt, msgArgs...)
}

// Float32Bounded checks that got lies within the bounded interval defined by
// minV and maxV according to the chosen option.
//
// On failure, the test is reported with the optional msg values appended.
// Returns true if got is within bounds.
func (chk *Chk) Float32Bounded(
	got float32, option BoundedOption, minV, maxV float32, msg ...any,
) bool {
	inRange, want := inBoundedRange(got, option, minV, maxV)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(float32TypeName, got, want, msg...)
}

// Float32Unboundedf checks that got lies within the unbounded interval
// defined by bound and option.
//
// On failure, the test is reported with a formatted message built from msgFmt
// and msgArgs. Returns true if got is within bounds.
func (chk *Chk) Float32Unboundedf(
	got float32, option UnboundedOption, bound float32,
	msgFmt string, msgArgs ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWntf(float32TypeName, got, want, msgFmt, msgArgs...)
}

// Float32Unbounded checks that got lies within the unbounded interval defined
// by bound and option.
//
// On failure, the test is reported with optional msg values appended. Returns
// true if got is within bounds.
func (chk *Chk) Float32Unbounded(
	got float32, option UnboundedOption, bound float32, msg ...any,
) bool {
	inRange, want := inUnboundedRange(got, option, bound)
	if inRange {
		return true
	}

	chk.t.Helper()

	return chk.errGotWnt(float32TypeName, got, want, msg...)
}
