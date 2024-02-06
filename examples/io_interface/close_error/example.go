// Package example shows various test options.
package example

import (
	"io"
)

func closeFile(r io.Closer) error {
	// This example will attempt to read 10 bytes from r read until an error or
	// eof is returned.

	return r.Close()
}
