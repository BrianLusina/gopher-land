// Package trees contains types and struct forTrees
package trees

import "gopherland/pkg/types"

// Tree interface to operate on trees
type Tree[T types.Comparable] interface {
	NewNode()

	// Capacity returns the number of nodes in the tree
	Capacity() int

	// Height returns the size of the tree
	Height() int

	// Serialize serializes the tree
	Serialize() string

	// Deserialize deserializes the tree
	Deserialize() *TreeNode[T]

	Insert(value interface{})

	InorderTraversalIteratively()

	InorderMorrisTraversal()
}
