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
