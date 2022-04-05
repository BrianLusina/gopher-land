// Package prefix provides a trie data structure.
package prefix

import "gopherland/datastructures/trees/trie"

type (
	TrieNode struct {
		*trie.Node
	}
)

// NewPrefixTrieNode creates a new Trie with initialized Children.
func NewPrefixTrieNode() *TrieNode {
	return &TrieNode{
		&trie.Node{
			Children: make(map[rune]*trie.Node),
			IsLeaf:   false,
		},
	}
}

// InsertMultiple inserts multiple words in a Trie
func (t *TrieNode) InsertMultiple(words ...string) {
	for _, word := range words {
		t.Insert(word)
	}
}

// Insert a single word into the trie.
func (t *TrieNode) Insert(word string) {
	current := t.Node

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
func (t *TrieNode) Search(word string) bool {
	next, ok := t.Node, false
	for _, char := range word {
		next, ok = next.Children[char]
		if !ok {
			return false
		}
	}
	return next.IsLeaf
}

// StartsWith returns true if there is a previously inserted word that starts with the prefix.
func (t *TrieNode) StartsWith(prefix string) bool {
	next, ok := t.Node, false
	for _, char := range prefix {
		next, ok = next.Children[char]
		if !ok {
			return false
		}
	}
	return true
}
