package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func addIntegers(int1, int2 int) int {
	return int1 + int2
}

// Passing test.
func Test_PASS_IntegerWithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Int(
		addIntegers(2, 3), // Got.
		5,                 // Want.
	)
}

// Failing test.
func Test_FAIL_IntegerWithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	chk.Int(
		addIntegers(1237456, 1000), // Got.
		1237456,                    // Want.
	)
}
