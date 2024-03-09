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
	r := defFailFast
	v, ok := os.LookupEnv(EnvFailFast)
	if ok {
		cleanValue, passed := validateFailFast(v)
		if passed {
			r = cleanValue
		}
	}
	settingFailFast = r
}

func initBufferSize() {
	r := defBufferSize
	v, ok := os.LookupEnv(EnvBufferSize)
	if ok {
		cleanValue, passed := validateBufferSize(v)
		if passed {
			r = cleanValue
		}
	}
	settingBufferSize = r
}

func initPermDir() {
	r := defPermDir
	v, ok := os.LookupEnv(EnvPermDir)
	if ok {
		cleanValue, passed := validatePermDir(v)
		if passed {
			r = cleanValue
		}
	}
	settingPermDir = r
}

func initPermFile() {
	r := defPermFile
	v, ok := os.LookupEnv(EnvPermFile)
	if ok {
		cleanValue, passed := validatePermFile(v)
		if passed {
			r = cleanValue
		}
	}
	settingPermFile = r
}

func initPermExe() {
	r := defPermExe
	v, ok := os.LookupEnv(EnvPermExe)
	if ok {
		cleanValue, passed := validatePermExe(v)
		if passed {
			r = cleanValue
		}
	}
	settingPermExe = r
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
	r := defDiffChars
	v, ok := os.LookupEnv(EnvDiffChars)
	if ok {
		cleanValue, passed := validateMinRunString(v)
		if passed {
			r = cleanValue
		}
	}
	settingDiffChars = r
}

func initDiffSlice() {
	r := defDiffSlice
	v, ok := os.LookupEnv(EnvDiffSlice)
	if ok {
		cleanValue, passed := validateMinRunSlice(v)
		if passed {
			r = cleanValue
		}
	}
	settingDiffSlice = r
}

func initMarkWntOn() {
	r := defMarkWntOn
	v, ok := os.LookupEnv(EnvMarkWntOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkWntOn, defMarkWntOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkWntOn = r
}

func initMarkWntOff() {
	r := defMarkWntOff
	v, ok := os.LookupEnv(EnvMarkWntOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkWntOff, defMarkWntOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkWntOff = r
}

func initMarkGotOn() {
	r := defMarkGotOn
	v, ok := os.LookupEnv(EnvMarkGotOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkGotOn, defMarkGotOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkGotOn = r
}

func initMarkGotOff() {
	r := defMarkGotOff
	v, ok := os.LookupEnv(EnvMarkGotOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkGotOff, defMarkGotOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkGotOff = r
}

func initMarkMsgOn() {
	r := defMarkMsgOn
	v, ok := os.LookupEnv(EnvMarkMsgOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkMsgOn, defMarkMsgOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkMsgOn = r
}

func initMarkMsgOff() {
	r := defMarkMsgOff
	v, ok := os.LookupEnv(EnvMarkMsgOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkMsgOff, defMarkMsgOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkMsgOff = r
}

func initMarkInsOn() {
	r := defMarkInsOn
	v, ok := os.LookupEnv(EnvMarkInsOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkInsOn, defMarkInsOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkInsOn = r
}

func initMarkInsOff() {
	r := defMarkInsOff
	v, ok := os.LookupEnv(EnvMarkInsOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkInsOff, defMarkInsOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkInsOff = r
}

func initMarkDelOn() {
	r := defMarkDelOn
	v, ok := os.LookupEnv(EnvMarkDelOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkDelOn, defMarkDelOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkDelOn = r
}

func initMarkDelOff() {
	r := defMarkDelOff
	v, ok := os.LookupEnv(EnvMarkDelOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkDelOff, defMarkDelOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkDelOff = r
}

func initMarkChgOn() {
	r := defMarkChgOn
	v, ok := os.LookupEnv(EnvMarkChgOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkChgOn, defMarkChgOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkChgOn = r
}

func initMarkChgOff() {
	r := defMarkChgOff
	v, ok := os.LookupEnv(EnvMarkChgOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkChgOff, defMarkChgOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkChgOff = r
}

func initMarkSepOn() {
	r := defMarkSepOn
	v, ok := os.LookupEnv(EnvMarkSepOn)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkSepOn, defMarkSepOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkSepOn = r
}

func initMarkSepOff() {
	r := defMarkSepOff
	v, ok := os.LookupEnv(EnvMarkSepOff)
	if ok {
		cleanValue, passed := validateMark(v, EnvMarkSepOff, defMarkSepOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkSepOff = r
}
