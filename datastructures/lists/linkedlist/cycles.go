package linkedlist

func DetectCycle(head SinglyLinkedListNode) SinglyLinkedListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slowPointer := head
	fastPointer := head

	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			slowPointer = head

			for slowPointer != fastPointer {
				slowPointer = slowPointer.Next
				fastPointer = fastPointer.Next
			}
			return slowPointer
		}
	}
	return nil
}
