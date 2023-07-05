package graphs

import (
	"gopherland/pkg/types"
)

// Vertex represents a node/vertex in a graph
type Vertex[T types.Comparable] struct {
	// ID of the node
	ID string

	// Data of the node
	Data T

	// metadata is a map of metadata to attach to a node
	metadata map[any]any

	// degree is the number of edges connected to this vertex
	degree int

	// adjacentVertices is a slice of this vertices neighbors
	adjacentVertices map[T]*Vertex[T]
}

// NewVertex creates a new node with the given data
func NewVertex[T types.Comparable](data T) *Vertex[T] {
	return &Vertex[T]{
		Data:             data,
		adjacentVertices: map[T]*Vertex[T]{},
		metadata:         map[any]any{},
	}
}

func (v *Vertex[T]) SetMetadata(data map[any]any) {
	v.metadata = data
}

func (v *Vertex[T]) GetMetadata(data map[any]any) {
	v.metadata = data
}

// AddAdjacentVertex adds a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) AddAdjacentVertex(vertex Vertex[T]) {
	if _, ok := v.adjacentVertices[vertex.Data]; !ok {
		v.adjacentVertices[vertex.Data] = &vertex

		// we increase the degree because as long as this vertex has neighbors, it has edges incoming or outgoing which means we
		// increase the degree to reflect the number of increasing edges implied by adding a neighbor
		v.degree++
	}
}

// RemoveAdjacentVertex removes a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) RemoveAdjacentVertex(vertex Vertex[T]) {
	if _, ok := v.adjacentVertices[vertex.Data]; ok {
		delete(v.adjacentVertices, vertex.Data)
		// we decrease the degree because we are removing a neighbor, which means an edge has also been removed
		// therefore decreasing the degree
		v.degree--
	}
}

// GetDegree gets the degree of this vertex
func (v *Vertex[T]) GetDegree() int {
	return v.degree
}

// GetNeighbors gets the neighbors of this vertex
func (v *Vertex[T]) GetNeighbors() []*Vertex[T] {
	neighbors := []*Vertex[T]{}

	for _, v := range v.adjacentVertices {
		neighbors = append(neighbors, v)
	}

	return neighbors
}
