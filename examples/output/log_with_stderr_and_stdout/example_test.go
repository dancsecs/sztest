package example

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLogWithStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureLogWithStderrAndStdout(t)
	defer chk.Release()

	log.Print("logged line 1")

	fmt.Fprintln(os.Stderr, "stderr line 1")

	log.Print("logged line 2")

	fmt.Fprintln(os.Stderr, "stderr line 2")

	fmt.Println("stdout line 1")

	// Log output as expected? (either chk.Log or chk.Stderr)
	chk.Log(
		"logged line 1",
		"stderr line 1",
		"logged line 2",
		"stderr line 2",
	)

	chk.Stdout(
		"stdout line 1",
	)
}

// Failing test.
func Test_FAIL_CaptureLogWithStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureLogWithStderrAndStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Process all checks in this test.

	log.Print("logged ONLY In Got")
	log.Print("logged SAME Line 1")

	fmt.Fprintln(os.Stderr, "this stderr line will be the same")

	log.Printf("logged CHANGED: This is first.")
	log.Print("logged SAME Line 2")
	log.Printf("logged CHANGED: This will be second. (Missing in want)")

	// Log output as expected? (either chk.Log or chk.Stderr)
	chk.Stderr(
		"logged SAME Line 1",
		"logged CHANGED: (Missing in got) This will be first.",
		"this stderr line will be the same",
		"logged SAME Line 2",
		"logged CHANGED: This is second.",
		"logged ONLY in want",
	)

	chk.Stdout(
		"stdout line 2",
	)
}
