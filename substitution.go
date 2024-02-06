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
	"regexp"
)

// Builtin replacement regular expressions.
const (
	SubTimestamp = `\d?\d:\d?\d:\d?\d(?:\.?\d{1,9})?`
	SubDuration  = `\d+(?:\.?\d+)?(?:ns|us|µs|ms|s|m|h)`
)

type substitution struct {
	re     *regexp.Regexp
	subStr string
}

// AddSub compiles and adds a new regexp and substitute string.
func (chk *Chk) AddSub(expr, subStr string) {
	re, err := regexp.Compile(expr)
	if err != nil {
		chk.t.Helper()
		chk.Fatalf(err.Error())
		return
	}
	chk.subs = append(chk.subs, substitution{
		re:     re,
		subStr: subStr,
	})
}

// subStr recursively applies all regexp substitution stored in in chk.subs
// to the supplied string.
func (chk *Chk) subStr(s string) string {
	beforeString := ""
	for beforeString != s {
		beforeString = s
		for _, re := range chk.subs {
			s = re.re.ReplaceAllString(s, re.subStr)
		}
	}
	return s
}
