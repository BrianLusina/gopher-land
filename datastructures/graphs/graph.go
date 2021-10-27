package graphs

type Node struct {
	Value int
}

func (t *Node) NewNode(value int) Node {
	return Node{
		Value: value,
	}
}
