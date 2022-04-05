package trie

type (
	// Trie represents a trie tree
	Trie struct {
		Root *Node
	}
)

// NewTrie creates a new Trie with initialized root.
func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

// Capacity returns the number of nodes in the Trie
func (t *Trie) Capacity() int {
	if t.Root == nil {
		return 0
	}

	current := t.Root

	r := 0
	for _, c := range current.Children {
		r += c.Capacity()
	}
	if current.IsLeaf {
		r++
	}
	return r
}

// Size returns the number of words in the Trie
func (t *Trie) Size() int {
	if t.Root == nil {
		return 0
	}

	current := t.Root

	r := 0
	for _, c := range current.Children {
		r += c.Size()
	}
	if current.IsLeaf {
		r++
	}
	return r
}

// Compact will remove unnecessary nodes, reducing the capacity, returning true if node n itself should be removed.
func (t *Trie) Compact() (remove bool) {
	if t.Root == nil {
		return
	}

	current := t.Root

	for r, c := range current.Children {
		if t.Compact() {
			delete(c.Children, r)
		}
	}

	return !current.IsLeaf && len(current.Children) == 0
}
