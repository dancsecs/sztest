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

func tstChkFloat32(t *testing.T) {
	t.Run("Good", chkFloat32Test_Good)

	t.Run("Bad", chkFloat32Test_Bad)
	t.Run("BadMsg1", chkFloat32Test_Bad1)
	t.Run("BadMsg2", chkFloat32Test_Bad2)
	t.Run("BadMsg3", chkFloat32Test_Bad3)

	t.Run("Slice-Good", chkFloat32SliceTest_Good)
	t.Run("Slice-BadMsg1", chkFloat32SliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkFloat32SliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkFloat32SliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkFloat32SliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkFloat32SliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkFloat32SliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkFloat32SliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkFloat32SliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkFloat32SliceTest_BadMsg9)

	t.Run("Bounded", chkFloat32BoundedTest_All)
	t.Run("Unbounded", chkFloat32UnboundedTest_All)
}

func chkFloat32Test_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32(0.0, 0.0, 0.0)
	chk.Float32(0.0, 0.0, 0.0, "not ", "displayed")
	chk.Float32f(0.0, 0.0, 0.0, "not %s", "displayed")

	chk.Float32(0.02, 0.01, 0.1)
	chk.Float32(0.02, 0.01, 0.1, "not ", "displayed")
	chk.Float32f(0.02, 0.01, 0.1, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat32Test_Bad(t *testing.T) {
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

func chkFloat32Test_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Float32f(0.02, 0.01, 0.005, "This message will be displayed %s", "first")

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

func chkFloat32Test_Bad2(t *testing.T) {
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

func chkFloat32Test_Bad3(t *testing.T) {
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

func chkFloat32SliceTest_Good(t *testing.T) {
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
		[]float32{0.02, 0.06, -.07, 9.0}, []float32{0.01, 0.05, -0.08, 9.0}, 0.1,
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkFloat32SliceTest_BadMsg1(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg2(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg3(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg4(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg5(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg6(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg7(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg8(t *testing.T) {
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

func chkFloat32SliceTest_BadMsg9(t *testing.T) {
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

func chkFloat32BoundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := float32(33)
	max := float32(35)

	// Bad: Error displayed.
	chk.Float32Bounded(30, BoundedClosed, min, max)
	chk.Float32Bounded(31, BoundedClosed, min, max, "msg:", "31")
	chk.Float32Boundedf(32, BoundedClosed, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.Float32Bounded(33, BoundedClosed, min, max)
	chk.Float32Bounded(34, BoundedClosed, min, max, "not ", "displayed")
	chk.Float32Boundedf(35, BoundedClosed, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.Float32Bounded(36, BoundedClosed, min, max)

	const wntMsg = "out of bounds: [33,35] - { want | 33 <= want <= 35 }"
	const fName = "Float32Bounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "float32", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "float32", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "float32", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "float32", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkFloat32UnboundedTest_All(t *testing.T) {
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
	chk.Float32Unboundedf(130, UnboundedMinClosed, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: [128,MAX) - { want | want >= 128 }"
	const fName = "Float32Unbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "125", fName, "float32", ""),
		chkOutNumericUnbounded_(wntMsg, "126", fName, "float32", "msg:126"),
		chkOutNumericUnboundedf(wntMsg, "127", fName, "float32", "msg:127"),

		chkOutRelease(),
	)
}
