package singlylinkedlist

import "gopherland/datastructures/list"

// mergeTwoLists merges 2 sorted linked lists & returns one sorted linked list
func mergeTwoSortedLists[T ~int](nodeOne, nodeTwo *list.Node[T]) *list.Node[T] {
	if nodeOne == nil {
		return nodeTwo
	}
	if nodeTwo == nil {
		return nodeOne
	}

	var temp *list.Node[T]

	if nodeOne.Data < nodeTwo.Data {
		temp = nodeOne
		nodeOne = nodeOne.Next
	} else {
		temp = nodeTwo
		nodeTwo = nodeTwo.Next
	}

	current := temp

	for nodeOne != nil && nodeTwo != nil {
		if nodeOne.Data > nodeTwo.Data {
			current.Next = nodeTwo
			nodeTwo = nodeTwo.Next
		} else {
			current.Next = nodeOne
			nodeOne = nodeOne.Next
		}
		current = current.Next
	}

	if nodeOne != nil {
		current.Next = nodeOne
	}

	if nodeTwo != nil {
		current.Next = nodeTwo
	}

	return temp
}
