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
	"fmt"
	"reflect"
)

const (
	nilTypeName    = "nil interface"
	notNilTypeName = "not " + nilTypeName
)

// Nilf reports whether got is nil, including cases where got is an interface
// holding a typed nil pointer. It reports a failure through chk’s testingT
// if got is non-nil. The message is formatted according to msgFmt and msgArgs.
func (chk *Chk) Nilf(got any, msgFmt string, msgArgs ...any) bool {
	if got == nil || reflect.ValueOf(got).IsNil() {
		return true
	}

	chk.t.Helper()
	chk.Error(
		errMsgHeader(notNilTypeName, fmt.Sprintf(msgFmt, msgArgs...)),
	)

	return false
}

// Nil reports whether got is nil, including cases where got is an interface
// holding a typed nil pointer. It reports a failure through chk’s testingT
// if got is non-nil. An optional message may be provided.
func (chk *Chk) Nil(got any, msg ...any) bool {
	if got == nil || reflect.ValueOf(got).IsNil() {
		return true
	}

	chk.t.Helper()
	chk.Error(
		errMsgHeader(notNilTypeName, msg...),
	)

	return false
}

// NotNilf reports whether got is non-nil. Unlike Nil, this treats an
// interface holding a typed nil pointer as nil. It reports a failure through
// chk’s testingT if got is nil. The message is formatted according to msgFmt
// and msgArgs.
func (chk *Chk) NotNilf(got any, msgFmt string, msgArgs ...any) bool {
	if got != nil && !reflect.ValueOf(got).IsNil() {
		return true
	}

	chk.t.Helper()
	chk.Error(
		errMsgHeader(nilTypeName, fmt.Sprintf(msgFmt, msgArgs...)),
	)

	return false
}

// NotNil reports whether got is non-nil. Unlike Nil, this treats an interface
// holding a typed nil pointer as nil. It reports a failure through chk’s
// testingT if got is nil. An optional message may be provided.
func (chk *Chk) NotNil(got any, msg ...any) bool {
	if got != nil && !reflect.ValueOf(got).IsNil() {
		return true
	}

	chk.t.Helper()
	chk.Error(errMsgHeader(nilTypeName, msg...))

	return false
}
