package stack

import (
	"gopherland/datastructures/stack"
	"gopherland/pkg/utils"

	"golang.org/x/exp/constraints"
)

type MinStack[T constraints.Ordered] struct {
	stack.Stack[T]
	currentMinimum T
}

// NewMinStack creates a new MinStack
func NewMinStack[T constraints.Ordered](capacity int) *MinStack[T] {
	return &MinStack[T]{
		Stack:          stack.NewStack[T](capacity),
		currentMinimum: utils.Zero[T](),
	}
}

// Push adds an item to the top of the stack
func (s *MinStack[T]) Push(val T) error {
	err := s.Stack.Push(val)

	if err != nil {
		return err
	}

	if val < s.currentMinimum {
		s.currentMinimum = val
	}

	return nil
}

// Pop removes an item from the top of the stack. If no item is in the stack an error is returned
func (s *MinStack[T]) Pop() (T, error) {
	data, err := s.Stack.Pop()

	if err != nil {
		return utils.Zero[T](), err
	}

	if s.size() == 0 {
		s.currentMinimum = utils.Zero[T]()
	} else if s.currentMinimum == data {
		tempMin, err := s.Peek()

		if err != nil {
			return utils.Zero[T](), err
		}

		for _, item := range s.GetItems() {
			if item < tempMin {
				tempMin = item
			}
		}

		s.currentMinimum = tempMin
	}

	return data, nil
}

// Peek check for the top of the item in the stack without removing it, Returns an error if the stack is empty
func (s *MinStack[T]) Peek() (T, error) {
	return s.Stack.Peek()
}

// GetMin returns the minimum value in the stack
func (s *MinStack[T]) GetMin() T {
	return s.currentMinimum
}

// size returns the current size of the stack
func (s *MinStack[T]) size() int {
	return s.Stack.Size()
}
