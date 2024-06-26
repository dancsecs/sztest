/*
   Golang test helper library: sztest.
   Copyright (C) 2023, 2024 Leslie Dancsecs

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

// SetPermDir changes the os.FileMode used when creating
// directories and returns the current value.
func (chk *Chk) SetPermDir(p os.FileMode) os.FileMode {
	lastPerm := settingPermDir
	settingPermDir = p

	return lastPerm
}

// SetPermFile changes the os.FileMode used when creating
// directories and returns the current value.
func (chk *Chk) SetPermFile(p os.FileMode) os.FileMode {
	lastPerm := settingPermFile
	settingPermFile = p

	return lastPerm
}

// SetPermExe changes the os.FileMode used when creating
// directories and returns the current value.
func (chk *Chk) SetPermExe(p os.FileMode) os.FileMode {
	lastPerm := settingPermExe
	settingPermExe = p

	return lastPerm
}

// SetTmpDir changes the root directory used when creating
// directories and returns the current value.
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

// CreateTmpFile removes and creates the named directory with the provided
// permissions.
func (chk *Chk) CreateTmpFile(data []byte) string {
	chk.t.Helper()

	return chk.CreateTmpFileIn("", data)
}

// CreateTmpFileIn removes and creates a tmp file in the provided path.
func (chk *Chk) CreateTmpFileIn(path string, data []byte) string {
	chk.t.Helper()

	return chk.CreateTmpFileAs(path, "", data)
}

// CreateTmpFileAs removes and creates the named file in the provided path.
func (chk *Chk) CreateTmpFileAs(path, fName string, data []byte) string {
	chk.t.Helper()

	return chk.createFile(
		path,
		fName,
		data,
		settingPermFile,
	)
}

// CreateTmpUnixScript removes and creates the named directory with the
// provided permissions.
func (chk *Chk) CreateTmpUnixScript(lines []string) string {
	chk.t.Helper()
	chk.CreateTmpDir()

	return chk.CreateTmpUnixScriptIn(
		"",
		lines,
	)
}

// CreateTmpUnixScriptIn removes and creates the generated script name with the
// provided permissions.
func (chk *Chk) CreateTmpUnixScriptIn(path string, lines []string) string {
	chk.t.Helper()

	return chk.CreateTmpUnixScriptAs(
		path,
		"",
		lines,
	)
}

// CreateTmpUnixScriptAs removes and creates the named script with the
// provided permissions.
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

// CreateTmpSubDir creates a temporary test directory using the test functions
// name in the /tmp directory.
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

// CreateTmpDir creates a temporary test directory using the test functions
// name in the /tmp directory.
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
