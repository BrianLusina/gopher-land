package graphs

// Edge represents an edge in a graph
type Edge struct {
	// NodeOne is 1 node in the edge
	NodeOne *Node
	// NodeTwo is the other node in the edge
	NodeTwo *Node
	// edge is possibly another edge connected to this edge. This could be the case for hypergraphs
	edges []*Edge
	// Weight is the weight of the edge, ir 0, this is considered an unweighted edge in a graph
	Weight int
}

// NewEdge creates a new edge between two nodes with a weight
func NewEdge(nodeOne, nodeTwo *Node, weight int) *Edge {
	if nodeOne.Neighbors == nil {
		nodeOne.Neighbors = map[any]*Node{}
		nodeOne.Neighbors[nodeTwo.Data] = nodeTwo
	} else {
		// is nodeTwo already in nodeOne's Neighbors?
		if _, ok := nodeOne.Neighbors[nodeTwo.Data]; !ok {
			nodeOne.Neighbors[nodeTwo.Data] = nodeTwo
		}
	}

	if nodeTwo.Neighbors == nil {
		nodeTwo.Neighbors = map[any]*Node{}
		nodeTwo.Neighbors[nodeOne.Data] = nodeOne
	} else {
		if _, ok := nodeTwo.Neighbors[nodeOne.Data]; !ok {
			nodeTwo.Neighbors[nodeOne.Data] = nodeOne
		}
	}

	// increment the degree of both nodes
	nodeOne.Degree++
	nodeTwo.Degree++

	return &Edge{
		NodeOne: nodeOne,
		NodeTwo: nodeTwo,
		Weight:  weight,
		edges:   []*Edge{},
	}
}

// IsSelfEdge returns true if an edge's 2 nodes are the same
func (e *Edge) IsSelfEdge() bool {
	return e.NodeOne == e.NodeTwo
}

// AddEdge adds an edges to an edge
func (e *Edge) AddEdge(edge ...*Edge) {
	e.edges = append(e.edges, edge...)
}

// GetEdges returns all the edges in an edge if any
func (e *Edge) GetEdges() []*Edge {
	return e.edges
}

// GetWeight returns the weight of the edge
func (e *Edge) GetWeight() int {
	return e.Weight
}

// GetNodes returns the nodes connected by the edge
// this includes the nodes in the edges if any
func (e *Edge) GetNodes() []*Node {
	if len(e.edges) == 0 {
		return []*Node{e.NodeOne, e.NodeTwo}
	}
	nodes := []*Node{e.NodeOne, e.NodeTwo}

	for _, edge := range e.edges {
		nodes = append(nodes, edge.GetNodes()...)
	}

	return nodes
}

// IsWeighted returns true if the edge has a weight
func (e *Edge) IsWeighted() bool {
	return e.Weight != 0
}
