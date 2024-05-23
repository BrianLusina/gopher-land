package list

import "fmt"

// Node basic node of a linked list
type Node[T comparable] struct {
	key  any
	Data T
	Next *Node[T]
}

// NewNode creates a new node for a linked list
func NewNode[T comparable](data T, opts ...NodeOption[T]) *Node[T] {
	node := &Node[T]{Data: data, key: data}

	for _, opt := range opts {
		opt(node)
	}

	return node
}

// SetKey sets a key to the node
func (n *Node[T]) SetKey(key any) {
	n.key = key
}

// GetKey gets a key for the node
func (n *Node[T]) Key() any {
	return n.key
}

// String representation of a node
func (n *Node[T]) String() string {
	return fmt.Sprintf("Node(data=%v, key=%v)", n.Data, n.key)
}

// Compare compares this node to another node
func (n *Node[T]) Compare(other *Node[T]) bool {
	if n.key == other.key || n.Data == other.Data {
		return true
	}
	return false
}
