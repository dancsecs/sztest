package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func fromFloat(f float32) float32 {
	return f / 2.0
}

// Passing test.
func Test_PASS_Float32WithUnformattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	const valueToTest = 1.234567
	chk.Float32(
		fromFloat(valueToTest),                  // Got.
		0.617,                                   // Want.
		0.001,                                   // Tolerance.
		"function fromFloat(", valueToTest, ")", // Additional message.
	)
}

// Failing test.
func Test_Fail_Float32WithUnformattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	const valueToTest = 2.468024
	chk.Float32(
		fromFloat(valueToTest),                  // Got.
		1.2356,                                  // Want.
		0.0005,                                  // Tolerance.
		"function fromFloat(", valueToTest, ")", // Additional message.
	)
}
