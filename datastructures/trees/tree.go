// Package trees contains types and struct forTrees
package trees

// Tree interface to operate on trees
type Tree interface {
	NewNode()
	// Capacity returns the number of nodes in the tree
	Capacity() int
	// Size returns the size of the tree
	Size() int
	Insert(value interface{})
	InorderTraversalIteratively()
	InorderMorrisTraversal()
}
