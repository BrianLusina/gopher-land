package findclosestvalue

import (
	"gopherland/datastructures/trees/binary"
	"gopherland/math/utils"
)

func findClosestValueInBst(node *binary.BinaryTreeNode[int], target int) int {
	if node == nil {
		return 0
	}
	if node.Data == target {
		return node.Data
	}

	closestValue := node.Data
	minDiff := utils.AbsDiff(target, node.Data)
	current := node

	for current != nil {
		currentDiff := utils.AbsDiff(target, current.Data)

		if currentDiff < minDiff {
			minDiff = currentDiff
			closestValue = current.Data
		}

		if current.Data == target {
			return current.Data
		}

		if target < current.Data {
			current = current.Left()
		} else {
			current = current.Right()
		}
	}

	return closestValue
}
