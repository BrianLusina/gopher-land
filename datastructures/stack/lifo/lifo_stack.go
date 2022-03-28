package stack

import (
	"fmt"
	"gopherland/datastructures/stack"
)

type LifoStack struct {
	items    []interface{}
	capacity int
}

// NewLifoStack creates a new stack with a given capacity
func NewLifoStack(capacity int) LifoStack {
	var items []interface{}

	return LifoStack{
		items:    items,
		capacity: capacity,
	}
}

// Push adds an item to the stack
func (s *LifoStack) Push(val interface{}) error {
	if !s.isFull() {
		s.items = append(s.items, val)
		return nil
	} else {
		return stack.ErrorFullStack
	}
}

// Pop removes the top item from the stack
func (s *LifoStack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, stack.ErrorEmptyStack
	} else {
		data := s.items[s.size()-1]
		items := s.items[:s.size()-1]

		s.items = items
		return data, nil
	}
}

// Peek checks on the top item in the stack without removing it
func (s *LifoStack) Peek() (interface{}, error) {
	if s.isEmpty() {
		return nil, stack.ErrorEmptyStack
	}
	return s.items[s.size()-1], nil
}

// String prints the stack
func (s *LifoStack) String() string {
	return fmt.Sprintf("%v", s.items)
}

// isFull checks if the stack has reached capacity
func (s *LifoStack) isFull() bool {
	return s.size() == s.capacity
}

// isEmpty checks if the stack is empty
func (s *LifoStack) isEmpty() bool {
	return s.size() == 0
}

// size gets the current size of the stack
func (s *LifoStack) size() int {
	return len(s.items)
}

// cap gets the current capacity of the stack
func (s *LifoStack) cap() int {
	return cap(s.items)
}

// GetItems returns the items in the stack
func (s *LifoStack) GetItems() []interface{} {
	return s.items
}
