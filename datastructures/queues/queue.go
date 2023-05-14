package queues

import (
	"errors"
)

var (
	ErrorEmptyQueue = errors.New("queue is empty")
	ErrorFullQueue  = errors.New("queue is full")
)

type Queue[T any] interface {
	Enqueue(data T) error
	Dequeue() (T, error)
	IsEmpty() bool
	IsFull() bool
	Peek() (T, error)
	Size() int
}
