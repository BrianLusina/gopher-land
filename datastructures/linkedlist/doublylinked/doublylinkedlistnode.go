package doublylinkedlist

// Node DoublyLinkedListNode of a Doubly Linked List
type Node struct {
	Data       interface{}
	Next, Prev *Node
}

// NewDoublyLinkedListNode creates a new DoublyLinkedList Node
func NewDoublyLinkedListNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}
}
