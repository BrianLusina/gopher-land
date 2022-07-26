package linkedlist

import "fmt"

// Node basic node of a linked list
type Node[T comparable] struct {
	key  any
	Data T
	Next *Node[T]
}

// NewNode creates a new node for a linked list
func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{Data: data}
}

// SetKey sets a key to the node
func (n *Node[T]) SetKey(key any) {
	n.key = key
}

// GetKey gets a key for the node
func (n *Node[T]) GetKey() any {
	return n.key
}

// String representation of a node
func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Data)
}
