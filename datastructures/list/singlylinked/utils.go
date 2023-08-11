package singlylinkedlist

import (
	"gopherland/datastructures/list"
	"gopherland/pkg/types"
	"gopherland/pkg/utils"
	pkgUtils "gopherland/pkg/utils"
)

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

// MaximumPairSum returns the maximum twin sum of a node and its twin, where a node's twin is at the index (n-1-i) where n is the
// number of nodes in the linked list.
// For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only
// nodes with twins for n = 4.
func MaximumPairSum[T types.Comparable](head *list.Node[T]) T {
	if head == nil {
		return utils.Zero[T]()
	}

	current := head
	values := []T{}

	for current != nil {
		values = append(values, current.Data)
		current = current.Next
	}

	var maximumSum T
	left := 0
	right := len(values) - 1

	for left < right {
		p := values[left] + values[right]
		maximumSum = pkgUtils.Max(maximumSum, p)
		left += 1
		right -= 1
	}

	return maximumSum
}

func MaximumPairSumStack[T types.Comparable](head *list.Node[T]) T {
	if head == nil {
		return utils.Zero[T]()
	}

	current := head
	stack := []T{}

	for current != nil {
		stack = append(stack, current.Data)
		current = current.Next
	}

	current = head
	var maximumSum T
	size := len(stack)
	count := 1

	for count < size/2 {
		topValue := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		maximumSum = pkgUtils.Max(maximumSum, current.Data+topValue)
		current = current.Next
		count++
	}

	return maximumSum
}

func MaximumPairSumReverseInPlace[T types.Comparable](head *list.Node[T]) T {
	if head == nil {
		return utils.Zero[T]()
	}

	slow := head
	fast := head

	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var nextNode *list.Node[T]
	var prev *list.Node[T]

	for slow != nil {
		nextNode = slow.Next
		slow.Next = prev
		prev = slow
		slow = nextNode
	}

	start := head
	var maximumSum T

	for prev != nil {
		maximumSum = pkgUtils.Max(maximumSum, start.Data+prev.Data)
		prev = prev.Next
		start = start.Next
	}

	return maximumSum
}
