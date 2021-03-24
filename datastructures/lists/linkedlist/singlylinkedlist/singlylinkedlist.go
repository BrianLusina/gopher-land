package linkedlist

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
