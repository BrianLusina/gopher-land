package vertex

import (
	"gopherland/pkg/types"
)

type (
	// Vertex represents a node/vertex in a graph
	Vertex[T types.Comparable] struct {
		// id of the node
		id string

		// data of the node
		data T

		// metadata is a map of metadata to attach to a node
		metadata map[any]any

		// degree is the number of edges connected to this vertex
		degree int

		// adjacentVertices is a map of this vertices neighbors
		adjacentVertices map[string]*Vertex[T]

		// TODO: add field to keep track of nth degree neighbors of this vertex, for example, 2nd degree neighbor, which means that in order to get to
		// the 2nd degree vertex, one has to move through 1 other vertex so, A->B->C, where C is a 2nd degree neighbor of A & vice versa.
	}
)

// NewVertex creates a new node with the given data
func NewVertex[T types.Comparable](data T, options ...VertexOption[T]) *Vertex[T] {
	vertex := &Vertex[T]{
		data:             data,
		adjacentVertices: map[string]*Vertex[T]{},
		metadata:         map[any]any{},
	}

	for _, option := range options {
		option(vertex)
	}

	return vertex
}

func (v *Vertex[T]) SetMetadata(data map[any]any) {
	v.metadata = data
}

func (v *Vertex[T]) GetMetadata(data map[any]any) {
	v.metadata = data
}

// AddAdjacentVertex adds a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) AddAdjacentVertex(vertex Vertex[T]) {
	if _, ok := v.adjacentVertices[vertex.id]; !ok {
		v.adjacentVertices[vertex.id] = &vertex

		// we increase the degree because as long as this vertex has neighbors, it has edges incoming or outgoing which means we
		// increase the degree to reflect the number of increasing edges implied by adding a neighbor
		v.degree++
	}
}

// RemoveAdjacentVertex removes a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) RemoveAdjacentVertex(vertex Vertex[T]) {
	if _, ok := v.adjacentVertices[vertex.id]; ok {
		delete(v.adjacentVertices, vertex.id)
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

// Data retruns the data contained in this vertex
func (v *Vertex[T]) Data() T {
	return v.data
}
