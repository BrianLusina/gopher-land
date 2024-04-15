package maxstack

import (
	stack "gopherland/datastructures/stack"
	dynamicstack "gopherland/datastructures/stack/dynamic"
	"gopherland/pkg/utils"

	"golang.org/x/exp/constraints"
)

// MaxStack is a stack that keeps track of the maximum value in the stack.
type MaxStack[T constraints.Ordered] struct {
	stack.Stack[T]
	currentMax T
}

// New creates a new max stack
func New[T constraints.Ordered]() *MaxStack[T] {
	ds := dynamicstack.New[T]()

	return &MaxStack[T]{
		Stack: ds,
	}
}

// Push adds a new item to the stop of the stack
func (ms *MaxStack[T]) Push(val T) error {
	err := ms.Stack.Push(val)
	if err != nil {
		return err
	}
	if val > ms.currentMax {
		ms.currentMax = val
	}
	return nil
}

// Pop removes an item from the stop of the stack
func (s *MaxStack[T]) Pop() (T, error) {
	data, err := s.Stack.Pop()

	if err != nil {
		return utils.Zero[T](), err
	}

	if s.Size() == 0 {
		s.currentMax = utils.GetZeroValue[T]()
	} else if s.currentMax == data {
		tempMax, err := s.Peek()

		if err != nil {
			return utils.GetZeroValue[T](), err
		}

		for _, item := range s.Items() {
			if item > tempMax {
				tempMax = item
			}
		}

		s.currentMax = tempMax
	}
	return data, nil
}

// Peek retrieves the top item from the stack without removing it
func (s *MaxStack[T]) Peek() (T, error) {
	return s.Stack.Peek()
}

// GetMax gets the max item in the stack
func (s *MaxStack[T]) Max() T {
	return s.currentMax
}

func (s *MaxStack[T]) Size() int {
	return s.Stack.Size()
}
