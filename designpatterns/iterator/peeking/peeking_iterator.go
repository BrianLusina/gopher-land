package peeking

import "gopherland/designpatterns/iterator"

type PeekingIterator[T int] struct {
	iterator *iterator.Iterator[T]
	stack    []T
}

func NewPeekingIterator(iter *iterator.Iterator[int]) *PeekingIterator[int] {
	return &PeekingIterator[int]{
		iterator: iter,
		stack:    make([]int, 0),
	}
}

func (pi *PeekingIterator[int]) hasNext() bool {
	if len(pi.stack) > 0 || pi.iterator.HasNext() {
		return true
	}
	return false
}

func (pi *PeekingIterator[int]) next() int {
	if len(pi.stack) > 0 {
		top := pi.stack[0]
		pi.stack = pi.stack[1:]
		return top
	}
	item := pi.iterator.Next()
	return item
}

func (pi *PeekingIterator[int]) peek() int {
	if len(pi.stack) == 0 {
		pi.stack = append(pi.stack, pi.iterator.Next())
	}
	return pi.stack[0]
}
