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
	"os"
)

//nolint:goCheckNoGlobals // Ok - initialized by init function.
var (
	defTmpDir         = os.TempDir()
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

// ReloadSettings re-initializes the settings maintaining global configuration
// that control defaults such as permissions, temporary directories, diff
// granularity, and ANSI markup styles. Each setting can be overridden by
// environment variables, or falls back to a built-in default if unset. Tests
// can reload the settings explicitly with ReloadSettings().
//
// Most of these values are surfaced through accessor functions (e.g.,
// SettingPermFile(), SettingDiffChars(), SettingMarkWntOn()) so that
// code and tests always consult the resolved value rather than reading
// environment variables directly. This makes behavior deterministic and
// consistent across environments.
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

// SettingDiffChars returns the minimum number of consecutive matching
// characters required within a line for sztest to treat regions of
// `got` and `wnt` strings as identical when computing diffs. In effect,
// this defines the size of the "diff window" horizontally across a line.
// Smaller values increase sensitivity but may produce noisier diffs.
func SettingDiffChars() int {
	return settingDiffChars
}

// SettingDiffSlice returns the minimum number of consecutive matching lines
// required within two slices for sztest to treat regions as identical when
// computing diffs. This is the vertical "diff window" size. Lower values
// highlight finer-grained changes, higher values collapse noise and make
// large blocks of similarity clearer.
func SettingDiffSlice() int {
	return settingDiffSlice
}

// SettingMarkWntOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkWntOn() string {
	return settingMarkWntOn
}

// SettingMarkWntOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkWntOn(), delimiting where the highlight
// decoration stops.
func SettingMarkWntOff() string {
	return settingMarkWntOff
}

// SettingMarkGotOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkGotOn() string {
	return settingMarkGotOn
}

// SettingMarkGotOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkGotOn(), delimiting where the highlight
// decoration stops.
func SettingMarkGotOff() string {
	return settingMarkGotOff
}

// SettingMarkMsgOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkMsgOn() string {
	return settingMarkMsgOn
}

// SettingMarkMsgOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkMsgOn(), delimiting where the highlight
// decoration stops.
func SettingMarkMsgOff() string {
	return settingMarkMsgOff
}

// SettingMarkInsOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkInsOn() string {
	return settingMarkInsOn
}

// SettingMarkInsOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkInsOn(), delimiting where the highlight
// decoration stops.
func SettingMarkInsOff() string {
	return settingMarkInsOff
}

// SettingMarkDelOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkDelOn() string {
	return settingMarkDelOn
}

// SettingMarkDelOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkDelOn(), delimiting where the highlight
// decoration stops.
func SettingMarkDelOff() string {
	return settingMarkDelOff
}

// SettingMarkChgOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkChgOn() string {
	return settingMarkChgOn
}

// SettingMarkChgOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkChgOn(), delimiting where the highlight
// decoration stops.
func SettingMarkChgOff() string {
	return settingMarkChgOff
}

// SettingMarkSepOn returns the resolved "wanted value start" marker string.
// This may be an ANSI escape sequence or plain text decoration, and is used
// when highlighting differences in test output. A blank string disables
// markup for this element.
func SettingMarkSepOn() string {
	return settingMarkSepOn
}

// SettingMarkSepOff returns the resolved "wanted value end" marker string.
// This complements SettingMarkSepOn(), delimiting where the highlight
// decoration stops.
func SettingMarkSepOff() string {
	return settingMarkSepOff
}
