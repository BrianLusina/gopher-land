package binary

import (
	"fmt"
	"gopherland/datastructures/trees"
	"gopherland/pkg/types"
	"strings"
)

// BinaryTreeNodeOption allows setting options to a Binary TreeNode
type BinaryTreeNodeOption[T types.Comparable] func(*BinaryTreeNode[T])

// Left sets the left node
func Left[T types.Comparable](left *BinaryTreeNode[T]) BinaryTreeNodeOption[T] {
	return func(node *BinaryTreeNode[T]) {
		node.left = left
	}
}

// Right sets the right node
func Right[T types.Comparable](right *BinaryTreeNode[T]) BinaryTreeNodeOption[T] {
	return func(node *BinaryTreeNode[T]) {
		node.right = right
	}
}

// BinaryTreeNode represent a Binary Tree Node in a Binary Tree
type BinaryTreeNode[T types.Comparable] struct {
	trees.TreeNode[T]
	left, right *BinaryTreeNode[T]
}

// NewBinaryTreeNode creates a new BinaryTreeNode
func NewBinaryTreeNode[T types.Comparable](data T, opts ...BinaryTreeNodeOption[T]) *BinaryTreeNode[T] {
	node := trees.NewTreeNode(data)

	binaryTreeNode := &BinaryTreeNode[T]{
		TreeNode: *node,
		left:     nil,
		right:    nil,
	}

	for _, opt := range opts {
		opt(binaryTreeNode)
	}

	return binaryTreeNode
}

// Left retrieves the left node of a binary tree node
func (node *BinaryTreeNode[T]) Left() *BinaryTreeNode[T] {
	return node.left
}

// SetLeft sets the left node of a binary tree node
func (node *BinaryTreeNode[T]) SetLeft(n *BinaryTreeNode[T]) {
	node.left = n
}

// Right retrieves the right node of a binary tree node
func (node *BinaryTreeNode[T]) Right() *BinaryTreeNode[T] {
	return node.right
}

// SetRight sets the right node of a binary tree node
func (node *BinaryTreeNode[T]) SetRight(n *BinaryTreeNode[T]) {
	node.right = n
}

// Serialize serializes a node into a string
func (node *BinaryTreeNode[T]) Serialize() string {
	if node == nil {
		return ""
	}

	stack := []*BinaryTreeNode[T]{node}
	nodeValues := []any{node.Data}

	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			if stack[i].left != nil {
				stack = append(stack, stack[i].left)
			}
			if stack[i].right != nil {
				stack = append(stack, stack[i].right)
			}
			add(&nodeValues, stack[i].left)
			add(&nodeValues, stack[i].right)
		}
		stack = stack[l:]
	}

	joined := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nodeValues...)), ","), "[]")
	return joined
}

func add[T types.Comparable](c *[]any, node *BinaryTreeNode[T]) {
	if node == nil {
		*c = append(*c, "nil")
	} else {
		*c = append(*c, node.Data)
	}
}

func (node *BinaryTreeNode[T]) InorderTraversalIteravely() (result []T) {
	var stack []*BinaryTreeNode[T]
	current := node

	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, current.Data)
		current = current.right
	}
	return
}

func (node *BinaryTreeNode[T]) InorderTraversalRecurse(root *BinaryTreeNode[T]) (result []T) {
	if root != nil {
		if root.left != nil {
			node.InorderTraversalRecurse(node.left)
		}

		result = append(result, root.Data)

		if root.right != nil {
			node.InorderTraversalRecurse(node.right)
		}
	}

	return
}

func (node *BinaryTreeNode[T]) InOrderMorrisTraversal() (result []T) {
	current := node
	var pre *BinaryTreeNode[T]

	for current != nil {
		if current.left == nil {
			// add the current value of the node
			result = append(result, current.Data)
			// # Move to next Right node
			current = current.right
		} else {
			// # we have a Left subtree
			pre = current.left

			// # find Rightmost
			for pre.right != nil {
				pre = pre.right
			}

			// # put current after the pre node
			pre.right = current
			// # store current node
			temp := current
			// # move current to top of new tree
			current = current.left
			// # original current Left be None, avoid infinite loops
			temp.left = nil
		}
	}

	return
}

// // IsValidBst checks if a binary tree is valid is a valid BST
// func (node *BinaryTreeNode[T]) IsValidBst() bool {
// 	max := 4294967296
// 	min := -4294967296
// 	return IsBstRecursive(t, min, max)
// }

// PreOrderTraversal of a binary tree, returns values of each node
func (node *BinaryTreeNode[T]) PreOrderTraversal() (values []T) {

	if node == nil {
		return
	}

	var stack []*BinaryTreeNode[T]
	current := node

	for current != nil || len(stack) != 0 {
		for current != nil {
			values = append(values, node.Data)
			stack = append(stack, node)
			current = node.left
		}
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		current = current.right
	}

	return
}

// PostOrderTraversal of a binary tree, returns values of each node starting with Left most subtree, uses 2 stacks
// to keep track of values of nodes and pops them from one stack adding them to the other
func (node *BinaryTreeNode[T]) PostOrderTraversal() (values []T) {
	var stackOne []*BinaryTreeNode[T]
	var stackTwo []*BinaryTreeNode[T]

	if node == nil {
		return
	}

	stackOne = append(stackOne, node)

	for len(stackOne) != 0 {
		node := stackOne[len(stackOne)-1]
		stackOne = stackOne[:len(stackOne)-1]
		stackTwo = append(stackTwo, node)

		if node.left != nil {
			stackOne = append(stackOne, node.left)
		}

		if node.right != nil {
			stackOne = append(stackOne, node.right)
		}
	}

	for len(stackTwo) != 0 {
		node := stackTwo[len(stackTwo)-1]
		stackTwo = stackTwo[:len(stackTwo)-1]
		values = append(values, node.Data)
	}

	return
}

// SearchNode searches for a value in a BST by walking either Left or Right of the tree given the value is either
// less than or greater than current node respectively. This uses a recursive approach to find a node in a Tree
// if found, returns the node which is the subtree with that value if not found, returns nil
func (node *BinaryTreeNode[T]) SearchNode(n *BinaryTreeNode[T], val T) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	if val == node.Data {
		return node
	}

	if val < node.Data && node.left != nil {
		return node.SearchNode(node.left, val)
	}

	if val > node.Data && node.right != nil {
		return node.SearchNode(node.right, val)
	}

	return nil
}

// InsertNode inserts a BinaryTreeNode into the BST. Inserts it Left if the val is less than the current root
// inserts it Right if the val is greater than the current root. This operation is repeated recursively
func (node *BinaryTreeNode[T]) InsertNode(val T) *BinaryTreeNode[T] {
	if node == nil {
		return &BinaryTreeNode[T]{
			left:     nil,
			TreeNode: *trees.NewTreeNode(val),
		}
	}

	if val < node.Data && node.left != nil {
		node.InsertNode(val)
	} else if val <= node.Data {
		node.left = &BinaryTreeNode[T]{
			TreeNode: *trees.NewTreeNode(val),
		}
	} else if val > node.Data && node.right != nil {
		node.InsertNode(val)
	} else {
		node.right = &BinaryTreeNode[T]{
			TreeNode: *trees.NewTreeNode(val),
		}
	}
	return node
}

// Height returns the height of the Binary Search Tree
func (node *BinaryTreeNode[T]) Height() int {
	if node == nil {
		return 0
	}

	if node.left == nil && node.right == nil {
		return 0
	}

	height := 0
	var queue []*BinaryTreeNode[T]

	for {
		currentLevelNodes := len(queue)

		if currentLevelNodes == 0 {
			return height
		}

		height++

		for currentLevelNodes > 0 {
			node := queue[0]
			queue = queue[1 : len(queue)-1]

			if node.left != nil {
				queue = append(queue, node.left)
			}

			if node.right != nil {
				queue = append(queue, node.right)
			}

			currentLevelNodes--
		}
	}
}

// LowestCommonAncestor returns the LCA of 2 nodes
// Considering it is a BST, we can assume that this tree is a valid BST, we could also check for this
// If both of the values in the 2 nodes provided are greater than the root node, then we move to the Right.
// if the nodes are less than the root node, we move to the Left.
// If there is no root node, then we exit and return None, as no common ancestor could exist in such a case with
// no root node.
//
// Assumptions:
// - assumes that the node itself can also be an ancestor/descendant of itself
//
// Complexity Analysis:
// Time Complexity: O(h).
// The Time Complexity of the above solution is O(h), where h is the height of the tree.
//
// Space Complexity: O(1).
// The space complexity of the above solution is constant.
func (node *BinaryTreeNode[T]) LowestCommonAncestor(nodeOne, nodeTwo BinaryTreeNode[T]) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	// if any of the node values matches the data value for the root node, return the root node
	if node.Data == nodeOne.Data || node.Data == nodeTwo.Data {
		return node
	}

	for node != nil {
		// if both node_one and node_two are smaller than root, then LCA lies in the Left
		if node.Data > nodeOne.Data && node.Data > nodeTwo.Data {
			node = node.left
		} else if node.Data < nodeOne.Data && node.Data < nodeTwo.Data {
			// if both node_one and node_two are greater than root, then LCA lies in the Right
			node = node.right
		} else {
			break
		}
	}

	return node
}

type pair[T types.Comparable] struct {
	node *BinaryTreeNode[T]
	path string
}

// Paths returns all the paths of this binary tree from root to leaf node
func (node *BinaryTreeNode[T]) Paths() (res []string) {

	if node == nil {
		return res
	}

	stack := []pair[T]{{node, ""}}

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		path := item.path

		if !(node.left != nil || node.right != nil) {
			res = append(res, path+fmt.Sprintf("%v", node.Data))
		}

		if node.left != nil {
			stack = append(stack, pair[T]{node.left, path + fmt.Sprintf("%v", node.Data) + "->"})
		}

		if node.right != nil {
			stack = append(stack, pair[T]{node.right, path + fmt.Sprintf("%v", node.Data) + "->"})
		}
	}

	return
}

// Size returns the number of nodes in this Tree
func (treeNode *BinaryTreeNode[T]) Size() (counter int) {
	if treeNode == nil {
		return
	}

	counter++
	stack := []*BinaryTreeNode[T]{treeNode}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.left != nil {
			counter++
			stack = append(stack, node.left)
		}

		if node.right != nil {
			counter++
			stack = append(stack, node.right)
		}
	}

	return
}

// IsFull checks if a binary tree is a full binary tree.
// A Full Binary tree has a parent or internal nodes who have either 2 or 0 children
func (node *BinaryTreeNode[T]) IsFull() bool {
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

	return isFullHelper(node)
}

// Deserialize converts string data into a binary tree
func Deserialize[T string](data string) *BinaryTreeNode[string] {
	if len(data) == 0 {
		return nil
	}

	nodeValues := strings.Split(data, ",")

	idx := 0
	treeNode := deserializeHelper(&idx, nodeValues)
	return treeNode
}

func deserializeHelper[T string](idx *int, arr []string) *BinaryTreeNode[string] {
	if arr[*idx] == "#" {
		return nil
	}
	root := NewBinaryTreeNode(arr[*idx])
	*idx++
	root.left = deserializeHelper(idx, arr)
	*idx++
	root.right = deserializeHelper(idx, arr)
	return root
}
