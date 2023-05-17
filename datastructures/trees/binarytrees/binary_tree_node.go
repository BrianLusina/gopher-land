package binarytrees

import (
	"fmt"
	"gopherland/datastructures/trees"
	"gopherland/pkg/types"
	"strings"
)

// BinaryTreeNode represent a Binary Tree Node in a Binary Tree
// when performing comparisons, it is important to cast this interface to a concrete type
type BinaryTreeNode[T types.Comparable] struct {
	trees.TreeNode[T]
	Left, Right *BinaryTreeNode[T]
}

// NewBinaryTreeNode creates a new BinaryTreeNode
func NewBinaryTreeNode[T types.Comparable](data T) *BinaryTreeNode[T] {
	node := trees.NewTreeNode[T](data)

	return &BinaryTreeNode[T]{
		TreeNode: *node,
		Left:     nil,
		Right:    nil,
	}
}

func (node *BinaryTreeNode[T]) Serialize() string {
	if node == nil {
		return ""
	}

	stack := []*BinaryTreeNode[T]{node}
	nodeValues := []any{node.Data}

	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
			add(&nodeValues, stack[i].Left)
			add(&nodeValues, stack[i].Right)
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

func (t *BinaryTreeNode[T]) InorderTraversalIteravely() (result []T) {
	var stack []*BinaryTreeNode[T]
	current := t

	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result = append(result, current.Data)
		current = current.Right
	}
	return
}

func (t *BinaryTreeNode[T]) InorderTraversalRecurse(root *BinaryTreeNode[T]) (result []T) {
	if root != nil {
		if root.Left != nil {
			t.InorderTraversalRecurse(t.Left)
		}

		result = append(result, root.Data)

		if root.Right != nil {
			t.InorderTraversalRecurse(t.Right)
		}
	}

	return
}

func (t *BinaryTreeNode[T]) inOrderMorrisTraversal() (result []T) {
	current := t
	var pre *BinaryTreeNode[T]

	for current != nil {
		if current.Left == nil {
			// add the current value of the node
			result = append(result, current.Data)
			// # Move to next Right node
			current = current.Right
		} else {
			// # we have a Left subtree
			pre = current.Left

			// # find Rightmost
			for pre.Right != nil {
				pre = pre.Right
			}

			// # put current after the pre node
			pre.Right = current
			// # store current node
			temp := current
			// # move current to top of new tree
			current = current.Left
			// # original current Left be None, avoid infinite loops
			temp.Left = nil
		}
	}

	return
}

// // IsValidBst checks if a binary tree is valid is a valid BST
// func (t *BinaryTreeNode[T]) IsValidBst() bool {
// 	max := 4294967296
// 	min := -4294967296
// 	return IsBstRecursive(t, min, max)
// }

// PreOrderTraversal of a binary tree, returns values of each node
func (t *BinaryTreeNode[T]) PreOrderTraversal() (values []T) {

	if t == nil {
		return
	}

	var stack []*BinaryTreeNode[T]
	current := t

	for current != nil || len(stack) != 0 {
		for current != nil {
			values = append(values, t.Data)
			stack = append(stack, t)
			current = t.Left
		}
		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		current = current.Right
	}

	return
}

// PostOrderTraversal of a binary tree, returns values of each node starting with Left most subtree, uses 2 stacks
// to keep track of values of nodes and pops them from one stack adding them to the other
func (t *BinaryTreeNode[T]) PostOrderTraversal() (values []T) {
	var stackOne []*BinaryTreeNode[T]
	var stackTwo []*BinaryTreeNode[T]

	if t == nil {
		return
	}

	stackOne = append(stackOne, t)

	for len(stackOne) != 0 {
		node := stackOne[len(stackOne)-1]
		stackOne = stackOne[:len(stackOne)-1]
		stackTwo = append(stackTwo, node)

		if node.Left != nil {
			stackOne = append(stackOne, node.Left)
		}

		if node.Right != nil {
			stackOne = append(stackOne, node.Right)
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
func (t *BinaryTreeNode[T]) SearchNode(node *BinaryTreeNode[T], val T) *BinaryTreeNode[T] {
	if node == nil {
		return nil
	}

	if val == node.Data {
		return node
	}

	if val < node.Data && node.Left != nil {
		return t.SearchNode(node.Left, val)
	}

	if val > node.Data && node.Right != nil {
		return t.SearchNode(node.Right, val)
	}

	return nil
}

// InsertNode inserts a BinaryTreeNode into the BST. Inserts it Left if the val is less than the current root
// inserts it Right if the val is greater than the current root. This operation is repeated recursively
func (t *BinaryTreeNode[T]) InsertNode(val T) *BinaryTreeNode[T] {
	if t == nil {
		return &BinaryTreeNode[T]{
			Left:     nil,
			TreeNode: *trees.NewTreeNode(val),
		}
	}

	if val < t.Data && t.Left != nil {
		t.InsertNode(val)
	} else if val <= t.Data {
		t.Left = &BinaryTreeNode[T]{
			TreeNode: *trees.NewTreeNode(val),
		}
	} else if val > t.Data && t.Right != nil {
		t.InsertNode(val)
	} else {
		t.Right = &BinaryTreeNode[T]{
			TreeNode: *trees.NewTreeNode(val),
		}
	}
	return t
}

// Height returns the height of the Binary Search Tree
func (t *BinaryTreeNode[T]) Height() int {
	if t == nil {
		return 0
	}

	if t.Left == nil && t.Right == nil {
		return 0
	}

	height := 0
	var queue []*BinaryTreeNode[T]

	for true {
		currentLevelNodes := len(queue)

		if currentLevelNodes == 0 {
			return height
		}

		height++

		for currentLevelNodes > 0 {
			node := queue[0]
			queue = queue[1 : len(queue)-1]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			currentLevelNodes--
		}
	}
	return 0
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
func (t *BinaryTreeNode[T]) LowestCommonAncestor(nodeOne, nodeTwo BinaryTreeNode[T]) *BinaryTreeNode[T] {
	if t == nil {
		return nil
	}

	// if any of the node values matches the data value for the root node, return the root node
	if t.Data == nodeOne.Data || t.Data == nodeTwo.Data {
		return t
	}

	for t != nil {
		// if both node_one and node_two are smaller than root, then LCA lies in the Left
		if t.Data > nodeOne.Data && t.Data > nodeTwo.Data {
			t = t.Left
		} else if t.Data < nodeOne.Data && t.Data < nodeTwo.Data {
			// if both node_one and node_two are greater than root, then LCA lies in the Right
			t = t.Right
		} else {
			break
		}
	}

	return t
}

// Paths returns all the paths of this binary tree from root to leaf node
func (t *BinaryTreeNode[T]) Paths() (res []string) {
	type Pair struct {
		node *BinaryTreeNode[T]
		path string
	}

	if t == nil {
		return res
	}

	stack := []Pair{{t, ""}}

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		path := item.path

		if !(node.Left != nil || node.Right != nil) {
			res = append(res, path+fmt.Sprintf("%v", node.Data))
		}

		if node.Left != nil {
			stack = append(stack, Pair{node.Left, path + fmt.Sprintf("%v", node.Data) + "->"})
		}

		if node.Right != nil {
			stack = append(stack, Pair{node.Right, path + fmt.Sprintf("%v", node.Data) + "->"})
		}
	}

	return
}

// Size returns the number of nodes in this Tree
func (t *BinaryTreeNode[T]) Size() (counter int) {
	if t == nil {
		return
	}

	counter++
	var stack []*BinaryTreeNode[T]

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			counter++
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			counter++
			stack = append(stack, node.Right)
		}
	}

	return
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
	root.Left = deserializeHelper(idx, arr)
	*idx++
	root.Right = deserializeHelper(idx, arr)
	return root
}
