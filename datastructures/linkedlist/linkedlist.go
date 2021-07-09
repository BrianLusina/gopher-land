package linkedlist

// iLinkedList is an interface for a linked list
type ILinkedList interface {
	Append(interface{})
	Prepend(interface{})
	DeleteHead()
	DeleteTail()
	DeleteAtPosition(position int)
	DeleteNode(node Node)
	DeleteNodeByData(data interface{})
	SwapNodes(dataOne, dataTwo interface{})
	SwapNodesAtKthAndKPlusOne(k int)
	Reverse()
	String() string
}

// LinkedList is a linked list
type LinkedList struct {
	Head *Node
}

func NewNode(data interface{}) Node {
	return Node{Data: data}
}

// NewLinkedList creates a new LinkedList
func NewLinkedList() LinkedList {
	return LinkedList{}
}
