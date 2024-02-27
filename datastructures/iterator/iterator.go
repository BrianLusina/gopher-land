package iterator

import "gopherland/pkg/utils"

type Iterator[T comparable] struct {
	index int
	data  []T
}

// New creates a new iterator
func New[T comparable](data []T) *Iterator[T] {
	return &Iterator[T]{
		index: 0,
		data:  data,
	}
}

// HasNext checks if there are still items to iterate over
func (i *Iterator[T]) HasNext() bool {
	if i.index < len(i.data) {
		return true
	}
	return false
}

// Next returns the next item in the iterator
func (i *Iterator[T]) Next() T {
	if i.HasNext() {
		item := i.data[i.index]
		i.index++
		return item
	}
	return utils.GetZeroValue[T]()
}
