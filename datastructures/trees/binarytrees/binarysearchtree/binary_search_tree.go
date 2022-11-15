package binarysearchtree

import (
	"errors"
	"gopherland/datastructures/trees/binarytrees"
	"math"
	"strconv"
	"strings"
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

// Serialize returns a string representation of the binary search tree
func (bst *BinarySearchTree) Serialize() []string {
	root := bst.root

	var output []string

	if root == nil {
		return output
	}

	var values []any
	var preOrderTraversal func(node *binarytrees.BinaryTreeNode)
	preOrderTraversal = func(node *binarytrees.BinaryTreeNode) {
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

func buildTree(nodeData []any, min, max int32) *binarytrees.BinaryTreeNode {
	startIndex := 0
	if startIndex == len(nodeData) {
		return nil
	}

	val := nodeData[startIndex].(int32)
	if val < min || val > max {
		return nil
	}

	node := binarytrees.NewBinaryTreeNode(val)
	startIndex++
	node.Left = buildTree(nodeData, min, val)
	node.Right = buildTree(nodeData, val, max)

	return node
}

func castToType(data any) interface{} {
	switch data.(type) {
	case string:
		i, err := strconv.Atoi(data.(string))
		if err == nil {
			return i
		}
		return nil
	case int:
		return data.(int)
	case int32:
		return data.(int32)
	case int64:
		return data.(int64)
	case float32:
		return data.(float32)
	case float64:
		return data.(float64)
	default:
		return data
	}
}

// Deserialize deserializes a string representation of a binary search tree
func (bst *BinarySearchTree) Deserialize(data string) {
	if len(data) == 0 {
		return
	}

	var nodeValues []any
	values := strings.Fields(data)
	for _, value := range values {
		nodeValues = append(nodeValues, value)
	}

	var buildTree func(nodeData []any, min, max int) *binarytrees.BinaryTreeNode
	startIndex := 0
	buildTree = func(nodeData []any, min, max int) *binarytrees.BinaryTreeNode {
		if startIndex == len(nodeData) {
			return nil
		}

		data := nodeData[startIndex]
		val := castToType(data)

		if val.(int) < min || val.(int) > max {
			return nil
		}

		node := binarytrees.NewBinaryTreeNode(val)
		startIndex++
		node.Left = buildTree(nodeData, min, val.(int))
		node.Right = buildTree(nodeData, val.(int), max)

		return node
	}

	root := buildTree(nodeValues, math.MinInt, math.MaxInt)
	bst.root = root
}

func (bst *BinarySearchTree) IsValid() bool {
	type Data struct {
		lowerBound int
		node       *binarytrees.BinaryTreeNode
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

		if node.Data.(int) <= lowerBound || node.Data.(int) >= upperBound {
			return false
		}

		if node.Left != nil {
			stack = append(stack, Data{lowerBound: lowerBound, node: node.Left, upperBound: node.Data.(int)})
		}

		if node.Right != nil {
			stack = append(stack, Data{lowerBound: node.Data.(int), node: node.Right, upperBound: upperBound})
		}
	}

	return true
}
