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

// AddSub registers a regexp pattern and its replacement string to normalize
// variable output before assertions.
//
// Each AddSub call compiles expr and stores it with subStr for later use.
// During comparison, substitutions are applied recursively across captured
// output and string assertions (e.g., Str, Err, Panic, Stdout, Stderr, Log)
// until no further matches remain.
//
// This is useful for masking nondeterministic values such as timestamps,
// memory addresses, or counters. Compilation failures cause an immediate
// fatal error. Currently subStr does not support regexp submatches, but this
// is planned for future versions.
func (chk *Chk) AddSub(expr, subStr string) {
	re, err := regexp.Compile(expr)
	if err != nil {
		chk.t.Helper()
		chk.Fatalf("%v", err)

		return
	}

	chk.subs = append(chk.subs, substitution{
		re:     re,
		subStr: subStr,
	})
}

// subStr recursively applies all regexp substitution stored in chk.subs
// to the supplied string.
func (chk *Chk) subStr(lines string) string {
	beforeLines := ""
	for beforeLines != lines {
		beforeLines = lines

		for _, re := range chk.subs {
			lines = re.re.ReplaceAllString(lines, re.subStr)
		}
	}

	return lines
}
