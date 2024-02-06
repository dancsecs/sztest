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
