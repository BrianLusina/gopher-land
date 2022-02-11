package linkedlist

import "errors"

// iLinkedList is an interface for a linked list
type ILinkedList interface {
	Append(interface{})
	Prepend(interface{})
	DeleteHead()
	DeleteTail() (node interface{}, err error)
	DeleteAtPosition(position int)
	DeleteNode(node interface{})
	DeleteNodeByData(data interface{})
	SwapNodes(dataOne, dataTwo interface{})
	SwapNodesAtKthAndKPlusOne(k int)
	Reverse()
	String() string
}

// LinkedList is a linked list
type LinkedList struct {
	Head *Node
}

var (
	ErrEmptyList = errors.New("List is empty")
)

func NewNode(data interface{}) Node {
	return Node{Data: data}
}

// NewLinkedList creates a new LinkedList
func NewLinkedList() LinkedList {
	return LinkedList{}
}
