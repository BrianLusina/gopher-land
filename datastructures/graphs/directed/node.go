package directed

import "gopherland/datastructures/graphs"

type Node struct {
	graphs.Node

	// inDegree is the number of edges pointing to the node
	inDegree int

	// outDegree is the number of edges pointing away from the node
	outDegree int
}

// NewNode creates a new node with the given data
func NewNode(data any) *Node {
	node := graphs.NewNode(data)
	return &Node{
		Node:      *node,
		inDegree:  0,
		outDegree: 0,
	}
}
