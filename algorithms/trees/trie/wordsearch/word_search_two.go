package wordsearch

import (
	"fmt"
	"gopherland/datastructures/trees/trie"
)

// FindStrings finds all the words in the grid and returns them as a list.
func FindStrings(grid [][]string, words []string) []string {
	if len(grid) == 0 || len(grid[0]) == 0 || len(words) == 0 {
		return []string{}
	}

	trieTree := trie.NewTrie()
	for _, word := range words {
		trieTree.Insert(word)
	}

	rowCount, colCount := len(grid), len(grid[0])
	result := []string{}
	visited := map[string]bool{}
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	isValid := func(row, col int) bool {
		return row >= 0 && row < rowCount && col >= 0 && col < colCount
	}

	var dfs func(row, col int, node *trie.Node, path string)

	dfs = func(row, col int, node *trie.Node, path string) {
		if node.IsLeaf {
			result = append(result, path)
			node.IsLeaf = false // Avoid duplicates
			trieTree.Remove(path)
		}
		// mark the current cell as visited
		visited[fmt.Sprintf("%d,%d", row, col)] = true

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if isValid(newRow, newCol) && !visited[fmt.Sprintf("%d,%d", newRow, newCol)] {
				nextChar := grid[newRow][newCol]
				if node.Children[rune(nextChar[0])] != nil {
					dfs(newRow, newCol, node.Children[rune(nextChar[0])], path+nextChar)
				}
			}
		}

		// unmark the current cell as visited
		delete(visited, fmt.Sprintf("%d,%d", row, col))
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			char := (grid[i][j])
			if node := trieTree.Root.Children[(rune(char[0]))]; node != nil {
				dfs(i, j, node, char)
			}
		}
	}

	return result
}
