package linkedlist

import "fmt"

// Node basic node of a linked list
type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

func NewNode[T comparable](data T) *Node[T] {
	return &Node[T]{Data: data}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Data)
}
