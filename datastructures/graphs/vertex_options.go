package graphs

import "gopherland/pkg/types"

// VertexOption is the option to apply to a Vertex
type VertexOption[T types.Comparable] func(*Vertex[T])

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

func WithProperties[T types.Comparable](metadata map[any]any) VertexOption[T] {
	return func(v *Vertex[T]) {
		v.properties = metadata
	}
}
