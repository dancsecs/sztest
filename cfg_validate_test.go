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
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	invalidOkBool        = "unexpected %s ok value: got: %t want: %t"
	invalidBool          = "unexpected %s: got: %t want: %t"
	invalidPerm          = "unexpected %s: got: %4.4o want: %4.4o"
	invalidString        = "unexpected %s:\ngot: %q\nwant: %q"
	invalidInt           = "unexpected %s: got: %d want: %d"
	invalidCaptureLength = "unexpected %s log output length: got: %d  want: %d"
)

func testConfigValidate(t *testing.T) {
	t.Run("FailFast", testConfigValidateFailFast)
	t.Run("PermDir", testConfigValidatePermDir)
	t.Run("PermFile", testConfigValidatePermFile)
	t.Run("PermExe", testConfigValidatePermExe)
	t.Run("TmpDir", testConfigValidateTmpDir)
	t.Run("Color", testConfigValidateColor)
	t.Run("MinRunString", testConfigValidateMinRunString)
	t.Run("MinRunSlice", testConfigValidateMinRunSlice)
	t.Run("BufferSize", testConfigValidateBufferSize)
}

func testConfigValidateFailFast(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "fail_fast"

	failFastValue, ok := validateFailFast("true")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName, ok, true)
	}
	if !failFastValue {
		t.Fatalf(invalidBool, jsonName, failFastValue, true)
	}

	failFastValue, ok = validateFailFast("false")
	if !ok {
		t.Fatalf(invalidBool, jsonName, ok, true)
	}
	if failFastValue {
		t.Fatalf(invalidBool, jsonName, failFastValue, false)
	}

	failFastValue, ok = validateFailFast("invalid")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}
	if failFastValue {
		t.Fatalf(invalidBool, jsonName, failFastValue, false)
	}

	lines := strings.Split(buf.String(), "\n")
	wLineLength := 2
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}
	wLine := fmt.Sprintf(
		errMsg, EnvFailFast, "invalid", validFailFast, settingFailFast,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

func testConfigValidatePermDir(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "perm_dir"

	permDirValue, ok := validatePermDir("0712")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			true,
		)
	}
	if permDirValue != 0o0712 {
		t.Fatalf(invalidPerm, jsonName,
			permDirValue,
			0o0712,
		)
	}

	permDirValue, ok = validatePermDir("0900")
	if ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			false,
		)
	}
	if permDirValue != 0 {
		t.Fatalf(invalidPerm, jsonName,
			permDirValue,
			0,
		)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 2
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvPermDir, "0900", validPermDir, settingPermDir,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

func testConfigValidatePermFile(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "perm_file"

	permFileValue, ok := validatePermFile("0612")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			true,
		)
	}
	if permFileValue != 0o0612 {
		t.Fatalf(invalidPerm, jsonName,
			permFileValue,
			0o0612,
		)
	}

	permFileValue, ok = validatePermFile("0700")
	if ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			false,
		)
	}
	if permFileValue != 0 {
		t.Fatalf(invalidPerm, jsonName,
			permFileValue,
			0,
		)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 2
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvPermFile, "0700", validPermFile, settingPermFile,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

func testConfigValidatePermExe(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "perm_exe"

	permExeValue, ok := validatePermExe("0712")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			true,
		)
	}
	if permExeValue != 0o0712 {
		t.Fatalf(invalidPerm, jsonName,
			permExeValue,
			0o0712,
		)
	}

	permExeValue, ok = validatePermExe("0800")
	if ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			false,
		)
	}
	if permExeValue != 0 {
		t.Fatalf(invalidPerm, jsonName,
			permExeValue,
			0,
		)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 2
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvPermExe, "0800", validPermExe, settingPermExe,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

//nolint:cyclop // Ok.
func testConfigValidateTmpDir(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "tmp_dir"

	newTmpPath := filepath.Join(settingTmpDir, "newTmpForValidate")

	err := os.RemoveAll(newTmpPath)
	if err != nil {
		t.Fatal("could not clear test directory: " + err.Error())
	}

	tmpDirValue, ok := validateTmpDir(newTmpPath)
	if ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			false,
		)
	}
	if tmpDirValue != "" {
		t.Fatalf(invalidString, jsonName,
			tmpDirValue,
			"",
		)
	}

	err = os.Mkdir(newTmpPath, 0o0700)
	if err != nil {
		t.Fatalf("could not setup tmp dir test: " + err.Error())
	}
	defer func() {
		_ = os.Remove(newTmpPath)
	}()

	tmpDirValue, ok = validateTmpDir(newTmpPath)
	if !ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			true,
		)
	}
	if tmpDirValue != newTmpPath {
		t.Fatalf(invalidString, jsonName,
			tmpDirValue,
			newTmpPath,
		)
	}

	// Not a directory.
	badDir := filepath.Join(newTmpPath, "badTmpDir")
	err = os.WriteFile(badDir, []byte{}, 0o0600)
	if err != nil {
		t.Fatal("could not create file to be bad tmp dir")
	}
	defer func() {
		_ = os.Remove(badDir)
	}()

	tmpDirValue, ok = validateTmpDir(badDir)
	if ok {
		t.Fatalf(invalidOkBool, jsonName,
			ok,
			false,
		)
	}
	if tmpDirValue != "" {
		t.Fatalf(invalidString, jsonName,
			tmpDirValue,
			"",
		)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 3
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvTmpDir, newTmpPath, validTmpDir, settingTmpDir,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
	wLine = fmt.Sprintf(
		errMsg, EnvTmpDir, badDir, validTmpDir, settingTmpDir,
	)
	if !strings.Contains(lines[1], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

//nolint:cyclop,gocognit // Ok.
func testConfigValidateColor(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "color_"
	const jsonChg = "chg"
	const jsonNameChg = jsonName + jsonChg

	testColor := func(s, envVarName, def, exp string) {
		t.Helper()
		v, ok := validateMark(s, envVarName, def)
		if !ok {
			t.Fatalf(invalidOkBool, envVarName, ok, true)
		}
		if v != exp {
			t.Fatalf(invalidString, envVarName, v, exp)
		}
	}

	testColor("black", "chg", settingMarkChgOn, clrBlack)
	testColor("red", "del", settingMarkDelOn, clrRed)
	testColor("green", "InsOn", settingMarkInsOn, clrGreen)
	testColor("yellow", "GotOn", settingMarkGotOn, clrYellow)
	testColor("blue", "WntOn", settingMarkWntOn, clrBlue)
	testColor("magenta", "sep", settingMarkSepOn, clrMagenta)
	testColor("cyan", "chg", settingMarkChgOn, clrCyan)
	testColor("white", "chg", settingMarkChgOn, clrWhite)

	testColor("hi-black", "chg", settingMarkChgOn, clrHiBlack)
	testColor("hi-red", "del", settingMarkDelOn, clrHiRed)
	testColor("hi-green", "InsOn", settingMarkInsOn, clrHiGreen)
	testColor("hi-yellow", "GotOn", settingMarkGotOn, clrHiYellow)
	testColor("hi-blue", "WntOn", settingMarkWntOn, clrHiBlue)
	testColor("hi-magenta", "sep", settingMarkSepOn, clrHiMagenta)
	testColor("hi-cyan", "chg", settingMarkChgOn, clrHiCyan)
	testColor("hi-white", "chg", settingMarkChgOn, clrHiWhite)

	testColor("bk-black", "chg", settingMarkChgOn, clrBkBlack)
	testColor("bk-red", "del", settingMarkDelOn, clrBkRed)
	testColor("bk-green", "InsOn", settingMarkInsOn, clrBkGreen)
	testColor("bk-yellow", "GotOn", settingMarkGotOn, clrBkYellow)
	testColor("bk-blue", "WntOn", settingMarkWntOn, clrBkBlue)
	testColor("bk-magenta", "sep", settingMarkSepOn, clrBkMagenta)
	testColor("bk-cyan", "chg", settingMarkChgOn, clrBkCyan)
	testColor("bk-white", "chg", settingMarkChgOn, clrBkWhite)

	testColor("bk-hi-black", "chg", settingMarkChgOn, clrBkHiBlack)
	testColor("bk-hi-red", "del", settingMarkDelOn, clrBkHiRed)
	testColor("bk-hi-green", "InsOn", settingMarkInsOn, clrBkHiGreen)
	testColor("bk-hi-yellow", "GotOn", settingMarkGotOn, clrBkHiYellow)
	testColor("bk-hi-blue", "WntOn", settingMarkWntOn, clrBkHiBlue)
	testColor("bk-hi-magenta", "sep", settingMarkSepOn, clrBkHiMagenta)
	testColor("bk-hi-cyan", "chg", settingMarkChgOn, clrBkHiCyan)
	testColor("bk-hi-white", "chg", settingMarkChgOn, clrBkHiWhite)

	testColor("default", "chg", clrOff, clrOff)
	testColor("bold", "chg", settingMarkChgOn, clrBold)
	testColor("italic", "chg", settingMarkChgOn, clrItalic)
	testColor("underline", "chg", settingMarkChgOn, clrUnderline)
	testColor("reverse", "chg", settingMarkChgOn, clrReverse)
	testColor("strikeout", "chg", settingMarkChgOn, clrStrikeout)

	markValue, ok := validateMark("", EnvMarkChgOn, settingMarkChgOn)
	if !ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, jsonNameChg, markValue, "")
	}

	markValue, ok = validateMark("_and_", EnvMarkChgOn, settingMarkChgOn)
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, false)
	}
	if markValue != "" {
		t.Fatalf(invalidString, jsonNameChg, markValue, "")
	}

	markValue, ok = validateMark("blue_And_blue", EnvMarkChgOn, settingMarkChgOn)
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, jsonNameChg, markValue, "")
	}

	markValue, ok = validateMark("bk-blue_AND_bk-blue", EnvMarkChgOn, settingMarkChgOn)
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, jsonNameChg, markValue, "")
	}

	markValue, ok = validateMark("bold_and_bold", EnvMarkChgOn, settingMarkChgOn)
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, jsonNameChg, markValue, "")
	}

	markValue, ok = validateMark("custom_aNd_custom", EnvMarkChgOn, "")
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, EnvMarkChgOn, markValue, "")
	}

	markValue, ok = validateMark("blue_aNd_default", EnvMarkChgOn, "")
	if ok {
		t.Fatalf(invalidOkBool, EnvMarkChgOn, ok, true)
	}
	if markValue != "" {
		t.Fatalf(invalidString, EnvMarkChgOn, markValue, "")
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 13
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonNameChg, len(lines), wLineLength)
	}

	wLine := "missing (empty) attribute"
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "_and_", validColor, settingMarkChgOn,
	)
	if !strings.Contains(lines[1], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}

	wLine = "foreground color redefined"
	if !strings.Contains(lines[2], wLine) {
		t.Fatalf(invalidString, jsonName, lines[2], wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "blue_And_blue", validColor, settingMarkChgOn,
	)
	if !strings.Contains(lines[3], wLine) {
		t.Fatalf(invalidString, jsonName, lines[3], wLine)
	}

	wLine = "background color redefined"
	if !strings.Contains(lines[4], wLine) {
		t.Fatalf(invalidString, jsonName, lines[4], wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "bk-blue_AND_bk-blue", validColor, settingMarkChgOn,
	)
	if !strings.Contains(lines[5], wLine) {
		t.Fatalf(invalidString, jsonName, lines[5], wLine)
	}

	wLine = "style redefined"
	if !strings.Contains(lines[6], wLine) {
		t.Fatalf(invalidString, jsonName, lines[6], wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "bold_and_bold", validColor, settingMarkChgOn,
	)
	if !strings.Contains(lines[7], wLine) {
		t.Fatalf(invalidString, jsonName, lines[7], wLine)
	}

	wLine = "custom mark redefined"
	if !strings.Contains(lines[8], wLine) {
		t.Fatalf(invalidString, jsonName, lines[8], wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "custom_aNd_custom", validColor, "",
	)
	if !strings.Contains(lines[9], wLine) {
		t.Fatalf(invalidString, jsonName, lines[9], wLine)
	}

	wLine = "default style must be defined by itself"
	if !strings.Contains(lines[10], wLine) {
		t.Fatalf(invalidString, jsonName, lines[10], wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvMarkChgOn, "blue_aNd_default", validColor, "",
	)
	if !strings.Contains(lines[11], wLine) {
		t.Fatalf(invalidString, jsonName, lines[11], wLine)
	}
}

//nolint:cyclop // Ok.
func testConfigValidateMinRunString(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "min_run_string"

	runStringValue, ok := validateMinRunString("0")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}
	if runStringValue != 0 {
		t.Fatalf(invalidInt, jsonName, runStringValue, 0)
	}

	runStringValue, ok = validateMinRunString("6")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}
	if runStringValue != 0 {
		t.Fatalf(invalidInt, jsonName, runStringValue, 0)
	}

	runStringValue, ok = validateMinRunString("1")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName, ok, true)
	}
	if runStringValue != 1 {
		t.Fatalf(invalidInt, jsonName, runStringValue, 1)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 3
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvDiffChars, "0", validMinRunString, settingDiffChars,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvDiffChars, "6", validMinRunString, settingDiffChars,
	)
	if !strings.Contains(lines[1], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

//nolint:cyclop // Ok.
func testConfigValidateMinRunSlice(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "min_run_slice"

	runSliceValue, ok := validateMinRunSlice("0")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}
	if runSliceValue != 0 {
		t.Fatalf(invalidInt, jsonName, runSliceValue, 0)
	}

	runSliceValue, ok = validateMinRunSlice("6")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}
	if runSliceValue != 0 {
		t.Fatalf(invalidInt, jsonName, runSliceValue, 0)
	}

	runSliceValue, ok = validateMinRunSlice("5")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName, ok, true)
	}
	if runSliceValue != 5 {
		t.Fatalf(invalidInt, jsonName, runSliceValue, 5)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 3
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvDiffSlice, "0", validMinRunSlice, settingDiffSlice,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}

	wLine = fmt.Sprintf(
		errMsg, EnvDiffSlice, "6", validMinRunSlice, settingDiffSlice,
	)
	if !strings.Contains(lines[1], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}

func testConfigValidateBufferSize(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	const jsonName = "output_buffer_size"

	bufSizeValue, ok := validateBufferSize("-1")
	if ok {
		t.Fatalf(invalidOkBool, jsonName, ok, false)
	}

	if bufSizeValue != 0 {
		t.Fatalf(invalidInt, jsonName, bufSizeValue, 0)
	}

	bufSizeValue, ok = validateBufferSize("15000")
	if !ok {
		t.Fatalf(invalidOkBool, jsonName, ok, true)
	}

	if bufSizeValue != 15000 {
		t.Fatalf(invalidInt, jsonName, bufSizeValue, 15000)
	}

	lines := strings.Split(buf.String(), "\n")

	wLineLength := 2
	if len(lines) != wLineLength || lines[wLineLength-1] != "" {
		t.Fatalf(invalidCaptureLength, jsonName, len(lines), wLineLength)
	}

	wLine := fmt.Sprintf(
		errMsg, EnvBufferSize, "-1", validBufferSize, settingBufferSize,
	)
	if !strings.Contains(lines[0], wLine) {
		t.Fatalf(invalidString, jsonName, buf.String(), wLine)
	}
}
