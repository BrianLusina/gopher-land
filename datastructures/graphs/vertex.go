package graphs

import (
	"gopherland/datastructures/graphs/edges"
	"gopherland/pkg/types"
)

type (
	// Vertex represents a node/vertex in a graph
	Vertex[T types.Comparable] struct {
		// id of the node
		id string

		// data of the node
		data T

		// properties is a map of properties to attach to a node
		properties map[any]any

		// adjacentVertices is a map of this vertices neighbors
		adjacentVertices map[string]*Vertex[T]

		// edges is a map of edges connected to this vertex
		edges  map[string]*edges.Edge[T]
		degree int
	}
)

// NewVertex creates a new node/vertex with the given data
func NewVertex[T types.Comparable](id string, data T, options ...VertexOption[T]) *Vertex[T] {
	vertex := &Vertex[T]{
		id:               id,
		data:             data,
		adjacentVertices: map[string]*Vertex[T]{},
		properties:       map[any]any{},
		edges:            map[string]*edges.Edge[T]{},
	}

	for _, option := range options {
		option(vertex)
	}

	return vertex
}

func (v *Vertex[T]) SetProperty(key any, value any) {
	v.properties[key] = value
}

func (v *Vertex[T]) GetProperty(key any) (any, bool) {
	value, ok := v.properties[key]
	return value, ok
}

func (v *Vertex[T]) GetProperties() map[any]any {
	return v.properties
}

// AddAdjacentVertex adds a vertex as a "neighbor" or adjacent vertex to this vertex
func (v *Vertex[T]) AddAdjacentVertex(adjacentVertex Vertex[T]) {
	if _, ok := v.adjacentVertices[adjacentVertex.id]; !ok {
		v.adjacentVertices[adjacentVertex.id] = &adjacentVertex
		adjacentVertex.AddAdjacentVertex(*v)

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
