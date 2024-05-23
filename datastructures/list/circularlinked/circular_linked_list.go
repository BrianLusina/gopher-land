package circularlinked

import (
	"gopherland/datastructures/list"
)

// CircularLinkedList is a circular linked list structure
type CircularLinkedList[T comparable] struct {
	// Head is the head of the circular linked list
	Head *list.Node[T]
}

// New creates a new circular linked list
func New[T comparable](opts ...Option[T]) *CircularLinkedList[T] {
	cll := new(CircularLinkedList[T])

	for _, opt := range opts {
		opt(cll)
	}

	return cll
}

// HeadNode retrieves the Head node of a linked list
func (c *CircularLinkedList[T]) HeadNode() *list.Node[T] {
	return c.Head
}

func (c *CircularLinkedList[T]) All(yield func(T) bool) {
	current := c.Head

	for current != nil {
		if !yield(current.Data) {
			return
		}
		if current.Next == c.Head {
			break
		}
		current = current.Next
	}
}

// Append adds a new data item at the end of the linked list.
func (c *CircularLinkedList[T]) Append(data T) {
	node := list.NewNode(data)
	if c.Head == nil {
		c.Head = node
		node.Next = c.Head
	} else {
		current := c.Head
		for current.Next != c.Head {
			current = current.Next
		}
		current.Next = node
		node.Next = c.Head
	}
}

// Prepend adds a new node to the circular linked list making it the head of the list
func (c *CircularLinkedList[T]) Prepend(data T) {
	newNode := list.NewNode(data)
	current := c.Head
	newNode.Next = c.Head

	if c.Head == nil {
		newNode.Next = newNode
	} else {
		for current.Next != c.Head {
			current = current.Next
		}
		current.Next = newNode
	}
	c.Head = newNode
}

// Array implements list.LinkedList.
func (c *CircularLinkedList[T]) Array() []interface{} {
	panic("unimplemented")
}

// DeleteAtPosition implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteAtPosition(position int) {
	panic("unimplemented")
}

// DeleteMiddle implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteMiddle() (any, error) {
	panic("unimplemented")
}

// DeleteNode implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteNode(node list.Node[T]) {
	if c.Head != nil {
		// if the head node's key matches the key we are looking for
		if c.Head.Compare(&node) {
			// set the current pointer to the head node. This will be used to track the last node as the pointer
			// moves through the list
			current := c.Head
			// move through the list until we reach the pointer that points back to the head node.
			for current.Next != c.Head {
				current = current.Next
			}

			// if the head node equals the next node, that means that this linked list has a length of 1, i.e. just 1
			// node. The head node can be set to None
			if c.Head == c.Head.Next {
				c.Head = nil
			} else {
				// set the current pointer to point to the current head's next
				current.Next = c.Head.Next
				// set the head to now become the next node
				c.Head = c.Head.Next
			}
		} else {
			// we have a situation where the head node's key is not equal to the head node, therefore, we need to
			// traverse the list to find the first node whose key matches the given key. Setting current to the head
			// node acts as the pointer that we keep track of
			current := c.Head
			// previous pointer helps to keep track of the previous node as we traverse, it is initially set to None
			var previous *list.Node[T]

			// we iterate through the linked list as long as the next pointer of the current head is not equal to
			// the head node. This is to avoid an infinite loop as this is a circular linked list.
			for current.Next != c.Head {
				// we set the previous pointer to the current node to keep track of the node before we reset the
				// current pointer to the next node
				previous = current
				// move the current pointer to the next node
				current = current.Next
				// if the current node's key is equal to the key we are searching for
				if current.Compare(&node) {
					// we set the previous node's next pointer to point to the current node's next pointer.
					// Essentially removing the current node from the list
					previous.Next = current.Next
					// set the current node to the current's next node
					current = current.Next
				}
			}
		}
	}
}

// DeleteNodeByKey deletes a given node by a key
func (c *CircularLinkedList[T]) DeleteNodeByKey(key any) {
	if c.Head != nil {
		// if the head node's key matches the key we are looking for
		if c.Head.Key() == key {
			// set the current pointer to the head node. This will be used to track the last node as the pointer
			// moves through the list
			current := c.Head
			// move through the list until we reach the pointer that points back to the head node.
			for current.Next != c.Head {
				current = current.Next
			}

			// if the head node equals the next node, that means that this linked list has a length of 1, i.e. just 1
			// node. The head node can be set to None
			if c.Head == c.Head.Next {
				c.Head = nil
			} else {
				// set the current pointer to point to the current head's next
				current.Next = c.Head.Next
				// set the head to now become the next node
				c.Head = c.Head.Next
			}
		} else {
			// we have a situation where the head node's key is not equal to the head node, therefore, we need to
			// traverse the list to find the first node whose key matches the given key. Setting current to the head
			// node acts as the pointer that we keep track of
			current := c.Head
			// previous pointer helps to keep track of the previous node as we traverse, it is initially set to None
			var previous *list.Node[T]

			// we iterate through the linked list as long as the next pointer of the current head is not equal to
			// the head node. This is to avoid an infinite loop as this is a circular linked list.
			for current.Next != c.Head {
				// we set the previous pointer to the current node to keep track of the node before we reset the
				// current pointer to the next node
				previous = current
				// move the current pointer to the next node
				current = current.Next
				// if the current node's key is equal to the key we are searching for
				if current.Key() == key {
					// we set the previous node's next pointer to point to the current node's next pointer.
					// Essentially removing the current node from the list
					previous.Next = current.Next
					// set the current node to the current's next node
					current = current.Next
				}
			}
		}
	}
}

/**
* SplitList splits a circular linked list into two halves and returns the two halves in a tuple. If the size is 0, i.e. no
* nodes are in this linked list, then it returns None. If the size is 1, then the first portion of the tuple, at
* index 0 will be the head of this circular linked list, while the second portion will be None.
 */
func (c *CircularLinkedList[T]) SplitList() (*CircularLinkedList[T], *CircularLinkedList[T]) {
	size := c.Length()

	if size == 0 {
		return nil, nil
	}

	if size == 1 {
		return c, nil
	}

	mid := size / 2
	count := 0

	var previous *list.Node[T]
	current := c.Head

	for current != nil && count < mid {
		count++
		previous = current
		current = current.Next
	}

	previous.Next = c.Head

	secondList := &CircularLinkedList[T]{}
	for current.Next != c.Head {
		secondList.Append(current.Data)
		current = current.Next
	}

	secondList.Append(current.Data)

	return c, secondList
}

// DeleteTail implements list.LinkedList.
func (c *CircularLinkedList[T]) DeleteTail() (node interface{}, err error) {
	panic("unimplemented")
}

// GetMiddleNode implements list.LinkedList.
func (c *CircularLinkedList[T]) GetMiddleNode() (node list.Node[T], err error) {
	panic("unimplemented")
}

// GetNthNode implements list.LinkedList.
func (c *CircularLinkedList[T]) GetNthNode(position int) (node list.Node[T], err error) {
	panic("unimplemented")
}

// Length implements list.LinkedList.
func (c *CircularLinkedList[T]) Length() int {
	current := c.Head
	count := 0
	for current != nil {
		count++
		current = current.Next
		if current == c.Head {
			break
		}
	}
	return count
}

// MaximumPairSum implements list.LinkedList.
func (c *CircularLinkedList[T]) MaximumPairSum() int {
	panic("unimplemented")
}

// OddEvenList implements list.LinkedList.
func (c *CircularLinkedList[T]) OddEvenList() *list.Node[T] {
	panic("unimplemented")
}

// Pop implements list.LinkedList.
func (c *CircularLinkedList[T]) Pop() (interface{}, error) {
	panic("unimplemented")
}

// Reverse implements list.LinkedList.
func (c *CircularLinkedList[T]) Reverse() {
	panic("unimplemented")
}

// ReverseGroups implements list.LinkedList.
func (c *CircularLinkedList[T]) ReverseGroups(k int) {
	panic("unimplemented")
}

// Rotate implements list.LinkedList.
func (c *CircularLinkedList[T]) Rotate(k int) {
	panic("unimplemented")
}

// String implements list.LinkedList.
func (c *CircularLinkedList[T]) String() string {
	panic("unimplemented")
}

// SwapNodes implements list.LinkedList.
func (c *CircularLinkedList[T]) SwapNodes(dataOne interface{}, dataTwo interface{}) {
	panic("unimplemented")
}

// SwapNodesAtKthAndKPlusOne implements list.LinkedList.
func (c *CircularLinkedList[T]) SwapNodesAtKthAndKPlusOne(k int) {
	panic("unimplemented")
}
