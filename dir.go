/*
   Golang test helper library: sztest.
   Copyright (C) 2023-2025 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package sztest

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// SetPermDir updates the default os.FileMode used for creating
// directories in temporary test setups, returning the previous
// value. The setting applies only to the current test. If not
// explicitly set, the default is taken from the SZTEST_PERM_DIR
// environment variable or falls back to 0o0700.
func (chk *Chk) SetPermDir(p os.FileMode) os.FileMode {
	lastPerm := settingPermDir
	settingPermDir = p

	return lastPerm
}

// SetPermFile updates the default os.FileMode used for creating
// regular files in temporary test setups, returning the previous
// value. The setting applies only to the current test. If not
// explicitly set, the default is taken from the SZTEST_PERM_FILE
// environment variable or falls back to 0o0600.
func (chk *Chk) SetPermFile(p os.FileMode) os.FileMode {
	lastPerm := settingPermFile
	settingPermFile = p

	return lastPerm
}

// SetPermExe updates the default os.FileMode used for creating
// executable files in temporary test setups, returning the
// previous value. The setting applies only to the current test.
// If not explicitly set, the default is taken from the
// SZTEST_PERM_EXE environment variable or falls back to 0o0700.
func (chk *Chk) SetPermExe(p os.FileMode) os.FileMode {
	lastPerm := settingPermExe
	settingPermExe = p

	return lastPerm
}

// SetTmpDir overrides the root directory used when creating
// temporary files and directories, returning the previous value.
// The setting applies only to the current test. By default, the
// root is taken from the SZTEST_TMP_DIR environment variable or
// falls back to /tmp.
func (chk *Chk) SetTmpDir(dir string) string {
	lastTmpDir := settingTmpDir

	if dir == "" {
		initTmpDir()

		dir = settingTmpDir
	} else if !filepath.IsAbs(dir) {
		initTmpDir()

		dir = filepath.Join(settingTmpDir, dir)
	}

	fi, err := os.Stat(dir)

	if err != nil || !fi.IsDir() {
		chk.t.Helper()
		chk.Error("invalid directory: ", dir)
	} else {
		settingTmpDir = dir
	}

	return lastTmpDir
}

func removeTestDir(dir string) error {
	fi, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err == nil && !fi.IsDir() {
		err = fmt.Errorf("%w: %q", ErrInvalidDirectory, dir)
	}

	if err == nil {
		err = os.Chmod(dir, settingPermDir)
	}

	if err == nil {
		err = os.RemoveAll(dir)
	}

	return err //nolint:wrapcheck // Ok.
}

func removeTestFile(path string) error {
	fi, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err == nil && fi.IsDir() {
		err = fmt.Errorf("%w: %q", ErrInvalidFile, path)
	}

	if err == nil {
		err = os.Chmod(path, settingPermFile)
	}

	if err == nil {
		err = os.Remove(path)
	}

	return err //nolint:wrapcheck // Ok.
}

func (chk *Chk) setupPath(tmpDir, dir, fName string) (string, string) {
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(tmpDir, dir)
	}

	if fName == "" {
		id := chk.nextTmpID
		chk.nextTmpID++
		fName = fmt.Sprint("tmpFile", id, ".tmp")
	}

	return dir, fName
}

func (chk *Chk) createFile(
	dir, fName string,
	data []byte,
	perm os.FileMode,
) string {
	chk.t.Helper()

	tmpDir := chk.CreateTmpDir()
	dir, fName = chk.setupPath(tmpDir, dir, fName)

	dirStat, err := os.Stat(dir)
	if err != nil || !dirStat.IsDir() {
		err = fmt.Errorf("%w: %q", ErrInvalidDirectory, dir)
	}

	path := filepath.Join(dir, fName)
	if err == nil {
		err = removeTestFile(path)
	}

	if err == nil {
		err = os.WriteFile(path, data, perm)
		if err == nil {
			chk.PushPreReleaseFunc(func() error {
				if chk.faultCount == 0 && !chk.keepTmpFiles {
					return removeTestFile(path)
				}

				return nil
			})
		}
	}

	if err != nil {
		chk.Error("createFile cause: ", err)
	}

	return path
}

// CreateTmpFile removes any existing file of the same name and creates a new
// file in the test’s root temporary directory, writing the provided data.
// File permissions follow the current file mode setting. Relative paths are
// placed under the root directory; absolute paths are not allowed here.
// Unless KeepTmpFiles is called, the file is removed automatically when the
// test completes successfully.
func (chk *Chk) CreateTmpFile(data []byte) string {
	chk.t.Helper()

	return chk.CreateTmpFileIn("", data)
}

// CreateTmpFileIn removes any existing file and creates a new file in the
// specified path, writing the provided data. The path may be relative (in
// which case it is resolved under the test’s root directory) or absolute,
// but absolute paths must begin with the test’s root directory. File
// permissions follow the current file mode setting. Unless KeepTmpFiles
// is called, the file is removed automatically when the test completes
// successfully.
func (chk *Chk) CreateTmpFileIn(path string, data []byte) string {
	chk.t.Helper()

	return chk.CreateTmpFileAs(path, "", data)
}

// CreateTmpFileAs removes any existing file and creates a new file with
// the specified name in the provided path, writing the given data. The
// path rules are the same as for CreateTmpFileIn, and file permissions
// follow the current file mode setting. Unless KeepTmpFiles is called,
// the file is removed automatically when the test completes successfully.
func (chk *Chk) CreateTmpFileAs(path, fName string, data []byte) string {
	chk.t.Helper()

	return chk.createFile(
		path,
		fName,
		data,
		settingPermFile,
	)
}

// CreateTmpUnixScript removes any existing file and creates a new Unix
// script in the test’s root temporary directory with the provided lines.
// Script permissions follow the current executable mode setting. Relative
// paths are placed under the root directory. Unless KeepTmpFiles is called,
// the script is removed automatically when the test completes successfully.
func (chk *Chk) CreateTmpUnixScript(lines []string) string {
	chk.t.Helper()
	chk.CreateTmpDir()

	return chk.CreateTmpUnixScriptIn(
		"",
		lines,
	)
}

// CreateTmpUnixScriptIn removes any existing file and creates a new Unix
// script in the specified path with the provided lines. Paths may be relative
// (resolved under the test’s root directory) or absolute (must begin with the
// root directory). Script permissions follow the current executable mode
// setting. Unless KeepTmpFiles is called, the script is removed automatically
// when the test completes successfully.
func (chk *Chk) CreateTmpUnixScriptIn(path string, lines []string) string {
	chk.t.Helper()

	return chk.CreateTmpUnixScriptAs(
		path,
		"",
		lines,
	)
}

// CreateTmpUnixScriptAs removes any existing file and creates a new Unix
// script with the given name in the specified path, writing the provided
// lines. Paths follow the same rules as CreateTmpUnixScriptIn. Script
// permissions follow the current executable mode setting. Unless KeepTmpFiles
// is called, the script is removed automatically when the test completes
// successfully.
func (chk *Chk) CreateTmpUnixScriptAs(
	path, fName string,
	lines []string,
) string {
	chk.t.Helper()

	// Strips away any leading blank lines and leading spaces on the first line
	// generally the #!/shell opening line
	cleanScript := ""
	trimFirst := 0

	const spaceCutouts = " \t"

	for _, entry := range lines {
		for _, l := range strings.Split(entry, "\n") {
			cleanLine := strings.TrimRight(l, spaceCutouts)
			if cleanScript == "" { //nolint:nestif // Ok.
				if cleanLine != "" {
					// found first non blank line
					ccl := strings.TrimLeft(cleanLine, spaceCutouts)
					if len(ccl) < 3 || ccl[0:3] != "#!/" {
						chk.Error(
							"invalid unix script:  " +
								"first line must start with '#!/' " +
								"after optional whitespace",
						)

						return ""
					}

					if len(ccl) < len(cleanLine) {
						trimFirst = len(cleanLine) - len(ccl)
					}

					cleanScript = ccl + "\n"
				}
			} else {
				if trimFirst > 0 && len(cleanLine) >= trimFirst {
					cleanScript += cleanLine[trimFirst:] + "\n"
				} else {
					cleanScript += cleanLine + "\n"
				}
			}
		}
	}

	return chk.createFile(
		path,
		fName,
		[]byte(cleanScript),
		settingPermExe,
	)
}

// CreateTmpSubDir creates one or more subdirectories under the test’s
// root temporary directory. Each argument is appended to the path using
// os.PathSeparator. The full absolute path to the final subdirectory is
// returned. Existing directories in the chain are reused, so multiple
// calls with a shared parent will not overwrite each other, e.g.:
//
//	pathA := chk.CreateTmpSubDir("parent", "pathA")
//	pathB := chk.CreateTmpSubDir("parent", "pathB")
//
// Both calls reuse the "parent" directory while creating separate
// child subdirectories. Unless KeepTmpFiles is called, all created
// directories are automatically removed if the test completes
// without errors.
func (chk *Chk) CreateTmpSubDir(subDirs ...string) string {
	var pathElements []string

	if !path.IsAbs(subDirs[0]) {
		pathElements = append(pathElements, chk.CreateTmpDir())
	}

	pathElements = append(pathElements, subDirs...)
	fullPath := filepath.Join(pathElements...)

	err := os.MkdirAll(fullPath, settingPermDir)
	if err != nil {
		chk.t.Helper()
		chk.Error("createTmpSubDir caused: ", err)
	}

	return fullPath
}

// CreateTmpDir creates the root temporary directory for the current test,
// named after the test function and placed under the configured root
// (defaulting to SZTEST_TMP_DIR or /tmp). If the directory already exists,
// it is left unchanged and the absolute path is returned. Unless
// KeepTmpFiles is called, the directory and its contents are automatically
// removed when the test finishes without errors.
func (chk *Chk) CreateTmpDir() string {
	var err error

	path := filepath.Join(settingTmpDir, chk.Name())

	if !chk.tmpDirCreated { //nolint:nestif // Ok.
		chk.t.Helper()

		err = removeTestDir(path)
		if err == nil {
			err = os.Mkdir(path, settingPermDir)
			if err == nil {
				chk.PushPreReleaseFunc(func() error {
					if chk.faultCount == 0 && !chk.keepTmpFiles {
						return removeTestDir(path)
					}

					return nil
				})
			}
		}
	}

	if err != nil {
		chk.Error("createTmpDir caused: ", err)
	}

	chk.tmpDirCreated = true

	return path
}
