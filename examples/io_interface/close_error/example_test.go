package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.SetCloseError(errors.New("the close error"))

	chk.Err(closeFile(chk), "the close error")
}
