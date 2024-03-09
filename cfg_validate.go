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
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	errMsg = "error environment variable: %s (got: %s want: %s default: %v)"
	base8  = 8
	base10 = 10
	bits64 = 64
)

const (
	validFailFast     = "true | false"
	validPermDir      = "07oo"
	validPermFile     = "06oo"
	validPermExe      = "07oo"
	validTmpDir       = "valid directory"
	validColor        = "valid color, style or custom"
	validMinRunString = "1 <= x <= 5"
	validMinRunSlice  = "1 <= x <= 5"
	validBufferSize   = "x >= 1000"
)

func validateFailFast(rawSetting string) (bool, bool) {
	switch strings.ToUpper(strings.TrimSpace(rawSetting)) {
	case "TRUE":
		return true, true
	case "FALSE":
		return false, true
	}

	log.Printf(errMsg, EnvFailFast,
		rawSetting,
		validFailFast,
		defFailFast,
	)

	return false, false
}

func valPerm(s, prefix string) (os.FileMode, bool) {
	if len(s) == 4 && strings.HasPrefix(s, prefix) {
		v, err := strconv.ParseInt(s, base8, bits64)
		if err == nil {
			return os.FileMode(v), true
		}
	}

	return 0, false
}

func validatePermDir(rawSetting string) (os.FileMode, bool) {
	permission, ok := valPerm(rawSetting, "07")
	if !ok {
		log.Printf(errMsg, EnvPermDir,
			rawSetting,
			validPermDir,
			settingPermDir,
		)
	}

	return permission, ok
}

func validatePermFile(rawSetting string) (os.FileMode, bool) {
	permission, ok := valPerm(rawSetting, "06")
	if !ok {
		log.Printf(errMsg, EnvPermFile,
			rawSetting,
			validPermFile,
			settingPermFile,
		)
	}

	return permission, ok
}

func validatePermExe(rawSetting string) (os.FileMode, bool) {
	permission, ok := valPerm(rawSetting, "07")
	if !ok {
		log.Printf(errMsg, EnvPermExe,
			rawSetting,
			validPermExe,
			settingPermExe,
		)
	}

	return permission, ok
}

func validateTmpDir(rawSetting string) (string, bool) {
	stat, err := os.Stat(rawSetting)
	if err != nil || !stat.IsDir() {
		log.Printf(errMsg, EnvTmpDir, rawSetting, validTmpDir, settingTmpDir)

		return "", false
	}

	return rawSetting, true
}

//nolint:gochecknoglobals // Ok.
var markStyles = map[string]string{
	"DEFAULT":   clrOff,
	"BOLD":      clrBold,
	"ITALIC":    clrItalic,
	"UNDERLINE": clrUnderline,
	"REVERSE":   clrReverse,
	"STRIKEOUT": clrStrikeout,
}

//nolint:gochecknoglobals // Ok.
var markFG = map[string]string{
	"BLACK":      clrBlack,
	"HI-BLACK":   clrHiBlack,
	"RED":        clrRed,
	"HI-RED":     clrHiRed,
	"GREEN":      clrGreen,
	"HI-GREEN":   clrHiGreen,
	"YELLOW":     clrYellow,
	"HI-YELLOW":  clrHiYellow,
	"BLUE":       clrBlue,
	"HI-BLUE":    clrHiBlue,
	"MAGENTA":    clrMagenta,
	"HI-MAGENTA": clrHiMagenta,
	"CYAN":       clrCyan,
	"HI-CYAN":    clrHiCyan,
	"WHITE":      clrWhite,
	"HI-WHITE":   clrHiWhite,
}

//nolint:gochecknoglobals // Ok.
var markBG = map[string]string{
	"BK-BLACK":      clrBkBlack,
	"BK-HI-BLACK":   clrBkHiBlack,
	"BK-RED":        clrBkRed,
	"BK-HI-RED":     clrBkHiRed,
	"BK-GREEN":      clrBkGreen,
	"BK-HI-GREEN":   clrBkHiGreen,
	"BK-YELLOW":     clrBkYellow,
	"BK-HI-YELLOW":  clrBkHiYellow,
	"BK-BLUE":       clrBkBlue,
	"BK-HI-BLUE":    clrBkHiBlue,
	"BK-MAGENTA":    clrBkMagenta,
	"BK-HI-MAGENTA": clrBkHiMagenta,
	"BK-CYAN":       clrBkCyan,
	"BK-HI-CYAN":    clrBkHiCyan,
	"BK-WHITE":      clrBkWhite,
	"BK-HI-WHITE":   clrBkHiWhite,
}

//nolint:funlen,cyclop // ok
func validateMark(colors, envVarName, defaultColor string) (string, bool) {
	var (
		found bool
		clr   string
		sty   string
	)

	ok := true
	isDefault := false
	foregroundColor := ""
	backgroundColor := ""
	custom := ""
	styles := make(map[string]bool)

	colors = strings.TrimSpace(colors)
	if colors == "" {
		// SZTEST env variable was there and set to nothing ""
		// so we honor that override.  Env variable not there and validate
		// does not get called but default is set.
		return "", true
	}

	splitOnWordANDRegExp := regexp.MustCompile(`_[A|a][N|n][D|d]_`)

	for _, clrEntry := range splitOnWordANDRegExp.Split(colors, -1) {
		uClrEntry := strings.ToUpper(strings.TrimSpace(clrEntry))
		if uClrEntry == "" {
			ok = false

			log.Print("missing (empty) attribute")

			break
		}

		clr, found = markFG[uClrEntry]

		if found {
			if foregroundColor != "" {
				ok = false

				log.Print("foreground color redefined")

				break
			}

			foregroundColor = clr

			continue
		}

		clr, found = markBG[uClrEntry]

		if found {
			if backgroundColor != "" {
				ok = false

				log.Print("background color redefined")

				break
			}

			backgroundColor = clr

			continue
		}

		sty, found = markStyles[uClrEntry]

		if found {
			if styles[sty] {
				ok = false

				log.Print("style redefined")

				break
			}

			styles[sty] = true

			if uClrEntry == "DEFAULT" {
				isDefault = true
			}

			continue
		}

		if custom != "" {
			ok = false

			log.Print("custom mark redefined")

			break
		}

		custom = clrEntry

		continue
	}

	if ok && isDefault && foregroundColor+backgroundColor+custom != "" {
		ok = false

		log.Print("default style must be defined by itself")
	}

	if ok {
		style := ""
		for k := range styles {
			style += k
		}

		return foregroundColor + backgroundColor + style + custom, ok
	}

	log.Printf(errMsg, envVarName,
		colors,
		validColor,
		defaultColor,
	)

	return "", false
}

func validateMinRunString(rawSetting string) (int, bool) {
	minRun64, err := strconv.ParseInt(rawSetting, base10, bits64)
	if err != nil || minRun64 < 1 || minRun64 > 5 {
		log.Printf(errMsg, EnvDiffChars,
			rawSetting,
			validMinRunString,
			defDiffChars,
		)

		return 0, false
	}

	return int(minRun64), true
}

func validateMinRunSlice(rawSetting string) (int, bool) {
	minRun64, err := strconv.ParseInt(rawSetting, base10, bits64)
	if err != nil || minRun64 < 1 || minRun64 > 5 {
		log.Printf(errMsg, EnvDiffSlice,
			rawSetting,
			"1 <= x <= 5",
			defDiffSlice,
		)

		return 0, false
	}

	return int(minRun64), true
}

func validateBufferSize(rawSetting string) (int, bool) {
	bufSize64, err := strconv.ParseInt(rawSetting, base10, bits64)

	if err != nil || bufSize64 < 1000 {
		log.Printf(errMsg, EnvBufferSize,
			rawSetting,
			validBufferSize,
			defBufferSize,
		)

		return 0, false
	}

	return int(bufSize64), true
}
