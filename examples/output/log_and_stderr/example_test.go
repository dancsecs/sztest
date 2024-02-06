package example

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogAndStderr(t *testing.T) {
	chk := sztest.CaptureLogAndStderr(t)
	defer chk.Release()

	log.Print("logged line 1")
	log.Print("logged line 2")

	fmt.Fprintln(os.Stderr, "stderr line 1")
	fmt.Fprintln(os.Stderr, "stderr line 2")

	// Log output as expected?
	chk.Log(
		"logged line 1",
		"logged line 2",
	)

	chk.Stderr(
		"stderr line 1",
		"stderr line 2",
	)
}

// Failing test.
func Test_FAIL_CaptureLogAndStserr(t *testing.T) {
	chk := sztest.CaptureLogAndStderr(t)
	defer chk.Release()

	chk.FailFast(false) // Process all checks in this test.

	log.Print("logged ONLY In Got")
	log.Print("logged SAME Line 1")
	log.Printf("logged CHANGED: This is first.")
	log.Print("logged SAME Line 2")
	log.Printf("logged CHANGED: This will be second. (Missing in want)")

	fmt.Fprintln(os.Stderr, "this line will be different")

	// Log output as expected?
	chk.Log(
		"logged SAME Line 1",
		"logged CHANGED: (Missing in got) This will be first.",
		"logged SAME Line 2",
		"logged CHANGED: This is second.",
		"logged ONLY in want",
	)

	chk.Stderr(
		"this line will not be the same",
	)
}
