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

// ANSI terminal color/style escape codes.
const (
	clrOff       = "\x1b[0m"
	clrBold      = "\x1b[1m"
	clrItalic    = "\x1b[3m"
	clrUnderline = "\x1b[4m"
	clrReverse   = "\x1b[7m"
	clrStrikeout = "\x1b[9m"

	// Basic foreground colors.
	clrBlack   = "\x1b[30m"
	clrRed     = "\x1b[31m"
	clrGreen   = "\x1b[32m"
	clrYellow  = "\x1b[33m"
	clrBlue    = "\x1b[34m"
	clrMagenta = "\x1b[35m"
	clrCyan    = "\x1b[36m"
	clrWhite   = "\x1b[37m"

	// Bright (intense) foreground colors.
	clrHiBlack   = "\x1b[90m"
	clrHiRed     = "\x1b[91m"
	clrHiGreen   = "\x1b[92m"
	clrHiYellow  = "\x1b[93m"
	clrHiBlue    = "\x1b[94m"
	clrHiMagenta = "\x1b[95m"
	clrHiCyan    = "\x1b[96m"
	clrHiWhite   = "\x1b[97m"

	// Basic background colors.
	clrBkBlack   = "\x1b[40m"
	clrBkRed     = "\x1b[41m"
	clrBkGreen   = "\x1b[42m"
	clrBkYellow  = "\x1b[43m"
	clrBkBlue    = "\x1b[44m"
	clrBkMagenta = "\x1b[45m"
	clrBkCyan    = "\x1b[46m"
	clrBkWhite   = "\x1b[47m"

	// Bright (intense) background colors.
	clrBkHiBlack   = "\x1b[100m"
	clrBkHiRed     = "\x1b[101m"
	clrBkHiGreen   = "\x1b[102m"
	clrBkHiYellow  = "\x1b[103m"
	clrBkHiBlue    = "\x1b[104m"
	clrBkHiMagenta = "\x1b[105m"
	clrBkHiCyan    = "\x1b[106m"
	clrBkHiWhite   = "\x1b[107m"
)
