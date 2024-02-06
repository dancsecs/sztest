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

// Passing test.
func Test_ArgsAndFlags_SingleGoodFlag(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	args := []string{
		"program/name",
		"-s",
		"str from arg",
	}
	chk.SetupArgsAndFlags(args)

	main()

	chk.Str(processedArgument, "Received: str from arg")
}
