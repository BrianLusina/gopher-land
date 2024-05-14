package circularlinked

import "gopherland/datastructures/list"

// CircularLinkedList is a circular linked list structure
type CircularLinkedList[T comparable] struct {
	// Head is the head of the circular linked list
	Head *list.Node[T]
}

// New creates a new circular linked list
func New[T comparable](opts ...Option[T]) list.LinkedList[T] {
	cll := new(CircularLinkedList[T])

	for _, opt := range opts {
		opt(cll)
	}

	return cll
}

// HeadNode retrieves the Head node of a linked list
func (c *CircularLinkedList[T]) HeadNode() *list.Node[T] {
	return c.Head
}

func (c *CircularLinkedList[T]) All(yield func(T) bool) {
	current := c.Head

	for current != nil {
		if !yield(current.Data) {
			return
		}
		if current.Next == c.Head {
			break
		}
		current = current.Next
	}
}

// Append adds a new data item at the end of the linked list.
func (c *CircularLinkedList[T]) Append(data T) {
	node := list.NewNode(data)
	if c.Head == nil {
		c.Head = node
		node.Next = c.Head
	} else {
		current := c.Head
		for current.Next != c.Head {
			current = current.Next
		}
		current.Next = node
		node.Next = c.Head
	}
}

// Array implements list.LinkedList.
func (c *CircularLinkedList[T]) Array() []interface{} {
	panic("unimplemented")
}

// DeleteAtPosition implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteAtPosition(position int) {
	panic("unimplemented")
}

// DeleteMiddle implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteMiddle() (any, error) {
	panic("unimplemented")
}

// DeleteNode implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteNode(node interface{}) {
	panic("unimplemented")
}

// DeleteNodeByData implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteNodeByData(data interface{}) {
	panic("unimplemented")
}

// DeleteTail implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteTail() (node interface{}, err error) {
	panic("unimplemented")
}

// GetMiddleNode implements list.LinkedList.
func (c *CircularLinkedList[T]) GetMiddleNode() (node list.Node[T], err error) {
	panic("unimplemented")
}

// GetNthNode implements list.LinkedList.
func (c *CircularLinkedList[T]) GetNthNode(position int) (node list.Node[T], err error) {
	panic("unimplemented")
}

// Length implements list.LinkedList.
func (c *CircularLinkedList[T]) Length() int {
	panic("unimplemented")
}

// MaximumPairSum implements list.LinkedList.
func (c *CircularLinkedList[T]) MaximumPairSum() int {
	panic("unimplemented")
}

// OddEvenList implements list.LinkedList.
func (c *CircularLinkedList[T]) OddEvenList() *list.Node[T] {
	panic("unimplemented")
}

// Pop implements list.LinkedList.
func (c *CircularLinkedList[T]) Pop() (interface{}, error) {
	panic("unimplemented")
}

// Prepend implements list.LinkedList.
func (c *CircularLinkedList[T]) Prepend(interface{}) {
	panic("unimplemented")
}

// Reverse implements list.LinkedList.
func (c *CircularLinkedList[T]) Reverse() {
	panic("unimplemented")
}

// ReverseGroups implements list.LinkedList.
func (c *CircularLinkedList[T]) ReverseGroups(k int) {
	panic("unimplemented")
}

// Rotate implements list.LinkedList.
func (c *CircularLinkedList[T]) Rotate(k int) {
	panic("unimplemented")
}

// String implements list.LinkedList.
func (c *CircularLinkedList[T]) String() string {
	panic("unimplemented")
}

// SwapNodes implements list.LinkedList.
func (c *CircularLinkedList[T]) SwapNodes(dataOne interface{}, dataTwo interface{}) {
	panic("unimplemented")
}

// SwapNodesAtKthAndKPlusOne implements list.LinkedList.
func (c *CircularLinkedList[T]) SwapNodesAtKthAndKPlusOne(k int) {
	panic("unimplemented")
}
