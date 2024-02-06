package example

import (
	"os"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TemporaryDirectory(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	tmpDir := chk.CreateTmpDir()

	stat, err := os.Stat(tmpDir)
	chk.NoErr(err)
	chk.NotNil(stat)
	chk.True(stat.IsDir())
}
