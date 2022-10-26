package singlylinkedlist

import (
	"errors"
	"fmt"
	"gopherland/datastructures/list"
)

// LinkedList data structure represents a singly linked list.
type LinkedList[T comparable] struct {
	// Head of the list
	Head *list.Node[T]
	// size of the list
	size int
}

// NewLinkedList creates a new Singly LinkedList
func NewLinkedList[T comparable]() *LinkedList[T] {
	return new(LinkedList[T])
}

// Append adds a new Node at the end of a linked list
func (sll *LinkedList[T]) Append(val T) {
	node := list.NewNode(val)

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
func (sll *LinkedList[T]) Prepend(val T) {
	node := list.NewNode(val)
	node.Next = sll.Head
	sll.Head = node
}

// Pop removes a node at the end of a linked list & returns its value.
func (sll *LinkedList[T]) Pop() (any, error) {
	if sll.Head == nil {
		return nil, list.ErrEmptyList
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

func (sll *LinkedList[T]) Array() []any {
	var values []any
	current := sll.Head
	for ; current != nil; current = current.Next {
		values = append(values, current.Data)
	}
	return values
}

func (sll LinkedList[T]) String() string {
	current := sll.Head
	result := ""

	for current.Next != nil {
		result += fmt.Sprintf("%v -> ", current.Data)
	}

	result += "nil"

	return result
}

func (sll *LinkedList[T]) DeleteNode(node any) {
	panic("implement me")
}

func (sll *LinkedList[T]) DeleteNodeByData(data any) {
	panic("implement me")
}

func (sll *LinkedList[T]) DeleteTail() (any, error) {
	panic("implement me")
}

func (sll *LinkedList[T]) SwapNodesAtKthAndKPlusOne(k int) {
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
func (sll *LinkedList[T]) DeleteAtBeg() any {
	if sll.Head == nil {
		return -1
	}

	current := sll.Head
	sll.Head = current.Next

	return current.Data
}

// Count counts the number of nodes in a Linked List
func (sll *LinkedList[T]) Count() (count int) {
	current := sll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (sll *LinkedList[T]) Display() {
	for cur := sll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome checks if a linked list is a palindrome
func (sll *LinkedList[T]) IsPalindrome() bool {
	if sll.Head == nil {
		return false
	}

	if sll.Head.Next == nil {
		return true
	}

	current := sll.Head
	stack := []any{}

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
func (sll *LinkedList[T]) DetectCycle() *list.Node[T] {
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
func (sll *LinkedList[T]) GetNthNode(position int) (n *list.Node[T], err error) {
	if position < 0 {
		return nil, errors.New("position less than 0")
	}

	if sll.Head == nil {
		return nil, list.ErrEmptyList
	}

	if position == 0 && sll.Head != nil {
		return sll.Head, nil
	}

	current := sll.Head

	for i := 0; i < position; i++ {
		current = current.Next
	}

	if current == nil {
		return nil, list.ErrInvalidIndex
	}

	return current, nil
}

// GetMiddleNode returns the middle node of the list
func (sll *LinkedList[T]) GetMiddleNode() (*list.Node[T], error) {
	if sll.Head == nil {
		return nil, list.ErrEmptyList
	}

	fast, slow := sll.Head, sll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow, nil
}

func (sll *LinkedList[T]) DeleteNodesByData(data any) *list.Node[T] {
	dummyHead := &list.Node[T]{}
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
func (sll *LinkedList[T]) RemoveDuplicates() *list.Node[T] {
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
func (sll *LinkedList[T]) PairwiseSwap() *list.Node[T] {
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
func (sll *LinkedList[T]) SwapNodesAtKthAndKPlus1(k int) *LinkedList[T] {
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
func (sll *LinkedList[T]) SwapNodes(dataOne, dataTwo any) {
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
func (sll *LinkedList[T]) Reverse() {
	if sll.Head == nil {
		return
	}

	var prev, next *list.Node[T]

	var current = sll.Head

	for current != nil {
		// get the next node
		next = current.Next

		// change next of current, where the actual reversing happens
		current.Next = prev

		// move prev and current one step forward
		prev = current
		current = next
	}

	sll.Head = prev
}

// Rotate rotates a linked list by k nodes counter-clockwise
func (sll *LinkedList[T]) Rotate(k int) {
	if k == 0 {
		return
	}

	if sll.Head == nil {
		return
	}

	current := sll.Head
	count := 1

	// move pointer k positions until we reach the kth node
	for count < k && current != nil {
		current = current.Next
		count++
	}

	// if we don't have a kth node(current is nil), k is greather than or equal to
	// count of nodes in linked list. So no need to rotate
	if current == nil {
		return
	}

	// store the kth node
	kthNode := current

	// move pointer to the end of the linked list
	for current.Next != nil {
		current = current.Next
	}

	// change next of current node to point to previous head
	current.Next = sll.Head

	// change head to k+1th node
	sll.Head = kthNode.Next

	kthNode.Next = nil
}

// ReverseGroup reverses a group of nodes in a linked list
func (sll *LinkedList[T]) ReverseGroups(k int) {
	if k == 0 {
		return
	}

	if sll.Head == nil {
		return
	}

	current := sll.Head
	var next *list.Node[T]
	var prev *list.Node[T]
	count := 0

	for current != nil && count < k {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		count++
	}

	if next != nil {
		sll.Head.Next = next
	}

	return

}

func (sll *LinkedList[T]) DeleteAtPosition(position int) (*list.Node[T], error) {
	if sll.Head == nil {
		return nil, list.ErrEmptyList
	}

	if position == 0 {
		sll.Head = sll.Head.Next
		return sll.Head, nil
	}

	if position > sll.Length() {
		return nil, list.ErrInvalidIndex
	}

	if sll.Head.Next == nil {
		return nil, list.ErrInvalidIndex
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
		return nil, list.ErrInvalidIndex
	}

	prev.Next = next
	return current, nil
}

// AddAtPosition adds a node at the specified position in the LinkedList.
// if position is greater than the length of the LinkedList, the node should not be added
// if the position is equal to the length of the LinkedList, the node should be added at the end of the LinkedList
func (sll *LinkedList[T]) AddAtPosition(position int, data T) {
	node := list.NewNode(data)

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
	var prev, next *list.Node[T]

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

func (sll LinkedList[T]) Length() int {
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
