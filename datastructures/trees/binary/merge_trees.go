package binary

import "gopherland/pkg/types"

// MergeTrees merges 2 binary trees into a single binary tree
func MergeTrees[T types.Comparable](root1 *BinaryTreeNode[T], root2 *BinaryTreeNode[T]) *BinaryTreeNode[T] {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	r2 := root2.Data
	r1 := root1.Data

	r1 += r2

	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)

	return root1
}
