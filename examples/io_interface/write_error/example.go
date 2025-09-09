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
