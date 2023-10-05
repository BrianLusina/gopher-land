package peeking

import "gopherland/designpatterns/iterator"

type PeekingIterator[T comparable] struct {
	iterator *iterator.Iterator[T]
	stack    []T
}

func NewPeekingIterator[T comparable](data []T) *PeekingIterator[T] {
	iter := iterator.New(data)

	return &PeekingIterator[T]{
		iterator: iter,
		stack:    make([]T, 0),
	}
}

func (pi *PeekingIterator[T]) hasNext() bool {
	if len(pi.stack) > 0 || pi.iterator.HasNext() {
		return true
	}
	return false
}

func (pi *PeekingIterator[T]) next() T {
	if len(pi.stack) > 0 {
		top := pi.stack[0]
		pi.stack = pi.stack[1:]
		return top
	}
	item := pi.iterator.Next()
	return item
}

func (pi *PeekingIterator[T]) peek() T {
	if len(pi.stack) == 0 {
		pi.stack = append(pi.stack, pi.iterator.Next())
	}
	return pi.stack[0]
}
