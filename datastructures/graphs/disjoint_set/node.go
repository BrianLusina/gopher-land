package disjointset

import "gopherland/datastructures/graphs"

type Node struct {
	graphs.Node
	parent *Node
	rank   int
}

// NewNode returns a new node in a disjoint set
func NewNode(data any) *Node {
	node := graphs.NewNode(data)
	n := &Node{
		Node: *node,
		rank: 0,
	}
	n.parent = n
	return n
}

func (n *Node) GetParent() *Node {
	return n.parent
}

func (n *Node) GetRank() int {
	return n.rank
}
