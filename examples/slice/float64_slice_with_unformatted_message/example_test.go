package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func scale(vector []float64, factor float64) []float64 {
	scaledVector := make([]float64, len(vector))
	for i, v := range vector {
		scaledVector[i] = v * factor
	}
	return scaledVector
}

// Passing test.
func Test_PASS_Float64SliceWithUnformattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testVector := []float64{1, 2, 3, 4, 5, 6}

	chk.Float64Slice(
		[]float64{2, 4, 6, 8, 10, 12},               // Got.
		scale(testVector, 2),                        // Want.
		1.0,                                         // Tolerance.
		"function scale(", testVector, ", ", 0, ")", // Additional message.
	)
}

// Failing test.
func Test_Fail_Float64SliceWithUnformattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testVector := []float64{1, 2, 3, 4, 5, 6}

	chk.Float64Slice(
		scale(testVector, 2),                        // Got.
		[]float64{2, 4, 6, 8.1, 10, 12},             // Want.
		0.01,                                        // Tolerance.
		"function scale(", testVector, ", ", 0, ")", // Additional message.
	)
}
