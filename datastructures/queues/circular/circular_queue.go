package circular

import (
	"gopherland/datastructures/queues"
	"gopherland/pkg/utils"
)

// circularQueue implementation with an underlying slice
type circularQueue[T any] struct {
	queue      []T
	maxsize    int
	head, tail int
}

// NewQueue creates a new Circular Queue
func NewQueue[T any](maxsize int) *circularQueue[T] {
	data := make([]T, 0, maxsize)

	return &circularQueue[T]{
		queue:   data,
		maxsize: maxsize,
		head:    -1,
		tail:    -1,
	}
}

// Enqueue adds a new data element to the queue
func (q *circularQueue[T]) Enqueue(data T) error {
	if q.IsFull() {
		return queues.ErrorFullQueue
	} else if q.head == -1 {
		q.head = 0
		q.tail = 0
		q.queue = append(q.queue, data)
	} else {
		q.tail = (q.tail + 1) % q.maxsize
		q.queue = append(q.queue, data)
	}
	return nil
}

// Dequeue operation removes an item from the front of the queue
func (cq *circularQueue[T]) Dequeue() (T, error) {
	if cq.IsEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	} else if cq.head == cq.tail {
		data := cq.queue[cq.head]
		cq.head = -1
		cq.tail = -1
		cq.queue = cq.queue[1:]
		return data, nil
	} else {
		data := cq.queue[cq.head]
		cq.head = (cq.head + 1) % cq.maxsize
		cq.queue = cq.queue[1:]
		return data, nil
	}
}

// Peek returns the first item from the front of the queue without removing it
func (cq *circularQueue[T]) Peek() T {
	return cq.queue[0]
}

// IsEmpty checks if the queue is empty
func (cq *circularQueue[T]) IsEmpty() bool {
	return cq.head == -1
}

// IsFull checks if the queue is full
func (cq *circularQueue[T]) IsFull() bool {
	return (cq.tail+1)%cq.maxsize == cq.head
}

// Size returns the length of the queue
func (cq *circularQueue[T]) Size() int {
	return len(cq.queue)
}
