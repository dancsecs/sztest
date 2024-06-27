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

// SetArgs invokes the current arguments in os.Args and
// flag.CommandLine.  Package variable os.Args is set to the provided arguments
// and a new flag set is assigned to flag.CommandLine and is ready to use.
// Original values are restores with the chk object is released.
func (chk *Chk) SetArgs(progName string, args ...string) {
	if progName == "" {
		progName = "unspecifiedProgram"
	}

	savedOsArgs := os.Args
	savedFlagCmdLine := flag.CommandLine

	chk.PushPreReleaseFunc(func() error {
		os.Args = savedOsArgs
		flag.CommandLine = savedFlagCmdLine

		return nil
	})

	fullArgs := append([]string{progName}, args...)

	os.Args = fullArgs
	flag.CommandLine = flag.NewFlagSet(progName, flag.PanicOnError)
}

// CaptureFlagUsage is a convenience function that captures the output
// of the provided *flag.FlagSet.
func (*Chk) CaptureFlagUsage(flagSet *flag.FlagSet) string {
	buf := strings.Builder{}
	origOut := flagSet.Output()

	defer func() {
		flagSet.SetOutput(origOut)
	}()
	flagSet.SetOutput(&buf)

	flagSet.Usage()

	return buf.String()
}
