package josephuscircle

import (
	"gopherland/datastructures/list"
	"gopherland/datastructures/list/circularlinked"
)

func josephusCircle[T comparable](circularLinkedList circularlinked.CircularLinkedList[T], step int) list.Node[T] {
	current := circularLinkedList.Head

	length := circularLinkedList.Length()
	for length > 1 {
		count := 1
		for count != step {
			current = current.Next
			count++
		}

		circularLinkedList.DeleteNode(*current)
		current = current.Next
		length--
	}
	return *current
}
