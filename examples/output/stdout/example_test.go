package example

import (
	"fmt"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_CaptureStdout(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	arg1 := "1"
	arg2 := "2"

	fmt.Printf("Line %s\n", arg1)
	fmt.Printf("Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)
	// Stdout output as expected?
	chk.Stdout(
		"Line {{arg1}}",
		"Line {{arg2}}",
	)
}

// Failing test.
func Test_FAIL_CaptureStdout(t *testing.T) {
	chk := sztest.CaptureStdout(t)
	defer chk.Release()

	arg1 := "1"
	arg2 := "2"

	fmt.Printf("Line %s\n", arg1)
	fmt.Print("Missing in want\n")
	fmt.Printf("Line %s\n", arg2)

	chk.AddSub("{{arg1}}", arg1)
	chk.AddSub("{{arg2}}", arg2)
	// Stdout output as expected?
	chk.Stdout(
		"Line {{arg1}}",
		"Line {{arg2}}",
	)
}
