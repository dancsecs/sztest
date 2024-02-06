// Package example shows various test options.
package example

import (
	"io"
)

func seekFile(w io.WriteSeeker, pos int64) (int64, error) {
	// This example will attempt to read 10 bytes from r read until an error or
	// eof is returned.

	return w.Seek(pos, io.SeekStart)
}
