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
func (chk *Chk) SetTmpDir(d string) string {
	lastTmpDir := settingTmpDir
	if d == "" {
		initTmpDir()
		d = settingTmpDir
	} else if !filepath.IsAbs(d) {
		initTmpDir()
		d = filepath.Join(settingTmpDir, d)
	}

	fi, err := os.Stat(d)

	if err != nil || !fi.IsDir() {
		chk.t.Helper()
		chk.Error("invalid directory: ", d)
	} else {
		settingTmpDir = d
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

func (chk *Chk) createFile(path string, data []byte, perm os.FileMode) string {
	chk.t.Helper()
	id := chk.nextTmpID
	chk.nextTmpID++
	fName := filepath.Join(path, fmt.Sprint("tmpFile", id, ".tmp"))

	pathStat, err := os.Stat(path)
	if err != nil || !pathStat.IsDir() {
		err = fmt.Errorf("%w: %q", ErrInvalidDirectory, path)
	}
	if err == nil {
		err = removeTestFile(fName)
	}
	if err == nil {
		err = os.WriteFile(fName, data, perm)
		if err == nil {
			chk.PushPreReleaseFunc(func() error {
				if chk.faultCount == 0 && !chk.keepTmpFiles {
					return removeTestFile(fName)
				}
				return nil
			})
		}
	}
	if err != nil {
		chk.Error("createFile cause: ", err)
	}
	return fName
}

// CreateTmpFile removes and creates the named directory with the provided
// permissions.
func (chk *Chk) CreateTmpFile(data []byte) string {
	chk.t.Helper()
	chk.CreateTmpDir()
	return chk.CreateTmpFileIn(filepath.Join(settingTmpDir, chk.Name()), data)
}

// CreateTmpFileIn removes and creates the named file in the provided path.
func (chk *Chk) CreateTmpFileIn(path string, data []byte) string {
	chk.t.Helper()
	chk.CreateTmpDir()
	return chk.createFile(
		path,
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
		filepath.Join(settingTmpDir, chk.Name()),
		lines,
	)
}

// CreateTmpUnixScriptIn removes and creates the named directory with the
// provided permissions.
func (chk *Chk) CreateTmpUnixScriptIn(path string, lines []string) string {
	chk.t.Helper()

	// Strips away any leading blank lines and leading spaces on the first line
	// generally the #!/shell opening line
	cleanScript := ""
	trimFirst := 0

	const spaceCutouts = " \t"
	for _, entry := range lines {
		for _, l := range strings.Split(entry, "\n") {
			cl := strings.TrimRight(l, spaceCutouts)
			if cleanScript == "" {
				if cl != "" {
					// found first non blank line
					ccl := strings.TrimLeft(cl, spaceCutouts)
					if len(ccl) < 3 || ccl[0:3] != "#!/" {
						chk.Error(
							"invalid unix script:  " +
								"first line must start with '#!/' after optional whitespace",
						)
						return ""
					}
					if len(ccl) < len(cl) {
						trimFirst = len(cl) - len(ccl)
					}
					cleanScript = ccl + "\n"
				}
			} else {
				if trimFirst > 0 && len(cl) >= trimFirst {
					cleanScript += cl[trimFirst:] + "\n"
				} else {
					cleanScript += cl + "\n"
				}
			}
		}
	}
	return chk.createFile(
		path,
		[]byte(cleanScript),
		settingPermExe,
	)
}

// CreateTmpSubDir creates a temporary test directory using the test functions
// name in the /tmp directory.
func (chk *Chk) CreateTmpSubDir(d ...string) string {
	var pathElements []string
	if !path.IsAbs(d[0]) {
		pathElements = append(pathElements, chk.CreateTmpDir())
	}
	pathElements = append(pathElements, d...)
	fullPath := filepath.Join(pathElements...)
	err := os.MkdirAll(fullPath, settingPermDir)

	if err != nil {
		chk.t.Helper()
		chk.Error("createTmpSubDir cause: ", err)
	}
	return fullPath
}

// CreateTmpDir creates a temporary test directory using the test functions
// name in the /tmp directory.
func (chk *Chk) CreateTmpDir() string {
	var err error
	path := filepath.Join(settingTmpDir, chk.Name())
	if !chk.tmpDirCreated {
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
		chk.Error("createTmpDir cause: ", err)
	}
	chk.tmpDirCreated = true
	return path
}
