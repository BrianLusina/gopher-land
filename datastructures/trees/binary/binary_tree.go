// package binary contains types and struct for Binary Trees
package binary

import (
	"fmt"
	"gopherland/math/utils"
	"gopherland/pkg/types"
	"strings"
)

// BinaryTree represents a binary tree
type BinaryTree[T types.Comparable] struct {
	root *BinaryTreeNode[T]
}

// NewBinaryTree creates a new binary tree with root as nil
func NewBinaryTree[T types.Comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// func IsBstRecursive[T types.Comparable](node *BinaryTreeNode[T], min, max T) bool {
// 	if node == nil {
// 		return true
// 	}

// 	if node.Data < min || node.Data > max {
// 		return false
// 	}

// 	return IsBstRecursive(node.Left, min, node.Data-1) && IsBstRecursive(node.Right, node.Data+1, max)
// }

// Serialize converts this tree into string representation
func (tree *BinaryTree[T]) Serialize() string {
	if tree.root == nil {
		return ""
	}

	stack := []*BinaryTreeNode[T]{}
	stack = append(stack, tree.root)
	nodeValues := []any{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			nodeValues = append(nodeValues, nil)
		} else {
			nodeValues = append(nodeValues, node.Data)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}

	joined := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nodeValues...)), ","), "[]")

	return joined
}

// IsFull checks if a binary tree is a full binary tree.
// A Full Binary tree has a parent or internal nodes who have either 2 or 0 children
func (tree *BinaryTree[T]) IsFull() bool {
	var isFullHelper func(root *BinaryTreeNode[T]) bool

	isFullHelper = func(root *BinaryTreeNode[T]) bool {
		// if root is nil, this is subtree is not a full binary tree and a this point this is the base case

		if root == nil {
			return false
		}

		// if this root node of this subtree has neither a left nor a right node, then by definition this
		// binary subtree is a full binary tree
		if root.Left == nil && root.Right == nil {
			return true
		}

		if root.Left != nil && root.Right != nil {
			return isFullHelper(root.Left) && isFullHelper(root.Right)
		}

		return false
	}

	return isFullHelper(tree.root)
}

// IsComplete checks if a binary tree is a complete binary tree
func (tree *BinaryTree[T]) IsComplete() bool {
	if tree.root == nil {
		return true
	}

	nodeCount := tree.root.Size()

	var isCompleteHelper func(root *BinaryTreeNode[T], index int) bool

	isCompleteHelper = func(root *BinaryTreeNode[T], index int) bool {
		if root == nil {
			return true
		}

		if index >= nodeCount {
			return false
		}

		return isCompleteHelper(root.Left, 2*index+1) && isCompleteHelper(root.Right, 2*index+2)
	}

	index := 0

	return isCompleteHelper(tree.root, index)
}

// GetDepth returns the depth of the tree
func (tree *BinaryTree[T]) GetDepth() int {
	depth := 0
	node := tree.root
	for node != nil {
		depth++
		node = node.Left
	}

	return depth
}

// IsPerfect checks if a binary tree is perfect
func (tree *BinaryTree[T]) IsPerfect() bool {
	if tree.root == nil {
		return false
	}

	if tree.root.Left == nil && tree.root.Right == nil {
		return true
	}

	depth := tree.GetDepth()

	var isPerfectHelper func(root *BinaryTreeNode[T], level int) bool

	isPerfectHelper = func(root *BinaryTreeNode[T], level int) bool {
		if root == nil {
			return true
		}

		if root.Left == nil && root.Right == nil {
			return depth == level+1
		}

		if root.Left == nil || root.Right == nil {
			return false
		}

		return isPerfectHelper(root.Left, level+1) && isPerfectHelper(root.Right, level+1)
	}

	return isPerfectHelper(tree.root, 0)
}

// IsBalanced checks if a binary tree is balanced.
func (tree *BinaryTree[T]) IsBalanced() bool {
	if tree.root == nil {
		return true
	}

	var isBalancedHelper func(root *BinaryTreeNode[T]) bool

	isBalancedHelper = func(root *BinaryTreeNode[T]) bool {
		leftHeight := 0
		rightHeight := 0

		if root == nil {
			return true
		}

		l := isBalancedHelper(root.Left)
		r := isBalancedHelper(root.Right)

		if utils.AbsDiff(leftHeight, rightHeight) <= 1 {
			return l && r
		}

		return false
	}

	return isBalancedHelper(tree.root)
}
