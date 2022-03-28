package stack

import (
	lifoStack "gopherland/datastructures/stack/lifo"
	"math"
)

// MaxStack is a stack that keeps track of the maximum value in the stack.
type MaxStack struct {
	lifoStack.LifoStack
	Maximum float64
}

// NewMaxStack creates a new max stack
func NewMaxStack(capacity int) *MaxStack {
	max := math.Inf(-1)

	return &MaxStack{
		LifoStack: lifoStack.NewLifoStack(capacity),
		Maximum:   max,
	}
}

// Push adds a new item to the stop of the stack
func (s *MaxStack) Push(val interface{}) error {
	err := s.LifoStack.Push(val)
	if err != nil {
		return err
	}
	if float64(val.(int)) > s.Maximum {
		s.Maximum = float64(val.(int))
	}
	return nil
}

// Pop removes an item from the stop of the stack
func (s *MaxStack) Pop() (interface{}, error) {
	data, err := s.LifoStack.Pop()

	if err != nil {
		return nil, err
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
func (s *MaxStack) Peek() (interface{}, error) {
	return s.LifoStack.Peek()
}

// GetMax gets the max item in the stack
func (s *MaxStack) GetMax() int {
	return int(s.Maximum)
}
