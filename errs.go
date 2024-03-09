package sztest

import "errors"

// Exported errors.
var (
	ErrInvalidLastArg    = errors.New("invalid last arg error")
	ErrInvalidDirectory  = errors.New("invalid directory")
	ErrInvalidFile       = errors.New("invalid file")
	ErrReadPastEndOfData = errors.New("read past end of data")
	ErrForcedOutOfSpace  = errors.New("forced out of space")
)
