package iterator

import "gopherland/pkg/utils"

type Iterator[T comparable] struct {
	index int
	Items []T
}

func (i *Iterator[T]) HasNext() bool {
	if i.index < len(i.Items) {
		return true
	}
	return false
}

func (i *Iterator[T]) Next() T {
	if i.HasNext() {
		item := i.Items[i.index]
		i.index++
		return item
	}
	return utils.GetZeroValue[T]()
}
