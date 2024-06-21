package trees

import "gopherland/pkg/types"

// TreeNode represent a Node in a Tree
type TreeNode[T types.Comparable] struct {
	Data     T
	Key      any
	Children []*TreeNode[T]
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
