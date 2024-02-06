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
	"flag"
	"os"
	"strings"
)

// SetupArgsAndFlags saves the current arguments in os.Args and
// flag.CommandLine.  Package variable os.Args is set to the provided arguments
// and a new flag set is assigned to flag.CommandLine and is ready to use.
// Original values are restores with the chk object is released.
func (chk *Chk) SetupArgsAndFlags(args []string) *flag.FlagSet {
	if len(args) < 1 {
		args = []string{"unspecifiedProgram"}
	}
	savedOsArgs := os.Args
	savedFlagCmdLine := flag.CommandLine
	chk.PushPreReleaseFunc(func() error {
		os.Args = savedOsArgs
		flag.CommandLine = savedFlagCmdLine
		return nil
	})
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(args[0], flag.PanicOnError)
	return flag.CommandLine
}

// CaptureFlagUsage is a convenience function that captures the output
// of the provided *flag.FlagSet.
func (*Chk) CaptureFlagUsage(f *flag.FlagSet) string {
	buf := strings.Builder{}
	origOut := f.Output()
	defer func() {
		f.SetOutput(origOut)
	}()
	f.SetOutput(&buf)

	f.Usage()
	return buf.String()
}
