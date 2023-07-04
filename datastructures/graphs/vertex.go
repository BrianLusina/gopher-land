package graphs

import "gopherland/pkg/types"

// Vertex represents a node/vertex in a graph
type Vertex[T types.Comparable] struct {
	// ID of the node
	ID string

	// Data of the node
	Data T

	// metadata is a map of metadata to attach to a node
	Metadata map[any]any

	// incomingEdges is a slice of incoming edges. Edges pointing to this node
	incomingEdges []Edge

	// outgoingEdges is a slice of outgoing edges. Edges pointing away from this node
	outgoingEdges []Edge
}

// NewVertex creates a new node with the given data
func NewVertex[T types.Comparable](data T) *Vertex[T] {
	return &Vertex[T]{
		Data: data,
	}
}

func (v *Vertex[T]) SetIncomingEdges(edges []Edge) {
	v.incomingEdges = edges
}

func (v *Vertex[T]) SetOutgoingEdges(edges []Edge) {
	v.outgoingEdges = edges
}

func (v *Vertex[T]) SetMetadata(data map[any]any) {
	v.Metadata = data
}

// AddAdjacentVertex adds a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) AddAdjacentVertex(vertex Vertex[T]) {

}
