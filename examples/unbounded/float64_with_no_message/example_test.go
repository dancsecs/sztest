package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func fromFloat(f float64) float64 {
	return f / 2.0
}

// Passing test.
func Test_PASS_BoundedFloat64WithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	const valueToTest = 1.234567
	chk.Float64Unbounded(fromFloat(valueToTest), sztest.UnboundedMinOpen, -2.0)
}

// Failing test.
func Test_Fail_BoundedFloat64WithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	const valueToTest = 1.234567
	chk.Float64Unbounded(fromFloat(valueToTest), sztest.UnboundedMinOpen, 2.0)
}
