package linkedlist

import "fmt"

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
