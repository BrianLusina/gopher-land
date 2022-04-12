package linkedlist

// Node basic node of a linked list
type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}
