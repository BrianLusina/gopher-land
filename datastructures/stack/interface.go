package stack

// stackInterface describes the methods that are available to a stack
type StackInterface interface {
	// Push adds an item to the top of the stack
	Push(interface{})

	// Pop removes an item from the stop of the stack
	Pop() interface{}

	// Peek gets the top item from the stack without removing it
	Peek() interface{}
}
