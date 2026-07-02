package threadedbinary

import (
	"gopherland/datastructures/trees/binary"
	"gopherland/datastructures/trees/binary/bst"
	"gopherland/pkg/types"
)

type ThreadedBinarySearchTree[T types.Comparable] struct {
	bst.BinarySearchTree[T]
}

func NewThreadedBinarySearchTree[T types.Comparable]() *ThreadedBinarySearchTree[T] {
	return &ThreadedBinarySearchTree[T]{
		BinarySearchTree: *bst.NewBinarySearchTree[T](),
	}
}

func (tree *ThreadedBinarySearchTree[T]) Insert(data T) error {
	if tree.BinarySearchTree.Root() == nil {
		tree.SetRoot(binary.NewBinaryTreeNode(data))
		return nil
	}

	current := tree.Root()
	for {
		switch {
		case data == current.Data:
			return nil // no duplicates

		case data < current.Data:
			if !current.LeftThread() && current.Left() != nil {
				current = current.Left()
				continue
			}
			newNode := binary.NewBinaryTreeNode(data)
			newNode.SetLeft(current.Left()) // inherit cur's predecessor
			newNode.SetLeftThread(true)
			newNode.SetRight(current) // cur is the new successor
			newNode.SetRightThread(true)
			current.SetLeft(newNode)
			current.SetLeftThread(false)
			return nil

		default: // value > cur.Value
			if !current.RightThread() && current.Right() != nil {
				current = current.Right()
				continue
			}
			newNode := binary.NewBinaryTreeNode(data)
			newNode.SetRight(current.Right()) // inherit cur's successor
			newNode.SetRightThread(true)
			newNode.SetLeft(current) // cur is the new predecessor
			newNode.SetLeftThread(true)
			current.SetRight(newNode)
			current.SetRightThread(false)
			return nil
		}
	}
}

func (tree *ThreadedBinarySearchTree[T]) leftmost(node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
	if node == nil {
		return nil
	}
	for !node.LeftThread() && node.Left() != nil {
		node = node.Left()
	}
	return node
}

func (tree *ThreadedBinarySearchTree[T]) rightmost(node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
	if node == nil {
		return nil
	}
	for !node.RightThread() && node.Right() != nil {
		node = node.Right()
	}
	return node
}

func (tree *ThreadedBinarySearchTree[T]) successor(node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
	if node.RightThread() {
		return node.Right()
	}
	return tree.leftmost(node.Right())
}

func (tree *ThreadedBinarySearchTree[T]) predecessor(node *binary.BinaryTreeNode[T]) *binary.BinaryTreeNode[T] {
	if node.LeftThread() {
		return node.Left()
	}
	return tree.rightmost(node.Left())
}

// InorderTraversal does an iterative inorder traversal using threads only - no stack needed.
func (tree *ThreadedBinarySearchTree[T]) InorderTraversal() []T {
	result := []T{}
	cur := tree.leftmost(tree.Root())
	for cur != nil {
		result = append(result, cur.Data)
		cur = tree.successor(cur)
	}
	return result
}

// ReverseInorderTraversal walks predecessor threads right-to-left.
func (tree *ThreadedBinarySearchTree[T]) ReverseInorderTraversal() []T {
	result := []T{}
	cur := tree.rightmost(tree.Root())
	for cur != nil {
		result = append(result, cur.Data)
		cur = tree.predecessor(cur)
	}
	return result
}
