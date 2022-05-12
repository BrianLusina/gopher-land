package graphs

// Edge represents an edge in a graph
type Edge struct {
	NodeOne *Node
	NodeTwo *Node
	Weight  int
}

// NewEdge creates a new edge between two nodes with a weight
func NewEdge(nodeOne, nodeTwo Node, weight int) *Edge {
	if nodeOne.Neighbors == nil {
		nodeOne.Neighbors = map[any]*Node{}
		nodeOne.Neighbors[nodeTwo.Data] = &nodeTwo
	} else {
		// is nodeTwo already in nodeOne's Neighbors?
		if _, ok := nodeOne.Neighbors[nodeTwo.Data]; !ok {
			nodeOne.Neighbors[nodeTwo.Data] = &nodeTwo
		}
	}

	if nodeTwo.Neighbors == nil {
		nodeTwo.Neighbors = map[any]*Node{}
		nodeTwo.Neighbors[nodeOne.Data] = &nodeOne
	} else {
		if _, ok := nodeTwo.Neighbors[nodeOne.Data]; !ok {
			nodeTwo.Neighbors[nodeOne.Data] = &nodeOne
		}
	}

	return &Edge{
		NodeOne: &nodeOne,
		NodeTwo: &nodeTwo,
		Weight:  weight,
	}
}
