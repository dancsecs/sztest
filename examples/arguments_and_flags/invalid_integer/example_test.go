package example

import (
	"flag"
	"testing"

	"github.com/dancsecs/sztest"
)

func main() {
	var intValue int
	flag.IntVar(&intValue, "n", 10,
		"usage of int value",
	)
	flag.Parse()
}

func Test_ArgsAndFlags_InvalidInteger(t *testing.T) {
	chk := sztest.CaptureStderr(t)
	defer chk.Release()

	args := []string{
		"program/name",
		"-n",
		"thisIsNotAnInteger",
	}
	chk.SetupArgsAndFlags(args)

	chk.Panic(
		main,
		"invalid value \"thisIsNotAnInteger\" for flag -n: parse error",
	)

	chk.Stderr(
		"invalid value \"thisIsNotAnInteger\" for flag -n: parse error",
		"Usage of program/name:",
		"\\s -n int", // Note: initial leading space.
		"\\s   \tusage of int value (default 10)",
	)
}
