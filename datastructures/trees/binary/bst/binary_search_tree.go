package bst

import (
	"errors"
	"fmt"
	"gopherland/datastructures/trees/binary"
	"gopherland/math/utils"
	"gopherland/pkg/types"
	"strconv"
)

// BinarySearchTree is a binary tree that satisfies the binary tree interface
type BinarySearchTree[T types.Comparable] struct {
	root *binary.BinaryTreeNode[T]
}

// NewBinarySearchTree returns a new binary search tree
func NewBinarySearchTree[T types.Comparable]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

// Insert inserts a new node with the value into the binary search tree
func (bst *BinarySearchTree[T]) Insert(value T) {
	// we don't have a root
	if bst.root == nil {
		bst.root = binary.NewBinaryTreeNode(value)
		return
	}

	// insert helper function to insert a new value in the binary search tree
	var insert func(node *binary.BinaryTreeNode[T], value T)

	insert = func(node *binary.BinaryTreeNode[T], value T) {
		if value <= node.Data {
			if node.Left != nil {
				insert(node.Left(), value)
			} else {
				node.SetLeft(binary.NewBinaryTreeNode(value))
			}
		} else {
			if node.Right != nil {
				insert(node.Right(), value)
			} else {
				node.SetRight(binary.NewBinaryTreeNode(value))
			}
		}
	}

	insert(bst.root, value)
}

// Delete deletes a value from the BinarySearchTree
func (bst *BinarySearchTree[T]) Delete(value T) error {
	if bst.root == nil {
		return fmt.Errorf("tree has no root, can not delete value %v", value)
	}

	var lift func(node, nodeToDelete *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T]
	lift = func(node, nodeToDelete *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
		// if the current node has a left child, we recursively call this function to continue down the left subtree to find the successor node
		if node.Left != nil {
			node.SetLeft(lift(node.Left(), nodeToDelete))
			return node
		} else {
			// if the current node has no left child, that means the current node of this function is the successor node, and we take it value
			// and make it the new value of the node that we are deleting
			nodeToDelete.Data = node.Data
			// we return the successor node's right child to be now used as it's parent's left child
			return node.Right()
		}
	}

	// delete is a helper function that recurses over the subtrees of the root node to delete the value provided
	var delete func(val T, node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T]

	delete = func(val T, node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
		// base case is when we have hit the bottom of the tree, and the parent node has no children
		if node == nil {
			return nil
		} else if val < node.Data {
			// if the value we are deleting is less than the current node's value, we set the left child to be
			// the return value of a recursive call
			node.SetLeft(delete(val, node.Left()))

			// we return the current node and it's subtree if existent to be use as the new value of it's parent's left child
			return node
		} else if val > node.Data {
			// if the value we are deleting is greater than the current node's value, we set the right child to be
			// the return value of a recursive call
			node.SetRight(delete(val, node.Right()))
			return node
		} else if val == node.Data {
			// the current node contains the value we want to delete
			if node.Left == nil {
				// if the current node has no left child, we delete it by returning the right child and it's subtree if existent to be its parents
				// new subtree
				return node.Right()
			} else if node.Right == nil {
				// if the current node has no left or right child, this ends up being nil as per the first line of this function
				return node.Left()
			} else {
				// if the current node has 2 childre, we delte the current node by calling the lift function. which changes the current node's
				// value to the value of it's successor node
				node.SetRight(lift(node.Right(), node))
				return node
			}
		}
		return nil
	}

	delete(value, bst.root)
	return nil
}

func (bst *BinarySearchTree[T]) Size() int {
	if bst == nil {
		return 0
	} else {
		return 1 + bst.root.Left().Size() + bst.root.Right().Size()
	}
}

// getValues is a helper function that returns a slice of all the values in the tree
func getValues[T types.Comparable](node *binary.BinaryTreeNode[T]) []T {
	var result []T
	if node.Left != nil {
		result = append(getValues(node.Left()), result...)
	}
	result = append(result, node.Data)
	if node.Right != nil {
		result = append(result, getValues(node.Right())...)
	}
	return result
}

// Values returns a slice of all the values in the tree
func (bst *BinarySearchTree[T]) Values() ([]T, error) {
	var result []T
	if bst == nil || bst.root == nil {
		return result, errors.New("tree is empty")
	}

	root := bst.root
	result = getValues(root)

	return result, nil
}

// calculateDepth helper function to calculate the depth of a tree
func calculateDepth[T types.Comparable](node *binary.BinaryTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + utils.Max(calculateDepth(node.Left()), calculateDepth(node.Right()))
}

func (bst *BinarySearchTree[T]) Depth() int {
	if bst == nil || bst.root == nil {
		return 0
	}
	return calculateDepth(bst.root)
}

// Serialize returns a string representation of the binary search tree
func (bst *BinarySearchTree[T]) Serialize() []string {
	root := bst.root

	var output []string

	if root == nil {
		return output
	}

	var values []any
	var preOrderTraversal func(node *binary.BinaryTreeNode[T])
	preOrderTraversal = func(node *binary.BinaryTreeNode[T]) {
		if node == nil {
			return
		}

		values = append(values, node.Data)
		preOrderTraversal(node.Left())
		preOrderTraversal(node.Right())
	}

	preOrderTraversal(root)

	for _, value := range values {
		switch value.(type) {
		case int:
			v := strconv.Itoa(value.(int))
			output = append(output, v)
		case string:
			output = append(output, value.(string))
		}
	}

	return output
}

// func buildTree[T types.Comparable](nodeData []T, min, max int32) *binary.BinaryTreeNode[T] {
// 	startIndex := 0
// 	if startIndex == len(nodeData) {
// 		return nil
// 	}

// 	val := nodeData[startIndex]
// 	if val < min || val > max {
// 		return nil
// 	}

// 	node := binary.NewBinaryTreeNode(val)
// 	startIndex++
// 	node.Left = buildTree(nodeData, min, val)
// 	node.Right = buildTree(nodeData, val, max)

// 	return node
// }

// TODO: deserialize and serialize binary search tree
// Deserialize deserializes a string representation of a binary search tree
// func (bst *BinarySearchTree[T]) Deserialize(data string) {
// 	if len(data) == 0 {
// 		return
// 	}

// 	var nodeValues []T
// 	values := strings.Fields(data)
// 	for _, value := range values {
// 		nodeValues = append(nodeValues, value)
// 	}

// 	var buildTree func(nodeData []T, min, max int) *binary.BinaryTreeNode[T]
// 	startIndex := 0
// 	buildTree = func(nodeData []T, min, max int) *binary.BinaryTreeNode[T] {
// 		if startIndex == len(nodeData) {
// 			return nil
// 		}

// 		data := nodeData[startIndex]

// 		if data < min || data > max {
// 			return nil
// 		}

// 		node := binary.NewBinaryTreeNode(data)
// 		startIndex++
// 		node.Left = buildTree(nodeData, min, data)
// 		node.Right = buildTree(nodeData, data, max)

// 		return node
// 	}

// 	root := buildTree(nodeValues, math.MinInt, math.MaxInt)
// 	bst.root = root
// }

// FIXME: check validity of a binary search tree
// IsValid checks if a binary search tree is valid
// func (bst *BinarySearchTree[T]) IsValid() bool {
// 	type Data struct {
// 		lowerBound T
// 		node       *binary.BinaryTreeNode[T]
// 		upperBound T
// 	}

// 	if bst.root == nil {
// 		return true
// 	}

// 	stack := []Data{}
// 	stack = append(stack, Data{lowerBound: math.Inf(-1), node: bst.root, upperBound: math.Inf(1)})

// 	for len(stack) != 0 {
// 		data := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		node := data.node
// 		lowerBound := data.lowerBound
// 		upperBound := data.upperBound

// 		if node == nil {
// 			continue
// 		}

// 		if node.Data <= lowerBound || node.Data >= upperBound {
// 			return false
// 		}

// 		if node.Left != nil {
// 			stack = append(stack, Data{lowerBound: lowerBound, node: node.Left, upperBound: node.Data})
// 		}

// 		if node.Right != nil {
// 			stack = append(stack, Data{lowerBound: node.Data, node: node.Right, upperBound: upperBound})
// 		}
// 	}

// 	return true
// }
