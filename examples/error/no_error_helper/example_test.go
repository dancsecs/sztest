package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoErrorHelper(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := error(nil)

	chk.NoErr(err)
}

// Failing test.
func Test_FAIL_NoErrorHelper(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("unexpected error")

	chk.NoErr(err)
}
