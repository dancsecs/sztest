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

import "fmt"

// BlankPanicMessage represents an empty panic message received.
const BlankPanicMessage = "sztest.BlankPanicMessage"

//nolint:nonamedreturns // Required to set return in deferred function.
func (chk *Chk) runPanicTest(f func()) (panicMessage string) {
	defer func() {
		r := recover()
		if r != nil {
			panicMessage = fmt.Sprintf("%v", r)
			if panicMessage == "" {
				panicMessage = BlankPanicMessage
			}
		}
	}()
	chk.runningPanicFunction = true
	f()
	return
}

// NoPanicf simply invokes Panic with want set to "" and msg formatted.
func (chk *Chk) NoPanicf(gotF func(), msgFmt string, msgArgs ...any) bool {
	defer func() {
		chk.runningPanicFunction = false
	}()
	pm := chk.runPanicTest(gotF)
	if chk.isStringify(pm) == "" {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(pm, "", "panic", msgFmt, msgArgs...)
}

// NoPanic simply invokes Err with want set to "".
func (chk *Chk) NoPanic(gotF func(), msg ...any) bool {
	defer func() {
		chk.runningPanicFunction = false
	}()
	pm := chk.runPanicTest(gotF)
	if chk.isStringify(pm) == "" {
		return true
	}
	chk.t.Helper()
	return chk.errChk(pm, "", "panic", msg...)
}

// Panicf runs the supplied function and compares the panic value asserted
// to the supplied string.
func (chk *Chk) Panicf(
	gotF func(), want string, msgFmt string, msgArgs ...any,
) bool {
	defer func() {
		chk.runningPanicFunction = false
	}()
	pm := chk.runPanicTest(gotF)
	if chk.isStringify(pm) == chk.isStringify(want) {
		return true
	}
	chk.t.Helper()
	return chk.errChkf(pm, want, "panic", msgFmt, msgArgs...)
}

// Panic runs the supplied function and compares the panic value asserted to
// the supplied string.
func (chk *Chk) Panic(gotF func(), want string, msg ...any) bool {
	defer func() {
		chk.runningPanicFunction = false
	}()
	pm := chk.runPanicTest(gotF)
	if chk.isStringify(pm) == chk.isStringify(want) {
		return true
	}
	chk.t.Helper()
	return chk.errChk(pm, want, "panic", msg...)
}
