package linkedlist

// Node basic node of a linked list
type Node struct {
	Data interface{}
	Next *Node
}

// DoublyNode reprentation of a Doubly Linked List Node with 2 pointers
type DoublyNode struct {
	Data interface{}
	Prev *DoublyNode
	Next *DoublyNode
}
