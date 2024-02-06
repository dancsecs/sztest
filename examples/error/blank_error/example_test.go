package example

import (
	"errors"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_BlankError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("")

	chk.Err(
		// blank "" error will be replaced with "sztest.EmptyErrorMessage"
		err,
		sztest.BlankErrorMessage,
	)
}

// Failing test.
func Test_FAIL_BlankError(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	err := errors.New("")
	chk.Err(
		// blank "" error will be replaced with "sztest.EmptyErrorMessage"
		err,
		"Error message wanted",
	)
}
