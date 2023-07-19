package equalrowandcols

import (
	"gopherland/pkg/utils"
)

// below denotes solution using a hash map

func equalPairsHashmap(grid [][]int) int {
	count := 0
	gridSize := len(grid)
	rowCounter := map[string]int{}

	for _, row := range grid {
		rowString := utils.SliceToString(row, ",")

		if _, ok := rowCounter[rowString]; ok {
			rowCounter[rowString] += 1
		} else {
			rowCounter[rowString] = 1
		}
	}

	for column := 0; column < gridSize; column++ {
		columns := []int{}

		for row := 0; row < gridSize; row++ {
			columns = append(columns, grid[row][column])
		}

		colString := utils.SliceToString(columns, ",")
		if _, ok := rowCounter[colString]; ok {
			count += rowCounter[colString]
		}
	}

	return count
}

// below denotes solution using brute force

func equalPairsBruteForce(grid [][]int) int {
	count := 0
	gridSize := len(grid)

	// check each row r against each column c
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			match := true
			// Iterate over row r and column c.
			for i := 0; i < gridSize; i++ {
				if grid[r][i] != grid[i][c] {
					match = false
					break
				}
			}

			// If row r equals column c, increment count by 1.
			if match {
				count++
			}
		}
	}

	return count
}

// below denotes solution using a trie

type trieNode struct {
	count    int
	children map[int]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: map[int]*trieNode{},
	}
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{
		root: newTrieNode(),
	}
}

func (t *trie) insert(array []int) {
	current := t.root

	for _, item := range array {
		if _, ok := current.children[item]; !ok {
			current.children[item] = newTrieNode()
		}

		if child, ok := current.children[item]; ok {
			current = child
		} else {
			current = newTrieNode()
		}
	}
	current.count++
}

func (t *trie) search(array []int) int {
	current := t.root

	for _, item := range array {
		if child, ok := current.children[item]; ok {
			current = child
		} else {
			return 0
		}
	}
	return current.count
}

func equalPairsTrieNode(grid [][]int) int {
	trie := newTrie()
	count := 0
	n := len(grid)

	for _, row := range grid {
		trie.insert(row)
	}

	for c := 0; c < n; c++ {
		colArray := make([]int, n)

		for r := 0; r < n; r++ {
			colArray[r] = grid[r][c]
		}
		count += trie.search(colArray)
	}

	return count
}
