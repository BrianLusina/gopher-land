package binarysearchtree

import (
	"errors"
	"gopherland/datastructures/trees/binarytrees"
	"math"
)

// BinarySearchTree is a binary tree that satisfies the binary tree interface
type BinarySearchTree struct {
	root *binarytrees.BinaryTreeNode
}

// NewBinarySearchTree returns a new binary search tree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// insert helper function to insert a new value in the binary search tree
func insert(node *binarytrees.BinaryTreeNode, value any) {
	if value.(int) <= node.Data.(int) {
		if node.Left != nil {
			insert(node.Left, value)
		} else {
			node.Left = binarytrees.NewBinaryTreeNode(value)
		}
	} else {
		if node.Right != nil {
			insert(node.Right, value)
		} else {
			node.Right = binarytrees.NewBinaryTreeNode(value)
		}
	}
}

// Insert inserts a new node with the value into the binary search tree
func (bst *BinarySearchTree) Insert(value any) {
	// we don't have a root
	if bst.root == nil {
		bst.root = binarytrees.NewBinaryTreeNode(value)
		return
	}

	rootNode := bst.root
	insert(rootNode, value)
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

// calculateDepth helper function to calculate the depth of a tree
func calculateDepth(node *binarytrees.BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(calculateDepth(node.Left)), float64(calculateDepth(node.Right))))
}

func (bst *BinarySearchTree) Depth() int {
	if bst == nil || bst.root == nil {
		return 0
	}
	return calculateDepth(bst.root)
}

//func (bst *BinarySearchTree) Serialize() String {
//	var values []interface{}
//	root := bst.root
//
//	if root != nil {
//		return String{}
//	}
//
//}
