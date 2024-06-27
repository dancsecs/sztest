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
	"fmt"
	"os"
	"testing"
)

func tstChkArgsAndFlags(t *testing.T) {
	t.Run("SetupEmpty", chkArgsAndFlagsTestSetupEmpty)
	t.Run("SetupOneArg", chkArgsAndFlagsTestSetupOneArg)
	t.Run("GoodParseDefault", chkArgsAndFlagsTestGoodParseDefault)
	t.Run("GoodParse", chkArgsAndFlagsTestGoodParse)
	t.Run("GoodParseExtraArguments",
		chkArgsAndFlagsTestGoodParseExtraArguments,
	)
	t.Run("BadParseInteger", chkArgsAndFlagsTestBadParseInteger)
	t.Run("CaptureFlagUsage", chkArgsAndFlagsTestCaptureFlagUsage)
}

func chkArgsAndFlagsTestSetupEmpty(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	origArgs := os.Args
	origFlagCmdLine := flag.CommandLine

	defer func() {
		os.Args = origArgs
		flag.CommandLine = origFlagCmdLine
	}()

	chk.SetArgs("")

	chk.Str(flag.CommandLine.Name(), "unspecifiedProgram")
	chk.StrSlice(os.Args, []string{"unspecifiedProgram"})
}

func chkArgsAndFlagsTestSetupOneArg(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	origArgs := os.Args
	origFlagCmdLine := flag.CommandLine

	defer func() {
		os.Args = origArgs
		flag.CommandLine = origFlagCmdLine
	}()

	chk.SetArgs("progName")

	chk.Str(flag.CommandLine.Name(), "progName")
	chk.StrSlice(os.Args, []string{"progName"})
}

func chkArgsAndFlagsTestGoodParseDefault(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	var strValue string

	main := func() {
		flag.StringVar(&strValue, "strValue", "defaultStrValue",
			"usage of default string value",
		)
		flag.Parse()
	}

	chk.SetArgs("progName")

	chk.Str(strValue, "")

	chk.NoPanic(main)

	chk.Str(strValue, "defaultStrValue")

	chk.Str(flag.CommandLine.Name(), "progName")
	chk.StrSlice(os.Args, []string{"progName"})

	chk.Int(flag.NArg(), 0)
}

func chkArgsAndFlagsTestGoodParse(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	var strValue string

	main := func() {
		flag.StringVar(&strValue, "s", "defaultStrValue",
			"usage of default string value",
		)
		flag.Parse()
	}

	chk.SetArgs(
		"progName",
		"-s",
		"str from arg",
	)

	chk.Str(strValue, "")

	chk.NoPanic(main)

	chk.Str(strValue, "str from arg")

	chk.Str(flag.CommandLine.Name(), "progName")
	chk.StrSlice(os.Args, []string{"progName", "-s", "str from arg"})

	chk.Int(flag.NArg(), 0)
}

func chkArgsAndFlagsTestGoodParseExtraArguments(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	var strValue string

	main := func() {
		flag.StringVar(&strValue, "s", "defaultStrValue",
			"usage of default string value",
		)
		flag.Parse()
	}

	chk.SetArgs(
		"progName",
		"-s",
		"str from arg",
		"extra Arg",
	)

	chk.Str(strValue, "")

	chk.NoPanic(main)

	chk.Str(strValue, "str from arg")

	chk.Str(flag.CommandLine.Name(), "progName")
	chk.StrSlice(
		os.Args,
		[]string{"progName", "-s", "str from arg", "extra Arg"},
	)

	chk.Int(flag.NArg(), 1)
	chk.Str(flag.Arg(0), "extra Arg")
}

func chkArgsAndFlagsTestBadParseInteger(t *testing.T) {
	chk := CaptureStderr(t)
	defer chk.Release()

	var intValue int

	main := func() {
		flag.IntVar(&intValue, "n", 10,
			"usage of default int value",
		)
		flag.Parse()
	}

	chk.SetArgs(
		"progName",
		"-n",
		"NotANumber",
	)

	chk.Int(intValue, 0)

	chk.Panic(
		main,
		`invalid value "NotANumber" for flag -n: parse error`,
	)

	chk.Int(intValue, 0)

	chk.Str(flag.CommandLine.Name(), "progName")
	chk.StrSlice(os.Args, []string{"progName", "-n", "NotANumber"})

	chk.Int(flag.NArg(), 0)

	chk.Stderr("" +
		`invalid value "NotANumber" for flag -n: parse error` + "\n" +
		"Usage of progName:\n" +
		`\s -n int` + "\n" +
		`\s   ` + "\tusage of default int value (default 10)",
	)
}

func chkArgsAndFlagsTestCaptureFlagUsage(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.SetArgs("progname")

	cmdLine := flag.CommandLine

	cmdLine.Usage = func() {
		_, _ = fmt.Fprint(cmdLine.Output(), "usage message")
	}
	chk.Str(
		chk.CaptureFlagUsage(cmdLine),
		"usage message",
	)
}
