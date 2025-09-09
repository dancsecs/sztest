<!--- gotomd::Auto:: See github.com/dancsecs/gotomd **DO NOT MODIFY** -->

# IO Interface

- [Example: IO Read Error](#example-io-read-error)
- [Example: IO Write Error](#example-io-write-error)
- [Example: IO Read Seek Error](#example-io-read-seek-error)
- [Example: IO Write Seek Error](#example-io-write-seek-error)
- [Example: IO Close Error](#example-io-close-error)

[Contents](../../README.md#contents)

## Example: IO Read Error

<!--- gotomd::Bgn::file::./read_error/example.go -->
```bash
cat ./read_error/example.go
```

```go
// Package example shows various test options.
package example

import (
    "errors"
    "io"
)

func readFile(r io.Reader) (string, error) {
    // This example will attempt to read 10 bytes from r read until an error or
    // eof is returned.

    const bufSize = 10

    bytes := make([]byte, bufSize)
    c, err := r.Read(bytes)

    if err == nil && c < bufSize {
        return string(bytes), errors.New("not enough bytes")
    }

    if errors.Is(err, io.EOF) {
        return "", errors.New("unexpected EOF")
    }

    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
```
<!--- gotomd::End::file::./read_error/example.go -->

<!--- gotomd::Bgn::file::./read_error/example_test.go -->
```bash
cat ./read_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    // Read without any data set will cause immediate EOF.
    _, err := readFile(chk)
    chk.Err(err, "unexpected EOF")

    // Read all the data

    chk.SetIOReaderData("0123456789")

    str, err := readFile(chk)
    chk.NoErr(err)
    chk.Str(str, "0123456789")

    str, err = readFile(chk)
    chk.Str(str, "")
    chk.Err(err, "unexpected EOF")

    // Not enough data

    chk.SetIOReaderData("01234")
    str, err = readFile(chk)
    chk.Err(err, "not enough bytes")
    chk.Str(str, "01234\x00\x00\x00\x00\x00")

    // Fail after a certain number of bytes is read.
    chk.SetIOReaderData("01234567890")
    chk.SetIOReaderError(2, errors.New("only two bytes read"))
    str, err = readFile(chk)
    chk.Err(err, "not enough bytes")
    chk.Str(str, "01\x00\x00\x00\x00\x00\x00\x00\x00")

    str, err = readFile(chk)
    chk.Err(err, "only two bytes read")
    chk.Str(str, "")

    // Setup a direct error to be returned on next read.  No other data needs
    // to be setup

    chk.SetReadError(2962, errors.New("example error on returning 2962"))

    n, err := chk.Read(nil)
    chk.Int(n, 2962)
    chk.Err(err, "example error on returning 2962")

    chk.SetReadError(2963, nil) // no error just a forced count.

    n, err = chk.Read(nil)
    chk.NoErr(err)
    chk.Int(n, 2963)
}
```
<!--- gotomd::End::file::./read_error/example_test.go -->

<!--- gotomd::Bgn::tst::./read_error/package -->
```bash
go test -v -cover ./read_error
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/io&#xA0;&#x332;&#xA0;&#x332;interface/read&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./read_error/package -->

[Contents](../../README.md#contents)

## Example: IO Write Error

<!--- gotomd::Bgn::file::./write_error/example.go -->
```bash
cat ./write_error/example.go
```

```go
// Package example shows various test options.
package example

import (
    "io"
)

func writeFile(w io.Writer) (int, error) {
    // Attempt to write 10 characters to the io.Writer.

    n, err := w.Write([]byte("0123456789"))

    return n, err
}
```
<!--- gotomd::End::file::./write_error/example.go -->

<!--- gotomd::Bgn::file::./write_error/example_test.go -->
```bash
cat ./write_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    // Read without anything .
    c, err := writeFile(chk)
    chk.NoErr(err)
    chk.Int(c, 10)
    chk.Str(string(chk.GetIOWriterData()), "0123456789")

    chk.SetIOWriterError(8, errors.New("Run out of space after 8 chars"))
    c, err = writeFile(chk)
    chk.Err(err, "Run out of space after 8 chars")
    chk.Int(c, 8)
    chk.Str(string(chk.GetIOWriterData()), "01234567")

    // Just set a write error and count to be returned on the next write.

    chk.SetWriteError(37, errors.New("the write error"))
    c, err = writeFile(chk)

    chk.Err(err, "the write error")
    chk.Int(c, 37)
}
```
<!--- gotomd::End::file::./write_error/example_test.go -->

<!--- gotomd::Bgn::tst::./write_error/package -->
```bash
go test -v -cover ./write_error
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/io&#xA0;&#x332;&#xA0;&#x332;interface/write&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./write_error/package -->

[Contents](../../README.md#contents)

## Example: IO Read Seek Error

<!--- gotomd::Bgn::file::./read_seek_error/example.go -->
```bash
cat ./read_seek_error/example.go
```

```go
// Package example shows various test options.
package example

import (
    "io"
)

func seekFile(r io.ReadSeeker, pos int64) (int64, error) {
    // This example will attempt position the io.ReadSeeker to the position
    // provided.

    return r.Seek(pos, io.SeekStart)
}
```
<!--- gotomd::End::file::./read_seek_error/example.go -->

<!--- gotomd::Bgn::file::./read_seek_error/example_test.go -->
```bash
cat ./read_seek_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.SetSeekError(34, errors.New("past end of file"))

    p, err := seekFile(chk, 962)
    chk.Err(err, "past end of file")
    chk.Int64(p, 34)
}
```
<!--- gotomd::End::file::./read_seek_error/example_test.go -->

<!--- gotomd::Bgn::tst::./read_seek_error/package -->
```bash
go test -v -cover ./read_seek_error
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/io&#xA0;&#x332;&#xA0;&#x332;interface/read&#xA0;&#x332;&#xA0;&#x332;seek&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./read_seek_error/package -->

[Contents](../../README.md#contents)

## Example: IO Write Seek Error

<!--- gotomd::Bgn::file::./write_seek_error/example.go -->
```bash
cat ./write_seek_error/example.go
```

```go
// Package example shows various test options.
package example

import (
    "io"
)

func seekFile(w io.WriteSeeker, pos int64) (int64, error) {
    // This example will attempt position the io.WriteSeeker to the position
    // provided.

    return w.Seek(pos, io.SeekStart)
}
```
<!--- gotomd::End::file::./write_seek_error/example.go -->

<!--- gotomd::Bgn::file::./write_seek_error/example_test.go -->
```bash
cat ./write_seek_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.SetSeekError(34, errors.New("past end of file"))

    p, err := seekFile(chk, 962)
    chk.Err(err, "past end of file")
    chk.Int64(p, 34)
}
```
<!--- gotomd::End::file::./write_seek_error/example_test.go -->

<!--- gotomd::Bgn::tst::./write_seek_error/package -->
```bash
go test -v -cover ./write_seek_error
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/io&#xA0;&#x332;&#xA0;&#x332;interface/write&#xA0;&#x332;&#xA0;&#x332;seek&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./write_seek_error/package -->

[Contents](../../README.md#contents)

## Example: IO Close Error

<!--- gotomd::Bgn::file::./close_error/example.go -->
```bash
cat ./close_error/example.go
```

```go
// Package example shows various test options.
package example

import (
    "io"
)

func closeFile(r io.Closer) error {
    // This example will attempt to close the closer returning any error
    // generated by the close action.

    return r.Close()
}
```
<!--- gotomd::End::file::./close_error/example.go -->

<!--- gotomd::Bgn::file::./close_error/example_test.go -->
```bash
cat ./close_error/example_test.go
```

```go
package example

import (
    "errors"
    "testing"

    "github.com/dancsecs/sztest"
)

func Test_IoInterface_ReadError(t *testing.T) {
    chk := sztest.CaptureNothing(t)
    defer chk.Release()

    chk.SetCloseError(errors.New("the close error"))

    chk.Err(closeFile(chk), "the close error")
}
```
<!--- gotomd::End::file::./close_error/example_test.go -->

<!--- gotomd::Bgn::tst::./close_error/package -->
```bash
go test -v -cover ./close_error
```

$\small{\texttt{===&#xA0;&#x34F;&#xA0;&#x34F;RUN&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError}}$
<br>
$\small{\texttt{‒‒‒&#xA0;&#x34F;&#xA0;&#x34F;PASS:&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;Test&#xA0;&#x332;&#xA0;&#x332;IoInterface&#xA0;&#x332;&#xA0;&#x332;ReadError&#xA0;&#x34F;&#xA0;&#x34F;(0.0s)}}$
<br>
$\small{\texttt{PASS}}$
<br>
$\small{\texttt{coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
$\small{\texttt{ok&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;github.com/dancsecs/sztest/examples/io&#xA0;&#x332;&#xA0;&#x332;interface/close&#xA0;&#x332;&#xA0;&#x332;error&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;&#xA0;&#x34F;coverage:&#xA0;&#x34F;&#xA0;&#x34F;100.0&#xFE6A;&#xA0;&#x34F;&#xA0;&#x34F;of&#xA0;&#x34F;&#xA0;&#x34F;statements}}$
<br>
<!--- gotomd::End::tst::./close_error/package -->

[Contents](../../README.md#contents)
