package example

import (
	"fmt"
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureStderr(t *testing.T) {
	chk := sztest.CaptureStderr(t)
	defer chk.Release()

	arg1 := "1"
	arg2 := "2"

	fmt.Fprintf(os.Stderr, "Line %s\n", arg1)
	fmt.Fprintf(os.Stderr, "Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)
	// Stderr output as expected?
	chk.Stderr(
		"Line {{arg1}}",
		"Line {{arg2}}",
	)
}

// Failing test.
func Test_FAIL_CaptureStderr(t *testing.T) {
	chk := sztest.CaptureStderr(t)
	defer chk.Release()

	arg1 := "1"
	arg2 := "2"

	fmt.Fprintf(os.Stderr, "Line %s\n", arg1)
	fmt.Fprintf(os.Stderr, "Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)
	// Stderr output as expected?
	chk.Stderr(
		"Line {{arg1}}",
		"Missing in got",
		"Line {{arg2}}",
	)
}
