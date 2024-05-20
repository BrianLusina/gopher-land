package list

type NodeOption[T comparable] func(*Node[T])

func NodeWithKeyOption[T comparable](key any) NodeOption[T] {
	return func(n *Node[T]) {
		n.key = key
	}
}
