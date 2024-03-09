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

const errMsg = "error environment variable: %s (got: %s want: %s default: %v)"

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

func validateFailFast(s string) (bool, bool) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "TRUE":
		return true, true
	case "FALSE":
		return false, true
	}
	log.Printf(errMsg, EnvFailFast,
		s,
		validFailFast,
		defFailFast,
	)
	return false, false
}

func valPerm(s, prefix string) (os.FileMode, bool) {
	const base8 = 8
	const bits64 = 64
	if len(s) == 4 && strings.HasPrefix(s, prefix) {
		v, err := strconv.ParseInt(s, base8, bits64)
		if err == nil {
			return os.FileMode(v), true
		}
	}
	return 0, false
}

func validatePermDir(s string) (os.FileMode, bool) {
	p, ok := valPerm(s, "07")
	if !ok {
		log.Printf(errMsg, EnvPermDir,
			s,
			validPermDir,
			settingPermDir,
		)
	}
	return p, ok
}

func validatePermFile(s string) (os.FileMode, bool) {
	p, ok := valPerm(s, "06")
	if !ok {
		log.Printf(errMsg, EnvPermFile,
			s,
			validPermFile,
			settingPermFile,
		)
	}
	return p, ok
}

func validatePermExe(s string) (os.FileMode, bool) {
	p, ok := valPerm(s, "07")
	if !ok {
		log.Printf(errMsg, EnvPermExe,
			s,
			validPermExe,
			settingPermExe,
		)
	}
	return p, ok
}

func validateTmpDir(s string) (string, bool) {
	stat, err := os.Stat(s)
	if err != nil || !stat.IsDir() {
		log.Printf(errMsg, EnvTmpDir, s, validTmpDir, settingTmpDir)
		return "", false
	}
	return s, true
}

var markStyles = map[string]string{
	"DEFAULT":   clrOff,
	"BOLD":      clrBold,
	"ITALIC":    clrItalic,
	"UNDERLINE": clrUnderline,
	"REVERSE":   clrReverse,
	"STRIKEOUT": clrStrikeout,
}

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

//nolint:funlen // ok
func validateMark(colors, envVarName, defaultColor string) (string, bool) {
	ok := true
	found := false
	isDefault := false
	foregroundColor := ""
	backgroundColor := ""
	custom := ""
	styles := make(map[string]bool)
	clr := ""
	sty := ""

	colors = strings.TrimSpace(colors)
	if colors == "" {
		// SZTEST env variable was there and set to nothing ""
		// so we honor that override.  Env variable not there and validate
		// does not get called but default is set.
		return "", true
	}
	splitOnWordANDRegExp := regexp.MustCompile(`_[A|a][N|n][D|d]_`)
	for _, c := range splitOnWordANDRegExp.Split(colors, -1) {
		uC := strings.ToUpper(strings.TrimSpace(c))
		if uC == "" {
			ok = false
			log.Print("missing (empty) attribute")
			break
		}
		clr, found = markFG[uC]
		if found {
			if foregroundColor != "" {
				ok = false
				log.Print("foreground color redefined")
				break
			} else {
				foregroundColor = clr
				continue
			}
		}
		clr, found = markBG[uC]
		if found {
			if backgroundColor != "" {
				ok = false
				log.Print("background color redefined")
				break
			} else {
				backgroundColor = clr
				continue
			}
		}
		sty, found = markStyles[uC]
		if found {
			if styles[sty] {
				ok = false
				log.Print("style redefined")
				break
			} else {
				styles[sty] = true
				if uC == "DEFAULT" {
					isDefault = true
				}
				continue
			}
		}
		if custom != "" {
			ok = false
			log.Print("custom mark redefined")
			break
		} else {
			custom = c
			continue
		}
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

func validateMinRunString(s string) (int, bool) {
	const base10 = 10
	const bits64 = 64
	n, err := strconv.ParseInt(s, base10, bits64)
	if err != nil || n < 1 || n > 5 {
		log.Printf(errMsg, EnvDiffChars,
			s,
			validMinRunString,
			defDiffChars,
		)
		return 0, false
	}
	return int(n), true
}

func validateMinRunSlice(s string) (int, bool) {
	const base10 = 10
	const bits64 = 64
	n, err := strconv.ParseInt(s, base10, bits64)
	if err != nil || n < 1 || n > 5 {
		log.Printf(errMsg, EnvDiffSlice,
			s,
			"1 <= x <= 5",
			defDiffSlice,
		)
		return 0, false
	}
	return int(n), true
}

func validateBufferSize(s string) (int, bool) {
	const base10 = 10
	const bits64 = 64
	n, err := strconv.ParseInt(s, base10, bits64)
	if err != nil || n < 1000 {
		log.Printf(errMsg, EnvBufferSize,
			s,
			validBufferSize,
			defBufferSize,
		)
		return 0, false
	}
	return int(n), true
}
