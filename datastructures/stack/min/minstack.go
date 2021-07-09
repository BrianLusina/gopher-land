package stack

import (
	"gopherland/datastructures/stack"
	"math"
)

type minStackInterface interface {
	stack.StackInterface
	GetMin() int
}

type MinStack struct {
	stack.Stack
	Minimum float64
}

// NewMinStack
func NewMinStack(capacity int) minStackInterface {
	items := []interface{}{}
	min := math.Inf(1)

	return &MinStack{
		Stack: stack.Stack{
			Items:    items,
			Capacity: capacity,
		},
		Minimum: min,
	}
}

func (s *MinStack) Push(val interface{}) {
	s.Stack.Push(val)
	if float64(val.(int)) < s.Minimum {
		s.Minimum = float64(val.(int))
	}
}

func (s *MinStack) Pop() interface{} {
	data := s.Stack.Pop()

	if s.size() == 0 {
		s.Minimum = math.Inf(1)
	} else if s.Minimum == float64(data.(int)) {
		m := s.Stack.Items[0]

		for i, e := range s.Stack.Items {
			if i == 0 || e.(int) < m.(int) {
				m = e
			}
		}
		s.Minimum = float64(m.(int))
	}
	return data
}

func (s *MinStack) Peek() interface{} {
	return s.Stack.Peek()
}

func (s *MinStack) GetMin() int {
	return int(s.Minimum)
}

func (s *MinStack) size() int {
	return len(s.Stack.Items)
}
