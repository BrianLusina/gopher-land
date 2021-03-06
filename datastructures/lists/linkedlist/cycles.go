package linkedlist

func DetectCycle(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slowPointer := head
	fastPointer := head

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

	curr := head
	for curr != slowPointer {
		curr = curr.Next
		slowPointer = slowPointer.Next
	}

	return curr
}
