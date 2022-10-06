package graphs

// Node represents a node/vertex in a graph
type Node struct {
	// ID of the node
	ID string

	// Data of the node
	Data any

	// Properties is a map of metadata to attach to a node
	Properties map[any]any

	// IncomingEdges is a slice of incoming edges. Edges pointing to this node
	IncomingEdges []Edge

	// OutgoingEdges is a slice of outgoing edges. Edges pointing away from this node
	OutgoingEdges []Edge
}

// NewNode creates a new node with the given data
func NewNode(data any) *Node {
	return &Node{
		Data: data,
	}
}

func (n *Node) SetIncomingEdges(edges []Edge) {
	n.IncomingEdges = edges
}

func (n *Node) SetOutgoingEdges(edges []Edge) {
	n.OutgoingEdges = edges
}

func (n *Node) SetProperties(properties map[any]any) {
	n.Properties = properties
}
