package linkedlist

import "errors"

var (
	ErrEmptyList = errors.New("list is empty")
)

// ILinkedList is an interface for a linked list
type ILinkedList interface {
	// Append adds a new element to the end of the list
	Append(interface{})

	// Prepend adds a new element to the beginning of the list
	Prepend(interface{})

	// Pop removes the last element of the list and returns it
	Pop() (interface{}, error)

	// Array returns the list as an array
	Array() []interface{}

	DeleteTail() (node interface{}, err error)
	DeleteAtPosition(position int)
	DeleteNode(node interface{})
	DeleteNodeByData(data interface{})
	SwapNodes(dataOne, dataTwo interface{})
	SwapNodesAtKthAndKPlusOne(k int)
	// Reverse reverses the list
	Reverse()
	String() string
}

// LinkedList is a linked list
type LinkedList struct {
	Head *Node
}

// NewLinkedList creates a new LinkedList
func NewLinkedList() LinkedList {
	return LinkedList{}
}
