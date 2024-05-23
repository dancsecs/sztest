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
	"testing"
)

func tstChkEnv(t *testing.T) {
	t.Run("SetEnv without existing", chkEnvSetNonExistent)
	t.Run("SetEnv with existing", chkEnvSetExisting)
	t.Run("DelEnv without existing", chkEnvDelNonExistent)
	t.Run("DelEnv with existing", chkEnvDelExisting)
}

func chkEnvSetNonExistent(t *testing.T) {
	chk := CaptureNothing(t)

	tstName := "TST_SZTEST_ENV_VARIABLE"

	_, found := os.LookupEnv((tstName))
	chk.False(found) // Just make sure it does not exist.

	chk.SetEnv(tstName, "TEST_VALUE")

	foundValue, found := os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, "TEST_VALUE")

	chk.Release() // Must remove env variable.

	_, found = os.LookupEnv(tstName)
	chk.False(found) // Make sure it was removed.
}

func chkEnvSetExisting(t *testing.T) {
	chk := CaptureNothing(t)

	tstName := "TST_SZTEST_ENV_VARIABLE"
	origValue := "ORIGINAL VALUE"
	newValue := "NEW VALUE"

	_, found := os.LookupEnv((tstName))
	chk.False(found) // Just make sure it does not exist.

	defer func() {
		_ = os.Unsetenv(tstName)
	}()

	chk.NoErr(os.Setenv(tstName, origValue))

	foundValue, found := os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, origValue)

	chk.SetEnv(tstName, newValue)

	foundValue, found = os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, newValue)

	chk.Release() // Must reset env variable.

	foundValue, found = os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, origValue)
}

func chkEnvDelNonExistent(t *testing.T) {
	chk := CaptureNothing(t)

	tstName := "TST_SZTEST_ENV_VARIABLE"

	_, found := os.LookupEnv((tstName))
	chk.False(found) // Just make sure it does not exist.

	chk.DelEnv(tstName)

	_, found = os.LookupEnv((tstName))
	chk.False(found) // Just make sure it continues to not exist.

	chk.Release() // Will do nothing.

	_, found = os.LookupEnv(tstName)
	chk.False(found) // Make sure it still is not there for completeness.
}

func chkEnvDelExisting(t *testing.T) {
	chk := CaptureNothing(t)

	tstName := "TST_SZTEST_ENV_VARIABLE"
	origValue := "ORIGINAL VALUE"

	_, found := os.LookupEnv((tstName))
	chk.False(found) // Just make sure it does not exist.

	defer func() {
		_ = os.Unsetenv(tstName)
	}()

	chk.NoErr(os.Setenv(tstName, origValue))

	foundValue, found := os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, origValue)

	chk.DelEnv(tstName)

	_, found = os.LookupEnv((tstName))
	chk.False(found) // Make sure it removed.

	chk.Release() // Will do nothing.

	foundValue, found = os.LookupEnv((tstName))
	chk.True(found)
	chk.Str(foundValue, origValue) // Make sure it was restored.
}
