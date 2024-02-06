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

func test_config_init(t *testing.T) {
	t.Run("Defaults", testCfgInit_AllDefaults)
	t.Run("Overrides", testCfgInit_AllOverrides)
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
		capture(envFailFast),
		capture(envPermDir),
		capture(envPermFile),
		capture(envPermExe),
		capture(envTmpDir),
		capture(envMarkWntOn),
		capture(envMarkWntOff),
		capture(envMarkGotOn),
		capture(envMarkGotOff),
		capture(envMarkMsgOn),
		capture(envMarkMsgOff),
		capture(envMarkInsOn),
		capture(envMarkInsOff),
		capture(envMarkDelOn),
		capture(envMarkDelOff),
		capture(envMarkChgOn),
		capture(envMarkChgOff),
		capture(envMarkSepOn),
		capture(envMarkSepOff),
		capture(envDiffChars),
		capture(envDiffSlice),
		capture(envBufferSize),
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

func setupSysEnvVarTest(tmpDir string) error {
	const errMsg = "could not set: %s: %v"
	var err error
	if err = os.Setenv(envFailFast, "false"); err != nil {
		return fmt.Errorf(errMsg, envFailFast, err)
	}

	if err = os.Setenv(envBufferSize, "12345"); err != nil {
		return fmt.Errorf(errMsg, envBufferSize, err)
	}

	if err = os.Setenv(envPermDir, "0701"); err != nil {
		return fmt.Errorf(errMsg, envPermDir, err)
	}

	if err = os.Setenv(envPermFile, "0601"); err != nil {
		return fmt.Errorf(errMsg, envPermFile, err)
	}

	if err = os.Setenv(envPermExe, "0711"); err != nil {
		return fmt.Errorf(errMsg, envPermExe, err)
	}

	if err = os.Setenv(envTmpDir, tmpDir); err != nil {
		return fmt.Errorf(errMsg, envTmpDir, err)
	}

	if err = os.Setenv(envDiffChars, "2"); err != nil {
		return fmt.Errorf(errMsg, envDiffChars, err)
	}

	if err = os.Setenv(envDiffSlice, "4"); err != nil {
		return fmt.Errorf(errMsg, envDiffSlice, err)
	}

	if err = os.Setenv(envMarkWntOn, "<<WntOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkWntOn, err)
	}

	if err = os.Setenv(envMarkWntOff, "<<WntOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkWntOff, err)
	}

	if err = os.Setenv(envMarkGotOn, "<<GotOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkGotOn, err)
	}

	if err = os.Setenv(envMarkGotOff, "<<GotOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkGotOff, err)
	}

	if err = os.Setenv(envMarkMsgOn, "<<MsgOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkMsgOn, err)
	}

	if err = os.Setenv(envMarkMsgOff, "<<MsgOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkMsgOff, err)
	}

	if err = os.Setenv(envMarkInsOn, "<<InsOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkInsOn, err)
	}

	if err = os.Setenv(envMarkInsOff, "<<InsOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkInsOff, err)
	}

	if err = os.Setenv(envMarkDelOn, "<<DelOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkDelOn, err)
	}

	if err = os.Setenv(envMarkDelOff, "<<DelOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkDelOff, err)
	}

	if err = os.Setenv(envMarkChgOn, "<<ChgOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkChgOn, err)
	}

	if err = os.Setenv(envMarkChgOff, "<<ChgOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkChgOff, err)
	}

	if err = os.Setenv(envMarkSepOn, "<<SepOn>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkSepOn, err)
	}

	if err = os.Setenv(envMarkSepOff, "<<SepOff>>"); err != nil {
		return fmt.Errorf(errMsg, envMarkSepOff, err)
	}

	return nil
}

//nolint:gocyclo // Ok.
func testCfgInit_AllDefaults(t *testing.T) {
	const errMsg = "unexpected default for %s: Got: %v Wnt: %v"

	savedEnvVars := clearAndCaptureAll()
	defer restoreAll(savedEnvVars)

	initAll()

	if !settingFailFast || !SettingFailFast() {
		t.Fatalf(errMsg, envFailFast, settingFailFast, defFailFast)
	}

	if settingBufferSize != defBufferSize ||
		SettingBufferSize() != defBufferSize {
		t.Fatalf(errMsg, envBufferSize, settingBufferSize, defBufferSize)
	}

	if settingPermDir != defPermDir ||
		SettingPermDir() != defPermDir {
		t.Fatalf(errMsg, envPermDir, settingPermDir, defPermDir)
	}

	if settingPermFile != defPermFile ||
		SettingPermFile() != defPermFile {
		t.Fatalf(errMsg, envPermDir, settingPermFile, defPermFile)
	}

	if settingPermExe != defPermExe ||
		SettingPermExe() != defPermExe {
		t.Fatalf(errMsg, envPermExe, settingPermExe, defPermExe)
	}

	if settingTmpDir != defTmpDir ||
		SettingTmpDir() != defTmpDir {
		t.Fatalf(errMsg, envTmpDir, settingTmpDir, defTmpDir)
	}

	if settingDiffChars != defDiffChars ||
		SettingDiffChars() != defDiffChars {
		t.Fatalf(errMsg, envDiffChars, settingDiffChars, defDiffChars)
	}

	if settingDiffSlice != defDiffSlice ||
		SettingDiffSlice() != defDiffSlice {
		t.Fatalf(errMsg, envDiffSlice, settingDiffSlice, defDiffSlice)
	}

	if settingMarkWntOn != defMarkWntOn ||
		SettingMarkWntOn() != defMarkWntOn {
		t.Fatalf(errMsg, envMarkWntOn, settingMarkWntOn, defMarkWntOn)
	}

	if settingMarkWntOff != defMarkWntOff ||
		SettingMarkWntOff() != defMarkWntOff {
		t.Fatalf(errMsg, envMarkWntOff, settingMarkWntOff, defMarkWntOff)
	}

	if settingMarkGotOn != defMarkGotOn ||
		SettingMarkGotOn() != defMarkGotOn {
		t.Fatalf(errMsg, envMarkGotOn, settingMarkGotOn, defMarkGotOn)
	}

	if settingMarkGotOff != defMarkGotOff ||
		SettingMarkGotOff() != defMarkGotOff {
		t.Fatalf(errMsg, envMarkGotOff, settingMarkGotOff, defMarkGotOff)
	}

	if settingMarkMsgOn != defMarkMsgOn ||
		SettingMarkMsgOn() != defMarkMsgOn {
		t.Fatalf(errMsg, envMarkMsgOn, settingMarkMsgOn, defMarkMsgOn)
	}

	if settingMarkMsgOff != defMarkMsgOff ||
		SettingMarkMsgOff() != defMarkMsgOff {
		t.Fatalf(errMsg, envMarkMsgOff, settingMarkMsgOff, defMarkMsgOff)
	}

	if settingMarkInsOn != defMarkInsOn ||
		SettingMarkInsOn() != defMarkInsOn {
		t.Fatalf(errMsg, envMarkInsOn, settingMarkInsOn, defMarkInsOn)
	}

	if settingMarkInsOff != defMarkInsOff ||
		SettingMarkInsOff() != defMarkInsOff {
		t.Fatalf(errMsg, envMarkInsOff, settingMarkInsOff, defMarkInsOff)
	}

	if settingMarkDelOn != defMarkDelOn ||
		SettingMarkDelOn() != defMarkDelOn {
		t.Fatalf(errMsg, envMarkDelOn, settingMarkDelOn, defMarkDelOn)
	}

	if settingMarkDelOff != defMarkDelOff ||
		SettingMarkDelOff() != defMarkDelOff {
		t.Fatalf(errMsg, envMarkDelOff, settingMarkDelOff, defMarkDelOff)
	}

	if settingMarkChgOn != defMarkChgOn ||
		SettingMarkChgOn() != defMarkChgOn {
		t.Fatalf(errMsg, envMarkChgOn, settingMarkChgOn, defMarkChgOn)
	}

	if settingMarkChgOff != defMarkChgOff ||
		SettingMarkChgOff() != defMarkChgOff {
		t.Fatalf(errMsg, envMarkChgOff, settingMarkChgOff, defMarkChgOff)
	}

	if settingMarkSepOn != defMarkSepOn ||
		SettingMarkSepOn() != defMarkSepOn {
		t.Fatalf(errMsg, envMarkSepOn, settingMarkSepOn, defMarkSepOn)
	}

	if settingMarkSepOff != defMarkSepOff ||
		SettingMarkSepOff() != defMarkSepOff {
		t.Fatalf(errMsg, envMarkSepOff, settingMarkSepOff, defMarkSepOff)
	}
}

//nolint:gocyclo // Ok.
func testCfgInit_AllOverrides(t *testing.T) {
	const errMsg = "unexpected override for %s: Got: %v Wnt: %v"

	savedEnvVars := clearAndCaptureAll()
	defer restoreAll(savedEnvVars)

	// Used as another directory directory to test Tmp Dir overrides.
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
		t.Fatalf(errMsg, envFailFast, settingFailFast, false)
	}

	if settingPermDir != 0701 ||
		SettingPermDir() != 0701 {
		t.Fatalf(errMsg, envPermDir, settingPermDir, 0701)
	}

	if settingPermFile != 0601 ||
		SettingPermFile() != 0601 {
		t.Fatalf(errMsg, envPermFile, settingPermFile, 0601)
	}

	if settingPermExe != 0711 ||
		SettingPermExe() != 0711 {
		t.Fatalf(errMsg, envPermExe, settingPermExe, 0711)
	}

	if settingTmpDir != userHomeDir ||
		SettingTmpDir() != userHomeDir {
		t.Fatalf(errMsg, envTmpDir, settingTmpDir, userHomeDir)
	}

	if settingMarkWntOn != "<<WntOn>>" ||
		SettingMarkWntOn() != "<<WntOn>>" {
		t.Fatalf(errMsg, envMarkWntOn, settingMarkWntOn, "<<WntOn>>")
	}

	if settingMarkWntOff != "<<WntOff>>" ||
		SettingMarkWntOff() != "<<WntOff>>" {
		t.Fatalf(errMsg, envMarkWntOff, settingMarkWntOff, "<<WntOff>>")
	}

	if settingMarkGotOn != "<<GotOn>>" ||
		SettingMarkGotOn() != "<<GotOn>>" {
		t.Fatalf(errMsg, envMarkGotOn, settingMarkGotOn, "<<GotOn>>")
	}

	if settingMarkGotOff != "<<GotOff>>" ||
		SettingMarkGotOff() != "<<GotOff>>" {
		t.Fatalf(errMsg, envMarkGotOff, settingMarkGotOff, "<<GotOff>>")
	}

	if settingMarkMsgOn != "<<MsgOn>>" ||
		SettingMarkMsgOn() != "<<MsgOn>>" {
		t.Fatalf(errMsg, envMarkMsgOn, settingMarkMsgOn, "<<MsgOn>>")
	}

	if settingMarkMsgOff != "<<MsgOff>>" ||
		SettingMarkMsgOff() != "<<MsgOff>>" {
		t.Fatalf(errMsg, envMarkMsgOff, settingMarkMsgOff, "<<MsgOff>>")
	}

	if settingMarkInsOn != "<<InsOn>>" ||
		SettingMarkInsOn() != "<<InsOn>>" {
		t.Fatalf(errMsg, envMarkInsOn, settingMarkInsOn, "<<InsOn>>")
	}

	if settingMarkInsOff != "<<InsOff>>" ||
		SettingMarkInsOff() != "<<InsOff>>" {
		t.Fatalf(errMsg, envMarkInsOff, settingMarkInsOff, "<<InsOff>>")
	}

	if settingMarkDelOn != "<<DelOn>>" ||
		SettingMarkDelOn() != "<<DelOn>>" {
		t.Fatalf(errMsg, envMarkDelOn, settingMarkDelOn, "<<DelOn>>")
	}

	if settingMarkDelOff != "<<DelOff>>" ||
		SettingMarkDelOff() != "<<DelOff>>" {
		t.Fatalf(errMsg, envMarkDelOff, settingMarkDelOff, "<<DelOff>>")
	}

	if settingMarkChgOn != "<<ChgOn>>" ||
		SettingMarkChgOn() != "<<ChgOn>>" {
		t.Fatalf(errMsg, envMarkChgOn, settingMarkChgOn, "<<ChgOn>>")
	}

	if settingMarkChgOff != "<<ChgOff>>" ||
		SettingMarkChgOff() != "<<ChgOff>>" {
		t.Fatalf(errMsg, envMarkChgOff, settingMarkChgOff, "<<ChgOff>>")
	}

	if settingMarkSepOn != "<<SepOn>>" ||
		SettingMarkSepOn() != "<<SepOn>>" {
		t.Fatalf(errMsg, envMarkSepOn, settingMarkSepOn, "<<SepOn>>")
	}

	if settingMarkSepOff != "<<SepOff>>" ||
		SettingMarkSepOff() != "<<SepOff>>" {
		t.Fatalf(errMsg, envMarkSepOff, settingMarkSepOff, "<<SepOff>>")
	}

	if settingDiffChars != 2 ||
		SettingDiffChars() != 2 {
		t.Fatalf(errMsg, envDiffChars, settingDiffChars, 2)
	}

	if settingDiffSlice != 4 ||
		SettingDiffSlice() != 4 {
		t.Fatalf(errMsg, envDiffSlice, settingDiffSlice, 4)
	}

	if settingBufferSize != 12345 ||
		SettingBufferSize() != 12345 {
		t.Fatalf(errMsg, envBufferSize, settingBufferSize, 12345)
	}
}
