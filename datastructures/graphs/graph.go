package graphs

type Graph interface {
	// AddNode adds a new node to the graph
	AddNode(nodeOne Node) error
	// GetNode returns a node from the graph
	GetNode(data any) *Node
	// Size returns the number of nodes in the graph
	Size() int
}
