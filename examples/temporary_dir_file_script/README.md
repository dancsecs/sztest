<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# Temporary directories, files, scripts

- [Example: Temporary Directory](#example-temporary-directory)
- [Example: Temporary File](#example-temporary-file)
- [Example: Temporary Unix Script](#example-temporary-unix-script)
- [Example: Temporary Directory Tree](#example-temporary-directory-tree)

[Contents](../../README.md#contents)

## Example: Temporary Directory

<!--- gotomd::Bgn::file::./temporary_directory/example_test.go -->
```bash
cat ./temporary_directory/example_test.go
```

```go
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
```
<!--- gotomd::End::file::./temporary_directory/example_test.go -->

<!--- gotomd::Bgn::tst::./temporary_directory/package -->
```bash
go test -v -cover ./temporary_directory
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;TemporaryDirectory}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;TemporaryDirectory&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/temporary&#x332;dir&#x332;file&#x332;script/temporary&#x332;directory&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./temporary_directory/package -->

[Contents](../../README.md#contents)

## Example: Temporary File

<!--- gotomd::Bgn::file::./temporary_file/example_test.go -->
```bash
cat ./temporary_file/example_test.go
```

```go
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
```
<!--- gotomd::End::file::./temporary_file/example_test.go -->

<!--- gotomd::Bgn::tst::./temporary_file/package -->
```bash
go test -v -cover ./temporary_file
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;TemporaryFile}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;TemporaryFile&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/temporary&#x332;dir&#x332;file&#x332;script/temporary&#x332;file&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./temporary_file/package -->

[Contents](../../README.md#contents)

## Example: Temporary Unix Script

<!--- gotomd::Bgn::file::./temporary_unix_script/example_test.go -->
```bash
cat ./temporary_unix_script/example_test.go
```

```go
package example

import (
    "os/exec"
    "testing"

    "github.com/dancsecs/sztest"
)

// Passing test.
func Test_PASS_TemporaryUnixScript(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    unixExeFile := chk.CreateTmpUnixScript([]string{`
    #!/bin/bash

    echo "Running script $0"
    `,
    })

    cmd := exec.Command(unixExeFile)
    out, err := cmd.Output()
    chk.NoErr(err)
    chk.NotNil(out)
    chk.Str(string(out), "Running script "+unixExeFile+"\n")
}
```
<!--- gotomd::End::file::./temporary_unix_script/example_test.go -->

<!--- gotomd::Bgn::tst::./temporary_unix_script/package -->
```bash
go test -v -cover ./temporary_unix_script
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;TemporaryUnixScript}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;TemporaryUnixScript&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/temporary&#x332;dir&#x332;file&#x332;script/temporary&#x332;unix&#x332;script&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./temporary_unix_script/package -->

[Contents](../../README.md#contents)

## Example: Temporary Directory Tree

<!--- gotomd::Bgn::file::./temporary_directory_tree/example_test.go -->
```bash
cat ./temporary_directory_tree/example_test.go
```

```go
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
```
<!--- gotomd::End::file::./temporary_directory_tree/example_test.go -->

<!--- gotomd::Bgn::tst::./temporary_directory_tree/package -->
```bash
go test -v -cover ./temporary_directory_tree
```

$\small{\texttt{===&#xa0;RUN&#xa0;&#xa0;&#xa0;Test&#x332;PASS&#x332;TemporaryDirectoryTree}}$
<br>
$\small{\texttt{---&#xa0;PASS:&#xa0;Test&#x332;PASS&#x332;TemporaryDirectoryTree&#xa0;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xa0;[no&#xa0;statements]}}$
<br>
$\small{\texttt{ok&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;&#xa0;github.com/dancsecs/sztest/examples/temporary&#x332;dir&#x332;file&#x332;script/temporary&#x332;directory&#x332;tree&#xa0;&#xa0;&#xa0;&#xa0;coverage:&#xa0;[no&#xa0;statements]}}$
<br>
<!--- gotomd::End::tst::./temporary_directory_tree/package -->

[Contents](../../README.md#contents)
