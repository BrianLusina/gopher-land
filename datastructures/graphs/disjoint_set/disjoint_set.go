package disjointset

type DisjointSet struct {
	nodes []*Node
	size  int
}

func NewDisjointSet(node ...Node) *DisjointSet {
	nodes := make([]*Node, len(node))
	size := len(node)

	for i, n := range node {
		nodes[i] = &n
	}

	return &DisjointSet{
		size:  size,
		nodes: nodes,
	}
}

// Union takes two nodes and makes them part of the same set
// set with bigger rank should be parent, so that the
// disjoint set tree will be more flat.
func (d *DisjointSet) Union(nodeOne, nodeTwo *Node) {
	rootX := d.Find(nodeOne)
	rootY := d.Find(nodeTwo)

	if rootX != rootY {
		if rootX.rank > rootY.rank {
			rootY.parent = rootX
		} else if rootX.rank < rootY.rank {
			rootX.parent = rootY
		} else {
			rootX.parent = rootY
			rootY.rank++
		}
	}
}

// Find returns the root node of the node
func (d *DisjointSet) Find(node *Node) *Node {
	if node != node.parent {
		node.parent = d.Find(node.parent)
	}
	return node.parent
}

func (d *DisjointSet) Connected(nodeOne, nodeTwo *Node) bool {
	return d.Find(nodeOne) == d.Find(nodeTwo)
}

func (d *DisjointSet) Size() int {
	return d.size
}
