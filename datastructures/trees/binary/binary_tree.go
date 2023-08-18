// package binary contains types and struct for Binary Trees
package binary

import (
	"fmt"
	"gopherland/math/utils"
	"gopherland/pkg/types"
	pkgUtils "gopherland/pkg/utils"
	"strings"
)

// BinaryTree represents a binary tree
type BinaryTree[T types.Comparable] struct {
	root *BinaryTreeNode[T]
}

// NewBinaryTree creates a new binary tree with root as nil
func NewBinaryTree[T types.Comparable](root *BinaryTreeNode[T]) *BinaryTree[T] {
	if root != nil {
		return &BinaryTree[T]{
			root: root,
		}
	}
	return &BinaryTree[T]{}
}

// func IsBstRecursive[T types.Comparable](node *BinaryTreeNode[T], min, max T) bool {
// 	if node == nil {
// 		return true
// 	}

// 	if node.Data < min || node.Data > max {
// 		return false
// 	}

// 	return IsBstRecursive(node.left, min, node.Data-1) && IsBstRecursive(node.right, node.Data+1, max)
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
			stack = append(stack, node.right)
			stack = append(stack, node.left)
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
		if root.left == nil && root.right == nil {
			return true
		}

		if root.left != nil && root.right != nil {
			return isFullHelper(root.left) && isFullHelper(root.right)
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

		return isCompleteHelper(root.left, 2*index+1) && isCompleteHelper(root.right, 2*index+2)
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
		node = node.left
	}

	return depth
}

// IsPerfect checks if a binary tree is perfect
func (tree *BinaryTree[T]) IsPerfect() bool {
	if tree.root == nil {
		return false
	}

	if tree.root.left == nil && tree.root.right == nil {
		return true
	}

	depth := tree.GetDepth()

	var isPerfectHelper func(root *BinaryTreeNode[T], level int) bool

	isPerfectHelper = func(root *BinaryTreeNode[T], level int) bool {
		if root == nil {
			return true
		}

		if root.left == nil && root.right == nil {
			return depth == level+1
		}

		if root.left == nil || root.right == nil {
			return false
		}

		return isPerfectHelper(root.left, level+1) && isPerfectHelper(root.right, level+1)
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

		l := isBalancedHelper(root.left)
		r := isBalancedHelper(root.right)

		if utils.AbsDiff(leftHeight, rightHeight) <= 1 {
			return l && r
		}

		return false
	}

	return isBalancedHelper(tree.root)
}

// Height Returns the height of a Tree or the maximum depth of the tree. A Tree's height is the number of links from its root to the furthest leaf
func (tree *BinaryTree[T]) Height() int {
	var heightHelper func(root *BinaryTreeNode[T]) int
	heightHelper = func(root *BinaryTreeNode[T]) int {
		if root == nil {
			return 0
		}
		return utils.Max(heightHelper(root.left), heightHelper(root.right)) + 1
	}

	return heightHelper(tree.root)
}

// LeafSimilar returns true if this tree has similar leaf value sequence to another tree.
// For example: If this tree has nodes = [3,5,1,6,2,9,8,null,null,7,4] and other tree has nodes =
// [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]. Then the leaf value sequence of both is [6,7,4,9,8] which is similar
func (tree *BinaryTree[T]) LeafSimilar(other *BinaryTree[T]) bool {
	if (tree.root == nil && other.root != nil) || (tree.root != nil && other.root == nil) {
		return false
	}

	leaves1 := []T{}
	leaves2 := []T{}

	var dfs func(node *BinaryTreeNode[T], leafValues *[]T)
	dfs = func(node *BinaryTreeNode[T], leafValues *[]T) {
		if node != nil {
			if node.left == nil && node.right == nil {
				*leafValues = append(*leafValues, node.Data)
			}
			dfs(node.left, leafValues)
			dfs(node.right, leafValues)
		}
	}

	dfs(tree.root, &leaves1)
	dfs(other.root, &leaves2)
	return pkgUtils.EqualSlices(leaves1, leaves2)
}

// Finds the number of good nodes in a tree. A good node is a node in which in the path from root to the node there
// are no nodes with a value greater than it
func (tree *BinaryTree[T]) CountGoodNodes() int {
	if tree.root == nil {
		return 0
	}

	// root is regarded as good
	if tree.root.left == nil && tree.root.right == nil {
		return 1
	}

	var goodNodesHelper func(node *BinaryTreeNode[T], data T) int
	goodNodesHelper = func(node *BinaryTreeNode[T], data T) int {
		if node != nil {
			nodeCount := goodNodesHelper(node.left, pkgUtils.Max[T](data, node.Data)) + goodNodesHelper(node.right, pkgUtils.Max[T](data, node.Data))
			if node.Data >= data {
				nodeCount++
			}
			return nodeCount
		}
		return 0
	}

	return goodNodesHelper(tree.root, tree.root.Data)
}
