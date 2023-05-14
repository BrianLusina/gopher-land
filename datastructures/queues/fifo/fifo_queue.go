package fifo

import (
	"gopherland/datastructures/queues"
	"gopherland/pkg/utils"
)

type FifoQueue[T any] struct {
	items   []T
	maxsize int
}

func NewFifoQueue[T any](maxsize int) *FifoQueue[T] {
	data := make([]T, 0, maxsize)
	return &FifoQueue[T]{items: data, maxsize: maxsize}
}

// Enqueue adds an item to the back of the queue
func (q *FifoQueue[T]) Enqueue(data T) error {
	if q.IsFull() {
		return queues.ErrorFullQueue
	}
	q.items = append(q.items, data)
	return nil
}

// Dequeue removes an item from the front of the queue
func (q *FifoQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, nil
}

// Peek returns the front element of the queue without removing it
func (q *FifoQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	}
	return q.items[0], nil
}

// isEmpty returns true if the queue is empty
func (q *FifoQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// isEmpty returns true if the queue is empty
func (q *FifoQueue[T]) IsFull() bool {
	return q.Size() == q.maxsize
}

// Items returns a slice of the items in the queue
func (q *FifoQueue[T]) Items() []T {
	return q.items
}

// Len returns the length of the queue or size of items in the queue
func (q *FifoQueue[T]) Size() int {
	return len(q.items)
}
