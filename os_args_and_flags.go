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
	"flag"
	"os"
	"strings"
)

// SetArgs replaces the process arguments used by os.Args and the default
// flag.CommandLine. It sets os.Args to progName followed by args, and creates
// a new flag set configured with flag.PanicOnError. The original arguments and
// flag.CommandLine are restored when chk.Release() is called.
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

// CaptureFlagUsage captures and returns the usage output of the supplied
// *flag.FlagSet as a string. This is useful for verifying custom flag
// definitions and usage messages in tests.
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
