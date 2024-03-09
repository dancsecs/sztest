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

/*
   Master test file.

   This package is aimed exclusively at providing a helpful framework to
   simplify the creation and maintenance of common go test patterns while
   providing tools to quickly decipher exactly what part of test was
   unexpected.  Therefore features are testing in a strict order and used in
   subsequent testing within this package.  Keeping this in mind it is
   suggested that your go tests utilize the -failfast flag or move to the first
   error listed to avoid investigating superfluous issues.
*/

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

const tstErrMinStr = 3

// tstMarkupFunc defines the function type used to decorate error messages
// to assist deciphering the problem between the got and the want.
type tstMarkupFunc func(string, any, any) string

// Gets updated once markup and diff functions are fully tested.
var errGotWnt tstMarkupFunc //nolint:gochecknoglobals // Ok.

func saveThenSetupDefaultEnvironment() func() {
	var restoreFunc func()

	orig := make(map[string]string)

	for _, e := range os.Environ() {
		setting := strings.SplitN(e, "=", 2)
		if len(setting) == 2 &&
			strings.HasPrefix(setting[0], "SZTEST_") &&
			setting[0] != "SZTEST_TMP_DIR" {
			// Save env variable for restoration.
			orig[setting[0]] = setting[1]
			// Remove env override
			if err := os.Unsetenv(setting[0]); err != nil {
				log.Printf("Could not unsetEnv(%q): %v", setting[0], err)
			}
		}
	}

	restoreFunc = func() {
		for k, v := range orig {
			if err := os.Setenv(k, v); err != nil {
				log.Printf("Could not SetEnv(%q,%q): %v", k, v, err)
			}
		}
	}

	return restoreFunc
}

//nolint:paralleltest // Ok single test drives all tests in sequence.
func TestSzTest(t *testing.T) {
	// Set initial testing got/want function to plain
	errGotWnt = errMarkupFuncNone

	resEnv := saveThenSetupDefaultEnvironment()
	defer resEnv()

	// Test configuration overrides.
	t.Run("Config Validate", testConfigValidate)
	t.Run("Config Init", testConfigInit)

	// Test underlying markup first to assist in subsequent testing.
	t.Run("DiffMarkup", testDiffMarkupPrerequisites)
	t.Run("DiffFmt", testDiffFmt)
	t.Run("Diff", testDiffPrerequisites)

	// Update to default markup processing that has been tested above.
	errGotWnt = errMarkupFuncDefault

	// Test the replacement interface for Google's testing.T object.
	t.Run("TestingInterface", testSzTestingPrerequisites)

	// Test core and generic chk methods to enable use in further testing.
	t.Run("chkCore", tstChkCore)
	t.Run("chkGenerics", tstChkGeneric)

	t.Run("chkData", chkData)
	t.Run("chkInterface", chkInterface)

	t.Run("chkLogging", tstChkLogging)
	t.Run("chkSubstitution", tstChkSubstitution)

	t.Run("chkDir", tstChkDir)
	t.Run("chkIoClose", tstChkIoClose)
	t.Run("chkIoReader", tstChkIoReader)
	t.Run("chkIoWriter", tstChkIoWriter)
	t.Run("chkIoSeek", tstChkIoSeek)

	t.Run("chkArgsAndFlags", tstChkArgsAndFlags)

	t.Run("chkClock", tstChkClock)
}

func chkData(t *testing.T) {
	t.Run("Bool", tstChkBool)
	t.Run("Byte", tstChkByte)
	t.Run("Complex64", tstChkComplex64)
	t.Run("Complex128", tstChkComplex128)
	t.Run("Float32", tstChkFloat32)
	t.Run("Float64", tstChkFloat64)
	t.Run("Int", tstChkInt)
	t.Run("Int8", tstChkInt8)
	t.Run("Int16", tstChkInt16)
	t.Run("Int32", tstChkInt32)
	t.Run("Int64", tstChkInt64)
	t.Run("Rune", tstChkRune)
	t.Run("String", tstChkString)
	t.Run("Uint", tstChkUint)
	t.Run("Uint8", tstChkUint8)
	t.Run("Uint16", tstChkUint16)
	t.Run("Uint32", tstChkUint32)
	t.Run("Uint64", tstChkUint64)
	t.Run("Uintptr", tstChkUintptr)
}

func chkInterface(t *testing.T) {
	t.Run("Any", tstChkAny)
	t.Run("time.Duration", tstChkDur)
	t.Run("Err", tstChkErr)
	t.Run("ErrLast", tstChkErrLast)
	t.Run("Panic", tstChkPanic)
}

func errMarkupFuncNone(area string, got, wnt any) string {
	cleanGot := fmt.Sprintf("%v", got)
	cleanWnt := fmt.Sprintf("%v", wnt)
	prefix := ""

	if strings.Count(cleanGot, "\n") > 1 && strings.Count(cleanWnt, "\n") > 1 {
		prefix = "\n"
	}

	return commonMsgPrefix + area + "\n" +
		"GOT: " + prefix + cleanGot + "\n" +
		"WNT: " + prefix + cleanWnt
}

func errMarkupFuncDefault(area string, got, wnt any) string {
	return commonMsgPrefix + area + "\n" +
		resolveMarksForDisplay(
			gotWntDiff(
				fmt.Sprintf("%v", got),
				fmt.Sprintf("%v", wnt),
				tstErrMinStr,
			),
		)
}
