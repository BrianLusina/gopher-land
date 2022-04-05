package trie

type (
	// Node represents a node in a trie tree
	Node struct {
		Children map[rune]*Node
		IsLeaf   bool
	}
)

// NewTrieNode creates a new Trie with initialized Children.
func NewTrieNode() *Node {
	return &Node{
		Children: make(map[rune]*Node),
		IsLeaf:   false,
	}
}

// Compact will remove unnecessary nodes, reducing the capacity, returning true if node n itself should be removed.
func (n *Node) Compact() (remove bool) {

	for r, c := range n.Children {
		if c.Compact() {
			delete(n.Children, r)
		}
	}
	return !n.IsLeaf && len(n.Children) == 0
}

// remove lazily a word from the Trie node, no node is actually removed.
func (n *Node) remove(s string) {
	if len(s) == 0 {
		return
	}
	next, ok := n, false
	for _, c := range s {
		next, ok = next.Children[c]
		if !ok {
			// word cannot be found - we're done !
			return
		}
	}
	next.IsLeaf = false
}

// Remove zero, one or more words lazily from the Trie, no node is actually removed.
func (n *Node) Remove(s ...string) {
	for _, ss := range s {
		n.remove(ss)
	}
}

// Capacity returns the number of nodes in the Trie
func (n *Node) Capacity() int {
	r := 0
	for _, c := range n.Children {
		r += c.Capacity()
	}
	if n.IsLeaf {
		r++
	}
	return r
}

// Size returns the number of words in the Trie
func (n *Node) Size() int {
	r := 0
	for _, c := range n.Children {
		r += c.Size()
	}
	if n.IsLeaf {
		r++
	}
	return r
}
