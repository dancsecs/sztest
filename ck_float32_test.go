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
	"testing"
)

func tstChkFloat32(t *testing.T) {
	t.Run("Good", chkFloat32TestGood)

	t.Run("Bad", chkFloat32TestBad)
	t.Run("BadMsg1", chkFloat32TestBad1)
	t.Run("BadMsg2", chkFloat32TestBad2)
	t.Run("BadMsg3", chkFloat32TestBad3)

	t.Run("Slice-Good", chkFloat32SliceTestGood)
	t.Run("Slice-BadMsg1", chkFloat32SliceTestBadMsg1)
	t.Run("Slice-BadMsg2", chkFloat32SliceTestBadMsg2)
	t.Run("Slice-BadMsg3", chkFloat32SliceTestBadMsg3)
	t.Run("Slice-BadMsg4", chkFloat32SliceTestBadMsg4)
	t.Run("Slice-BadMsg5", chkFloat32SliceTestBadMsg5)
	t.Run("Slice-BadMsg6", chkFloat32SliceTestBadMsg6)
	t.Run("Slice-BadMsg7", chkFloat32SliceTestBadMsg7)
	t.Run("Slice-BadMsg8", chkFloat32SliceTestBadMsg8)
	t.Run("Slice-BadMsg9", chkFloat32SliceTestBadMsg9)

	t.Run("Bounded", chkFloat32BoundedTestAll)
	t.Run("Unbounded", chkFloat32UnboundedTestAll)
}

func chkFloat32TestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32(0.0, 0.0, 0.0)
	chk.Float32(0.0, 0.0, 0.0, "not ", "displayed")
	chk.Float32f(0.0, 0.0, 0.0, "not %s", "displayed")

	chk.Float32(0.02, 0.01, 0.1)
	chk.Float32(0.02, 0.01, 0.1, "not ", "displayed")
	chk.Float32f(0.02, 0.01, 0.1, "not %s", "displayed")

	chk.Float32f(float32(math.NaN()), float32(math.NaN()),
		0, "not %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat32TestBad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32(0, -0.01, 0.005)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float32",
			chkOutCommonMsg("", float32TypeString(0.005)),
			g(markAsChg("0", "-0.01", DiffGot)),
			w(markAsChg("0", "-0.01", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat32TestBad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32f(
		0.02, 0.01, 0.005, "This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float32f",
			chkOutCommonMsg(
				"This message will be displayed first",
				float32TypeString(0.005),
			),
			g("0.0"+markAsChg("2", "1", DiffGot)),
			w("0.0"+markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat32TestBad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32f(
		-0.02, -0.01, 0.005,
		"This message will be displayed %s", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float32f",
			chkOutCommonMsg(
				"This message will be displayed second",
				float32TypeString(0.005),
			),
			g("-0.0"+markAsChg("2", "1", DiffGot)),
			w("-0.0"+markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat32TestBad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32(0, -0.01, 0.005, "This message will be displayed ", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Float32",
			chkOutCommonMsg(
				"This message will be displayed third",
				float32TypeString(0.005),
			),
			g(markAsChg("0", "-0.01", DiffGot)),
			w(markAsChg("0", "-0.01", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestGood(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice(
		[]float32{}, []float32{}, 0.0,
		"This message will NOT be displayed",
	)
	chk.Float32Slice(
		[]float32{0.0}, []float32{0.0}, 0.0,
		"This message will NOT be displayed",
	)
	chk.Float32Slice(
		[]float32{0.02}, []float32{0.01}, 0.1,
		"This message will NOT be displayed",
	)
	chk.Float32Slice(
		[]float32{0.02, 0.06}, []float32{0.01, 0.05}, 0.1,
		"This message will NOT be displayed",
	)
	chk.Float32Slice(
		[]float32{0.02, 0.06, -.07}, []float32{0.01, 0.05, -0.08}, 0.1,
		"This message will NOT be displayed",
	)

	chk.Float32Slicef(
		[]float32{0.02, 0.06, -.07, 9.0},
		[]float32{0.01, 0.05, -0.08, 9.0}, 0.1,
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slicef(
		[]float32{0.02}, []float32{}, 0.0,
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]float32", "Float32Slicef",
			"This message will be displayed first",
			chkOutLnGot("0", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice(
		[]float32{}, []float32{0.01}, 0.0,
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]float32", "Float32Slice",
			"This message will be displayed second",
			chkOutLnWnt("0", "0.01"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slicef(
		[]float32{0.01, 0.02}, []float32{0.01}, 0.0,
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]float32", "Float32Slicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnGot("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice(
		[]float32{0.01}, []float32{0.01, 0.02}, 0.0,
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]float32", "Float32Slice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnWnt("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slicef(
		[]float32{0.01, 0.02}, []float32{0.02, 0.02}, 0.005,
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float32TypeString(0.005), "Float32Slicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "0.01"),
			chkOutLnSame("1", "0", "0.02"),
			chkOutLnWnt("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice(
		[]float32{0.02, 0.02}, []float32{0.01, 0.02}, 0.005,
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float32TypeString(0.005), "Float32Slice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "0.01"),
			chkOutLnSame("0", "1", "0.02"),
			chkOutLnGot("1", "0.02"),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slicef(
		[]float32{0.01, 0.02}, []float32{0.01, 0.03}, 0.005,
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float32TypeString(0.005), "Float32Slicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnChanged("1", "1", "0.0"+markAsChg("2", "3", DiffMerge)),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice(
		[]float32{0.01, 0.03}, []float32{0.01, 0.02}, 0.005,
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float32TypeString(0.005), "Float32Slice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "0.01"),
			chkOutLnChanged("1", "1", "0.0"+markAsChg("3", "2", DiffMerge)),
		),
		chkOutRelease(),
	)
}

func chkFloat32SliceTestBadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32Slice([]float32{0.01, 0.03}, []float32{0.01, 0.02}, 0.005)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]"+float32TypeString(0.005), "Float32Slice",
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

func chkFloat32BoundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	minV := float32(33)
	maxV := float32(35)

	// Bad: Error displayed.
	chk.Float32Bounded(30, BoundedClosed, minV, maxV)
	chk.Float32Bounded(31, BoundedClosed, minV, maxV, "msg:", "31")
	chk.Float32Boundedf(32, BoundedClosed, minV, maxV, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Float32Bounded(33, BoundedClosed, minV, maxV)
	chk.Float32Bounded(34, BoundedClosed, minV, maxV, "not ", "displayed")
	chk.Float32Boundedf(35, BoundedClosed, minV, maxV, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Float32Bounded(36, BoundedClosed, minV, maxV)

	const (
		wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
		fName  = "Float32Bounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded(wntMsg, "30", fName, float32TypeName, ""),
		chkOutNumericBounded(wntMsg, "31", fName, float32TypeName, "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, float32TypeName, "msg:32"),

		chkOutNumericBounded(wntMsg, "36", fName, float32TypeName, ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkFloat32UnboundedTestAll(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := float32(128)

	// Bad: Error displayed.
	chk.Float32Unbounded(125, UnboundedMinClosed, bound)
	chk.Float32Unbounded(126, UnboundedMinClosed, bound, "msg:", "126")
	chk.Float32Unboundedf(127, UnboundedMinClosed, bound, "msg:%d", 127)

	// Good:  No error displayed.
	chk.Float32Unbounded(128, UnboundedMinClosed, bound)
	chk.Float32Unbounded(129, UnboundedMinClosed, bound, "not ", "displayed")
	chk.Float32Unboundedf(
		130, UnboundedMinClosed, bound, "not %s", "displayed",
	)

	const (
		wntMsg = "out of bounds: [128,MAX) - { want | want >= 128 }"
		fName  = "Float32Unbounded"
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded(wntMsg, "125", fName, float32TypeName, ""),
		chkOutNumericUnbounded(
			wntMsg, "126", fName, float32TypeName, "msg:126",
		),
		chkOutNumericUnboundedf(
			wntMsg, "127", fName, float32TypeName, "msg:127",
		),

		chkOutRelease(),
	)
}
