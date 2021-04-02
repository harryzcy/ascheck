package remotecheck

import (
	"errors"
)

var (
	// ErrNotFound is returned when an app is not found.
	ErrNotFound = errors.New("app not found")
)
