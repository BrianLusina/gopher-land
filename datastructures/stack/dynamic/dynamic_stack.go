package dynamicstack

import (
	"fmt"
	"gopherland/datastructures/stack"
	"gopherland/pkg/utils"
)

type DynamicStack[T any] struct {
	items []T
}

// New creates a new dynamic stack with unbounded capacity
func New[T any]() stack.Stack[T] {
	var items []T

	return &DynamicStack[T]{
		items: items,
	}
}

// Push adds an item to the stack
func (s *DynamicStack[T]) Push(val T) error {
	s.items = append(s.items, val)
	return nil
}

// Pop removes the top item from the stack
func (s *DynamicStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return utils.Zero[T](), stack.ErrorEmptyStack
	} else {
		data := s.items[s.Size()-1]
		items := s.items[:s.Size()-1]

		s.items = items
		return data, nil
	}
}

// Peek checks on the top item in the stack without removing it
func (s *DynamicStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return utils.Zero[T](), stack.ErrorEmptyStack
	}
	return s.items[s.Size()-1], nil
}

// String prints the stack
func (s *DynamicStack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// IsEmpty checks if the stack is empty
func (s *DynamicStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size gets the current Size of the stack
func (s *DynamicStack[T]) Size() int {
	return len(s.items)
}

// GetItems returns the items in the stack
func (s *DynamicStack[T]) Items() []T {
	return s.items
}
