// Package binarytrees contains types and struct forTrees
package binarytrees

import "strconv"

// BinaryTreeNode represent a BinaryTreeNode in a BinarySearchTree
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTreeIterator struct {
	stack []*BinaryTreeNode
}

// NewBinaryTree creates a new Binary Tree Node
func (t *BinaryTreeNode) NewBinaryTree(data int) BinaryTreeNode {
	return BinaryTreeNode{
		Data: data,
	}
}

func (t *BinaryTreeNode) inorderTraversalIteravely() (result []int) {
	stack := []*BinaryTreeNode{}
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

func (t *BinaryTreeNode) inorderTraversalRecurse(root *BinaryTreeNode) (result []int) {
	if root != nil {
		if root.Left != nil {
			t.inorderTraversalRecurse(t.Left)
		}

		result = append(result, root.Data)

		if root.Right != nil {
			t.inorderTraversalRecurse(t.Right)
		}
	}

	return
}

func (t *BinaryTreeNode) inOrderMorrisTraversal() (result []int) {
	current := t
	var pre *BinaryTreeNode

	for current != nil {
		if current.Left == nil {
			// add the current value of the node
			result = append(result, current.Data)
			// # Move to next right node
			current = current.Right
		} else {
			// # we have a left subtree
			pre = current.Left

			// # find rightmost
			for pre.Right != nil {
				pre = pre.Right
			}

			// # put current after the pre node
			pre.Right = current
			// # store current node
			temp := current
			// # move current to top of new tree
			current = current.Left
			// # original current left be None, avoid infinite loops
			temp.Left = nil
		}
	}

	return
}

// IsValidBst checks if a binary tree is valid
func (t *BinaryTreeNode) IsValidBst() bool {
	panic("Implement me!")
}

// PreOrderTraversal of a binary tree, returns values of each node
func (t *BinaryTreeNode) PreOrderTraversal() (values []int) {

	if t == nil {
		return
	}

	stack := []*BinaryTreeNode{}
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

// PostorderTraversal of a binary tree, returns values of each node starting with left most subtree, uses 2 stacks
// to keep track of values of nodes and pops them from one stack adding them to the other
func (t *BinaryTreeNode) PostOrderTraversal() (values []int) {
	stackOne := []*BinaryTreeNode{}
	stackTwo := []*BinaryTreeNode{}

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

// SearchNode searches for a value in a BST by walking either left or right of the tree given the value is either
// less than or greater than current node respectively. This uses a recursive approach to find a node in a Tree
// if found, returns the node which is the subtree with that value if not found, returns nil
func (t *BinaryTreeNode) SearchNode(node *BinaryTreeNode, val int) *BinaryTreeNode {
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

// InsertNode inserts a BinaryTreeNode into the BST. Inserts it left if the val is less than the current root
// inserts it right if the val is greater than the current root. This operation is repeated recursively
func (root *BinaryTreeNode) InsertNode(val int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{
			Data: val,
		}
	}

	if val < root.Data && root.Left != nil {
		root.InsertNode(val)
	} else if val <= root.Data {
		root.Left = &BinaryTreeNode{
			Data: val,
		}
	} else if val > root.Data && root.Right != nil {
		root.InsertNode(val)
	} else {
		root.Right = &BinaryTreeNode{
			Data: val,
		}
	}
	return root
}

// Height returns the height of the Binary Search Tree
func (root *BinaryTreeNode) Height() int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 0
	}

	height := 0
	queue := []*BinaryTreeNode{}

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
// If both of the values in the 2 nodes provided are greater than the root node, then we move to the right.
// if the nodes are less than the root node, we move to the left.
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
func (root *BinaryTreeNode) LowestCommonAncestor(nodeOne, nodeTwo BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	// if any of the node values matches the data value for the root node, return the root node
	if root.Data == nodeOne.Data || root.Data == nodeTwo.Data {
		return root
	}

	for root != nil {
		// if both node_one and node_two are smaller than root, then LCA lies in the left
		if root.Data > nodeOne.Data && root.Data > nodeTwo.Data {
			root = root.Left
		} else if root.Data < nodeOne.Data && root.Data < nodeTwo.Data {
			// if both node_one and node_two are greater than root, then LCA lies in the right
			root = root.Right
		} else {
			break
		}
	}

	return root
}

// Paths returns all the paths of this binary tree from root to leaf node
func (root *BinaryTreeNode) Paths() (res []string) {
	type Pair struct {
		node *BinaryTreeNode
		path string
	}

	if root == nil {
		return res
	}

	stack := []Pair{{root, ""}}

	for len(stack) != 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		path := item.path

		if !(node.Left != nil || node.Right != nil) {
			res = append(res, path+strconv.Itoa(node.Data))
		}

		if node.Left != nil {
			stack = append(stack, Pair{node.Left, path + strconv.Itoa(node.Data) + "->"})
		}

		if node.Right != nil {
			stack = append(stack, Pair{node.Right, path + strconv.Itoa(node.Data) + "->"})
		}
	}

	return
}
