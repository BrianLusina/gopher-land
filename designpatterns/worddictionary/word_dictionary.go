package worddictionary

import (
	"gopherland/datastructures/trees/trie"
	"gopherland/datastructures/trees/trie/prefix"
)

type WordDictionary struct {
	*prefix.Trie
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{
		prefix.NewPrefixTrie(),
	}
}

func (wd *WordDictionary) AddWord(word string) {
	wd.Insert(word)
}

func (wd *WordDictionary) Search(word string) bool {
	res := false

	dfs(wd.Root, word, &res)
	return res
}

func dfs(node *trie.Node, word string, res *bool) {
	if len(word) == 0 {
		if node.IsLeaf {
			*res = true
		}
		return
	}
	if word[0] == '.' {
		for _, child := range node.Children {
			dfs(child, word[1:], res)
		}
	} else {
		node = node.Children[rune(word[0])]
		if node == nil {
			return
		}
		dfs(node, word[1:], res)
	}
}
