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

import "testing"

func tstChkByte(t *testing.T) {
	t.Run("Good", chkByteTest_Good)

	t.Run("Bad", chkByteTest_Bad)
	t.Run("BadMsg1", chkByteTest_Bad1)
	t.Run("BadMsg2", chkByteTest_Bad2)
	t.Run("BadMsg3", chkByteTest_Bad3)

	t.Run("Slice-Good", chkByteSliceTest_Good)
	t.Run("Slice-BadMsg1", chkByteSliceTest_BadMsg1)
	t.Run("Slice-BadMsg2", chkByteSliceTest_BadMsg2)
	t.Run("Slice-BadMsg3", chkByteSliceTest_BadMsg3)
	t.Run("Slice-BadMsg4", chkByteSliceTest_BadMsg4)
	t.Run("Slice-BadMsg5", chkByteSliceTest_BadMsg5)
	t.Run("Slice-BadMsg6", chkByteSliceTest_BadMsg6)
	t.Run("Slice-BadMsg7", chkByteSliceTest_BadMsg7)
	t.Run("Slice-BadMsg8", chkByteSliceTest_BadMsg8)
	t.Run("Slice-BadMsg9", chkByteSliceTest_BadMsg9)

	t.Run("Bounded", chkByteBoundedTest_All)
	t.Run("Unbounded", chkByteUnboundedTest_All)
}

func chkByteTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Byte(0, 0)
	chk.Byte(1, 1, "not ", "displayed")
	chk.Bytef(2, 2, "not %s", "displayed")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkByteTest_Bad(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Byte(2, 1)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Byte",
			chkOutCommonMsg("", "byte"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkByteTest_Bad1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bytef(2, 1, "This message will be displayed %s", "first")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Bytef",
			chkOutCommonMsg("This message will be displayed first", "byte"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkByteTest_Bad2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Byte(2, 1, "This message will be displayed ", "second")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Byte",
			chkOutCommonMsg("This message will be displayed second", "byte"),
			g(markAsChg("2", "1", DiffGot)),
			w(markAsChg("2", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkByteTest_Bad3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.Bytef(0, 1, "This message will be displayed %s", "third")

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsError(
			"Bytef",
			chkOutCommonMsg("This message will be displayed third", "byte"),
			g(markAsChg("0", "1", DiffGot)),
			w(markAsChg("0", "1", DiffWant)),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_Good(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice(
		[]byte{}, []byte{},
		"This message will NOT be displayed",
	)
	chk.ByteSlice(
		[]byte{0}, []byte{0},
		"This message will NOT be displayed",
	)
	chk.ByteSlice(
		[]byte{2, 6, 7}, []byte{2, 6, 7},
		"This message will NOT be displayed",
	)

	chk.ByteSlicef(
		[]byte{2, 6, 7, 9}, []byte{2, 6, 7, 9},
		"This message will NOT be %s", "displayed",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg1(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlicef(
		[]byte{2}, []byte{},
		"This message will be displayed %s", "first",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 0,
			"[]byte", "ByteSlicef",
			"This message will be displayed first",
			chkOutLnGot("0", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg2(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice(
		[]byte{}, []byte{1},
		"This message will be displayed ", "second",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			0, 1,
			"[]byte", "ByteSlice",
			"This message will be displayed second",
			chkOutLnWnt("0", "1"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg3(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlicef(
		[]byte{1, 2}, []byte{1},
		"This message will be displayed %s", "third",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 1,
			"[]byte", "ByteSlicef",
			"This message will be displayed third",
			chkOutLnSame("0", "0", "1"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg4(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice(
		[]byte{1}, []byte{1, 2},
		"This message will be displayed ", "fourth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			1, 2,
			"[]byte", "ByteSlice",
			"This message will be displayed fourth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg5(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlicef(
		[]byte{1, 2}, []byte{2, 2},
		"This message will be displayed %s", "fifth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]byte", "ByteSlicef",
			"This message will be displayed fifth",
			chkOutLnGot("0", "1"),
			chkOutLnSame("1", "0", "2"),
			chkOutLnWnt("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg6(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice(
		[]byte{2, 2}, []byte{1, 2},
		"This message will be displayed ", "sixth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]byte", "ByteSlice",
			"This message will be displayed sixth",
			chkOutLnWnt("0", "1"),
			chkOutLnSame("0", "1", "2"),
			chkOutLnGot("1", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg7(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlicef(
		[]byte{1, 2}, []byte{1, 3},
		"This message will be displayed %s", "seventh",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]byte", "ByteSlicef",
			"This message will be displayed seventh",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "2", "3"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg8(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice(
		[]byte{1, 3}, []byte{1, 2},
		"This message will be displayed ", "eighth",
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2,
			"[]byte", "ByteSlice",
			"This message will be displayed eighth",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

func chkByteSliceTest_BadMsg9(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.ByteSlice([]byte{1, 3}, []byte{1, 2})

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutIsSliceError(
			false,
			2, 2, "[]byte", "ByteSlice", "",
			chkOutLnSame("0", "0", "1"),
			chkOutLnChanged("1", "1", "3", "2"),
		),
		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Bounded
/////////////////////////////////////////////

func chkByteBoundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	min := byte(32)
	max := byte(36)

	// Bad: Error displayed.
	chk.ByteBounded(30, BoundedOpen, min, max)
	chk.ByteBounded(31, BoundedOpen, min, max, "msg:", "31")
	chk.ByteBoundedf(32, BoundedOpen, min, max, "msg:%d", 32)

	// Good:  No error displayed.
	chk.ByteBounded(33, BoundedOpen, min, max)
	chk.ByteBounded(34, BoundedOpen, min, max, "not ", "displayed")
	chk.ByteBoundedf(35, BoundedOpen, min, max, "not %s", "displayed")

	// Bad: Error displayed.
	chk.ByteBounded(36, BoundedOpen, min, max)

	const wntMsg = "out of bounds: (32,36) - { want | 32 < want < 36 }"
	const fName = "ByteBounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericBounded_(wntMsg, "30", fName, "byte", ""),
		chkOutNumericBounded_(wntMsg, "31", fName, "byte", "msg:31"),
		chkOutNumericBoundedf(wntMsg, "32", fName, "byte", "msg:32"),

		chkOutNumericBounded_(wntMsg, "36", fName, "byte", ""),

		chkOutRelease(),
	)
}

//////////////////////////////////////////////
// Unbounded
/////////////////////////////////////////////

func chkByteUnboundedTest_All(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	bound := byte(128)

	// Bad: Error displayed.
	chk.ByteUnbounded(126, UnboundedMinOpen, bound)
	chk.ByteUnbounded(127, UnboundedMinOpen, bound, "msg:", "127")
	chk.ByteUnboundedf(128, UnboundedMinOpen, bound, "msg:%d", 128)

	// Good:  No error displayed.
	chk.ByteUnbounded(129, UnboundedMinOpen, bound)
	chk.ByteUnbounded(130, UnboundedMinOpen, bound, "not ", "displayed")
	chk.ByteUnboundedf(131, UnboundedMinOpen, bound, "not %s", "displayed")

	const wntMsg = "out of bounds: (128,MAX) - { want | want > 128 }"
	const fName = "ByteUnbounded"

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),

		chkOutNumericUnbounded_(wntMsg, "126", fName, "byte", ""),
		chkOutNumericUnbounded_(wntMsg, "127", fName, "byte", "msg:127"),
		chkOutNumericUnboundedf(wntMsg, "128", fName, "byte", "msg:128"),

		chkOutRelease(),
	)
}
