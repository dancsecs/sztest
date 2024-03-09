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

const (
	nilStr      = ""
	errTypeName = "err"
)

// BlankErrorMessage represents an empty panic message received.
const BlankErrorMessage = "sztest.BlankErrorMessage"

func (chk *Chk) errPrepareWant(want string) string {
	if want == "" {
		return nilStr
	}

	return chk.isStringify(want)
}

func (chk *Chk) errPrepareWantSlice(want []string) []string {
	if len(want) == 0 {
		return nil
	}

	r := make([]string, len(want))
	for i, s := range want {
		r[i] = chk.errPrepareWant(s)
	}

	return r
}

func (chk *Chk) errPrepareGot(got error) string {
	if got == nil {
		return nilStr
	}

	errMsg := got.Error()
	if errMsg == "" {
		return BlankErrorMessage
	}

	return chk.isStringify(errMsg)
}

func (chk *Chk) errPrepareGotSlice(got []error) []string {
	if len(got) == 0 {
		return nil
	}

	r := make([]string, len(got))
	for i, err := range got {
		r[i] = chk.errPrepareGot(err)
	}

	return r
}

// NoErrf simply invokes Err with want set to "" and msg formatted.
func (chk *Chk) NoErrf(got error, msgFmt string, msgArgs ...any) bool {
	if chk.errPrepareGot(got) == nilStr {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(chk.errPrepareGot(got), nilStr, errTypeName, msgFmt, msgArgs...)
}

// NoErr simply invokes Err with want set to "".
func (chk *Chk) NoErr(got error, msg ...any) bool {
	if chk.errPrepareGot(got) == nilStr {
		return true
	}

	chk.t.Helper()

	return chk.errChk(chk.errPrepareGot(got), nilStr, errTypeName, msg...)
}

// Errf compare the gotten error against the wanted error string.
// If either  "" or "<nil>" is wanted the error should be a nil.
func (chk *Chk) Errf(
	got error, want string, msgFmt string, msgArgs ...any,
) bool {
	if chk.errPrepareGot(got) == chk.errPrepareWant(want) {
		return true
	}

	chk.t.Helper()

	return chk.errChkf(
		chk.errPrepareGot(got),
		chk.errPrepareWant(want),
		errTypeName,
		msgFmt, msgArgs...,
	)
}

// Err compare the gotten error against the wanted error string.
// If either  "" or "<nil>" is wanted the error should be a nil.
func (chk *Chk) Err(got error, want string, msg ...any) bool {
	if chk.errPrepareGot(got) == chk.errPrepareWant(want) {
		return true
	}

	chk.t.Helper()

	return chk.errChk(
		chk.errPrepareGot(got), chk.errPrepareWant(want), errTypeName, msg...,
	)
}

// ErrSlicef compare the gotten error against the wanted error string.
// If either  "" or "<nil>" is wanted the error should be a nil.
func (chk *Chk) ErrSlicef(
	got []error, want []string, msgFmt string, msgArgs ...any,
) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = chk.errPrepareGot(got[i]) == chk.errPrepareWant(want[i])
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlicef(
		chk, chk.errPrepareGotSlice(got), chk.errPrepareWantSlice(want), errTypeName,
		defaultCmpFunc[string], msgFmt, msgArgs...,
	)
}

// ErrSlice compare the gotten error against the wanted error string.
// If either  "" or "<nil>" is wanted the error should be a nil.
func (chk *Chk) ErrSlice(
	got []error, want []string, msg ...any,
) bool {
	l := len(got)
	equal := l == len(want)

	for i := 0; equal && i < l; i++ {
		equal = chk.errPrepareGot(got[i]) == chk.errPrepareWant(want[i])
	}

	if equal {
		return true
	}

	chk.t.Helper()

	return errSlice(chk,
		chk.errPrepareGotSlice(got), chk.errPrepareWantSlice(want), errTypeName,
		defaultCmpFunc[string], msg...,
	)
}
