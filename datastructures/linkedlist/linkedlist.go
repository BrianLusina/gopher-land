package linkedlist

type LinkedList interface {
	Append(interface{})
	Prepend(interface{})
	DeleteHead() Node
	DeleteTail() Node
	DeleteAtPosition(position int) Node
	DeleteNode(node Node)
	DeleteNodeByData(data interface{})
	SwapNodesAtKthAndKPlusOne(k int)
	Print() string
}

func Display(ll LinkedList) string {
	return ll.Print()
}
