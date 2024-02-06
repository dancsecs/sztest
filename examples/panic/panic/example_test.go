package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_Panic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			panic("expected panic message")
		},
		"expected panic message",
	)
}

// Failing test.
func Test_FAIL_Panic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	// Failure.  The invoked function's panic message
	// is not what is expected.
	chk.Panic(
		func() {
			panic("this is the panic generated")
		},
		"this is the panic wanted",
	)
}
