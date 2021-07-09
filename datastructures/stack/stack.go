package stack

type Stack struct {
	Items    []interface{}
	Capacity int
}

// NewStack creates a new stack with a given capacity
func NewStack(capacity int) Stack {
	stack := []interface{}{}

	return Stack{
		Items:    stack,
		Capacity: capacity,
	}
}

// Push adds an item to the stack
func (s *Stack) Push(val interface{}) {
	if !s.isFull() {
		s.Items = append(s.Items, val)
	} else {
		panic("Stack Max Capacity Reached")
	}
}

// Pop removes the top item from the stack
func (s *Stack) Pop() interface{} {
	if s.isEmpty() {
		panic("Stack is empty")
	} else {
		data := s.Items[s.size()-1]
		stack := s.Items[:s.size()-1]

		s.Items = stack
		return data
	}
}

// Peek checks on the top item in the stack without removing it
func (s *Stack) Peek() interface{} {
	if s.isEmpty() {
		panic("Stack is empty")
	}
	return s.Items[s.size()-1]
}

// isFull checks if the stack has reached capacity
func (s *Stack) isFull() bool {
	return s.size() == s.Capacity
}

// isEmpty checks if the stack is empty
func (s *Stack) isEmpty() bool {
	return s.size() == 0
}

// size gets the current size of the stack
func (s *Stack) size() int {
	return len(s.Items)
}
