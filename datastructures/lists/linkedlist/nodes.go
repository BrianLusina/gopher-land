package linkedlist

// Node basic node of a linked list
type Node struct {
	data int32
	next *Node
}

// SinglyLinkedListNode of a Singly Linked List
type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

// DoublyLinkedListNode reprentation of a Doubly Linked List Node with 2 pointers
type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}
