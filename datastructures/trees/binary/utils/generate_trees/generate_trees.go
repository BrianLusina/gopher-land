package utils

import (
	"gopherland/datastructures/trees"
	"gopherland/datastructures/trees/binary"
)

// generator generates all possible Binary Trees given a range of numbers
func generator[T int](start, end int) []*binary.BinaryTreeNode[int] {
	var bsts []*binary.BinaryTreeNode[int]

	if start > end {
		bsts = append(bsts, nil)
		return bsts
	}

	if start == end {
		node := &binary.BinaryTreeNode[int]{
			TreeNode: *trees.NewTreeNode(start),
		}
		bsts = append(bsts, node)
		return bsts
	}

	for i := start; i < end+1; i++ {
		leftSubtree := generator(start, i-1)
		rightSubtree := generator(i+1, end)

		for j := range leftSubtree {
			left := leftSubtree[j]

			for k := range rightSubtree {
				right := rightSubtree[k]
				node := &binary.BinaryTreeNode[int]{
					TreeNode: *trees.NewTreeNode(i),
				}
				node.Left = left
				node.Right = right
				bsts = append(bsts, node)
			}
		}
	}
	return bsts
}

// GenerateTrees Generates n Unique Binary Search Trees given a number n
func GenerateTrees[T int](n int) []*binary.BinaryTreeNode[int] {
	return generator[int](1, n)
}
