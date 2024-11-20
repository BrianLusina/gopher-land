package cpucaches

type node struct {
	value int
	next  *node
}

func sumLinkedList(n *node) int {
	var total int
	for n != nil {
		total += n.value
		n = n.next
	}
	return total
}
