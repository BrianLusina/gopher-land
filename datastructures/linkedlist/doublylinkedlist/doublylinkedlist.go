package doublylinkedlist

import (
	"errors"
	"fmt"
)

// DoublyLinkedListNode of a Singly Linked List
type DoublyLinkedListNode struct {
	Data interface{}
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

// LinkedList data structure
type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
}

func (l DoublyLinkedList) Print() string {
	return fmt.Sprintf("%s", l.Head.Data)
}

func CreateList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// NewNode creates a new LinkedList Node
func NewNode(data interface{}) *DoublyLinkedListNode {
	return &DoublyLinkedListNode{data, nil, nil}
}

// Prepend adds a new node to the beggining of the list
func (ll *DoublyLinkedList) Prepend(val interface{}) {
	node := NewNode(val)
	node.Next = ll.Head
	ll.Head = node
}

// Append adds a new Node at the end of a linked list
func (ll *DoublyLinkedList) Append(val interface{}) {
	node := NewNode(val)

	if ll.Head == nil {
		ll.Head = node
		return
	}

	current := ll.Head
	for ; current.Next != nil; current = current.Next {
	}

	current.Next = node
}

// DeleteAtBeg removes a node at the beggining of a linked list & returns its value.
// Returns -1 if the list is empty
func (ll *DoublyLinkedList) DeleteAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	current := ll.Head
	ll.Head = current.Next

	return current.Data
}

// DeleteAtEnd removes a node at the end of a linked list & returns its value.
// Returns -1 if the list is empty
func (ll *DoublyLinkedList) DeleteAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}

	current := ll.Head
	for ; current.Next.Next != nil; current = current.Next {
	}

	data := current.Next.Data
	current.Next = nil
	return data
}

// Count counts the number of nodes in a Linked List
func (ll *DoublyLinkedList) Count() (count int) {
	current := ll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (ll *DoublyLinkedList) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome
func (ll *DoublyLinkedList) IsPalindrome() bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Next == nil {
		return true
	}

	current := ll.Head
	stack := []interface{}{}

	for current != nil {
		stack = append(stack, current.Data)
		current = current.Next
	}

	current = ll.Head

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
func (ll *DoublyLinkedList) DetectCycle() *DoublyLinkedListNode {
	if ll.Head == nil || ll.Head.Next == nil {
		return nil
	}

	slowPointer := ll.Head
	fastPointer := ll.Head

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

	curr := ll.Head
	for curr != slowPointer {
		curr = curr.Next
		slowPointer = slowPointer.Next
	}

	return curr
}

// GetNthNode gets the nth node in a linked list
func (ll *DoublyLinkedList) GetNthNode(position int) (n *DoublyLinkedListNode, err error) {
	if position < 0 {
		return nil, errors.New("Position less than 0")
	}

	if ll.Head == nil {
		return nil, errors.New("List is empty")
	}

	current := ll.Head

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
func (ll *DoublyLinkedList) DeleteNodeAtPosition(position int) (*DoublyLinkedListNode, error) {
	if position < 0 {
		errMessage := fmt.Sprintf("Invalid Index position given. Index is {%d}, expected position >= 0", position)
		return nil, errors.New(errMessage)
	}

	// Nothing to delete here
	if ll.Head == nil {
		return nil, nil
	}

	if ll.Head != nil && position == 0 {
		current := ll.Head
		ll.Head = nil
		return current, nil
	}

	current := ll.Head

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

// RemoveDuplicates removes duplicates from a DoublyLinkedList
// This assumes the linked list is sorted in ascending order
func (ll *DoublyLinkedList) RemoveDuplicates() *DoublyLinkedListNode {
	head := ll.Head

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
func (ll *DoublyLinkedList) PairwiseSwap() *DoublyLinkedListNode {
	head := ll.Head

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
func (ll *DoublyLinkedList) SwapNodesAtKthAndKPlus1(k int) *DoublyLinkedList {
	a, b := ll.Head, ll.Head

	for i := 1; i < k; i++ {
		a = a.Next
	}

	node := a
	a = a.Next

	for a != nil {
		a, b = a.Next, b.Next
	}

	node.Data, b.Data = b.Data, node.Data

	return ll
}

// SwapNodes swaps the data items of 2 nodes in the Linkedlist
// Swaps two nodes based on the data they contain. We search through the LinkedList looking for the data item in
// each node. Once the first is found, we keep track of it and move on until we find the next data item. Once that
// is found, we swap the two nodes' data items.
// If we can't find the first data item nor the second. No need to perform swap. If the 2 data items are similar
// no need to perform swap as well.
// If the LinkedList is empty (i.e. has no head node), return, no need to swap when we have no LinkedList :)
func (ll *DoublyLinkedList) SwapNodes(dataOne, dataTwo interface{}) {
	if ll.Head == nil {
		return
	}

	if dataOne == dataTwo {
		return
	}

	currentOne := ll.Head
	currentTwo := ll.Head

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
func (ll *DoublyLinkedList) Reverse() {
	if ll.Head == nil {
		return
	}

	var prev, Next *DoublyLinkedListNode

	var current = ll.Head

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	ll.Head = prev
}
