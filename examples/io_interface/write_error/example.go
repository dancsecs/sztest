// Package example shows various test options.
package example

import (
	"io"
)

func writeFile(w io.Writer) (int, error) {
	n, err := w.Write([]byte("0123456789"))

	return n, err
}
