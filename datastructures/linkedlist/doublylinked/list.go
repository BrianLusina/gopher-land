package doublylinkedlist

import "container/list"

func insertListElements(n int) *list.List { // add elements in list from 1 to n
	lst := list.New()
	head := lst.PushFront(1)
	current := head

	for i := 2; i <= n; i++ {
		lst.InsertAfter(i, current)
		current = current.Next()
	}
	return lst
}
