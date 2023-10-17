package ternary

import (
	"fmt"
	"gopherland/pkg/types"
	"gopherland/pkg/utils"
	"strings"
)

type (
	// TernaryTree represents a ternary tree data structure
	TenaryTree[T types.Comparable] struct {
		// root is the root of the tree
		root *TernaryTreeNode[T]
	}

	// TenaryTreeOption represents an option that can be applied to the tree
	TenaryTreeOption[T types.Comparable] func(*TenaryTree[T])
)

// NewTernaryTree creates a new ternary tree
func NewTernaryTree[T types.Comparable](opts ...TenaryTreeOption[T]) *TenaryTree[T] {
	tree := &TenaryTree[T]{}

	for _, opt := range opts {
		opt(tree)
	}

	return tree
}

// Root sets the root of the tree
func Root[T types.Comparable](node *TernaryTreeNode[T]) TenaryTreeOption[T] {
	return func(tree *TenaryTree[T]) {
		tree.root = node
	}
}

// Paths return all the paths of the Tree from root node to leaf node
func (t *TenaryTree[T]) Paths() []string {
	var dfs func(*TernaryTreeNode[T], []string, *[]string)

	dfs = func(root *TernaryTreeNode[T], paths []string, result *[]string) {
		if utils.All(root.children, func(data *TernaryTreeNode[T]) bool { return data == nil }) {
			*result = append(*result, strings.Join(paths, "->")+"->"+fmt.Sprintf("%v", root.data))
			return
		}

		for _, child := range root.children {
			if child != nil {
				paths = append(paths, fmt.Sprintf("%v", root.data))
				dfs(child, paths, result)
				paths = paths[:len(paths)-1]
			}
		}
	}

	res := []string{}
	if t.root != nil {
		dfs(t.root, []string{}, &res)
	}
	return res
}
