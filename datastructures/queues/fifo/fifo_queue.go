package fifo

import (
	"gopherland/datastructures/queues"
)

type FifoQueue struct {
	items []interface{}
}

func NewFifoQueue() *FifoQueue {
	// new built in function can be used or struct literal
	// return &FifoQueue{}
	return new(FifoQueue)
}

// Enqueue adds an item to the back of the queue
func (q *FifoQueue) Enqueue(x interface{}) {
	q.items = append(q.items, x)
}

// Dequeue removes an item from the front of the queue
func (q *FifoQueue) Dequeue() (interface{}, error) {
	if q.isEmpty() {
		return nil, queues.ErrorEmptyQueue
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, nil
}

// Peek returns the front element of the queue without removing it
func (q *FifoQueue) Peek() (interface{}, error) {
	if q.isEmpty() {
		return nil, queues.ErrorEmptyQueue
	}
	return q.items[0], nil
}

// isEmpty returns true if the queue is empty
func (q *FifoQueue) isEmpty() bool {
	return q.Len() == 0
}

// Items returns a slice of the items in the queue
func (q *FifoQueue) Items() []interface{} {
	return q.items
}

// Len returns the length of the queue or size of items in the queue
func (q *FifoQueue) Len() int {
	return len(q.items)
}
