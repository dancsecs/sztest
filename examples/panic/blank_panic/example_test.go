package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_BlankPanic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			panic("") // Blank panic message.
		},
		sztest.BlankPanicMessage, // Panic without message is expected.
		// Message returned in place of an empty string representing
		// that an empty ("") panic was issued by the function.
	)
}

// Failing test.
func Test_FAIL_BlankPanic(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Panic(
		func() {
			panic("") // Empty panic message will be flagged.
		},
		"",
		// sztest.BlankPanicMessage will be returned.
	)
}
