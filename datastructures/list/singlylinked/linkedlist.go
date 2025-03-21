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
	size int
}

// New creates a new Singly LinkedList
func New[T comparable]() *LinkedList[T] {
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

func (sll *LinkedList[T]) String() string {
	current := sll.Head
	result := ""

	for current.Next != nil {
		result += fmt.Sprintf("%v -> ", current.Data)
	}

	result += "nil"

	return result
}

func (sll *LinkedList[T]) DeleteNodeByPosition(position int) (*list.Node[T], error) {
	if sll.Head == nil {
		return nil, list.ErrEmptyList
	}

	current := sll.Head
	if position == 0 {
		sll.Head = current.Next
		return current, nil
	}

	var prev list.Node[T]
	count := 0

	for current != nil && count != position {
		prev = *current
		current = current.Next
		count += 1
	}

	if current == nil {
		return nil, list.ErrInvalidIndex
	}

	prev.Next = current.Next
	return current, nil
}

func (sll *LinkedList[T]) DeleteNodeByData(data any) {
	current := sll.Head

	// If the data we are deleting is at the head, then change the head to the next node in the linked list
	// and return
	if current != nil && current.Data == data {
		sll.Head = current.Next
		return
	}

	// this will be used to keep track of the previous node of the node to delete
	var previous list.Node[T]

	// we move the pointer down the LinkedList until we find the Node whose data matches what we want to delete
	for current != nil && current.Data != data {
		previous = *current
		current = current.Next
	}

	//if there is no node that matches the condition above, we exit
	if current == nil {
		return
	}

	// re-assign the pointers of the nodes around the node to delete. That is, moving the previous node's next
	// pointer to the current node's next pointer. This essentially 'deletes' the node by the data attribute
	previous.Next = current.Next
	return
}

func (sll *LinkedList[T]) DeleteTail() (any, error) {
	panic("implement me")
}

func (sll *LinkedList[T]) DeleteMiddle() *list.Node[T] {
	if sll.Head == nil || sll.Head.Next == nil {
		return nil
	}

	nodeCount := sll.Length()
	middleIndex := nodeCount / 2

	current := sll.Head

	for i := 0; i < middleIndex-1; i++ {
		current = current.Next
	}

	middleNode := current.Next
	current.Next = current.Next.Next

	return middleNode
}

func (sll *LinkedList[T]) DeleteMiddle2Pointers() *list.Node[T] {
	if sll.Head == nil || sll.Head.Next == nil {
		return nil
	}

	slowPointer, fastPointer := sll.Head, sll.Head.Next.Next

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	middleNode := slowPointer.Next
	slowPointer.Next = slowPointer.Next.Next

	return middleNode
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

// CountOccurrences counts the number of occurrences of a data in a LinkedList. If the linked list is empty(no head). 0 is returned.
// otherwise the occurrences of the data element will be sought using the equality operator. This assumes that the
// data element in each node already implements this operator.
//
// Complexity:
// The assumption here is that n is the number of nodes in the linked list.
//
// Time O(n): This is because the algorithm iterates through each node in the linked list to find data values in
// each node that equal the provided data argument in the function. This is both for the worst and best case as
// each node in the linked list has to be checked
//
// Space O(1): no extra space is required other than the value being incremented for each node whose data element
// equals the provided data argument.
func (sll *LinkedList[T]) CountOccurrences(data T) int {
	current := sll.Head
	if current == nil {
		return 0
	}

	occurrences := 0

	for current != nil {
		if current.Data == data {
			occurrences++
		}
		current = current.Next
	}

	return occurrences
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
	stack := []T{}

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

// IsPalindromeTwoPointers checks if a linked list is a palindrome using two pointers
func (sll *LinkedList[T]) IsPalindromeTwoPointers() bool {
	if sll.Head == nil {
		return false
	}

	if sll.Head.Next == nil {
		return true
	}

	firstPointer := sll.Head
	lastPointer := sll.Head
	previous := []*list.Node[T]{}
	i := 0

	for lastPointer != nil {
		previous = append(previous, lastPointer)
		lastPointer = lastPointer.Next
		i++
	}
	lastPointer = previous[i-1]

	count := 0

	for count <= i/2+1 {
		if previous[(len(previous)-count)-1].Data != firstPointer.Data {
			return false
		}
		firstPointer = firstPointer.Next
		count++
	}

	return true
}

// IsPalindromeTwoPointers2 checks if a linked list is a palindrome using another variation of two pointers
// Checks to see if a Linked list is a Palindrome.
// Returns True if it is, false otherwise.
// Uses two pointers approach to check if a linked list is a palindrome. First it finds the middle of the list using
// two pointers a fast and a slow pointer and then reverses the second half of the list. Once the second half is
// reversed, it compares the first half and the reversed second half
//
// Complexity:
// We assume that n is the number of nodes in the linked list
//
// Time O(n): we traverse the linked list to check for the palindrome property.
//
// Space O(1): No extra space is used when traversing the linked list
func (sll *LinkedList[T]) IsPalindromeTwoPointers2() bool {
	// an empty linked list or a linked list with only one node is considered a palindrome
	if sll.Head == nil || sll.Head.Next == nil {
		return true
	}

	// find the middle of the list using fast and slow pointers. The fast pointer will have gotten to the end of the
	// the linked list and the slow pointer will be at the middle of the linked list
	fastPointer := sll.Head
	slowPointer := sll.Head
	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	var previous *list.Node[T]
	// reverse the second half of the list
	for slowPointer != nil {
		nxt := slowPointer.Next
		slowPointer.Next = previous
		previous = slowPointer
		slowPointer = nxt
	}

	// now prev is the head of the reversed second half
	// compare the first half and the reversed second half
	left, right := sll.Head, previous
	for right != nil {
		if left.Data != right.Data {
			return false
		}
		left = left.Next
		right = right.Next
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

	seen := map[T]bool{}
	current := head
	var previous *list.Node[T]

	for current != nil {
		_, ok := seen[current.Data]

		if ok {
			previous.Next = current
			current = nil
		} else {
			seen[current.Data] = true
			previous = current
		}
		current = previous.Next
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

	// if we don't have a kth node(current is nil), k is greater than or equal to
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

func (sll *LinkedList[T]) Length() int {
	if sll.Head == nil {
		return 0
	}

	var current = sll.Head
	var count int

	for current != nil {
		current = current.Next
		count++
	}

	sll.size = count

	return count
}

// MiddleNode returns the middle node in the singly linked list
func (sll *LinkedList[T]) MiddleNode() *list.Node[T] {
	if sll.Head == nil {
		return nil
	}

	slow_pointer := sll.Head
	fast_pointer := sll.Head

	for fast_pointer != nil && fast_pointer.Next != nil {
		slow_pointer = slow_pointer.Next
		fast_pointer = fast_pointer.Next.Next
	}

	return slow_pointer
}

// KthToLastNode returns the kth to the last node in the linked list
//
// For example given a linked list of:
// a -> b -> c -> d
//
// if k is the number 1 then d should be returned
// if k is the number 2 then c should be returned
// if k is the number 3 then b should be returned
// if k is the number 4 then a should be returned
// if k exceeds the size of the list then nil and an error are returned
//
// Complexity analysis:
//
// Time Complexity: O(n) where n is the number of nodes in the linked list. We traverse the linked list to get to the the Kth to the last node.
// We are using 2 pointers 1 to keep track of the end node and the other that will be a distance of k from the first pointer. Since we have to
// traverse up to the last node, the worst case is O(n) as k could be the very last position of the last node.
//
// Space complexity is O(1), no extra memory is being used other than the pointers on the linked list
func (sll *LinkedList[T]) KthToLastNode(k int) (*list.Node[T], error) {
	if k < 1 {
		return nil, fmt.Errorf("k %d is less than 1", k)
	}

	leftNode := sll.Head
	rightNode := sll.Head

	for i := 0; i < k-1; i++ {
		if rightNode.Next == nil {
			return nil, fmt.Errorf("k %d is larger than the length of the linked list", k)
		}
		rightNode = rightNode.Next
	}

	for rightNode.Next != nil {
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}

	return leftNode, nil
}

// OddEvenList returns the linked list grouping all nodes with odd indices together followed by nodes with even indices
func (sll *LinkedList[T]) OddEvenList() *list.Node[T] {
	if sll.Head == nil || sll.Head.Next == nil {
		return sll.Head
	}

	odd := sll.Head
	even := sll.Head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return sll.Head
}

// InsertAfterNode inserts the data after the given prevNode if the node is in the linked list
func (sll *LinkedList[T]) InsertAfterNode(prevNode *list.Node[T], data T) {
	if prevNode == nil {
		return
	}

	// node to insert
	newNode := list.NewNode(data)
	// set the new node's next pointer to point to the prevNode's next pointer
	newNode.Next = prevNode.Next

	// set the prevNode's next pointer to point to the new node
	prevNode.Next = newNode
}

// MoveTailToHead moves the tail to the head of the linked list making it the new head node
// Uses two pointers where last pointer will be moved until it points to the last node in the linked list. The second pointer, previous, will
// point to the second last node in the linked list.
// Complexity Analysis:
// An assumption is made where n is the number of nodes in the linked list
// Time: O(n) as the the pointers have to be moved through each node in the linked list until both point to the last and second last nodes in the linked list
// Space O(1) as no extra space is incurred in the iteration. Only pointers are moved at the end to move the tail node to the head and make the second to last node
// the new tail
func (sll *LinkedList[T]) MoveTailToHead() {
	if sll.Head != nil && sll.Head.Next != nil {
		// pointer that is initially set to the head node of the linked list. This will be used to keep track of the last node in the linked list
		last := sll.Head
		// previous is a pointer initially set to nil that will be used to keep track of the second to last node in the linked list that will become
		// the new tail
		var previous *list.Node[T]

		// move the last pointer to the end of the linked list while the node has a next pointer. This will set the previous pointer to last while also
		// setting the last pointer to the last node. After this loop, the previous will be at the second last node while the last will be at the last node
		// in the linked list
		for last.Next != nil {
			previous = last
			last = last.Next
		}
		// set the last node's next pointer to point to the head node. Note that at this point in time this has become a circular linked list
		last.Next = sll.Head
		// set the previous'(second to last node) next pointer to nil, consequentially breaking the circular linked list and setting this node as the last(tail)
		// node in the linked list
		previous.Next = nil
		// set the new head of the linked list as the last node
		sll.Head = last
	}
}
