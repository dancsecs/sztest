package example

import (
	"log"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureLog(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	log.Print("Line 1")
	log.Print("Line 2")

	// Log output as expected?
	chk.Log(
		"Line 1",
		"Line 2",
	)
}

// Failing test.
func Test_FAIL_CaptureLog(t *testing.T) {
	chk := sztest.CaptureLog(t)
	defer chk.Release()

	log.Print("ONLY In Got")
	log.Print("SAME Line 1")
	log.Printf("CHANGED: This is first.")
	log.Print("SAME Line 2")
	log.Printf("CHANGED: This will be second. (Missing in want)")

	// Log output as expected?
	chk.Log(
		"SAME Line 1",
		"CHANGED: (Missing in got) This will be first.",
		"SAME Line 2",
		"CHANGED: This is second.",
		"ONLY in want",
	)
}
