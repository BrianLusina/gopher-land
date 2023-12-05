package vertex

import "gopherland/pkg/types"

// VertexOption is the option to apply to a Vertex
type VertexOption[T types.Comparable] func(*Vertex[T])

// WithID sets the ID on the vertext node
func WithID[T types.Comparable](id string) VertexOption[T] {
	return func(v *Vertex[T]) {
		v.id = id
	}
}

// WithAdjacentVertices adds adjacent vertices to this vertex
func WithAdjacentVertices[T types.Comparable](adjacentVertices map[string]*Vertex[T]) VertexOption[T] {
	return func(v *Vertex[T]) {
		for id, adjacentVertex := range adjacentVertices {
			if _, ok := v.adjacentVertices[id]; !ok {
				// add adjacent vertex as a neighbour
				v.adjacentVertices[id] = adjacentVertex

				// add this vertex as a neighbour to adjacent vertex
				adjacentVertex.adjacentVertices[v.id] = v
			}
		}
	}
}

func WithMetadata[T types.Comparable](metadata map[any]any) VertexOption[T] {
	return func(v *Vertex[T]) {
		v.metadata = metadata
	}
}
