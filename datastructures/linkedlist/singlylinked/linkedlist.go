package singlylinkedlist

import (
	"errors"
	"fmt"
	"gopherland/datastructures/linkedlist"
)

// LinkedList data structure represents a singly linked list.
type LinkedList struct {
	// Head of the list
	Head *linkedlist.Node
	// size of the list
	size int
}

// NewLinkedList creates a new LinkedList
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new Node at the end of a linked list
func (sll *LinkedList) Append(val interface{}) {
	node := linkedlist.NewNode(val)

	if sll.Head == nil {
		sll.Head = node
		return
	}

	current := sll.Head
	for ; current.Next != nil; current = current.Next {
	}

	current.Next = node
}

// Prepend adds a new node to the beginning of the list
func (sll *LinkedList) Prepend(val interface{}) {
	node := linkedlist.NewNode(val)
	node.Next = sll.Head
	sll.Head = node
}

// Pop removes a node at the end of a linked list & returns its value.
func (sll *LinkedList) Pop() (interface{}, error) {
	if sll.Head == nil {
		return nil, linkedlist.ErrEmptyList
	}

	current := sll.Head

	if current.Next == nil {
		data := current.Data
		current = nil
		return data, nil
	}

	for ; current.Next.Next != nil; current = current.Next {
	}

	data := current.Next.Data
	current.Next = nil
	return data, nil
}

func (sll *LinkedList) Array() []interface{} {
	var values []interface{}
	current := sll.Head
	for ; current != nil; current = current.Next {
		values = append(values, current.Data)
	}
	return values
}

func (sll LinkedList) String() string {
	return fmt.Sprintf("%s", sll.Head.Data)
}

func (sll *LinkedList) DeleteNode(node interface{}) {}

func (sll *LinkedList) DeleteNodeByData(data interface{}) {}

func (sll *LinkedList) DeleteTail() (interface{}, error) {
	panic("implement me")
}

func (sll *LinkedList) SwapNodesAtKthAndKPlusOne(k int) {
	if sll.Head == nil || sll.Head.Next == nil {
		return
	}

	current := sll.Head
	next := current.Next

	for i := 0; i < k; i++ {
		current = current.Next
		next = current.Next
	}

	if next == nil {
		return
	}

	current.Next = next.Next
	next.Next = current
	sll.Head = next
}

// DeleteAtBeg removes a node at the beggining of a linked list & returns its value.
// Returns -1 if the list is empty
func (sll *LinkedList) DeleteAtBeg() interface{} {
	if sll.Head == nil {
		return -1
	}

	current := sll.Head
	sll.Head = current.Next

	return current.Data
}

// Count counts the number of nodes in a Linked List
func (sll *LinkedList) Count() (count int) {
	current := sll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (sll *LinkedList) Display() {
	for cur := sll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome checks if a linked list is a palindrome
func (sll *LinkedList) IsPalindrome() bool {
	if sll.Head == nil {
		return false
	}

	if sll.Head.Next == nil {
		return true
	}

	current := sll.Head
	stack := []interface{}{}

	for current != nil {
		stack = append(stack, current.Data)
		current = current.Next
	}

	current = sll.Head

	for current != nil {
		data := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Data != data {
			return false
		}
		current = current.Next
	}

	return true
}

// DetectCycle detects cycles in a LinkedList & returns the Node that contains a cycle
func (sll *LinkedList) DetectCycle() *linkedlist.Node {
	if sll.Head == nil || sll.Head.Next == nil {
		return nil
	}

	slowPointer := sll.Head
	fastPointer := sll.Head

	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			break
		}
	}

	if fastPointer == nil || fastPointer.Next == nil {
		return nil
	}

	curr := sll.Head
	for curr != slowPointer {
		curr = curr.Next
		slowPointer = slowPointer.Next
	}

	return curr
}

// GetNthNode gets the nth node in a linked list
func (sll *LinkedList) GetNthNode(position int) (n *linkedlist.Node, err error) {
	if position < 0 {
		return nil, errors.New("position less than 0")
	}

	if sll.Head == nil {
		return nil, linkedlist.ErrEmptyList
	}

	current := sll.Head

	for i := 0; i < position; i++ {
		current = current.Next
	}

	if current == nil {
		return nil, linkedlist.ErrInvalidIndex
	}

	return current, nil
}

func (sll *LinkedList) DeleteNodesByData(data interface{}) *linkedlist.Node {
	dummyHead := &linkedlist.Node{Data: -1}
	current := dummyHead

	for current.Next != nil {

		if current.Next.Data == data {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummyHead.Next
}

// RemoveDuplicates removes duplicates from a LinkedList
// This assumes the linked list is sorted in ascending order
func (sll *LinkedList) RemoveDuplicates() *linkedlist.Node {
	head := sll.Head

	if head == nil || head.Next == nil {
		return head
	}

	current := head
	next := current.Next

	for next != nil {
		if next.Data == current.Data {
			current.Next = current.Next.Next
			next = current.Next
		} else {
			current = next
			next = current.Next
		}
	}
	return head
}

// PairwiseSwap swaps each pair in a linked list
// Swaps nodes in a linked list in pairs.
// The premise(idea) is to swap the data of each node with the data of the next node. This is while using
// an iterative approach
// Example:
// 1 -> 2 -> 3 -> 4
// becomes
// 2 -> 1 -> 4 -> 3
func (sll *LinkedList) PairwiseSwap() *linkedlist.Node {
	head := sll.Head

	// Nothing to do here
	if head == nil {
		return head
	}

	current := head

	// traverse only if there are at least 2 nodes left
	for current != nil && current.Next != nil {

		// if the 2 nodes have the same data value, no need to swap
		if current.Data == current.Next.Data {
			// move to the next pair
			current = current.Next.Next
		} else {
			temp := current.Data
			current.Data = current.Next.Data
			current.Next.Data = temp

			// move to next pair
			current = current.Next.Next
		}
	}

	return head
}

// SwapNodesAtKthAndKPlus1 SwapNodes swaps head of the linked list after swapping the values of the kth node from the beginning and the kth node
// from the end (the list is 1-indexed).
// E.g Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
// Output: [7,9,6,6,8,7,3,0,9,5]
func (sll *LinkedList) SwapNodesAtKthAndKPlus1(k int) *LinkedList {
	a, b := sll.Head, sll.Head

	for i := 1; i < k; i++ {
		a = a.Next
	}

	node := a
	a = a.Next

	for a != nil {
		a, b = a.Next, b.Next
	}

	node.Data, b.Data = b.Data, node.Data

	return sll
}

// SwapNodes swaps the data items of 2 nodes in the Linkedlist
// Swaps two nodes based on the data they contain. We search through the LinkedList looking for the data item in
// each node. Once the first is found, we keep track of it and move on until we find the next data item. Once that
// is found, we swap the two nodes' data items.
// If we can't find the first data item nor the second. No need to perform swap. If the 2 data items are similar
// no need to perform swap as well.
// If the LinkedList is empty (i.e. has no head node), return, no need to swap when we have no LinkedList :)
func (sll *LinkedList) SwapNodes(dataOne, dataTwo interface{}) {
	if sll.Head == nil {
		return
	}

	if dataOne == dataTwo {
		return
	}

	currentOne := sll.Head
	currentTwo := sll.Head

	for currentOne != nil && currentOne.Data != dataOne {
		currentOne = currentOne.Next
	}

	for currentTwo != nil && currentTwo.Data != dataTwo {
		currentTwo = currentTwo.Next
	}

	if currentOne == nil || currentTwo == nil {
		return
	}

	currentOne.Data, currentTwo.Data = currentTwo.Data, currentOne.Data
}

// Reverse a LinkedList. Making the head the tail and the tail the head
func (sll *LinkedList) Reverse() {
	if sll.Head == nil {
		return
	}

	var prev, Next *linkedlist.Node

	var current = sll.Head

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	sll.Head = prev
}

func (sll *LinkedList) DeleteAtPosition(position int) (*linkedlist.Node, error) {
	if sll.Head == nil {
		return nil, linkedlist.ErrEmptyList
	}

	if position == 0 {
		sll.Head = sll.Head.Next
		return sll.Head, nil
	}

	if position > sll.Length() {
		return nil, linkedlist.ErrInvalidIndex
	}

	if sll.Head.Next == nil {
		return nil, linkedlist.ErrInvalidIndex
	}

	index := 0
	current, prev, next := sll.Head, sll.Head, sll.Head

	for current != nil {
		if index == position {
			next = current.Next
			break
		}
		prev = current
		current = current.Next
		index++
	}

	if current == nil {
		return nil, linkedlist.ErrInvalidIndex
	}

	prev.Next = next
	return current, nil
}

// AddAtPosition adds a node at the specified position in the LinkedList.
// if position is greater than the length of the LinkedList, the node should not be added
// if the position is equal to the length of the LinkedList, the node should be added at the end of the LinkedList
func (sll *LinkedList) AddAtPosition(position int, data interface{}) {
	node := linkedlist.NewNode(data)

	if sll.Head == nil {
		sll.Head = node
		return
	}

	if position == 0 {
		sll.Prepend(data)
		return
	}

	if position > sll.Length() {
		return
	}

	if position == sll.Length() {
		sll.Append(data)
		return
	}

	if sll.Head.Next == nil {
		sll.Head.Next = node
		return
	}

	var current = sll.Head
	var prev, next *linkedlist.Node

	for current != nil && position != 0 {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		position--
	}

	prev.Next = node
	node.Next = current

	sll.Head = prev
}

func (sll LinkedList) Length() int {
	if sll.Head == nil {
		return 0
	}

	var current = sll.Head
	var count int

	for current != nil {
		current = current.Next
		count++
	}

	return count
}