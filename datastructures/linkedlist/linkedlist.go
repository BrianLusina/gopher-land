package linkedlist

import "errors"

var (
	ErrEmptyList    = errors.New("list is empty")
	ErrInvalidIndex = errors.New("invalid index")
)

// LinkedList is an interface for a linked list
type LinkedList[T comparable] interface {
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

	// GetNthNode returns the nth node in the list
	GetNthNode(position int) (node Node[T], err error)

	// GetMiddleNode returns the middle node of the list
	GetMiddleNode() (node Node[T], err error)

	DeleteNode(node interface{})

	DeleteNodeByData(data interface{})

	SwapNodes(dataOne, dataTwo interface{})

	SwapNodesAtKthAndKPlusOne(k int)

	// Reverse reverses the list
	Reverse()

	// Rotate rotates the list by k positions counter-clockwise
	Rotate(k int)

	// ReverseGroups reverses the list in groups of k elements
	ReverseGroups(k int)

	String() string

	// Length returns the length of the list
	Length() int
}
