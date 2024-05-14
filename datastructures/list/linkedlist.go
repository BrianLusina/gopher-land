package list

import (
	"errors"
)

var (
	ErrEmptyList    = errors.New("list is empty")
	ErrInvalidIndex = errors.New("invalid index")
)

// LinkedList is an interface for a linked list
type LinkedList[T comparable] interface {
	// HeadNode retrieves the Head node of a linked list
	HeadNode() *Node[T]

	// All can be used to construct an iterator
	All(yield func(T) bool)

	// Append adds a new element to the end of the list
	Append(T)

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

	DeleteMiddle() (any, error)

	SwapNodes(dataOne, dataTwo interface{})

	SwapNodesAtKthAndKPlusOne(k int)

	// Reverse reverses the list
	Reverse()

	// OddEvenList returns the linked list grouping all nodes with odd indices together followed by nodes with even indices
	OddEvenList() *Node[T]

	// MaximumPairSum returns the maximum twin sum of a node and its twin, where a node's twin is at the index (n-1-i) where n is the
	// number of nodes in the linked list.
	// For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only
	// nodes with twins for n = 4.
	MaximumPairSum() int

	// Rotate rotates the list by k positions counter-clockwise
	Rotate(k int)

	// ReverseGroups reverses the list in groups of k elements
	ReverseGroups(k int)

	String() string

	// Length returns the length of the list
	Length() int
}
