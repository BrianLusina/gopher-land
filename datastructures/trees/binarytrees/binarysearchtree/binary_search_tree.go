package binarysearchtree

import (
	"errors"
	"gopherland/datastructures/trees/binarytrees"
	"gopherland/math/utils"
	"gopherland/pkg/types"
	"math"
	"strconv"
	"strings"
)

// BinarySearchTree is a binary tree that satisfies the binary tree interface
type BinarySearchTree[T types.Comparable] struct {
	root *binarytrees.BinaryTreeNode[T]
}

// NewBinarySearchTree returns a new binary search tree
func NewBinarySearchTree[T types.Comparable]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

// insert helper function to insert a new value in the binary search tree
func insert[T types.Comparable](node *binarytrees.BinaryTreeNode[T], value T) {
	if value <= node.Data {
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
func (bst *BinarySearchTree[T]) Insert(value T) {
	// we don't have a root
	if bst.root == nil {
		bst.root = binarytrees.NewBinaryTreeNode(value)
		return
	}

	rootNode := bst.root
	insert(rootNode, value)
}

func (bst *BinarySearchTree[T]) Size() int {
	if bst == nil {
		return 0
	} else {
		return 1 + bst.root.Left.Size() + bst.root.Right.Size()
	}
}

// getValues is a helper function that returns a slice of all the values in the tree
func getValues[T types.Comparable](node *binarytrees.BinaryTreeNode[T]) []T {
	var result []T
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
func calculateDepth[T types.Comparable](node *binarytrees.BinaryTreeNode[T]) int {
	if node == nil {
		return 0
	}
	return 1 + utils.Max(calculateDepth(node.Left), calculateDepth(node.Right))
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
	var preOrderTraversal func(node *binarytrees.BinaryTreeNode[T])
	preOrderTraversal = func(node *binarytrees.BinaryTreeNode[T]) {
		if node == nil {
			return
		}

		values = append(values, node.Data)
		preOrderTraversal(node.Left)
		preOrderTraversal(node.Right)
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

func buildTree[T types.Comparable](nodeData []T, min, max int32) *binarytrees.BinaryTreeNode[T] {
	startIndex := 0
	if startIndex == len(nodeData) {
		return nil
	}

	val := nodeData[startIndex]
	if val < min || val > max {
		return nil
	}

	node := binarytrees.NewBinaryTreeNode(val)
	startIndex++
	node.Left = buildTree(nodeData, min, val)
	node.Right = buildTree(nodeData, val, max)

	return node
}

// Deserialize deserializes a string representation of a binary search tree
func (bst *BinarySearchTree[T]) Deserialize(data string) {
	if len(data) == 0 {
		return
	}

	var nodeValues []any
	values := strings.Fields(data)
	for _, value := range values {
		nodeValues = append(nodeValues, value)
	}

	var buildTree func(nodeData []T, min, max int) *binarytrees.BinaryTreeNode[T]
	startIndex := 0
	buildTree = func(nodeData []T, min, max int) *binarytrees.BinaryTreeNode[T] {
		if startIndex == len(nodeData) {
			return nil
		}

		data := nodeData[startIndex]

		if data < min || data > max {
			return nil
		}

		node := binarytrees.NewBinaryTreeNode(data)
		startIndex++
		node.Left = buildTree(nodeData, min, data)
		node.Right = buildTree(nodeData, data, max)

		return node
	}

	root := buildTree(nodeValues, math.MinInt, math.MaxInt)
	bst.root = root
}

func (bst *BinarySearchTree[T]) IsValid() bool {
	type Data struct {
		lowerBound int
		node       *binarytrees.BinaryTreeNode[T]
		upperBound int
	}

	if bst.root == nil {
		return true
	}

	stack := []Data{}
	stack = append(stack, Data{lowerBound: int(math.Inf(-1)), node: bst.root, upperBound: int(math.Inf(1))})

	for len(stack) != 0 {
		data := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := data.node
		lowerBound := data.lowerBound
		upperBound := data.upperBound

		if node == nil {
			continue
		}

		if node.Data <= lowerBound || node.Data >= upperBound {
			return false
		}

		if node.Left != nil {
			stack = append(stack, Data{lowerBound: lowerBound, node: node.Left, upperBound: node.Data})
		}

		if node.Right != nil {
			stack = append(stack, Data{lowerBound: node.Data, node: node.Right, upperBound: upperBound})
		}
	}

	return true
}
