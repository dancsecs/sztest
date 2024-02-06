package example

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TemporaryDirectoryTree(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	// Run manually in function.
	//  defer chk.Release()

	rootDir := chk.CreateTmpDir()

	appDir := chk.CreateTmpSubDir(rootDir, "myApp")

	data := chk.CreateTmpSubDir(appDir, "data")

	// Different ways to create children.
	data1 := chk.CreateTmpSubDir(data, "d1")
	data2 := chk.CreateTmpSubDir(appDir, "data", "d2")
	data3 := chk.CreateTmpSubDir(rootDir, "myApp", "data", "d3")
	data4 := chk.CreateTmpSubDir("myApp", "data", "d4")
	data5 := chk.CreateTmpSubDir("myApp/data/d5")
	dataX := chk.CreateTmpSubDir(rootDir, "myApp/data", "a/b/c")

	chkDirMade := func(got, wnt string) {
		chk.Str(got, wnt)
		// Make sure directory are there.
		stat, err := os.Stat(got)
		chk.NoErr(err)
		chk.True(stat.IsDir())
	}

	chkDirMade(appDir, filepath.Join(rootDir, "myApp"))
	chkDirMade(data1, filepath.Join(rootDir, "myApp", "data", "d1"))
	chkDirMade(data2, filepath.Join(rootDir, "myApp", "data", "d2"))
	chkDirMade(data3, filepath.Join(rootDir, "myApp", "data", "d3"))
	chkDirMade(data4, filepath.Join(rootDir, "myApp", "data", "d4"))
	chkDirMade(data5, filepath.Join(rootDir, "myApp", "data", "d5"))
	chkDirMade(dataX, filepath.Join(rootDir, "myApp", "data", "a", "b", "c"))

	// Manually run Release to purge all tmp files.
	chk.Release()

	chkDirGone := func(d string) {
		_, err := os.Stat(d)
		chk.Err(err, "stat "+d+": no such file or directory")
	}

	chkDirGone(appDir)
	chkDirGone(data1)
	chkDirGone(data2)
	chkDirGone(data3)
	chkDirGone(data4)
	chkDirGone(data5)
	chkDirGone(dataX)
}
