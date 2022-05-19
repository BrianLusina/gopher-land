package directed

import "gopherland/datastructures/graphs"

type Node struct {
	graphs.Node

	// InDegree is the number of edges pointing to the node
	InDegree int

	// OutDegree is the number of edges pointing away from the node
	OutDegree int
}

// NewNode creates a new node with the given data
func NewNode(data any) *Node {
	node := graphs.NewNode(data)
	return &Node{
		Node:      *node,
		InDegree:  0,
		OutDegree: 0,
	}
}
