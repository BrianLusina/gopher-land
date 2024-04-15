package stack

import "errors"

var (
	ErrorEmptyStack = errors.New("stack is empty")
	ErrorFullStack  = errors.New("stack is full")
)
