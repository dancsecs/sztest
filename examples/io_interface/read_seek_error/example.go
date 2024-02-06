// Package example shows various test options.
package example

import (
	"io"
)

func seekFile(r io.ReadSeeker, pos int64) (int64, error) {
	// This example will attempt to read 10 bytes from r read until an error or
	// eof is returned.

	return r.Seek(pos, io.SeekStart)
}
