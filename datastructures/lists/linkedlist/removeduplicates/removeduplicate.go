package linkedlist

func removeduplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil || head.next == nil {
		return head
	}

	current := head
	next := current.next

	for next != nil {
		if next.data == current.data {
			current.next = current.next.next
			next = current.next
		} else {
			current = next
			next = current.next
		}
	}
	return head
}
