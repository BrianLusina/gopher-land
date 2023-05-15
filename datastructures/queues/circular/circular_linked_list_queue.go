/*
*
Contains implementation of a circular linked list queue with an underlying implementation of a linked list
*
*/
package circular

import (
	"gopherland/datastructures/list"
	"gopherland/datastructures/queues"
	"gopherland/pkg/utils"
)

// circularLinkedListQueue implementation with an underlying slice
type circularLinkedListQueue[T comparable] struct {
	maxsize     int
	currentSize int
	head, tail  *list.Node[T]
}

// NewLinkedListQueue creates a new Circular Queue with an underlying implementation of a linked list
func NewLinkedListQueue[T comparable](maxsize int) *circularLinkedListQueue[T] {
	return &circularLinkedListQueue[T]{
		currentSize: 0,
		maxsize:     maxsize,
		head:        nil,
		tail:        nil,
	}
}

// Enqueue adds a new data element to the queue
func (cq *circularLinkedListQueue[T]) Enqueue(data T) error {
	if cq.IsFull() {
		return queues.ErrorFullQueue
	}
	node := list.NewNode(data)

	if cq.head == nil {
		cq.head = node
	} else {
		cq.tail.Next = node
	}
	cq.tail = node
	cq.tail.Next = cq.head
	cq.currentSize++
	return nil
}

// Dequeue operation removes an item from the front of the queue
func (cq *circularLinkedListQueue[T]) Dequeue() (T, error) {
	if cq.IsEmpty() {
		return utils.Zero[T](), queues.ErrorEmptyQueue
	}

	cq.currentSize--
	if cq.head == cq.tail {
		data := cq.head.Data
		cq.head = nil
		cq.tail = nil
		return data, nil
	} else {
		temp := cq.head
		data := temp.Data
		cq.head = cq.head.Next
		cq.tail.Next = cq.head
		return data, nil
	}
}

// Peek returns the first item from the front of the queue without removing it
func (cq *circularLinkedListQueue[T]) Peek() T {
	return cq.head.Data
}

// IsEmpty checks if the queue is empty
func (cq *circularLinkedListQueue[T]) IsEmpty() bool {
	return cq.currentSize == 0
}

// IsFull checks if the queue is full
func (cq *circularLinkedListQueue[T]) IsFull() bool {
	return cq.currentSize == cq.maxsize
}

// Size returns the length of the queue
func (cq *circularLinkedListQueue[T]) Size() int {
	return cq.currentSize
}
