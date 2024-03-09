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
	"fmt"
	"reflect"
)

const (
	nilTypeName    = "nil interface"
	notNilTypeName = "not " + nilTypeName
)

// Nilf checks that the interface value is nil with formatted msg.
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

// Nil checks that the interface value is nil.
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

// NotNilf checks that the interface value is nil with formatted msg.
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

// NotNil checks that the interface value is nil.
func (chk *Chk) NotNil(got any, msg ...any) bool {
	if got != nil && !reflect.ValueOf(got).IsNil() {
		return true
	}

	chk.t.Helper()
	chk.Error(errMsgHeader(nilTypeName, msg...))

	return false
}
