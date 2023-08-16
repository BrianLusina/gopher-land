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

	root1.left = MergeTrees(root1.left, root2.left)
	root1.right = MergeTrees(root1.right, root2.right)

	return root1
}
