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

// sortList sorts a linked list in ascending order given the head node
func sortList[T ~int](head *list.Node[T]) *list.Node[T] {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddleNode(head)
	left := sortList(head)
	right := sortList(mid)
	return mergeSortLists(left, right)
}

// getMiddleNode retrieves the middle node of a single linked list given the head node
func getMiddleNode[T comparable](node *list.Node[T]) *list.Node[T] {
	if node == nil {
		return nil
	}

	slow := node
	fast := slow.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	head := slow.Next
	slow.Next = nil
	return head
}

// mergeSortLists merges 2 unsorted linked lists into 1 sorted list
func mergeSortLists[T ~int](nodeOne, nodeTwo *list.Node[T]) *list.Node[T] {
	dummyHead := list.NewNode[T](0)
	tail := dummyHead

	for nodeOne != nil && nodeTwo != nil {
		if nodeOne.Data < nodeTwo.Data {
			tail.Next = nodeOne
			nodeOne = nodeOne.Next
		} else {
			tail.Next = nodeTwo
			nodeTwo = nodeTwo.Next
		}
		tail = tail.Next
	}

	if nodeOne != nil {
		tail.Next = nodeOne
	} else {
		tail.Next = nodeTwo
	}
	return dummyHead.Next
}
