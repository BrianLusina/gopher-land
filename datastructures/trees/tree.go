// Package trees contains types and struct forTrees
package trees

// Tree interface to operate on trees
type Tree interface {
	NewNode()
}

// TreeNode represent a Node in a Tree
type TreeNode struct {
	Left *TreeNode
	// Data, for now ints, until generics are out in Golang
	// Reference https://blog.golang.org/generics-next-step
	Data  int
	Right *TreeNode
}

// NewTree creates a new Tree Node
func (t *TreeNode) NewTree(data int) TreeNode {
	return TreeNode{
		Data: data,
	}
}

// NewNode returns a new TreeNode
func (t *TreeNode) NewNode(data int) TreeNode {
	return TreeNode{
		Data: data,
	}
}
