package utils

import (
	. "gopherland/datastructures/trees/binarytrees"
)

// generator generates all possible Binary Trees given a range of numbers
func generator(start, end int) []*BinaryTreeNode {
	bsts := []*BinaryTreeNode{}

	if start > end {
		bsts = append(bsts, nil)
		return bsts
	}

	if start == end {
		node := &BinaryTreeNode{
			Data: start,
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
				node := &BinaryTreeNode{
					Data: i,
				}
				node.Left = left
				node.Right = right
				bsts = append(bsts, node)
			}
		}
	}
	return bsts
}

// Generates n Unique Binary Search Trees given a number n
func GenerateTrees(n int) []*BinaryTreeNode {
	return generator(1, n)
}
