package trees

import "gopherland/pkg/types"

// TreeNode represent a Node in a Tree Each node contains data and provides additional properties such as a unique key, depth, and height.
type TreeNode[T types.Comparable] struct {
	Data     T
	Key      any
	Children []*TreeNode[T]
	// Height represents the height of a node in the tree.
	// The height of a node is defined as the number of edges or links on the longest path
	// from the node to a leaf. A leaf node has a height of 0. For nodes that are not part
	// of a tree or whose height cannot be determined, the value will be `null`.
	Height int

	// Depth represents the depth of a node in the tree.
	// The depth of a node is defined as the number of edges or links from the tree's root
	// to the node. For the root node, the depth is 0. For other nodes, it increases incrementally
	// moving downward in the tree structure.
	Depth int

	// Degree represents the degree of a node in a tree.
	// The degree of a node is defined as the number of immediate child nodes it has.
	// In the context of a tree structure, this represents the branching factor
	// of the node. For a leaf node, the degree is `0`, while for internal nodes,
	// it is equal to the number of connected child nodes.
	Degree int
}

// NewTreeNode returns a new TreeNode
func NewTreeNode[T types.Comparable](data T, opts ...TreeNodeOption[T]) *TreeNode[T] {
	node := &TreeNode[T]{
		Data:     data,
		Key:      data,
		Children: nil,
	}

	for _, opt := range opts {
		opt(node)
	}

	return node
}
