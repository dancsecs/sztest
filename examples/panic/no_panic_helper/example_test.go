package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoPanicHelper(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.NoPanic(
		func() {
			// Exit function without panicking.
		},
	)
}

// Failing test.
func Test_FAIL_NoPanicHelper(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.NoPanic(
		func() {
			panic("panic message") // Unexpected Panic
		},
	)
}
