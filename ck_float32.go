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

// Float32f compare the wanted boolean against the gotten bool invoking an
// error should they not match.
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

// Float32 compares the wanted boolean against the gotten bool invoking an
// error should they not match.
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

// Float32Slicef compare the wanted boolean against the gotten bool invoking an
// error should they not match.
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

// Float32Slice compare the wanted boolean against the gotten bool invoking an
// error should they not match.
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

	if tolerance == 0.0 {
		return num1 == num2
	}

	// Are a and b within tolerance t
	switch {
	case num1 < num2:
		return (num2 - num1) <= tolerance
	case num1 > num2:
		return (num1 - num2) <= tolerance
	}

	return true // a == b
}

////////////////////////////////////////////////////////////////
// Bounded and Unbounded Ranges.
////////////////////////////////////////////////////////////////

// Float32Boundedf checks value is within specified bounded range.
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

// Float32Bounded checks value is within specified bounded range.
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

// Float32Unboundedf checks value is within specified unbounded range.
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

// Float32Unbounded checks value is within specified unbounded range.
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
