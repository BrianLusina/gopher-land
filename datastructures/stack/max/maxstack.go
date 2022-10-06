package stack

import (
	lifoStack "gopherland/datastructures/stack/lifo"
	"gopherland/pkg/utils"
	"math"
)

// MaxStack is a stack that keeps track of the maximum value in the stack.
type MaxStack[T comparable] struct {
	lifoStack.LifoStack[T]
	Maximum T
}

// NewMaxStack creates a new max stack
func NewMaxStack[T comparable](capacity int) *MaxStack[T] {
	max := math.Inf(-1)

	return &MaxStack[T]{
		LifoStack: lifoStack.NewLifoStack[T](capacity),
		Maximum:   max,
	}
}

// Push adds a new item to the stop of the stack
func (s *MaxStack[T]) Push(val T) error {
	err := s.LifoStack.Push(val)
	if err != nil {
		return err
	}
	if val > s.Maximum {
		s.Maximum = val
	}
	return nil
}

// Pop removes an item from the stop of the stack
func (s *MaxStack[T]) Pop() (T, error) {
	data, err := s.LifoStack.Pop()

	if err != nil {
		return utils.Zero[T](), err
	}

	if len(s.LifoStack.GetItems()) == 0 {
		s.Maximum = math.Inf(1)
	} else if s.Maximum == float64(data.(int)) {
		m := s.LifoStack.GetItems()[0]

		for i, e := range s.LifoStack.GetItems() {
			if i == 0 || e.(int) > m.(int) {
				m = e
			}
		}
		s.Maximum = float64(m.(int))
	}
	return data, nil
}

// Peek retrieves the top item from the stack without removing it
func (s *MaxStack[T]) Peek() (T, error) {
	return s.LifoStack.Peek()
}

// GetMax gets the max item in the stack
func (s *MaxStack[T]) GetMax() T {
	return int(s.Maximum)
}
