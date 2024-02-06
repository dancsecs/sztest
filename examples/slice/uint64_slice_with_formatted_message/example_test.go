package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func scale(vector []uint64, factor uint64) []uint64 {
	scaledVector := make([]uint64, len(vector))
	for i, v := range vector {
		scaledVector[i] = v * factor
	}
	return scaledVector
}

// Passing test.
func Test_PASS_Uint64SliceWithFormattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testVector := []uint64{1, 2, 3, 4, 5, 6}

	chk.Uint64Slicef(
		scale(testVector, 2),                   // Got.
		[]uint64{2, 4, 6, 8, 10, 12},           // Wnt.
		"function scale(%v,%d)", testVector, 0, // Additional message.
	)
}

// Failing test.
func Test_Fail_Uint64SliceWithFormattedMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testVector := []uint64{1, 2, 3, 4, 5, 6}

	chk.Uint64Slicef(
		scale(testVector, 2),                   // Got.
		[]uint64{2, 4, 6, 9, 10, 12, 14},       // Wnt.
		"function scale(%v,%d)", testVector, 0, // Additional message.
	)
}
