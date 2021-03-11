// Package binarytrees contains types and struct forTrees
package binarytrees

// BinaryTreeNode represent a BinaryTreeNode in a BinarySearchTree
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// NewBinaryTree creates a new Binary Tree Node
func (t *BinaryTreeNode) NewBinaryTree(data int) BinaryTreeNode {
	return BinaryTreeNode{
		Data: data,
	}
}
