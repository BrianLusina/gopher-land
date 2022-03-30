// Package prefix provides a trie data structure.
package prefix

// Trie represents a node in a trie.
type Trie struct {
	children map[rune]*Trie
	isLeaf   bool
}

// NewTrie creates a new Trie with initialized children.
func NewTrie() *Trie {
	return &Trie{
		children: make(map[rune]*Trie),
		isLeaf:   false,
	}
}

// InsertMultiple inserts multiple words in a Trie
func (t *Trie) InsertMultiple(words ...string) {
	for _, word := range words {
		t.Insert(word)
	}
}

// Insert a single word into the trie.
func (t *Trie) Insert(word string) {
	current := t

	for _, char := range word {
		next, ok := current.children[char]
		if !ok {
			next = NewTrie()
			current.children[char] = next
		}
		current = next
	}
	current.isLeaf = true
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	next, ok := t, false
	for _, char := range word {
		next, ok = next.children[char]
		if !ok {
			return false
		}
	}
	return next.isLeaf
}

// StartsWith returns true if there is a previously inserted word that starts with the prefix.
func (t *Trie) StartsWith(prefix string) bool {
	next, ok := t, false
	for _, char := range prefix {
		next, ok = next.children[char]
		if !ok {
			return false
		}
	}
	return true
}

// Capacity returns the number of nodes in the Trie
func (t *Trie) Capacity() int {
	r := 0
	for _, c := range t.children {
		r += c.Capacity()
	}
	if t.isLeaf {
		r++
	}
	return r
}

// Size returns the number of words in the Trie
func (t *Trie) Size() int {
	r := 0
	for _, c := range t.children {
		r += c.Size()
	}
	if t.isLeaf {
		r++
	}
	return r
}

// remove lazily a word from the Trie node, no node is actually removed.
func (t *Trie) remove(s string) {
	if len(s) == 0 {
		return
	}
	next, ok := t, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			// word cannot be found - we're done !
			return
		}
	}
	next.isLeaf = false
}

// Remove zero, one or more words lazily from the Trie, no node is actually removed.
func (t *Trie) Remove(s ...string) {
	for _, ss := range s {
		t.remove(ss)
	}
}

// Compact will remove unnecessary nodes, reducing the capacity, returning true if node n itself should be removed.
func (t *Trie) Compact() (remove bool) {

	for r, c := range t.children {
		if c.Compact() {
			delete(t.children, r)
		}
	}
	return !t.isLeaf && len(t.children) == 0
}
