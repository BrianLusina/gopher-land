package binarysearchtree

import (
	"errors"
	"gopherland/datastructures/trees/binarytrees"
)

// BinarySearchTree is a binary tree that satisfies the binary tree interface
type BinarySearchTree struct {
	root *binarytrees.BinaryTreeNode
}

// NewBinarySearchTree returns a new binary search tree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert inserts a new node with the value into the binary search tree
func (bst *BinarySearchTree) Insert(value any) {
	// we don't have a root
	if bst.root == nil {
		bst.root = binarytrees.NewBinaryTreeNode(value)
		return
	}

	rootNode := bst.root

	if value.(int) <= rootNode.Data.(int) {
		if rootNode.Left == nil {
			newNode := binarytrees.NewBinaryTreeNode(value)
			rootNode.Left = newNode
			return
		} else {
			bst.Insert(value)
		}
	} else {
		if rootNode.Right == nil {
			newNode := binarytrees.NewBinaryTreeNode(value)
			rootNode.Right = newNode
			return
		} else {
			bst.Insert(value)
		}
	}
}

func (bst *BinarySearchTree) Size() int {
	if bst == nil {
		return 0
	} else {
		return 1 + bst.root.Left.Size() + bst.root.Right.Size()
	}
}

// getValues is a helper function that returns a slice of all the values in the tree
func getValues(node *binarytrees.BinaryTreeNode) []any {
	var result []interface{}
	if node.Left != nil {
		result = append(getValues(node.Left), result...)
	}
	result = append(result, node.Data)
	if node.Right != nil {
		result = append(result, getValues(node.Right)...)
	}
	return result
}

// Values returns a slice of all the values in the tree
func (bst *BinarySearchTree) Values() ([]interface{}, error) {
	var result []interface{}
	if bst == nil || bst.root == nil {
		return result, errors.New("tree is empty")
	}

	root := bst.root
	result = getValues(root)

	return result, nil
}
