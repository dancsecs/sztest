package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_StringWithFormattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Strf(
		"This string being tested.",
		"This string being tested.",
		// Optional formatted message.
		"%s message with %s information", "Formatted", "additional",
	)
}

// Failing test.
func Test_FAIL_StringWithFormattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Strf(
		"This (extra) is the Got string being tested.",
		"This is the Wnt string (missing) being tested.",
		// Optional formatted message.
		"%s message with %s information", "Formatted", "additional",
	)
}
