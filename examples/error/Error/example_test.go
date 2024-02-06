package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_Error(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("error condition")

	chk.Err(
		err,
		"error condition",
	)
}

// Failing test.
func Test_FAIL_Error(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("error condition generated")
	chk.Err(
		err,
		"error condition wanted",
	)
}
