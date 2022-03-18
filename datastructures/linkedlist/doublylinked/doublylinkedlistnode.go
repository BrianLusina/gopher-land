package doublylinkedlist

// DoublyLinkedListNode of a Doubly Linked List
type DoublyLinkedListNode struct {
	Data       interface{}
	Next, Prev *DoublyLinkedListNode
}

// NewDoublyLinkedListNode creates a new DoublyLinkedList Node
func NewDoublyLinkedListNode(data interface{}) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}
