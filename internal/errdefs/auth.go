package errdefs

import "errors"

var (
	ErrPasswordMismatch = errors.New("invalid password")
)
