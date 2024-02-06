package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.SetSeekError(34, errors.New("past end of file"))

	p, err := seekFile(chk, 962)
	chk.Err(err, "past end of file")
	chk.Int64(p, 34)
}
