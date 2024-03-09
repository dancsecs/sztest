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
	"log"
	"os"
	"testing"
)

func testConfigInit(t *testing.T) {
	t.Run("Defaults", testCfgInitAllDefaults)
	t.Run("Overrides", testCfgInitAllOverrides)
}

/*

To enable testing of the initialization's code interaction with system
environment variables some utility test functions are here introduced to
save/restore the actual system environment variables and function's to
clear and set each environment variable to known states.

*/

type sysEnvVar struct {
	id       string
	v        string
	declared bool
}

func capture(id string) *sysEnvVar {
	v, ok := os.LookupEnv(id)

	return &sysEnvVar{
		id:       id,
		v:        v,
		declared: ok,
	}
}

func captureAll() []*sysEnvVar {
	return []*sysEnvVar{
		capture(EnvFailFast),
		capture(EnvPermDir),
		capture(EnvPermFile),
		capture(EnvPermExe),
		capture(EnvTmpDir),
		capture(EnvMarkWntOn),
		capture(EnvMarkWntOff),
		capture(EnvMarkGotOn),
		capture(EnvMarkGotOff),
		capture(EnvMarkMsgOn),
		capture(EnvMarkMsgOff),
		capture(EnvMarkInsOn),
		capture(EnvMarkInsOff),
		capture(EnvMarkDelOn),
		capture(EnvMarkDelOff),
		capture(EnvMarkChgOn),
		capture(EnvMarkChgOff),
		capture(EnvMarkSepOn),
		capture(EnvMarkSepOff),
		capture(EnvDiffChars),
		capture(EnvDiffSlice),
		capture(EnvBufferSize),
	}
}

func (env *sysEnvVar) set() {
	err := os.Setenv(env.id, env.v)
	if err != nil {
		log.Print("could not set env var: ", env.id, " Got: ", err)
	}
}

func (env *sysEnvVar) clear() {
	err := os.Unsetenv(env.id)
	if err != nil {
		log.Print("could not clear env var: ", env.id, " Got: %v", err)
	}
}

func clearAndCaptureAll() []*sysEnvVar {
	envVars := captureAll()
	for _, e := range envVars {
		e.clear()
	}

	initAll()

	return envVars
}

func restoreAll(envVars []*sysEnvVar) {
	for _, e := range envVars {
		if e.declared {
			e.set()
		} else {
			e.clear()
		}
	}

	initAll()
}

//nolint:cyclop // Ok.
func setupSysEnvVarTest(tmpDir string) error {
	const errMsg = "could not set: %s: %v"

	var err error

	if err = os.Setenv(EnvFailFast, "false"); err != nil {
		return fmt.Errorf(errMsg, EnvFailFast, err)
	}

	if err = os.Setenv(EnvBufferSize, "12345"); err != nil {
		return fmt.Errorf(errMsg, EnvBufferSize, err)
	}

	if err = os.Setenv(EnvPermDir, "0701"); err != nil {
		return fmt.Errorf(errMsg, EnvPermDir, err)
	}

	if err = os.Setenv(EnvPermFile, "0601"); err != nil {
		return fmt.Errorf(errMsg, EnvPermFile, err)
	}

	if err = os.Setenv(EnvPermExe, "0711"); err != nil {
		return fmt.Errorf(errMsg, EnvPermExe, err)
	}

	if err = os.Setenv(EnvTmpDir, tmpDir); err != nil {
		return fmt.Errorf(errMsg, EnvTmpDir, err)
	}

	if err = os.Setenv(EnvDiffChars, "2"); err != nil {
		return fmt.Errorf(errMsg, EnvDiffChars, err)
	}

	if err = os.Setenv(EnvDiffSlice, "4"); err != nil {
		return fmt.Errorf(errMsg, EnvDiffSlice, err)
	}

	if err = os.Setenv(EnvMarkWntOn, "<<WntOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkWntOn, err)
	}

	if err = os.Setenv(EnvMarkWntOff, "<<WntOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkWntOff, err)
	}

	if err = os.Setenv(EnvMarkGotOn, "<<GotOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkGotOn, err)
	}

	if err = os.Setenv(EnvMarkGotOff, "<<GotOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkGotOff, err)
	}

	if err = os.Setenv(EnvMarkMsgOn, "<<MsgOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkMsgOn, err)
	}

	if err = os.Setenv(EnvMarkMsgOff, "<<MsgOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkMsgOff, err)
	}

	if err = os.Setenv(EnvMarkInsOn, "<<InsOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkInsOn, err)
	}

	if err = os.Setenv(EnvMarkInsOff, "<<InsOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkInsOff, err)
	}

	if err = os.Setenv(EnvMarkDelOn, "<<DelOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkDelOn, err)
	}

	if err = os.Setenv(EnvMarkDelOff, "<<DelOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkDelOff, err)
	}

	if err = os.Setenv(EnvMarkChgOn, "<<ChgOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkChgOn, err)
	}

	if err = os.Setenv(EnvMarkChgOff, "<<ChgOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkChgOff, err)
	}

	if err = os.Setenv(EnvMarkSepOn, "<<SepOn>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkSepOn, err)
	}

	if err = os.Setenv(EnvMarkSepOff, "<<SepOff>>"); err != nil {
		return fmt.Errorf(errMsg, EnvMarkSepOff, err)
	}

	return nil
}

//nolint:cyclop,gocognit // Ok.
func testCfgInitAllDefaults(t *testing.T) {
	const errMsg = "unexpected default for %s: Got: %v Wnt: %v"

	savedEnvVars := clearAndCaptureAll()
	defer restoreAll(savedEnvVars)

	initAll()

	if !settingFailFast || !SettingFailFast() {
		t.Fatalf(errMsg, EnvFailFast, settingFailFast, defFailFast)
	}

	if settingBufferSize != defBufferSize ||
		SettingBufferSize() != defBufferSize {
		t.Fatalf(errMsg, EnvBufferSize, settingBufferSize, defBufferSize)
	}

	if settingPermDir != defPermDir ||
		SettingPermDir() != defPermDir {
		t.Fatalf(errMsg, EnvPermDir, settingPermDir, defPermDir)
	}

	if settingPermFile != defPermFile ||
		SettingPermFile() != defPermFile {
		t.Fatalf(errMsg, EnvPermDir, settingPermFile, defPermFile)
	}

	if settingPermExe != defPermExe ||
		SettingPermExe() != defPermExe {
		t.Fatalf(errMsg, EnvPermExe, settingPermExe, defPermExe)
	}

	if settingTmpDir != defTmpDir ||
		SettingTmpDir() != defTmpDir {
		t.Fatalf(errMsg, EnvTmpDir, settingTmpDir, defTmpDir)
	}

	if settingDiffChars != defDiffChars ||
		SettingDiffChars() != defDiffChars {
		t.Fatalf(errMsg, EnvDiffChars, settingDiffChars, defDiffChars)
	}

	if settingDiffSlice != defDiffSlice ||
		SettingDiffSlice() != defDiffSlice {
		t.Fatalf(errMsg, EnvDiffSlice, settingDiffSlice, defDiffSlice)
	}

	if settingMarkWntOn != defMarkWntOn ||
		SettingMarkWntOn() != defMarkWntOn {
		t.Fatalf(errMsg, EnvMarkWntOn, settingMarkWntOn, defMarkWntOn)
	}

	if settingMarkWntOff != defMarkWntOff ||
		SettingMarkWntOff() != defMarkWntOff {
		t.Fatalf(errMsg, EnvMarkWntOff, settingMarkWntOff, defMarkWntOff)
	}

	if settingMarkGotOn != defMarkGotOn ||
		SettingMarkGotOn() != defMarkGotOn {
		t.Fatalf(errMsg, EnvMarkGotOn, settingMarkGotOn, defMarkGotOn)
	}

	if settingMarkGotOff != defMarkGotOff ||
		SettingMarkGotOff() != defMarkGotOff {
		t.Fatalf(errMsg, EnvMarkGotOff, settingMarkGotOff, defMarkGotOff)
	}

	if settingMarkMsgOn != defMarkMsgOn ||
		SettingMarkMsgOn() != defMarkMsgOn {
		t.Fatalf(errMsg, EnvMarkMsgOn, settingMarkMsgOn, defMarkMsgOn)
	}

	if settingMarkMsgOff != defMarkMsgOff ||
		SettingMarkMsgOff() != defMarkMsgOff {
		t.Fatalf(errMsg, EnvMarkMsgOff, settingMarkMsgOff, defMarkMsgOff)
	}

	if settingMarkInsOn != defMarkInsOn ||
		SettingMarkInsOn() != defMarkInsOn {
		t.Fatalf(errMsg, EnvMarkInsOn, settingMarkInsOn, defMarkInsOn)
	}

	if settingMarkInsOff != defMarkInsOff ||
		SettingMarkInsOff() != defMarkInsOff {
		t.Fatalf(errMsg, EnvMarkInsOff, settingMarkInsOff, defMarkInsOff)
	}

	if settingMarkDelOn != defMarkDelOn ||
		SettingMarkDelOn() != defMarkDelOn {
		t.Fatalf(errMsg, EnvMarkDelOn, settingMarkDelOn, defMarkDelOn)
	}

	if settingMarkDelOff != defMarkDelOff ||
		SettingMarkDelOff() != defMarkDelOff {
		t.Fatalf(errMsg, EnvMarkDelOff, settingMarkDelOff, defMarkDelOff)
	}

	if settingMarkChgOn != defMarkChgOn ||
		SettingMarkChgOn() != defMarkChgOn {
		t.Fatalf(errMsg, EnvMarkChgOn, settingMarkChgOn, defMarkChgOn)
	}

	if settingMarkChgOff != defMarkChgOff ||
		SettingMarkChgOff() != defMarkChgOff {
		t.Fatalf(errMsg, EnvMarkChgOff, settingMarkChgOff, defMarkChgOff)
	}

	if settingMarkSepOn != defMarkSepOn ||
		SettingMarkSepOn() != defMarkSepOn {
		t.Fatalf(errMsg, EnvMarkSepOn, settingMarkSepOn, defMarkSepOn)
	}

	if settingMarkSepOff != defMarkSepOff ||
		SettingMarkSepOff() != defMarkSepOff {
		t.Fatalf(errMsg, EnvMarkSepOff, settingMarkSepOff, defMarkSepOff)
	}
}

//nolint:cyclop,gocognit // Ok.
func testCfgInitAllOverrides(t *testing.T) {
	const errMsg = "unexpected override for %s: Got: %v Wnt: %v"

	savedEnvVars := clearAndCaptureAll()
	defer restoreAll(savedEnvVars)

	// Used as another directory to test Tmp Dir overrides.
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("could not retrieve user home directory: " + err.Error())
	}

	err = setupSysEnvVarTest(userHomeDir)
	if err != nil {
		t.Fatal("could not setupSysEnvVars with tmpHome directory: " +
			userHomeDir + ": " + err.Error())
	}

	initAll()

	if settingFailFast || SettingFailFast() {
		t.Fatalf(errMsg, EnvFailFast, settingFailFast, false)
	}

	if settingPermDir != 0o0701 ||
		SettingPermDir() != 0o0701 {
		t.Fatalf(errMsg, EnvPermDir, settingPermDir, 0o0701)
	}

	if settingPermFile != 0o0601 ||
		SettingPermFile() != 0o0601 {
		t.Fatalf(errMsg, EnvPermFile, settingPermFile, 0o0601)
	}

	if settingPermExe != 0o0711 ||
		SettingPermExe() != 0o0711 {
		t.Fatalf(errMsg, EnvPermExe, settingPermExe, 0o0711)
	}

	if settingTmpDir != userHomeDir ||
		SettingTmpDir() != userHomeDir {
		t.Fatalf(errMsg, EnvTmpDir, settingTmpDir, userHomeDir)
	}

	if settingMarkWntOn != "<<WntOn>>" ||
		SettingMarkWntOn() != "<<WntOn>>" {
		t.Fatalf(errMsg, EnvMarkWntOn, settingMarkWntOn, "<<WntOn>>")
	}

	if settingMarkWntOff != "<<WntOff>>" ||
		SettingMarkWntOff() != "<<WntOff>>" {
		t.Fatalf(errMsg, EnvMarkWntOff, settingMarkWntOff, "<<WntOff>>")
	}

	if settingMarkGotOn != "<<GotOn>>" ||
		SettingMarkGotOn() != "<<GotOn>>" {
		t.Fatalf(errMsg, EnvMarkGotOn, settingMarkGotOn, "<<GotOn>>")
	}

	if settingMarkGotOff != "<<GotOff>>" ||
		SettingMarkGotOff() != "<<GotOff>>" {
		t.Fatalf(errMsg, EnvMarkGotOff, settingMarkGotOff, "<<GotOff>>")
	}

	if settingMarkMsgOn != "<<MsgOn>>" ||
		SettingMarkMsgOn() != "<<MsgOn>>" {
		t.Fatalf(errMsg, EnvMarkMsgOn, settingMarkMsgOn, "<<MsgOn>>")
	}

	if settingMarkMsgOff != "<<MsgOff>>" ||
		SettingMarkMsgOff() != "<<MsgOff>>" {
		t.Fatalf(errMsg, EnvMarkMsgOff, settingMarkMsgOff, "<<MsgOff>>")
	}

	if settingMarkInsOn != "<<InsOn>>" ||
		SettingMarkInsOn() != "<<InsOn>>" {
		t.Fatalf(errMsg, EnvMarkInsOn, settingMarkInsOn, "<<InsOn>>")
	}

	if settingMarkInsOff != "<<InsOff>>" ||
		SettingMarkInsOff() != "<<InsOff>>" {
		t.Fatalf(errMsg, EnvMarkInsOff, settingMarkInsOff, "<<InsOff>>")
	}

	if settingMarkDelOn != "<<DelOn>>" ||
		SettingMarkDelOn() != "<<DelOn>>" {
		t.Fatalf(errMsg, EnvMarkDelOn, settingMarkDelOn, "<<DelOn>>")
	}

	if settingMarkDelOff != "<<DelOff>>" ||
		SettingMarkDelOff() != "<<DelOff>>" {
		t.Fatalf(errMsg, EnvMarkDelOff, settingMarkDelOff, "<<DelOff>>")
	}

	if settingMarkChgOn != "<<ChgOn>>" ||
		SettingMarkChgOn() != "<<ChgOn>>" {
		t.Fatalf(errMsg, EnvMarkChgOn, settingMarkChgOn, "<<ChgOn>>")
	}

	if settingMarkChgOff != "<<ChgOff>>" ||
		SettingMarkChgOff() != "<<ChgOff>>" {
		t.Fatalf(errMsg, EnvMarkChgOff, settingMarkChgOff, "<<ChgOff>>")
	}

	if settingMarkSepOn != "<<SepOn>>" ||
		SettingMarkSepOn() != "<<SepOn>>" {
		t.Fatalf(errMsg, EnvMarkSepOn, settingMarkSepOn, "<<SepOn>>")
	}

	if settingMarkSepOff != "<<SepOff>>" ||
		SettingMarkSepOff() != "<<SepOff>>" {
		t.Fatalf(errMsg, EnvMarkSepOff, settingMarkSepOff, "<<SepOff>>")
	}

	if settingDiffChars != 2 ||
		SettingDiffChars() != 2 {
		t.Fatalf(errMsg, EnvDiffChars, settingDiffChars, 2)
	}

	if settingDiffSlice != 4 ||
		SettingDiffSlice() != 4 {
		t.Fatalf(errMsg, EnvDiffSlice, settingDiffSlice, 4)
	}

	if settingBufferSize != 12345 ||
		SettingBufferSize() != 12345 {
		t.Fatalf(errMsg, EnvBufferSize, settingBufferSize, 12345)
	}
}
