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
