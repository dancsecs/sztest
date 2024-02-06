package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := error(nil)

	chk.Err(
		err,
		"", // Empty string represents nil error.  Can be calculated.
	)
}

// Failing test.
func Test_FAIL_NoError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("unexpected error")

	chk.Err(
		err,
		"", // Empty string represents nil error.  Can be calculated.
	)
}
