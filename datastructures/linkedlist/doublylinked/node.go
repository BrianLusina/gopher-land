package doublylinkedlist

import (
	"fmt"
	"gopherland/datastructures/linkedlist"
)

// Node DoublyLinkedListNode of a Doubly Linked List
type Node[T comparable] struct {
	*linkedlist.Node[T]
	Next, Prev *Node[T]
}

// NewNode creates a new DoublyLinkedList Node
func NewNode[T comparable](data T) *Node[T] {
	node := linkedlist.NewNode(data)
	return &Node[T]{
		Node: node,
		Next: nil,
		Prev: nil,
	}
}

func (n *Node[T]) String() string {
	if n.Prev != nil && n.Next != nil {
		return fmt.Sprintf("%v -> [%v] -> %v", n.Prev.Data, n.Data, n.Next.Data)
	} else if n.Prev == nil && n.Next != nil {
		return fmt.Sprintf("[%v] -> %v", n.Data, n.Next.Data)
	} else if n.Prev != nil && n.Next == nil {
		return fmt.Sprintf("%v -> [%v]", n.Prev.Data, n.Data)
	}
	return fmt.Sprintf("[%v]", n.Data)
}
