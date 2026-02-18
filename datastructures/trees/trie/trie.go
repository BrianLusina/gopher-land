package trie

import "slices"

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

func (t *Trie) Insert(word string) {
	current := t.Root
	for _, c := range word {
		if current.Children[c] == nil {
			current.Children[c] = NewTrieNode()
		}
		current = current.Children[c]
	}
	current.IsLeaf = true
}

func (t *Trie) Search(word string) bool {
	current := t.Root
	for _, c := range word {
		if current.Children[c] == nil {
			return false
		}
		current = current.Children[c]
	}
	return current.IsLeaf
}

// Remove removes a word from the Trie
func (t *Trie) Remove(word string) {
	current := t.Root
	type pair struct {
		node *Node
		char rune
	}
	childList := []pair{}

	for _, c := range word {
		childList = append(childList, pair{node: current, char: c})
		current = current.Children[c]
	}

	slices.Reverse(childList)
	for _, c := range childList {
		parent := c.node
		childChar := c.char
		target := parent.Children[childChar]

		if target.Children != nil {
			return
		}
		delete(parent.Children, childChar)
	}
}
