package stack

import (
	"fmt"
	"gopherland/pkg/utils"
)

type Stack[T comparable] struct {
	items    []T
	capacity int
}

// NewStack creates a new stack with a given capacity
func NewStack[T comparable](capacity int) Stack[T] {
	var items []T

	return Stack[T]{
		items:    items,
		capacity: capacity,
	}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(val T) error {
	if !s.isFull() {
		s.items = append(s.items, val)
		return nil
	} else {
		return ErrorFullStack
	}
}

// Pop removes the top item from the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return utils.Zero[T](), ErrorEmptyStack
	} else {
		data := s.items[s.Size()-1]
		items := s.items[:s.Size()-1]

		s.items = items
		return data, nil
	}
}

// Peek checks on the top item in the stack without removing it
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return utils.Zero[T](), ErrorEmptyStack
	}
	return s.items[s.Size()-1], nil
}

// String prints the stack
func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// isFull checks if the stack has reached capacity
func (s *Stack[T]) isFull() bool {
	return s.Size() == s.capacity
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size gets the current Size of the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// GetItems returns the items in the stack
func (s *Stack[T]) GetItems() []T {
	return s.items
}
