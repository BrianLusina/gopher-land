package peeking

import "gopherland/pkg/utils"

type Iterator[T comparable] struct {
	index int
	items []T
}

func (i *Iterator[T]) hasNext() bool {
	if i.index < len(i.items) {
		return true
	}
	return false
}

func (i *Iterator[T]) next() T {
	if i.hasNext() {
		item := i.items[i.index]
		i.index++
		return item
	}
	return utils.GetZeroValue[T]()
}
