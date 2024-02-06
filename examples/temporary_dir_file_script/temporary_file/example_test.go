package example

import (
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TemporaryFile(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	data := "25\n" +
		"50"

	filename := chk.CreateTmpFile([]byte(data))

	dataRead, err := os.ReadFile(filename)

	chk.NoErr(err)
	chk.NotNil(dataRead)
	chk.Str(string(dataRead), data)
}
