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
	"math"
	"strings"
)

type diffLnFmt struct {
	gOffset  int
	wOffset  int
	nbrWidth int
}

func newDiffLnFmt(gLen, wLen int) *diffLnFmt {
	var maxWidth int
	if gLen > wLen {
		maxWidth = int(math.Log10(float64(gLen))) + 1
	} else {
		maxWidth = int(math.Log10(float64(wLen))) + 1
	}

	if maxWidth < 1 {
		maxWidth = 1
	}

	return &diffLnFmt{
		nbrWidth: maxWidth,
		gOffset:  0,
		wOffset:  0,
	}
}

func (lnFmt *diffLnFmt) newOffset(g, w int) *diffLnFmt {
	return &diffLnFmt{
		gOffset:  g + lnFmt.gOffset,
		wOffset:  w + lnFmt.wOffset,
		nbrWidth: lnFmt.nbrWidth,
	}
}

func (lnFmt *diffLnFmt) fmtLnNbr(n int) string {
	if n < 0 {
		return strings.Repeat("-", lnFmt.nbrWidth)
	}

	return fmt.Sprintf("%*.*d", lnFmt.nbrWidth, lnFmt.nbrWidth, n)
}

func (lnFmt *diffLnFmt) same(g, w int, value any) string {
	return "" +
		lnFmt.fmtLnNbr(g+lnFmt.gOffset) +
		":" +
		lnFmt.fmtLnNbr(w+lnFmt.wOffset) +
		" " +
		fmt.Sprint(value)
}

func (lnFmt *diffLnFmt) changed(g, w int, line string) string {
	return "" +
		markAsChg(lnFmt.fmtLnNbr(g+lnFmt.gOffset), "", DiffGot) +
		":" +
		markAsChg("", lnFmt.fmtLnNbr(w+lnFmt.wOffset), DiffWant) +
		" " +
		line
}

func (lnFmt *diffLnFmt) justGot(g int, value any) string {
	return "" +
		markAsIns(lnFmt.fmtLnNbr(g+lnFmt.gOffset)) +
		":" +
		lnFmt.fmtLnNbr(-1) +
		" " +
		markAsIns(fmt.Sprint(value))
}

func (lnFmt *diffLnFmt) justWnt(w int, value any) string {
	return "" +
		lnFmt.fmtLnNbr(-1) +
		":" +
		markAsDel(lnFmt.fmtLnNbr(w+lnFmt.wOffset)) +
		" " +
		markAsDel(fmt.Sprint(value))
}
