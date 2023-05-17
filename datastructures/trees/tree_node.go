package trees

import "gopherland/pkg/types"

// TreeNode represent a Node in a Tree
type TreeNode[T types.Comparable] struct {
	Data     T
	Children []*TreeNode[T]
}

// NewTreeNode returns a new TreeNode
func NewTreeNode[T types.Comparable](data T) *TreeNode[T] {
	return &TreeNode[T]{
		Data:     data,
		Children: nil,
	}
}
