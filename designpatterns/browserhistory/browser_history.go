package browserhistory

import (
	dl "gopherland/datastructures/linkedlist/doublylinked"
)

type BrowserHistory struct {
	CurrentNode *dl.Node
}

func Constructor(homepage string) *BrowserHistory {
	currentNode := &dl.Node{Data: homepage}
	return &BrowserHistory{
		CurrentNode: currentNode,
	}
}

// Visit url from the current page. It clears up all the forward history.
func (b *BrowserHistory) Visit(url string) {
	node := &dl.Node{Data: url}
	current := b.CurrentNode

	if b.CurrentNode.Next == nil {
		node.Prev = b.CurrentNode
		b.CurrentNode = node
		current.Next = node
	} else {
		current.Next = node
		node.Prev = b.CurrentNode
		b.CurrentNode = node
	}
}

// Back moves steps back in history. If you can only return x steps in the history and steps > x,
// you will return only x steps. Return the current url after moving back in history at most steps
func (b *BrowserHistory) Back(steps int) string {
	current := b.CurrentNode
	for i := 0; i < steps; i++ {
		if current.Prev != nil {
			current = current.Prev
		}
	}
	b.CurrentNode = current
	return current.Data.(string)
}

// Forward moves steps forward in history. If you can only forward x steps in the history and steps > x,
// you will forward only x steps. Return the current url after forwarding in history at most steps.
func (b *BrowserHistory) Forward(steps int) string {
	current := b.CurrentNode
	for i := 0; i < steps; i++ {
		if current.Next != nil {
			current = current.Next
		}
	}
	b.CurrentNode = current
	return current.Data.(string)
}
