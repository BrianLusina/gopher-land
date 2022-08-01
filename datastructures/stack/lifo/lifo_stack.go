package stack

import (
	"fmt"
	"gopherland/datastructures/stack"
	"gopherland/pkg/utils"
)

type LifoStack[T comparable] struct {
	items    []T
	capacity int
}

// NewLifoStack creates a new stack with a given capacity
func NewLifoStack[T comparable](capacity int) LifoStack[T] {
	var items []T

	return LifoStack[T]{
		items:    items,
		capacity: capacity,
	}
}

// Push adds an item to the stack
func (s *LifoStack[T]) Push(val T) error {
	if !s.isFull() {
		s.items = append(s.items, val)
		return nil
	} else {
		return stack.ErrorFullStack
	}
}

// Pop removes the top item from the stack
func (s *LifoStack[T]) Pop() (T, error) {
	if s.isEmpty() {
		return utils.GetZeroValue[T](), stack.ErrorEmptyStack
	} else {
		data := s.items[s.size()-1]
		items := s.items[:s.size()-1]

		s.items = items
		return data, nil
	}
}

// Peek checks on the top item in the stack without removing it
func (s *LifoStack[T]) Peek() (T, error) {
	if s.isEmpty() {
		return utils.GetZeroValue[T](), stack.ErrorEmptyStack
	}
	return s.items[s.size()-1], nil
}

// String prints the stack
func (s *LifoStack[T]) String() string {
	return fmt.Sprintf("%v", s.items)
}

// isFull checks if the stack has reached capacity
func (s *LifoStack[T]) isFull() bool {
	return s.size() == s.capacity
}

// isEmpty checks if the stack is empty
func (s *LifoStack[T]) isEmpty() bool {
	return s.size() == 0
}

// size gets the current size of the stack
func (s *LifoStack[T]) size() int {
	return len(s.items)
}

// cap gets the current capacity of the stack
func (s *LifoStack[T]) cap() int {
	return cap(s.items)
}

// GetItems returns the items in the stack
func (s *LifoStack[T]) GetItems() []T {
	return s.items
}
