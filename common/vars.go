package common

import "errors"

var (
	// ErrNotFound is a generic error for when an entity does not exist
	ErrNotFound = errors.New("not found")
)
