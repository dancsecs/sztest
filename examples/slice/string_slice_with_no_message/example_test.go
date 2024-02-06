package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Example function being tested.
func addPrefix(list []string, prefix string) []string {
	prefixedList := make([]string, len(list))
	for i, entry := range list {
		prefixedList[i] = prefix + entry
	}
	return prefixedList
}

// Passing test.
func Test_PASS_StringSliceWithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testList := []string{"Alpha", "Bravo", "Charlie", "Delta"}

	chk.StrSlice(
		addPrefix(testList, "-->"), // Got.
		[]string{
			"-->Alpha",
			"-->Bravo",
			"-->Charlie",
			"-->Delta",
		}, // Want.
	)
}

// Failing test.
func Test_Fail_StringSliceWithNoMessage(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	testList := []string{"Alpha", "Bravo", "Charlie", "Delta"}

	chk.StrSlice(
		addPrefix(testList, "-->"), // Got.
		[]string{
			"-->Alpha",
			"-->Sheen",
			"-->Delta",
			"-->Echo",
		}, // Want.
	)
}
