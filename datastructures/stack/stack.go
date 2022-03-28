package stack

import "errors"

var (
	ErrorEmptyStack = errors.New("stack is empty")
	ErrorFullStack  = errors.New("stack is full")
)

// Stack describes the methods that are available to a stack
type Stack interface {
	// Push adds an item to the top of the stack
	Push(interface{}) error

	// Pop removes an item from the stop of the stack
	Pop() (interface{}, error)

	// Peek gets the top item from the stack without removing it
	Peek() (interface{}, error)
}

// MaxStack describes methods available to a MaxStack
type MaxStack interface {
	Stack
	GetMax() int
}

type MinStack interface {
	Stack
	GetMin() int
}
