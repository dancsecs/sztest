package example

import (
	"flag"
	"testing"

	"github.com/dancsecs/sztest"
)

var processedArgument string

func main() {
	var strValue string
	flag.StringVar(&strValue, "s", "defaultStrValue",
		"usage of default string value",
	)
	flag.Parse()

	processedArgument = "Received: " + strValue
}

func Test_ArgsAndFlags_InvalidFlag(t *testing.T) {
	chk := sztest.CaptureStderr(t)
	defer chk.Release()

	chk.SetArgs(
		"program/name",
		"-x",
		"str from arg",
	)

	chk.Panic(
		main,
		"flag provided but not defined: -x",
	)

	chk.Str(processedArgument, "") // Not processed.

	chk.Stderr(
		"flag provided but not defined: -x",
		"Usage of program/name:",
		"\\s -s string", // Note: initial leading space.
		"\\s   \tusage of default string value (default \"defaultStrValue\")",
	)
}
