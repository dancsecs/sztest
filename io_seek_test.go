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
	"errors"
	"io"
	"testing"
)

func tstChkIoSeek(t *testing.T) {
	t.Run("SeekError", chkIOSeekTestSetSeekError)
}

func chkIOSeekTestSetSeekError(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	p, err := chk.Seek(2, io.SeekStart)
	chk.Int64(p, 0)
	chk.NoErr(err)

	chk.SetSeekError(24, errors.New("the seek error"))
	p, err = chk.Seek(2, io.SeekStart)
	chk.Int64(p, 24)
	chk.Err(err, "the seek error")
}
