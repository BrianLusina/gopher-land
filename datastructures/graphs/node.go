package graphs

// Node represents a node in a graph
type Node struct {
	// Data of the node
	Data any
	// Neighbours is an adjacency list of the node. This contains a map
	// where the key is the Node's Data and the value is the Node
	Neighbors map[any]*Node
	// Degree is the number of edges connecting the node to other nodes
	// this applies to unweighted graphs
	Degree int
}

// NewNode creates a new node with the given data
func NewNode(data any) *Node {
	return &Node{
		Data:      data,
		Neighbors: make(map[any]*Node),
		Degree:    0,
	}
}
