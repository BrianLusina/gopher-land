package binarytrees

import "gopherland/datastructures/trees"

// BinaryTreeNode represent a Binary Tree Node in a Binary Tree
// when performing comparisons, it is important to cast this interface to a concrete type
type BinaryTreeNode struct {
	trees.TreeNode
	Left, Right *BinaryTreeNode
}

// NewBinaryTreeNode creates a new BinaryTreeNode
func NewBinaryTreeNode(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		TreeNode: trees.NewTreeNode(data),
		Left:     nil,
		Right:    nil,
	}
}
