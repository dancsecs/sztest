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
	"os"
)

//nolint:goCheckNoGlobals // Ok - initialized by init function.
var (
	defTmpDir         string = os.TempDir()
	settingFailFast   bool
	settingBufferSize int
	settingPermDir    os.FileMode
	settingPermFile   os.FileMode
	settingPermExe    os.FileMode
	settingTmpDir     string
	settingDiffChars  int
	settingDiffSlice  int
	settingMarkWntOn  string
	settingMarkWntOff string
	settingMarkGotOn  string
	settingMarkGotOff string
	settingMarkMsgOn  string
	settingMarkMsgOff string
	settingMarkInsOn  string
	settingMarkInsOff string
	settingMarkDelOn  string
	settingMarkDelOff string
	settingMarkChgOn  string
	settingMarkChgOff string
	settingMarkSepOn  string
	settingMarkSepOff string
)

// ReloadSettings re-initializes the settings permitting an application
// to force certain settings required for embedded runs.
func ReloadSettings() {
	initAll()
}

// SettingFailFast returns the default setting overridden by env settings.
func SettingFailFast() bool {
	return settingFailFast
}

// SettingBufferSize returns the default setting overridden by env settings.
func SettingBufferSize() int {
	return settingBufferSize
}

// SettingPermDir returns the default setting overridden by env settings.
func SettingPermDir() os.FileMode {
	return settingPermDir
}

// SettingPermFile returns the default setting overridden by env settings.
func SettingPermFile() os.FileMode {
	return settingPermFile
}

// SettingPermExe returns the default setting overridden by env settings.
func SettingPermExe() os.FileMode {
	return settingPermExe
}

// SettingTmpDir returns the default setting overridden by env settings.
func SettingTmpDir() string {
	return settingTmpDir
}

// SettingDiffChars returns the default setting overridden by env settings.
func SettingDiffChars() int {
	return settingDiffChars
}

// SettingDiffSlice returns the default setting overridden by env settings.
func SettingDiffSlice() int {
	return settingDiffSlice
}

// SettingMarkWntOn returns the default setting overridden by env settings.
func SettingMarkWntOn() string {
	return settingMarkWntOn
}

// SettingMarkWntOff returns the default setting overridden by env settings.
func SettingMarkWntOff() string {
	return settingMarkWntOff
}

// SettingMarkGotOn returns the default setting overridden by env settings.
func SettingMarkGotOn() string {
	return settingMarkGotOn
}

// SettingMarkGotOff returns the default setting overridden by env settings.
func SettingMarkGotOff() string {
	return settingMarkGotOff
}

// SettingMarkMsgOn returns the default setting overridden by env settings.
func SettingMarkMsgOn() string {
	return settingMarkMsgOn
}

// SettingMarkMsgOff returns the default setting overridden by env settings.
func SettingMarkMsgOff() string {
	return settingMarkMsgOff
}

// SettingMarkInsOn returns the default setting overridden by env settings.
func SettingMarkInsOn() string {
	return settingMarkInsOn
}

// SettingMarkInsOff returns the default setting overridden by env settings.
func SettingMarkInsOff() string {
	return settingMarkInsOff
}

// SettingMarkDelOn returns the default setting overridden by env settings.
func SettingMarkDelOn() string {
	return settingMarkDelOn
}

// SettingMarkDelOff returns the default setting overridden by env settings.
func SettingMarkDelOff() string {
	return settingMarkDelOff
}

// SettingMarkChgOn returns the default setting overridden by env settings.
func SettingMarkChgOn() string {
	return settingMarkChgOn
}

// SettingMarkChgOff returns the default setting overridden by env settings.
func SettingMarkChgOff() string {
	return settingMarkChgOff
}

// SettingMarkSepOn returns the default setting overridden by env settings.
func SettingMarkSepOn() string {
	return settingMarkSepOn
}

// SettingMarkSepOff returns the default setting overridden by env settings.
func SettingMarkSepOff() string {
	return settingMarkSepOff
}
