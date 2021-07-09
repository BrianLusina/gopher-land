package stack

import (
	"gopherland/datastructures/stack"
	"math"
)

// maxStackInterface is the interface that wraps the methods required
type maxStackInterface interface {
	stack.StackInterface
	GetMax() int
}

// MaxStack is a stack that keeps track of the maximum value in the stack.
type MaxStack struct {
	stack.Stack
	Maximum float64
}

// NewMaxStack creates a new max stack
func NewMaxStack(capacity int) maxStackInterface {
	items := []interface{}{}
	max := math.Inf(-1)

	return &MaxStack{
		Stack: stack.Stack{
			Items:    items,
			Capacity: capacity,
		},
		Maximum: max,
	}
}

// Push adds a new item to the stop of the stack
func (s *MaxStack) Push(val interface{}) {
	s.Stack.Push(val)
	if float64(val.(int)) > s.Maximum {
		s.Maximum = float64(val.(int))
	}
}

// Pop removes an item from the stop of the stack
func (s *MaxStack) Pop() interface{} {
	data := s.Stack.Pop()

	if len(s.Stack.Items) == 0 {
		s.Maximum = math.Inf(1)
	} else if s.Maximum == float64(data.(int)) {
		m := s.Stack.Items[0]

		for i, e := range s.Stack.Items {
			if i == 0 || e.(int) > m.(int) {
				m = e
			}
		}
		s.Maximum = float64(m.(int))
	}
	return data
}

// Peek retrieves the top item from the stack without removing it
func (s *MaxStack) Peek() interface{} {
	return s.Stack.Peek()
}

// GetMax gets the max item in the stack
func (s *MaxStack) GetMax() int {
	return int(s.Maximum)
}
