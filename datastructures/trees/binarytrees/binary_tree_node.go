package binarytrees

import (
	"fmt"
	"gopherland/datastructures/trees"
	"strings"
)

// BinaryTreeNode represent a Binary Tree Node in a Binary Tree
// when performing comparisons, it is important to cast this interface to a concrete type
type BinaryTreeNode struct {
	trees.TreeNode
	Left, Right *BinaryTreeNode
}

// NewBinaryTreeNode creates a new BinaryTreeNode
func NewBinaryTreeNode(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{
		TreeNode: trees.NewTreeNode(data),
		Left:     nil,
		Right:    nil,
	}
}

func (node *BinaryTreeNode) Serialize() string {
	if node == nil {
		return ""
	}

	stack := []*BinaryTreeNode{node}
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

func add(c *[]any, node *BinaryTreeNode) {
	if node == nil {
		*c = append(*c, "nil")
	} else {
		*c = append(*c, node.Data)
	}
}
