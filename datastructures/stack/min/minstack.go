package stack

import (
	lifoStack "gopherland/datastructures/stack/lifo"
	"math"
)

type MinStack struct {
	lifoStack.LifoStack
	Minimum float64
}

// NewMinStack creates a new MinStack
func NewMinStack(capacity int) *MinStack {
	min := math.Inf(1)

	return &MinStack{
		LifoStack: lifoStack.NewLifoStack(capacity),
		Minimum:   min,
	}
}

func (s *MinStack) Push(val any) error {
	err := s.LifoStack.Push(val)

	if err != nil {
		return err
	}

	if float64(val.(int)) < s.Minimum {
		s.Minimum = float64(val.(int))
	}

	return nil
}

func (s *MinStack) Pop() (any, error) {
	data, err := s.LifoStack.Pop()

	if err != nil {
		return nil, err
	}

	if s.size() == 0 {
		s.Minimum = math.Inf(1)
	} else if s.Minimum == float64(data.(int)) {
		m := s.LifoStack.GetItems()[0]

		for i, e := range s.LifoStack.GetItems() {
			if i == 0 || e.(int) < m.(int) {
				m = e
			}
		}
		s.Minimum = float64(m.(int))
	}
	return data, nil
}

func (s *MinStack) Peek() (any, error) {
	return s.LifoStack.Peek()
}

func (s *MinStack) GetMin() int {
	return int(s.Minimum)
}

func (s *MinStack) size() int {
	return len(s.LifoStack.GetItems())
}
