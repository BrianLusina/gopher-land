package linkedlist

// Node basic node of a linked list
type Node struct {
	Data int32
	Next *Node
}

// DoublyLinkedListNode reprentation of a Doubly Linked List Node with 2 pointers
type DoublyLinkedListNode struct {
	Data int32
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}
