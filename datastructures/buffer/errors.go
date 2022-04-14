package buffer

import "errors"

var (
	ErrEmptyBuffer = errors.New("buffer is empty")
	ErrFullBuffer  = errors.New("buffer is full")
)
