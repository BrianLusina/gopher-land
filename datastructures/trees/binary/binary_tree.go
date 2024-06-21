// package binary contains types and struct for Binary Trees
package binary

import (
	"fmt"
	"gopherland/datastructures/queues/fifo"
	"gopherland/datastructures/trees"
	"gopherland/math/utils"
	"gopherland/pkg/types"
	pkgUtils "gopherland/pkg/utils"
	"math"
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

// NewBinaryTreeFromSlice creates a binary tree from a slice
func NewBinaryTreeFromSlice[T types.Comparable](data []T) *BinaryTree[T] {
	if len(data) == 0 {
		return &BinaryTree[T]{}
	}

	// root node is the first value at index 0
	root := &BinaryTreeNode[T]{
		TreeNode: trees.TreeNode[T]{
			Data: data[0],
		},
	}

	var createTree func(node *BinaryTreeNode[T], index int, data []T)
	createTree = func(node *BinaryTreeNode[T], index int, data []T) {
		if node != nil {
			leftChildPosition := 2*index + 1
			rightChildPosition := 2*index + 2

			if leftChildPosition < len(data) {
				leftChildValue := data[leftChildPosition]
				if pkgUtils.IsZero[T](leftChildValue) {
					node.left = nil
				} else {
					node.left = NewBinaryTreeNode[T](leftChildValue)
				}

				createTree(node.left, leftChildPosition, data)
			}

			if rightChildPosition < len(data) {
				rightChildValue := data[rightChildPosition]
				if pkgUtils.IsZero[T](rightChildValue) {
					node.right = nil
				} else {
					node.right = NewBinaryTreeNode[T](rightChildValue)
				}

				createTree(node.right, rightChildPosition, data)
			}
		}
	}

	createTree(root, 0, data)
	return &BinaryTree[T]{
		root: root,
	}
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
	nodeValues := []string{}

	var dfs func(node *BinaryTreeNode[T])
	dfs = func(node *BinaryTreeNode[T]) {
		if node == nil {
			nodeValues = append(nodeValues, "nil")
			return
		}

		nodeValues = append(nodeValues, fmt.Sprint(node.Data))
		dfs(node.left)
		dfs(node.right)
	}

	dfs(tree.root)

	return strings.Join(nodeValues, ",")
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

// RightSideView returns a slice of nodes that can be seen from the right side
func (tree *BinaryTree[T]) RightSideView() []T {
	if tree.root == nil {
		return []T{}
	}

	// root is available, but no left nor right subtrees
	if tree.root.left == nil && tree.root.right == nil {
		return []T{tree.root.Data}
	}

	result := []T{}
	queue := []*BinaryTreeNode[T]{tree.root}

	for len(queue) != 0 {
		levelLength := len(queue)

		result = append(result, queue[0].Data)

		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.right != nil {
				queue = append(queue, node.right)
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
		}
	}

	return result
}

// MaxLevelSum Returns the smallest level x such that the sum of all the values of nodes at level x is maximal Returns Int maximum value at level x
func (tree *BinaryTree[T]) MaxLevelSum() int {
	if tree.root == nil {
		return 0
	}

	if tree.root.left == nil && tree.root.right == nil {
		return 1
	}

	maximumSum := tree.root.Data
	level := 0
	smallestLevel := 0

	levels := []*BinaryTreeNode[T]{tree.root}

	for len(levels) != 0 {
		level++
		sumAtCurrentLevel := pkgUtils.Zero[T]()

		for range levels {
			node := levels[0]
			levels = levels[1:]

			sumAtCurrentLevel += node.Data

			if node.left != nil {
				levels = append(levels, node.left)
			}

			if node.right != nil {
				levels = append(levels, node.right)
			}
		}

		if sumAtCurrentLevel > maximumSum {
			maximumSum = sumAtCurrentLevel
			smallestLevel = level
		}
	}

	return smallestLevel
}

// LevelOrderTraversal traverses the tree in a level order fashion returning the data on each node in a slice following the property
// of arrays where an element in the array with an index of i can be used to find the left and the right children using 2i+1 for the left node
// and 2i+2 for the right node and the parent node can be found with floor((i-1)/2)
func (tree *BinaryTree[T]) LevelOrderTraversal() []T {
	if tree.root == nil {
		return []T{}
	}

	queue := fifo.NewFifoQueue[*BinaryTreeNode[T]](math.MaxInt16)
	queue.Enqueue(tree.root)
	result := []T{}

	for !queue.IsEmpty() {
		p, _ := queue.Peek()
		result = append(result, p.Data)

		node, _ := queue.Dequeue()

		if node.left != nil {
			queue.Enqueue(node.left)
		}

		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

	return result
}
