package example

import (
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_GeneralForm(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Don't exit test on first failure.

	// Test panic condition.
	chk.Panic(
		func() {
			Process(0, "failing message")
		},
		"factor (0) out of bounds: \"failing message\"",
	)

	// Test valid operation.
	gotInt, gotStr, gotFloat := Process(2, "Hello")

	chk.Int(gotInt, 4)
	chk.Str(gotStr, "Processed: Hello")
	chk.Float64(gotFloat, 0.6666, 0.0005) // Tolerance (± 0.0005)

	// Check output.
	chk.Stdout(
		"Processing with factor: 2 and message: Hello",
	)

	// Check logging.
	chk.Log(chk.TrimAll(`
      Entered process(0, "failing message")
      factor (0) out of bounds: "failing message"
      Entered process(2, "Hello")
  `))
}

// Failing test.
func Test_FAIL_GeneralForm(t *testing.T) {
	chk := sztest.CaptureLogAndStdout(t)
	defer chk.Release()

	chk.FailFast(false) // Don't exit test on first failure.

	// Test panic condition.
	chk.Panic(
		func() {
			Process(0, "failing message")
		},
		"factor (0) out of bounds: \"wrong message\"",
	)

	// Test valid operation.
	gotInt, gotStr, gotFloat := Process(2, "Hello")

	chk.Int(gotInt, 3)
	chk.Str(gotStr, "Hello", "missing", " ", `"Processed"`, " ", "prefix")
	// Tolerance (± 0.0005)
	chk.Float64f(gotFloat, 0.6661, 0.0005, "off by %f", 0.6661-0.6665)

	// Check output.
	chk.Stdout(
		"Processing with factor: 2 and message: Processed Hello",
	)

	// Check logging.
	chk.Log(chk.TrimAll(`
      Entered process(0, "wrong message")
      factor (0) out of bounds: "wrong message"
      Entered process(2, "Processed Hello")
  `))
}
