package doublylinkedlist

import (
	"errors"
	"fmt"
	"gopherland/datastructures/linkedlist"
)

// LinkedList LinkedList data structure
type LinkedList[T comparable] struct {
	Head, Tail *Node[T]
	Len        int //current length of the list
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return new(LinkedList[T])
}

// Append adds a new Node at the end of a linked list
func (dll *LinkedList[T]) Append(val T) {
	node := NewNode(val)
	dll.Len++

	if dll.Head == nil {
		dll.Head = node
		return
	}

	current := dll.Head
	for ; current.Next != nil; current = current.Next {
	}

	current.Next = node
	node.Prev = current
}

func (dll *LinkedList[T]) String() string {
	var s string
	current := dll.Head

	for current != nil {
		s += fmt.Sprintf("%v <-> ", current.Data)
		current = current.Next
	}

	s += "nil"

	return s
}

// Size returns the size of the doubly linked list
func (dll *LinkedList[T]) Size() int {
	return dll.Len
}

// Front returns the head of the linked list or nil if empty
func (dll *LinkedList[T]) Front() *Node[T] {
	if dll.Len == 0 {
		return nil
	}
	return dll.Head
}

// Back returns the tail of the linked list or nil if empty
func (dll *LinkedList[T]) Back() *Node[T] {
	if dll.Len == 0 {
		return nil
	}
	return dll.Tail
}

// Prepend adds a new node to the beginning of the list
func (dll *LinkedList[T]) Prepend(val T) {
	node := NewNode(val)

	if dll.Head == nil {
		dll.Head = node
	}

	node.Next = dll.Head
	dll.Head.Prev = node
	dll.Head = node
	dll.Len++
}

func (dll *LinkedList[T]) Pop() {

}

// DeleteNode removes a node from the list
func (dll *LinkedList[T]) DeleteNode(node *Node[T]) {
	if dll.Head == nil {
		return
	}

	current := dll.Head

	for current != nil {
		if current.Data == node.Data {
			if node.Next != nil {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			} else {
				current.Prev.Next = nil
			}

			dll.Len--
			return
		}
		current = current.Next
	}
}

// DeleteNodeByData deletes a node by its data
func (dll *LinkedList[T]) DeleteNodeByData(data any) {
	if dll.Head == nil {
		return
	}

	current := dll.Head

	for current != nil {
		if current.Data == data {
			if current.Next != nil {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			} else {
				current.Prev.Next = nil
			}

			dll.Len--
			return
		}
		current = current.Next
	}
}

// DeleteTail removes the last node in a doubly linked list and returns it
func (dll *LinkedList[T]) DeleteTail() (*Node[T], error) {
	switch {
	case dll.Head == nil:
		return nil, linkedlist.ErrEmptyList

	case dll.Head.Next == nil:
		node := dll.Head
		dll.Head = nil
		dll.Len--
		return node, nil
	case dll.Head.Next != nil:
		current := dll.Head

		for current.Next.Next != nil {
			current = current.Next
		}

		dll.Len--
		node := current.Next
		node.Prev = nil
		current.Next = nil
		return node, nil
	default:
		return nil, errors.New("Failed to delete tail")
	}
	// in the instance of having a tail
	// switch {

	// 	// empty list
	// 	case l.Head == nil && l.Tail == nil:
	// 		return nil, ErrEmptyList

	// 	// linked list with 1 node, the head is the tail
	// 	case l.Head != nil && l.Tail != nil && l.Tail.prev == nil:
	// 		tail := l.Tail
	// 		l.Head = nil
	// 		l.Tail = nil
	// 		return tail.Val, nil

	// 	// linked list has more than 1 node
	// 	case l.Head != nil && l.Tail != nil && l.Tail.prev != nil:
	// 		tail := l.Tail
	// 		l.Tail.prev.next = nil
	// 		l.Tail = l.Tail.prev
	// 		return tail.Val, nil

	// 	default:
	// 		return nil, errors.New("unexpected error")
	// 	}
}

func (dll *LinkedList[T]) SwapNodesAtKthAndKPlusOne(k int) {
	if dll.Head == nil || dll.Head.Next == nil {
		return
	}

	current := dll.Head
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
	dll.Head = next
}

// DeleteAtBeg removes a node at the beggining of a linked list & returns its value.
func (dll *LinkedList[T]) DeleteAtBeg() any {
	if dll.Head == nil {
		return nil
	}

	current := dll.Head
	dll.Head = current.Next
	dll.Len--

	return current.Data
}

// DeleteAtEnd removes a node at the end of a linked list & returns its value.
// Returns -1 if the list is empty
func (dll *LinkedList[T]) DeleteAtEnd() any {
	if dll.Head == nil {
		return nil
	}

	current := dll.Head
	for ; current.Next.Next != nil; current = current.Next {
	}

	data := current.Next.Data
	current.Next = nil
	dll.Len--
	return data
}

// Count counts the number of nodes in a Linked List
func (dll *LinkedList[T]) Count() (count int) {
	current := dll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (dll *LinkedList[T]) Display() {
	for cur := dll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome
func (dll *LinkedList[T]) IsPalindrome() bool {
	if dll.Head == nil {
		return false
	}

	if dll.Head.Next == nil {
		return true
	}

	current := dll.Head
	stack := []any{}

	for current != nil {
		stack = append(stack, current.Data)
		current = current.Next
	}

	current = dll.Head

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
func (dll *LinkedList[T]) DetectCycle() *Node[T] {
	if dll.Head == nil || dll.Head.Next == nil {
		return nil
	}

	slowPointer := dll.Head
	fastPointer := dll.Head

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

	curr := dll.Head
	for curr != slowPointer {
		curr = curr.Next
		slowPointer = slowPointer.Next
	}

	return curr
}

// GetNthNode gets the nth node in a linked list
func (dll *LinkedList[T]) GetNthNode(position int) (n *Node[T], err error) {
	if position < 0 {
		return nil, errors.New("Position less than 0")
	}

	if dll.Head == nil {
		return nil, errors.New("List is empty")
	}

	current := dll.Head

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

// GetMiddleNode returns the middle node of the list
func (sll *LinkedList[T]) GetMiddleNode() (*Node[T], error) {
	if sll.Head == nil {
		return nil, linkedlist.ErrEmptyList
	}

	fast, slow := sll.Head, sll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow, nil
}

// DeleteNodeAtPosition deletes a node at the specified position and returns the deleted node
func (dll *LinkedList[T]) DeleteNodeAtPosition(position int) (*Node[T], error) {
	if position < 0 {
		errMessage := fmt.Sprintf("Invalid Index position given. Index is {%d}, expected position >= 0", position)
		return nil, errors.New(errMessage)
	}

	// Nothing to delete here
	if dll.Head == nil {
		return nil, nil
	}

	if dll.Head != nil && position == 0 {
		current := dll.Head
		dll.Head = nil
		dll.Len--
		return current, nil
	}

	current := dll.Head

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
	dll.Len--

	return node, nil
}

// RemoveDuplicates removes duplicates from a LinkedList
// This assumes the linked list is sorted in ascending order
func (dll *LinkedList[T]) RemoveDuplicates() *Node[T] {
	head := dll.Head

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
func (dll *LinkedList[T]) PairwiseSwap() *Node[T] {
	head := dll.Head

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
func (dll *LinkedList[T]) SwapNodesAtKthAndKPlus1(k int) *LinkedList[T] {
	a, b := dll.Head, dll.Head

	for i := 1; i < k; i++ {
		a = a.Next
	}

	node := a
	a = a.Next

	for a != nil {
		a, b = a.Next, b.Next
	}

	node.Data, b.Data = b.Data, node.Data

	return dll
}

// SwapNodes swaps the data items of 2 nodes in the Linkedlist
// Swaps two nodes based on the data they contain. We search through the LinkedList looking for the data item in
// each node. Once the first is found, we keep track of it and move on until we find the next data item. Once that
// is found, we swap the two nodes' data items.
// If we can't find the first data item nor the second. No need to perform swap. If the 2 data items are similar
// no need to perform swap as well.
// If the LinkedList is empty (i.e. has no head node), return, no need to swap when we have no LinkedList :)
func (dll *LinkedList[T]) SwapNodes(dataOne, dataTwo any) {
	if dll.Head == nil {
		return
	}

	if dataOne == dataTwo {
		return
	}

	currentOne := dll.Head
	currentTwo := dll.Head

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
func (dll *LinkedList[T]) Reverse() {
	if dll.Head == nil {
		return
	}

	var prev *Node[T]

	var current = dll.Head

	for current != nil {
		// copy pointer to next element before we overwrite current.Next
		next := current.Next

		// reverse the next pointer and previous
		current.Next = prev
		current.Prev = next

		// move prev and current forward
		prev = current
		current = next
	}

	dll.Head = prev
}

func (dll *LinkedList[T]) DeleteAtPosition(position int) {
	if dll.Head == nil {
		return
	}

	var prev, Next *Node[T]

	var current = dll.Head

	for current != nil {
		Next = current.Next
		current.Next = prev
		prev = current
		current = Next
	}

	dll.Head = prev
	dll.Len--
}
