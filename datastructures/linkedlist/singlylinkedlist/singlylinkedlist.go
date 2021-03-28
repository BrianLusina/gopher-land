package linkedlist

import (
	"errors"
	"fmt"
)

// SinglyLinkedListNode of a Singly Linked List
type SinglyLinkedListNode struct {
	Data interface{}
	Next *SinglyLinkedListNode
}

// LinkedList data structure
type SinglyLinkedList struct {
	Head *SinglyLinkedListNode
}

func CreateList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// NewNode creates a new LinkedList Node
func NewNode(data interface{}) *SinglyLinkedListNode {
	return &SinglyLinkedListNode{data, nil}
}

// Prepend adds a new node to the beggining of the list
func (ll *SinglyLinkedList) Prepend(val interface{}) {
	node := NewNode(val)
	node.Next = ll.Head
	ll.Head = node
}

// Append adds a new Node at the end of a linked list
func (ll *SinglyLinkedList) Append(val interface{}) {
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
func (ll *SinglyLinkedList) DeleteAtBeg() interface{} {
	if ll.Head == nil {
		return -1
	}

	current := ll.Head
	ll.Head = current.Next

	return current.Data
}

// DeleteAtEnd removes a node at the end of a linked list & returns its value.
// Returns -1 if the list is empty
func (ll *SinglyLinkedList) DeleteAtEnd() interface{} {
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
func (ll *SinglyLinkedList) Count() (count int) {
	current := ll.Head

	for current != nil {
		count++
	}

	return
}

// Display prints out the elements of the list.
func (ll *SinglyLinkedList) Display() {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}

	fmt.Print("\n")
}

// IsPalindrome
func (ll *SinglyLinkedList) IsPalindrome() bool {
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
func (ll *SinglyLinkedList) DetectCycle() *SinglyLinkedListNode {
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
func (ll *SinglyLinkedList) GetNthNode(position int) (n *SinglyLinkedListNode, err error) {
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
func (ll *SinglyLinkedList) DeleteNodeAtPosition(position int) (*SinglyLinkedListNode, error) {
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
