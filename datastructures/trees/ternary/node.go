package ternary

import (
	"fmt"
	"gopherland/pkg/types"
)

// TernaryTreeNode represents a node in a Ternary Tree
type (
	TernaryTreeNode[T types.Comparable] struct {
		data     T
		children []*TernaryTreeNode[T]
	}

	// TernaryTreeNodeOption allows setting options to a Binary TreeNode
	TernaryTreeNodeOption[T types.Comparable] func(*TernaryTreeNode[T])
)

// NewTernaryTreeNode creates a new Ternary tree node
func NewTernaryTreeNode[T types.Comparable](data T, opts ...TernaryTreeNodeOption[T]) *TernaryTreeNode[T] {
	ternaryTreeNode := &TernaryTreeNode[T]{
		data: data,
	}

	for _, opt := range opts {
		opt(ternaryTreeNode)
	}
	return ternaryTreeNode
}

// Children sets the children of the node
func Children[T types.Comparable](children []*TernaryTreeNode[T]) TernaryTreeNodeOption[T] {
	return func(node *TernaryTreeNode[T]) {
		if len(node.children) == 3 {
			panic(fmt.Errorf("node has reached children capacity of 3. current length : %d", len(node.children)))
		}
		node.children = children
	}
}

// Data retrieves the data of the node
func (n *TernaryTreeNode[T]) Data() T {
	return n.data
}

// Children retrieves the children of the node
func (n *TernaryTreeNode[T]) Children() []*TernaryTreeNode[T] {
	return n.children
}

// String returns the string representation of the node
func (n *TernaryTreeNode[T]) String() string {
	return fmt.Sprintf("TernaryTreeNode{data: %v, children: %v}", n.data, n.children)
}
