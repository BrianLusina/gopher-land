// dlqueue contains a fifo queue implementation using a double linked list
package dlqueue

import (
	dll "gopherland/datastructures/linkedlist/doublylinked"
	"gopherland/datastructures/queues"
	"gopherland/datastructures/queues/fifo"
)

type LinkedListQueue struct {
	fifo.FifoQueue
	ll *dll.DoublyLinkedList
}

func NewLinkedListQueue() *LinkedListQueue {
	list := dll.NewDoublyLinkedList()
	return &LinkedListQueue{
		ll: list,
	}
}

func (q *LinkedListQueue) Enqueue(x interface{}) {
	q.ll.Append(x)
}

func (q *LinkedListQueue) Dequeue() (interface{}, error) {
	if q.isEmpty() {
		return nil, queues.ErrorEmptyQueue
	}
	value := q.ll.Head.Data
	q.ll.DeleteAtBeg()
	return value, nil
}

func (q *LinkedListQueue) Peek() (interface{}, error) {
	if q.isEmpty() {
		return nil, queues.ErrorEmptyQueue
	}
	return q.ll.Head.Data, nil
}

func (q *LinkedListQueue) isEmpty() bool {
	return q.Len() == 0
}

func (q *LinkedListQueue) Items() []interface{} {
	if q.isEmpty() {
		return []interface{}{}
	}
	items := make([]interface{}, 0)
	current := q.ll.Head
	for current != nil {
		items = append(items, current.Data)
		current = current.Next
	}
	return items
}

func (q *LinkedListQueue) Len() int {
	return q.ll.Size()
}
