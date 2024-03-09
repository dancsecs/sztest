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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func tstChkDir(t *testing.T) {
	t.Run("RemoveTestDir", chkDirTestRemoveTestDir)
	t.Run("RemoveTestFile", chkDirTestRemoveTestFile)
	t.Run("SetDirPerm", chkDirTestSetDirPerm)
	t.Run("SetFilePerm", chkDirTestSetFilePerm)
	t.Run("SetPermExe", chkDirTestSetPermExe)
	t.Run("SetTmpDirEmpty", chkDirTestSetTmpDirEmpty)
	t.Run("SetTmpDirNotExists", chkDirTestSetTmpDirNotExists)
	t.Run("SetTmpDirNotDirectory", chkDirTestSetTmpDirNotDirectory)
	t.Run("SetTmpDirExtendExisting", chkDirTestSetTmpDirExtendExisting)
	t.Run("CreateDirNotExist", chkDirTestCreateDirNotExist)
	t.Run("CreateTmpDirEmpty", chkDirTestCreateTmpDirEmpty)
	t.Run("CreateTmpFileEmpty", chkDirTestCreateTmpFileEmpty)
	t.Run("CreateTmpFileEmptyKeepTmp", chkDirTestCreateTmpFileEmptyKeepTmp)
	t.Run(
		"CreateTmpFileInvalidDIrectory", chkDirTestCreateTmpFileInvalidDIrectory,
	)
	t.Run("CreateTmpUnixScriptEmpty", chkDirTestCreateTmpUnixScriptEmpty)
	t.Run("CreateTmpUnixScriptInvalid", chkDirTestCreateTmpUnixScriptInvalid)
	t.Run(
		"CreateTmpUnixScriptNoLeading", chkDirTestCreateTmpUnixScriptNoLeading,
	)
	t.Run(
		"CreateTmpUnixScriptWithLeading",
		chkDirTestCreateTmpUnixScriptWithLeading,
	)
	t.Run("CreateTmpSubDir", chkDirTestCreateTmpSubDir)
	t.Run("CreateRealTmpDir", chkDirTestCreateRealTmpDir)
}

func chkDirTestRemoveTestDir(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.KeepTmpFiles()

	tstDir := filepath.Join(settingTmpDir, "testDirectory")

	chk.NoErr(removeTestDir(tstDir))

	chk.NoErr(os.Mkdir(tstDir, settingPermDir))

	fileName := filepath.Join(tstDir, "fileNotDir")

	chk.NoErr(os.WriteFile(fileName, []byte{}, 0o0600))

	chk.Err(
		removeTestDir(fileName),
		ErrInvalidDirectory.Error()+": \""+fileName+"\"",
	)

	chk.NoErr(removeTestDir(tstDir))
}

func chkDirTestRemoveTestFile(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	chk.KeepTmpFiles()

	tstDir := filepath.Join(settingTmpDir, "testDirectory")

	chk.NoErr(removeTestDir(tstDir))

	chk.NoErr(os.Mkdir(tstDir, settingPermDir))

	chk.Err(
		removeTestFile(tstDir),
		ErrInvalidFile.Error()+": \""+tstDir+"\"",
	)

	fileName := filepath.Join(tstDir, "fileNotDir")

	// no error if file is not there
	chk.NoErr(removeTestFile(fileName))

	chk.NoErr(os.WriteFile(fileName, []byte{}, 0o0600))

	chk.NoErr(
		removeTestFile(fileName),
	)

	chk.NoErr(removeTestFile(fileName))

	chk.NoErr(removeTestDir(tstDir))
}

func chkDirTestSetDirPerm(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	newPerm := os.FileMode(0o0777)

	oldSettingPermDir := settingPermDir

	chk.Int(
		int(chk.SetPermDir(newPerm)),
		int(oldSettingPermDir),
	)

	chk.Int(int(settingPermDir), int(newPerm))

	chk.Int(
		int(chk.SetPermDir(settingPermDir)),
		int(newPerm),
	)
}

func chkDirTestSetFilePerm(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	newPerm := os.FileMode(0o0777)

	oldSettingPermFile := settingPermFile

	chk.Int(
		int(chk.SetPermFile(newPerm)),
		int(oldSettingPermFile),
	)

	chk.Int(int(settingPermFile), int(newPerm))

	chk.Int(
		int(chk.SetPermFile(settingPermFile)),
		int(newPerm),
	)
}

func chkDirTestSetPermExe(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	newPerm := os.FileMode(0o0777)

	olsSettingPermExe := settingPermExe
	chk.Int(
		int(chk.SetPermExe(newPerm)),
		int(olsSettingPermExe),
	)

	chk.Int(int(settingPermExe), int(newPerm))

	chk.Int(
		int(chk.SetPermExe(settingPermExe)),
		int(newPerm),
	)
}

func chkDirTestSetTmpDirEmpty(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	oldSettingTmpDir := settingTmpDir

	const dirTestValue = "TEMP_DIRECTORY"
	// Resets tmpDir back to default
	chk.Str(
		chk.SetTmpDir(""),
		settingTmpDir,
	)
	chk.Str(settingTmpDir, oldSettingTmpDir) // still the same

	// Set a test value

	settingTmpDir = dirTestValue // actual directory does not exist

	// Resets tmpDir back to default
	chk.Str(
		chk.SetTmpDir(""),
		dirTestValue,
	)
	chk.Str(settingTmpDir, oldSettingTmpDir) // still the same
}

func chkDirTestSetTmpDirNotExists(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	oldSettingTmpDir := settingTmpDir
	// Invalid Directory
	chk.Str(
		chk.SetTmpDir("/DOES/NOT/EXIST"),
		settingTmpDir,
	)
	chk.Str(settingTmpDir, oldSettingTmpDir) // no change

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("SetTmpDir"),
		chkOutError("invalid directory: /DOES/NOT/EXIST"),
		chkOutRelease(),
	)
}

func chkDirTestSetTmpDirNotDirectory(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	oldSettingTmpDir := settingTmpDir

	fName := filepath.Join(settingTmpDir, chk.Name())

	rootDir, _ := filepath.Split(fName)

	chk.NoErr(os.MkdirAll(rootDir, 0o0700))

	if chk.NoErr(os.WriteFile(fName, []byte{}, 0o0600)) {
		chk.PushPreReleaseFunc(
			func() error {
				return os.Remove(fName) //nolint:wrapcheck // Ok.
			},
		)
	}

	// Invalid Directory (a file)
	chk.Str(
		chk.SetTmpDir(fName),
		settingTmpDir,
	)
	chk.Str(settingTmpDir, oldSettingTmpDir) // no change

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutPush("Pre", ""),
		chkOutHelper("SetTmpDir"),
		chkOutError("invalid directory: "+fName),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestSetTmpDirExtendExisting(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	oldSettingTmpDir := settingTmpDir
	defer func() {
		settingTmpDir = oldSettingTmpDir
	}()

	const defaultDirPerm = os.FileMode(0o0700)

	fPath := filepath.Join(settingTmpDir, chk.Name())

	if chk.NoErr(os.Mkdir(fPath, defaultDirPerm)) {
		chk.PushPreReleaseFunc(
			func() error {
				return os.RemoveAll(fPath) //nolint:wrapcheck // Ok.
			},
		)
	}
	// Valid

	chk.Str(
		chk.SetTmpDir(chk.Name()),
		oldSettingTmpDir,
	)
	chk.Str(settingTmpDir, fPath) // extended
}

func chkDirTestCreateDirNotExist(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	oldSettingTmpDir := settingTmpDir
	defer func() {
		settingTmpDir = oldSettingTmpDir
	}()

	settingTmpDir = "/DOES_NOT_EXIST/"

	chk.CreateTmpDir()

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutError(
			"createTmpDir cause: mkdir /DOES_NOT_EXIST/Internal Testing Object: "+
				"no such file or directory",
		),
		chkOutRelease(),
	)
}

func chkDirTestCreateTmpDirEmpty(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dirName := chk.CreateTmpDir()

	chk.Str(dirName, filepath.Join(settingTmpDir, chk.Name()))

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpFileEmpty(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	_ = chk.CreateTmpDir()

	fileName := chk.CreateTmpFile([]byte{})

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpFile"),
		chkOutHelper("CreateTmpFileIn"),
		chkOutHelper("createFile"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpFileEmptyKeepTmp(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.KeepTmpFiles()

	tmpDir := chk.CreateTmpDir()

	fileName := chk.CreateTmpFile([]byte{})

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	chk.Release()

	chk.NoErr(removeTestFile(fileName))

	chk.NoErr(removeTestDir(tmpDir))

	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpFile"),
		chkOutHelper("CreateTmpFileIn"),
		chkOutHelper("createFile"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpFileInvalidDIrectory(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	chk.tmpDirCreated = true
	fileName := chk.CreateTmpFile([]byte{})

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpFile"),
		chkOutHelper("CreateTmpFileIn"),
		chkOutHelper("createFile"),
		chkOutError(
			"createFile cause: "+ErrInvalidDirectory.Error()+
				": \"/tmp/Internal Testing Object\"",
		),
		chkOutRelease(),
	)
}

func chkDirTestCreateTmpUnixScriptEmpty(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	_ = chk.CreateTmpDir()

	fileName := chk.CreateTmpUnixScript([]string{})

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpUnixScript"),
		chkOutHelper("CreateTmpUnixScriptIn"),
		chkOutHelper("createFile"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpUnixScriptInvalid(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	fileName := chk.CreateTmpUnixScript([]string{"First Line invalid"})

	chk.Str(fileName, "")

	chk.faultCount = 0 // Reset so tmp dir is properly deleted.
	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpUnixScript"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpUnixScriptIn"),
		chkOutError(
			"invalid unix script:  first line must start "+
				"with '#!/' after optional whitespace",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpUnixScriptNoLeading(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	fileName := chk.CreateTmpUnixScript(
		[]string{
			"#!/bin/bash\n\necho \"Hello, world!\"",
		},
	)

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	fData, err := os.ReadFile(fileName) //nolint:gosec // Ok.

	chk.NoErr(err)
	chk.StrSlice(strings.Split(string(fData), "\n"), []string{
		"#!/bin/bash",
		"",
		"echo \"Hello, world!\"",
		"",
	},
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpUnixScript"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpUnixScriptIn"),
		chkOutHelper("createFile"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpUnixScriptWithLeading(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	fileName := chk.CreateTmpUnixScript(
		[]string{
			`
		#!/bin/bash

		echo "Hello, world!"
		`,
		},
	)

	chk.Str(
		fileName,
		filepath.Join(settingTmpDir, chk.Name(), "tmpFile0.tmp"),
	)

	fData, err := os.ReadFile(fileName) //nolint:gosec // Ok.

	chk.NoErr(err)
	chk.StrSlice(strings.Split(string(fData), "\n"), []string{
		"#!/bin/bash",
		"",
		"echo \"Hello, world!\"",
		"",
		"",
	},
	)

	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpUnixScript"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpUnixScriptIn"),
		chkOutHelper("createFile"),
		chkOutPush("Pre", ""),
		chkOutRelease(),
		chkOutPush("Pre", "func2"),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateTmpSubDir(t *testing.T) {
	iT := new(iTst)
	chk := CaptureNothing(iT)
	iT.chk = chk

	dir := chk.CreateTmpDir()

	_ = chk.CreateTmpSubDir("abc")

	_ = chk.CreateTmpSubDir("abc", filepath.Join("def", "efg"))

	_, err := os.Stat(filepath.Join(dir, "abc"))
	chk.NoErr(err)

	_, err = os.Stat(filepath.Join(dir, "abc", "def"))
	chk.NoErr(err)

	_, err = os.Stat(filepath.Join(dir, "abc", "def", "efg"))
	chk.NoErr(err)

	// Make reference to root dir.
	_ = chk.CreateTmpSubDir("/../../../../../../../../../../here")

	chk.faultCount = 0 // Reset so tmp directory is properly deleted.
	chk.Release()
	iT.check(t,
		chkOutCapture("Nothing"),
		chkOutHelper("CreateTmpDir"),
		chkOutPush("Pre", ""),
		chkOutHelper("CreateTmpSubDir"),
		chkOutError(
			"createTmpSubDir cause: mkdir /here: read-only file system",
		),
		chkOutRelease(),
		chkOutPush("Pre", "func1"),
	)
}

func chkDirTestCreateRealTmpDir(t *testing.T) {
	chk := CaptureNothing(t)
	defer chk.Release()

	dir := chk.CreateTmpDir()

	rawName := t.Name()
	adjName := strings.ReplaceAll(rawName, string(os.PathSeparator), "-")

	rawPath := filepath.Join(settingTmpDir, rawName)
	adjPath := filepath.Join(settingTmpDir, adjName)

	chk.False(rawPath == adjPath, "should be a path as it is executed via t.Run")

	chk.Str(
		dir,
		adjPath,
	)
}
