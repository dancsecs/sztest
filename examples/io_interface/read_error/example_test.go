package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	// Read without anyset will cause immediate EOF.
	_, err := readFile(chk)
	chk.Err(err, "unexpected EOF")

	// Read all the data

	chk.SetIOReaderData("0123456789")

	str, err := readFile(chk)
	chk.NoErr(err)
	chk.Str(str, "0123456789")

	str, err = readFile(chk)
	chk.Str(str, "")
	chk.Err(err, "unexpected EOF")

	// Not enough data

	chk.SetIOReaderData("01234")
	str, err = readFile(chk)
	chk.Err(err, "not enough bytes")
	chk.Str(str, "01234\x00\x00\x00\x00\x00")

	// Fail after a certain number of bytes is read.
	chk.SetIOReaderData("01234567890")
	chk.SetIOReaderError(2, errors.New("only two bytes read"))
	str, err = readFile(chk)
	chk.Err(err, "not enough bytes")
	chk.Str(str, "01\x00\x00\x00\x00\x00\x00\x00\x00")

	str, err = readFile(chk)
	chk.Err(err, "only two bytes read")
	chk.Str(str, "")

	// Setup a direct error to be returned on next read.  No other data needs
	// to be setup

	chk.SetReadError(2962, errors.New("example error on returning 2962"))

	n, err := chk.Read(nil)
	chk.Int(n, 2962)
	chk.Err(err, "example error on returning 2962")

	chk.SetReadError(2963, nil) // no error just a forced count.

	n, err = chk.Read(nil)
	chk.NoErr(err)
	chk.Int(n, 2963)
}
