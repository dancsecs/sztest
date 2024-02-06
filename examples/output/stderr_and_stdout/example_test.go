package example

import (
	"fmt"
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// / Passing test.
func Test_PASS_CaptureStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureStderrAndStdout(t)
	defer chk.Release()

	arg1 := "1"
	arg2 := "2"

	fmt.Printf("Stdout: Line %s\n", arg1)
	fmt.Printf("Stdout: Line %s\n", arg2)

	fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg1)
	fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)

	// Stdout output as expected?
	chk.Stdout(
		"Stdout: Line {{arg1}}",
		"Stdout: Line {{arg2}}",
	)

	// Stderr output as expected?
	chk.Stderr(
		"Stderr: Line {{arg1}}",
		"Stderr: Line {{arg2}}",
	)
}

// Failing test.
func Test_FAIL_CaptureStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureStderrAndStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Process all checks in this test.

	arg1 := "1"
	arg2 := "2"

	fmt.Printf("Stdout: Line %s\n", arg1)
	fmt.Println("Stdout: Missing In Want")
	fmt.Printf("Stdout: Line %s\n", arg2)

	fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg1)
	fmt.Fprintf(os.Stderr, "Stderr: Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)

	// Stdout output as expected?
	chk.Stdout(
		"Stdout: Line {{arg1}}",
		"Stdout: Line {{arg2}}",
	)

	// Stderr output as expected?
	chk.Stderr(
		"Stderr: Line {{arg1}}",
		"StdErr: Missing in got",
		"Stderr: Line {{arg2}}",
	)
}
