package stack

// Stack describes the methods that are available to a stack
type Stack[T any] interface {
	// Push adds an item to the top of the stack
	Push(T) error

	// Pop removes an item from the stop of the stack
	Pop() (T, error)

	// Peek gets the top item from the stack without removing it
	Peek() (T, error)

	IsEmpty() bool

	Size() int

	Items() []T
}
