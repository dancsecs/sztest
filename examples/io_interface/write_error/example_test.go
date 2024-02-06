package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	// Read without anything .
	c, err := writeFile(chk)
	chk.NoErr(err)
	chk.Int(c, 10)
	chk.Str(string(chk.GetIOWriterData()), "0123456789")

	chk.SetIOWriterError(8, errors.New("Run out of space after 8 chars"))
	c, err = writeFile(chk)
	chk.Err(err, "Run out of space after 8 chars")
	chk.Int(c, 8)
	chk.Str(string(chk.GetIOWriterData()), "01234567")

	// Just set a write error and count to be returned on the next write.

	chk.SetWriteError(37, errors.New("the write error"))
	c, err = writeFile(chk)

	chk.Err(err, "the write error")
	chk.Int(c, 37)
}
