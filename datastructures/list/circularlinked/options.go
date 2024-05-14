package circularlinked

import "gopherland/datastructures/list"

type Option[T comparable] func(*CircularLinkedList[T])

// WithHeadOption adds a head node to a linked list
func WithHeadOption[T comparable](head *list.Node[T]) Option[T] {
	return func(cll *CircularLinkedList[T]) {
		cll.Head = head
	}
}
