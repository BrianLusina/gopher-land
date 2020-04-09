// Package binarytrees contains types and struct for Binary Trees
package binarytrees

// Tree represent a Node in a Binary Tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) NewTree(value int) Tree {
	return Tree{
		Value: value,
	}
}
