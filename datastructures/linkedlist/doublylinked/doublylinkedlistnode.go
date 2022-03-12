package doublylinkedlist

// DoublyLinkedListNode of a Doubly Linked List
type DoublyLinkedListNode struct {
	Data       interface{}
	Next, Prev *DoublyLinkedListNode
}
