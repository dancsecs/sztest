package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_NoPanic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			// Exit function without panicking.
		},
		"", // Permits a NoPanic wnt to be calculated by the test.
	)
}

// Failing test.
func Test_FAIL_NoPanic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			panic("this is the panic generated")
		},
		// Expecting no panic to be thrown
		"",
	)
}
