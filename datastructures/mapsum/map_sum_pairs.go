package mapsum

import (
	"strings"
)

// MapSumBruteForce is a map sum data structure that finds the sum of keys with a matching prefix uses a
// Hash Table combined with Brute-Force Search and String Matching.
//
// Time Complexity: Every insert operation is O(1). Every sum operation is O(N*P) where N is the number of items in the
// map, and P is the length of the input prefix.
//
// Space Complexity: The space used by map is linear in the size of all input key and val values combined.
type MapSumBruteForce struct {
	mapping map[string]int
}

func NewMapSumBruteForce() *MapSumBruteForce {
	return &MapSumBruteForce{
		mapping: map[string]int{},
	}
}

func (ms *MapSumBruteForce) Insert(key string, val int) {
	ms.mapping[key] = val
}

func (ms *MapSumBruteForce) Sum(prefix string) int {
	total := 0
	for k, v := range ms.mapping {
		if strings.HasPrefix(k, prefix) {
			total += v
		}
	}
	return total
}

/*
MapSumPrefix is a data structure for calculating prefix sums for given prefixes.
We can remember the answer for all possible prefixes in a HashMap score. When we get a new (key, val) pair, we
update every prefix of key appropriately: each prefix will be changed by delta = val - map[key], where map is the
previously associated value of key (zero if undefined.)

Time Complexity: Every insert operation is O(K^2), where K is the length of the key, as K strings are made of an
average length of K. Every sum operation is O(1).

Space Complexity: The space used by map is linear in the size of all input key and val values combined.
*/
type MapSumPrefix struct {
	mapping map[string]int
	score   map[string]int
}

func NewMapSumPrefix() *MapSumPrefix {
	return &MapSumPrefix{
		mapping: map[string]int{},
		score:   map[string]int{},
	}
}

func (ms *MapSumPrefix) Insert(key string, val int) {
	delta := val
	data, ok := ms.mapping[key]
	if ok {
		delta = val - data
	}
	ms.mapping[key] = val
	for i := range len(key) + 1 {
		prefix := key[:i]
		ms.score[prefix] += delta
	}
}

func (ms *MapSumPrefix) Sum(prefix string) int {
	result, ok := ms.score[prefix]
	if !ok {
		return 0
	}
	return result
}

/*
Since we are dealing with prefixes, a Trie (prefix tree) is a natural data structure to approach this problem. For
every node of the trie corresponding to some prefix, we will remember the desired answer (score) and store it at
this node. As in the approach of using a prefix has map, this involves modifying each node by delta = val - map[key].

Time Complexity: Every insert operation is O(K), where K is the length of the key. Every sum operation is O(K).
Space Complexity: The space used is linear in the size of the total input.
*/
type MapSumTrie struct {
	keys map[rune]*MapSumTrie
	sum  int
}

func NewMapSumTrie() *MapSumTrie {
	return &MapSumTrie{
		keys: make(map[rune]*MapSumTrie),
		sum:  0,
	}
}

func (ms *MapSumTrie) Insert(key string, val int) {
	pointer := ms

	for _, char := range key {
		if pointer.keys[char] == nil {
			newKey := NewMapSumTrie()
			pointer.keys[char] = newKey
		}

		pointer = pointer.keys[char]
	}

	pointer.sum = val
}

func (ms *MapSumTrie) Sum(prefix string) int {
	var dfs func(node *MapSumTrie) int
	dfs = func(node *MapSumTrie) int {
		res := node.sum

		for _, node := range node.keys {
			res += dfs(node)
		}

		return res
	}

	for _, char := range prefix {
		if ms.keys[char] == nil {
			return 0
		}

		ms = ms.keys[char]
	}

	return dfs(ms)
}
