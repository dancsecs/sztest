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
	envFailFast   = "SZTEST_FAIL_FAST"
	envBufferSize = "SZTEST_BUFFER_SIZE"
	envPermDir    = "SZTEST_PERM_DIR"
	envPermFile   = "SZTEST_PERM_FILE"
	envPermExe    = "SZTEST_PERM_EXE"
	envTmpDir     = "SZTEST_TMP_DIR"
	envDiffChars  = "SZTEST_DIFF_CHARS"
	envDiffSlice  = "SZTEST_DIFF_SLICE"
	envMarkWntOn  = "SZTEST_MARK_WNT_ON"
	envMarkWntOff = "SZTEST_MARK_WNT_OFF"
	envMarkGotOn  = "SZTEST_MARK_GOT_ON"
	envMarkGotOff = "SZTEST_MARK_GOT_OFF"
	envMarkMsgOn  = "SZTEST_MARK_MSG_ON"
	envMarkMsgOff = "SZTEST_MARK_MSG_OFF"
	envMarkInsOn  = "SZTEST_MARK_INS_ON"
	envMarkInsOff = "SZTEST_MARK_INS_OFF"
	envMarkDelOn  = "SZTEST_MARK_DEL_ON"
	envMarkDelOff = "SZTEST_MARK_DEL_OFF"
	envMarkChgOn  = "SZTEST_MARK_CHG_ON"
	envMarkChgOff = "SZTEST_MARK_CHG_OFF"
	envMarkSepOn  = "SZTEST_MARK_SEP_ON"
	envMarkSepOff = "SZTEST_MARK_SEP_OFF"
)

const (
	defFailFast   = true
	defBufferSize = 10_000
	defPermDir    = os.FileMode(0700)
	defPermFile   = os.FileMode(0600)
	defPermExe    = os.FileMode(0700)
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
	v, ok := os.LookupEnv(envFailFast)
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
	v, ok := os.LookupEnv(envBufferSize)
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
	v, ok := os.LookupEnv(envPermDir)
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
	v, ok := os.LookupEnv(envPermFile)
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
	v, ok := os.LookupEnv(envPermExe)
	if ok {
		cleanValue, passed := validatePermExe(v)
		if passed {
			r = cleanValue
		}
	}
	settingPermExe = r
}

func initTmpDir() {
	v, ok := os.LookupEnv(envTmpDir)
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
	v, ok := os.LookupEnv(envDiffChars)
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
	v, ok := os.LookupEnv(envDiffSlice)
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
	v, ok := os.LookupEnv(envMarkWntOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkWntOn, defMarkWntOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkWntOn = r
}

func initMarkWntOff() {
	r := defMarkWntOff
	v, ok := os.LookupEnv(envMarkWntOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkWntOff, defMarkWntOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkWntOff = r
}

func initMarkGotOn() {
	r := defMarkGotOn
	v, ok := os.LookupEnv(envMarkGotOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkGotOn, defMarkGotOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkGotOn = r
}

func initMarkGotOff() {
	r := defMarkGotOff
	v, ok := os.LookupEnv(envMarkGotOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkGotOff, defMarkGotOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkGotOff = r
}

func initMarkMsgOn() {
	r := defMarkMsgOn
	v, ok := os.LookupEnv(envMarkMsgOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkMsgOn, defMarkMsgOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkMsgOn = r
}

func initMarkMsgOff() {
	r := defMarkMsgOff
	v, ok := os.LookupEnv(envMarkMsgOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkMsgOff, defMarkMsgOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkMsgOff = r
}

func initMarkInsOn() {
	r := defMarkInsOn
	v, ok := os.LookupEnv(envMarkInsOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkInsOn, defMarkInsOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkInsOn = r
}

func initMarkInsOff() {
	r := defMarkInsOff
	v, ok := os.LookupEnv(envMarkInsOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkInsOff, defMarkInsOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkInsOff = r
}

func initMarkDelOn() {
	r := defMarkDelOn
	v, ok := os.LookupEnv(envMarkDelOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkDelOn, defMarkDelOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkDelOn = r
}

func initMarkDelOff() {
	r := defMarkDelOff
	v, ok := os.LookupEnv(envMarkDelOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkDelOff, defMarkDelOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkDelOff = r
}

func initMarkChgOn() {
	r := defMarkChgOn
	v, ok := os.LookupEnv(envMarkChgOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkChgOn, defMarkChgOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkChgOn = r
}

func initMarkChgOff() {
	r := defMarkChgOff
	v, ok := os.LookupEnv(envMarkChgOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkChgOff, defMarkChgOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkChgOff = r
}

func initMarkSepOn() {
	r := defMarkSepOn
	v, ok := os.LookupEnv(envMarkSepOn)
	if ok {
		cleanValue, passed := validateMark(v, envMarkSepOn, defMarkSepOn)
		if passed {
			r = cleanValue
		}
	}
	settingMarkSepOn = r
}

func initMarkSepOff() {
	r := defMarkSepOff
	v, ok := os.LookupEnv(envMarkSepOff)
	if ok {
		cleanValue, passed := validateMark(v, envMarkSepOff, defMarkSepOff)
		if passed {
			r = cleanValue
		}
	}
	settingMarkSepOff = r
}
