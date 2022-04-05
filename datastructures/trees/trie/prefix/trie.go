// Package prefix provides a trie data structure.
package prefix

import "gopherland/datastructures/trees/trie"

type (
	Trie struct {
		trie.Trie
	}
)

// NewPrefixTrie creates a new Trie with uninitialized root.
func NewPrefixTrie() *Trie {
	return &Trie{
		trie.Trie{
			Root: nil,
		},
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
	current := t.Root

	if current == nil {
		current = trie.NewTrieNode()
		t.Root = current
	}

	for _, char := range word {
		next, ok := current.Children[char]
		if !ok {
			next = trie.NewTrieNode()
			current.Children[char] = next
		}
		current = next
	}
	current.IsLeaf = true
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	current := t.Root

	if current == nil {
		return false
	}

	next, ok := current, false
	for _, char := range word {
		next, ok = next.Children[char]
		if !ok {
			return false
		}
	}
	return next.IsLeaf
}

// StartsWith returns true if there is a previously inserted word that starts with the prefix.
func (t *Trie) StartsWith(prefix string) bool {
	current := t.Root
	if current == nil {
		return false
	}

	next, ok := current, false
	for _, char := range prefix {
		next, ok = next.Children[char]
		if !ok {
			return false
		}
	}
	return true
}

// remove lazily a word from the Trie node, no node is actually removed.
func (t *Trie) remove(s string) {
	if len(s) == 0 {
		return
	}

	if t.Root == nil {
		return
	}

	current := t.Root

	next, ok := current, false
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
func (t *Trie) Remove(s ...string) {
	for _, ss := range s {
		t.remove(ss)
	}
}
