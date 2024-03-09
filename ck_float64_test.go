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
	"testing"
)

func tstChkFloat64(t *testing.T) {
	t.Run("Good", chkFloat64TestGood)

	t.Run("Bad", chkFloat64TestBad)
	t.Run("BadMsg1", chkFloat64TestBad1)
	t.Run("BadMsg2", chkFloat64TestBad2)
	t.Run("BadMsg3", chkFloat64TestBad3)

	t.Run("Slice-Good", chkFloat64SliceTestGood)
	t.Run("Slice-BadMsg1", chkFloat64SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkFloat64SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkFloat64SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkFloat64SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkFloat64SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkFloat64SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkFloat64SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkFloat64SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkFloat64SliceTestBadMsg9)

	t.Run("Bounded", chkFloat64BoundedTestAll)
	t.Run("Unbounded", chkFloat64UnboundedTestAll)
}

func chkFloat64TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64(0.0, 0.0, 0.0)
	chk.Float64(0.0, 0.0, 0.0, "not ", "displayed")
	chk.Float64f(0.0, 0.0, 0.0, "not %s", "displayed")

	chk.Float64(0.02, 0.01, 0.1)
	chk.Float64(0.02, 0.01, 0.1, "not ", "displayed")
	chk.Float64f(0.02, 0.01, 0.1, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat64TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64(0, -0.01, 0.005)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float64",
			chkOutCommonMsg("", float64TypeString(0.005)),
			g(markAsChg("0", "-0.01", DiffGot)),
			w(markAsChg("0", "-0.01", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat64TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64f(
		0.02, 0.01, 0.005, "This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float64f",
			chkOutCommonMsg(
				"This message will be displayed first",
				float64TypeString(0.005),
			),
			g("0.0"+markAsChg("2", "1", DiffGot)),
			w("0.0"+markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat64TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64f(
		-0.02, -0.01, 0.005,
		"This message will be displayed %s", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float64f",
			chkOutCommonMsg(
				"This message will be displayed second",
				float64TypeString(0.005),
			),
			g("-0.0"+markAsChg("2", "1", DiffGot)),
			w("-0.0"+markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat64TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64(0, -0.01, 0.005, "This message will be displayed ", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float64",
			chkOutCommonMsg(
				"This message will be displayed third",
				float64TypeString(0.005),
			),
			g(markAsChg("0", "-0.01", DiffGot)),
			w(markAsChg("0", "-0.01", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice(
		[]float64{}, []float64{}, 0.0,
		"This message will NOT be displayed",
	)
	chk.Float64Slice(
		[]float64{0.0}, []float64{0.0}, 0.0,
		"This message will NOT be displayed",
	)
	chk.Float64Slice(
		[]float64{0.02}, []float64{0.01}, 0.1,
		"This message will NOT be displayed",
	)
	chk.Float64Slice(
		[]float64{0.02, 0.06}, []float64{0.01, 0.05}, 0.1,
		"This message will NOT be displayed",
	)
	chk.Float64Slice(
		[]float64{0.02, 0.06, -.07}, []float64{0.01, 0.05, -0.08}, 0.1,
		"This message will NOT be displayed",
	)

	chk.Float64Slicef(
		[]float64{0.02, 0.06, -.07, 9.0},
		[]float64{0.01, 0.05, -0.08, 9.0},
		0.1,
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slicef(
		[]float64{0.02}, []float64{}, 0.0,
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]float64", "Float64Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice(
		[]float64{}, []float64{0.01}, 0.0,
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]float64", "Float64Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "0.01"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slicef(
		[]float64{0.01, 0.02}, []float64{0.01}, 0.0,
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]float64", "Float64Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnGot("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice(
		[]float64{0.01}, []float64{0.01, 0.02}, 0.0,
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]float64", "Float64Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnWnt("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slicef(
		[]float64{0.01, 0.02}, []float64{0.02, 0.02}, 0.005,
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float64TypeString(0.005), "Float64Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "0.01"),
			chkOutLnSame("1", "0", "0.02"),
			chkOutLnWnt("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice(
		[]float64{0.02, 0.02}, []float64{0.01, 0.02}, 0.005,
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float64TypeString(0.005), "Float64Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "0.01"),
			chkOutLnSame("0", "1", "0.02"),
			chkOutLnGot("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slicef(
		[]float64{0.01, 0.02}, []float64{0.01, 0.03}, 0.005,
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float64TypeString(0.005), "Float64Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnChanged("1", "1", "0.0"+markAsChg("2", "3", DiffMerge)),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice(
		[]float64{0.01, 0.03}, []float64{0.01, 0.02}, 0.005,
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float64TypeString(0.005), "Float64Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnChanged("1", "1", "0.0"+markAsChg("3", "2", DiffMerge)),
		),
		chkOutRelease(),
	)
}

func chkFloat64SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float64Slice([]float64{0.01, 0.03}, []float64{0.01, 0.02}, 0.005)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float64TypeString(0.005), "Float64Slice",
			"",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnChanged("1", "1", "0.0"+markAsChg("3", "2", DiffMerge)),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkFloat64BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := float64(33)
	max := float64(35)

	// Bad: Error displayed.
	chk.Float64Bounded(30, BoundedClosed, min, max)
	chk.Float64Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Float64Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Float64Bounded(33, BoundedClosed, min, max)
	chk.Float64Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Float64Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Float64Bounded(36, BoundedClosed, min, max)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Float64Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, float64TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, float64TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, float64TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, float64TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkFloat64UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := float64(128)

	// Bad: Error displayed.
	chk.Float64Unbounded(125, UnboundedMinClosed, bound)
	chk.Float64Unbounded(126, UnboundedMinClosed, bound, "msg:", "126")
	chk.Float64Unboundedf(127, UnboundedMinClosed, bound, "msg:%d", 127)

	// Good:  No error displayed.
	chk.Float64Unbounded(128, UnboundedMinClosed, bound)
	chk.Float64Unbounded(129, UnboundedMinClosed, bound, "not ", "displayed")
	chk.Float64Unboundedf(
		130, UnboundedMinClosed, bound, "not %s", "displayed",
	)

	const (
		wntMsg = "out of bounds: [128,MAX) - { want | want >= 128 }"
		fName  = "Float64Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "125", fName, float64TypeName, ""),
		chkOutNumericUnbounded(
			wntMsg, "126", fName, float64TypeName, "msg:126",
		),
		chkOutNumericUnboundedf(
			wntMsg, "127", fName, float64TypeName, "msg:127",
		),

		chkOutRelease(),
	)
}
