// Package binarytrees contains types and struct for Binary Trees
package binarytrees

import (
	"fmt"
	"gopherland/pkg/types"
	"strings"
)

// BinaryTree represents a binary tree
type BinaryTree[T types.Comparable] struct {
	root *BinaryTreeNode[T]
}

type BinaryTreeIterator[T types.Comparable] struct {
	stack []*BinaryTreeNode[T]
}

// NewBinaryTree creates a new binary tree with root as nil
func NewBinaryTree[T types.Comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// func IsBstRecursive[T types.Comparable](node *BinaryTreeNode[T], min, max T) bool {
// 	if node == nil {
// 		return true
// 	}

// 	if node.Data < min || node.Data > max {
// 		return false
// 	}

// 	return IsBstRecursive(node.Left, min, node.Data-1) && IsBstRecursive(node.Right, node.Data+1, max)
// }

// Serialize converts this tree into string representation
func (t *BinaryTree[T]) Serialize() string {
	if t.root == nil {
		return ""
	}

	stack := []*BinaryTreeNode[T]{}
	stack = append(stack, t.root)
	nodeValues := []any{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			nodeValues = append(nodeValues, nil)
		} else {
			nodeValues = append(nodeValues, node.Data)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}

	joined := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nodeValues...)), ","), "[]")

	return joined
}
