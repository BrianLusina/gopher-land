package fixedsizedstack

import (
	"fmt"
	"gopherland/datastructures/stack"
	"gopherland/pkg/utils"
)

type FixedSizedStack[T any] struct {
	items    []T
	capacity int
}

// New creates a new fixed sized stack with bounded capacity
func New[T any](capacity int) stack.Stack[T] {
	items := make([]T, 0, capacity)

	return &FixedSizedStack[T]{
		items:    items,
		capacity: capacity,
	}
}

// Push adds an item to the stack
func (s *FixedSizedStack[T]) Push(val T) error {
	if s.IsFull() {
		return stack.ErrorFullStack
	}
	s.items = append(s.items, val)
	return nil
}

// Pop removes the top item from the stack
func (s *FixedSizedStack[T]) Pop() (T, error) {
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
func (s *FixedSizedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return utils.Zero[T](), stack.ErrorEmptyStack
	}
	return s.items[s.Size()-1], nil
}

// String prints the stack
func (s *FixedSizedStack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// IsEmpty checks if the stack is empty
func (s *FixedSizedStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// IsFull checks if the stack is full
func (s *FixedSizedStack[T]) IsFull() bool {
	return s.Size() == s.capacity
}

// Size gets the current Size of the stack
func (s *FixedSizedStack[T]) Size() int {
	return len(s.items)
}

// GetItems returns the items in the stack
func (s *FixedSizedStack[T]) Items() []T {
	return s.items
}
