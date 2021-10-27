package linkedlist

import (
	"errors"
	"fmt"
	"gopherland/datastructures/linkedlist"
)

// iSinglyLinkedList is an interface for singly linked list
type iSinglyLinkedList interface {
	linkedlist.ILinkedList
}

// SinglyLinkedListNode of a Singly Linked List
type SinglyLinkedListNode struct {
	Data interface{}
	Next *SinglyLinkedListNode
}

// SinglyLinkedList data structure
type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
}

// NewSinglyLinkedListNode creates a new LinkedList Node
func NewSinglyLinkedListNode(data interface{}) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{
		Data: data,
		Next: nil,
	}
}

// NewSinglyLinkedListNode creates a new SinglyLinkedList
func NewSinglyLinkedList() iSinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l SinglyLinkedList) String() string {
	return fmt.Sprintf("%s", l.Head.Data)
}

func (ll *SinglyLinkedList) DeleteNode(node interface{}) {}

func (ll *SinglyLinkedList) DeleteNodeByData(data interface{}) {}

func (ll *SinglyLinkedList) DeleteTail() {}

func (ll *SinglyLinkedList) SwapNodesAtKthAndKPlusOne(k int) {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	current := ll.Head
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
	ll.Head = next
}

// Prepend adds a new node to the beggining of the list
func (sll *SinglyLinkedList) Prepend(val interface{}) {
	node := NewSinglyLinkedListNode(val)
	node.Next = sll.Head
	sll.Head = node
}

// Append adds a new Node at the end of a linked list
func (sll *SinglyLinkedList) Append(val interface{}) {
	node := NewSinglyLinkedListNode(val)

	if sll.Head == nil {
		sll.Head = node
		return
	}

	current := sll.Head
	for ; current.Next != nil; current = current.Next {
	}

	current.Next = node
}

func (sll *SinglyLinkedList) DeleteHead() {

}

// DeleteAtBeg removes a node at the beggining of a linked list & returns its value.
// Returns -1 if the list is empty
func (sll *SinglyLinkedList) DeleteAtBeg() interface{} {
	if sll.Head == nil {
		return -1
	}

	current := sll.Head
	sll.Head = current.Next

	return current.Data
}

// DeleteAtEnd removes a node at the end of a linked list & returns its value.
// Returns -1 if the list is empty
func (sll *SinglyLinkedList) DeleteAtEnd() interface{} {
	if sll.Head == nil {
		return -1
	}

	current := sll.Head
	for ; current.Next.Next != nil; current = current.Next {
	}

	data := current.Next.Data
	current.Next = nil
	return data
}

// Count counts the number of nodes in a Linked List
func (sll *SinglyLinkedList) Count() (count int) {
	current := sll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (sll *SinglyLinkedList) Display() {
	for cur := sll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome
func (sll *SinglyLinkedList) IsPalindrome() bool {
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
func (sll *SinglyLinkedList) DetectCycle() *SinglyLinkedListNode {
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
func (sll *SinglyLinkedList) GetNthNode(position int) (n *SinglyLinkedListNode, err error) {
	if position < 0 {
		return nil, errors.New("Position less than 0")
	}

	if sll.Head == nil {
		return nil, errors.New("List is empty")
	}

	current := sll.Head

	for current != nil {
		for x := 0; x < position; x++ {
			current = current.Next

			if current == nil {
				return nil, errors.New("Null node encountered")
			}
		}
	}

	return current, nil
}

// DeleteNodeAtPosition deletes a node at the specified position and returns the deleted node
func (sll *SinglyLinkedList) DeleteNodeAtPosition(position int) (*SinglyLinkedListNode, error) {
	if position < 0 {
		errMessage := fmt.Sprintf("Invalid Index position given. Index is {%d}, expected position >= 0", position)
		return nil, errors.New(errMessage)
	}

	// Nothing to delete here
	if sll.Head == nil {
		return nil, nil
	}

	if sll.Head != nil && position == 0 {
		current := sll.Head
		sll.Head = nil
		return current, nil
	}

	current := sll.Head

	for current != nil {
		for i := 0; i < position; i++ {
			current = current.Next

			if current == nil {
				return nil, fmt.Errorf("Invalid position %d specified, reached end of list.", position)
			}
		}
	}

	node := current
	current.Data = current.Next.Data
	current.Next = current.Next.Next

	return node, nil
}

func (sll *SinglyLinkedList) DeleteNodesByData(data interface{}) *SinglyLinkedListNode {
	dummyHead := &SinglyLinkedListNode{-1, nil}
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

// RemoveDuplicates removes duplicates from a SinglyLinkedList
// This assumes the linked list is sorted in ascending order
func (sll *SinglyLinkedList) RemoveDuplicates() *SinglyLinkedListNode {
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
func (sll *SinglyLinkedList) PairwiseSwap() *SinglyLinkedListNode {
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

// SwapNodes swaps head of the linked list after swapping the values of the kth node from the beginning and the kth node
// from the end (the list is 1-indexed).
// E.g Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
// Output: [7,9,6,6,8,7,3,0,9,5]
func (sll *SinglyLinkedList) SwapNodesAtKthAndKPlus1(k int) *SinglyLinkedList {
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
func (sll *SinglyLinkedList) SwapNodes(dataOne, dataTwo interface{}) {
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
func (sll *SinglyLinkedList) Reverse() {
	if sll.Head == nil {
		return
	}

	var prev, Next *SinglyLinkedListNode

	var current = sll.Head

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	sll.Head = prev
}

func (sll *SinglyLinkedList) DeleteAtPosition(position int) {
	if sll.Head == nil {
		return
	}

	var prev, Next *SinglyLinkedListNode

	var current = sll.Head

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	sll.Head = prev
}
