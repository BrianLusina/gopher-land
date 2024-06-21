package dlqueue

import (
	// dll "gopherland/datastructures/linkedlist/doublylinked"
	dll "gopherland/datastructures/list/doublylinked"
	"gopherland/datastructures/queues"
	"gopherland/pkg/types"
	"gopherland/pkg/utils"
)

// FifoDLLQueue is a FIFO Queue that uses a doubly linked list as the underlying abstract data type to store elements
type FifoDLLQueue[T types.Comparable] struct {
	ll *dll.DoublyLinkedList[T]
}

func NewFifoDLLQueue[T types.Comparable]() *FifoDLLQueue[T] {
	list := dll.New[T]()
	return &FifoDLLQueue[T]{
		ll: list,
	}
}

func (q *FifoDLLQueue[T]) Enqueue(x T) {
	q.ll.Append(x)
}

func (q *FifoDLLQueue[T]) Dequeue() (T, error) {
	if q.isEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	}
	value := q.ll.Head.Data
	q.ll.DeleteAtBeg()
	return value, nil
}

func (q *FifoDLLQueue[T]) Peek() (T, error) {
	if q.isEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	}
	return q.ll.Head.Data, nil
}

func (q *FifoDLLQueue[T]) isEmpty() bool {
	return q.Len() == 0
}

func (q *FifoDLLQueue[T]) Items() []T {
	if q.isEmpty() {
		return []T{}
	}
	items := make([]T, 0)
	current := q.ll.Head
	for current != nil {
		items = append(items, current.Data)
		current = current.Next
	}
	return items
}

func (q *FifoDLLQueue[T]) Len() int {
	return q.ll.Size()
}
