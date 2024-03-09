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

// Environment variable identifiers.
const (
	EnvFailFast   = "SZTEST_FAIL_FAST"
	EnvBufferSize = "SZTEST_BUFFER_SIZE"
	EnvPermDir    = "SZTEST_PERM_DIR"
	EnvPermFile   = "SZTEST_PERM_FILE"
	EnvPermExe    = "SZTEST_PERM_EXE"
	EnvTmpDir     = "SZTEST_TMP_DIR"
	EnvDiffChars  = "SZTEST_DIFF_CHARS"
	EnvDiffSlice  = "SZTEST_DIFF_SLICE"
	EnvMarkWntOn  = "SZTEST_MARK_WNT_ON"
	EnvMarkWntOff = "SZTEST_MARK_WNT_OFF"
	EnvMarkGotOn  = "SZTEST_MARK_GOT_ON"
	EnvMarkGotOff = "SZTEST_MARK_GOT_OFF"
	EnvMarkMsgOn  = "SZTEST_MARK_MSG_ON"
	EnvMarkMsgOff = "SZTEST_MARK_MSG_OFF"
	EnvMarkInsOn  = "SZTEST_MARK_INS_ON"
	EnvMarkInsOff = "SZTEST_MARK_INS_OFF"
	EnvMarkDelOn  = "SZTEST_MARK_DEL_ON"
	EnvMarkDelOff = "SZTEST_MARK_DEL_OFF"
	EnvMarkChgOn  = "SZTEST_MARK_CHG_ON"
	EnvMarkChgOff = "SZTEST_MARK_CHG_OFF"
	EnvMarkSepOn  = "SZTEST_MARK_SEP_ON"
	EnvMarkSepOff = "SZTEST_MARK_SEP_OFF"
)

const (
	defFailFast   = true
	defBufferSize = 10_000
	defPermDir    = os.FileMode(0o0700)
	defPermFile   = os.FileMode(0o0600)
	defPermExe    = os.FileMode(0o0700)
	defDiffChars  = 3
	defDiffSlice  = 1
	defMarkWntOn  = clrCyan
	defMarkWntOff = clrOff
	defMarkGotOn  = clrMagenta
	defMarkGotOff = clrOff
	defMarkMsgOn  = clrBold + clrItalic + clrUnderline
	defMarkMsgOff = clrOff
	defMarkInsOn  = clrGreen + clrReverse
	defMarkInsOff = clrOff
	defMarkDelOn  = clrRed + clrReverse
	defMarkDelOff = clrOff
	defMarkChgOn  = clrBlue + clrReverse
	defMarkChgOff = clrOff
	defMarkSepOn  = clrBkYellow
	defMarkSepOff = clrOff
)

//nolint:goCheckNoInits // Ok.
func init() {
	ReloadSettings()
}

func initAll() {
	initFailFast()
	initBufferSize()
	initPermDir()
	initPermFile()
	initPermExe()
	initTmpDir()
	initDiffChars()
	initDiffSlice()

	initMarkWntOn()
	initMarkWntOff()
	initMarkGotOn()
	initMarkGotOff()
	initMarkMsgOn()
	initMarkMsgOff()
	initMarkInsOn()
	initMarkInsOff()
	initMarkDelOn()
	initMarkDelOff()
	initMarkChgOn()
	initMarkChgOff()
	initMarkSepOn()
	initMarkSepOff()
}

func initFailFast() {
	result := defFailFast
	v, ok := os.LookupEnv(EnvFailFast)
	if ok {
		cleanValue, passed := validateFailFast(v)
		if passed {
			result = cleanValue
		}
	}
	settingFailFast = result
}

func initBufferSize() {
	result := defBufferSize
	v, ok := os.LookupEnv(EnvBufferSize)
	if ok {
		cleanValue, passed := validateBufferSize(v)
		if passed {
			result = cleanValue
		}
	}
	settingBufferSize = result
}

func initPermDir() {
	result := defPermDir
	v, ok := os.LookupEnv(EnvPermDir)
	if ok {
		cleanValue, passed := validatePermDir(v)
		if passed {
			result = cleanValue
		}
	}
	settingPermDir = result
}

func initPermFile() {
	result := defPermFile
	v, ok := os.LookupEnv(EnvPermFile)
	if ok {
		cleanValue, passed := validatePermFile(v)
		if passed {
			result = cleanValue
		}
	}
	settingPermFile = result
}

func initPermExe() {
	result := defPermExe
	v, ok := os.LookupEnv(EnvPermExe)
	if ok {
		cleanValue, passed := validatePermExe(v)
		if passed {
			result = cleanValue
		}
	}
	settingPermExe = result
}

func initTmpDir() {
	v, ok := os.LookupEnv(EnvTmpDir)
	if ok {
		cleanValue, passed := validateTmpDir(v)
		if passed {
			settingTmpDir = cleanValue

			return
		}
	}
	settingTmpDir = os.TempDir()
}

func initDiffChars() {
	result := defDiffChars
	v, ok := os.LookupEnv(EnvDiffChars)
	if ok {
		cleanValue, passed := validateMinRunString(v)
		if passed {
			result = cleanValue
		}
	}
	settingDiffChars = result
}

func initDiffSlice() {
	result := defDiffSlice
	v, ok := os.LookupEnv(EnvDiffSlice)
	if ok {
		cleanValue, passed := validateMinRunSlice(v)
		if passed {
			result = cleanValue
		}
	}
	settingDiffSlice = result
}

func initMarkWntOn() {
	result := defMarkWntOn
	v, ok := os.LookupEnv(EnvMarkWntOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkWntOn, defMarkWntOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkWntOn = result
}

func initMarkWntOff() {
	result := defMarkWntOff
	v, ok := os.LookupEnv(EnvMarkWntOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkWntOff, defMarkWntOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkWntOff = result
}

func initMarkGotOn() {
	result := defMarkGotOn
	v, ok := os.LookupEnv(EnvMarkGotOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkGotOn, defMarkGotOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkGotOn = result
}

func initMarkGotOff() {
	result := defMarkGotOff
	v, ok := os.LookupEnv(EnvMarkGotOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkGotOff, defMarkGotOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkGotOff = result
}

func initMarkMsgOn() {
	result := defMarkMsgOn
	v, ok := os.LookupEnv(EnvMarkMsgOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkMsgOn, defMarkMsgOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkMsgOn = result
}

func initMarkMsgOff() {
	result := defMarkMsgOff
	v, ok := os.LookupEnv(EnvMarkMsgOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkMsgOff, defMarkMsgOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkMsgOff = result
}

func initMarkInsOn() {
	result := defMarkInsOn
	v, ok := os.LookupEnv(EnvMarkInsOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkInsOn, defMarkInsOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkInsOn = result
}

func initMarkInsOff() {
	result := defMarkInsOff
	v, ok := os.LookupEnv(EnvMarkInsOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkInsOff, defMarkInsOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkInsOff = result
}

func initMarkDelOn() {
	result := defMarkDelOn
	v, ok := os.LookupEnv(EnvMarkDelOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkDelOn, defMarkDelOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkDelOn = result
}

func initMarkDelOff() {
	result := defMarkDelOff
	v, ok := os.LookupEnv(EnvMarkDelOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkDelOff, defMarkDelOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkDelOff = result
}

func initMarkChgOn() {
	result := defMarkChgOn
	v, ok := os.LookupEnv(EnvMarkChgOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkChgOn, defMarkChgOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkChgOn = result
}

func initMarkChgOff() {
	result := defMarkChgOff
	v, ok := os.LookupEnv(EnvMarkChgOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkChgOff, defMarkChgOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkChgOff = result
}

func initMarkSepOn() {
	result := defMarkSepOn
	v, ok := os.LookupEnv(EnvMarkSepOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkSepOn, defMarkSepOn)
		if passed {
			result = cleanValue
		}
	}
	settingMarkSepOn = result
}

func initMarkSepOff() {
	result := defMarkSepOff
	v, ok := os.LookupEnv(EnvMarkSepOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkSepOff, defMarkSepOff)
		if passed {
			result = cleanValue
		}
	}
	settingMarkSepOff = result
}
