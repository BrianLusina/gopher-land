package bst

import (
	"gopherland/datastructures/trees/binary"
	"gopherland/pkg/types"
	"slices"
)

// ConstructTree constructs a binary search tree from a slice of comparable values.
// It first sorts the slice, and then constructs the tree in a bottom-up manner.
// The tree is constructed by recursively constructing the left and right subtrees
// and setting the root node's left and right children accordingly.
// The resulting tree is a valid binary search tree with minimum height
// If the input slice is empty, the function returns nil.
func ConstructTree[T types.Comparable](data []T) *BinarySearchTree[T] {
	if len(data) == 0 {
		return nil
	}

	// sort the data, this will ensure that if the data is not sorted it is sorted in place
	slices.Sort(data)

	var constructTreeHelper func(left, right int) *binary.BinaryTreeNode[T]

	constructTreeHelper = func(left, right int) *binary.BinaryTreeNode[T] {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		node := binary.NewBinaryTreeNode(data[mid])
		// construct the left subtree
		node.SetLeft(constructTreeHelper(left, mid-1))
		// construct the right subtree
		node.SetRight(constructTreeHelper(mid+1, right))
		return node
	}

	root := constructTreeHelper(0, len(data)-1)
	bst := &BinarySearchTree[T]{root: root}
	return bst
}
