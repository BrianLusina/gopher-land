package binary

import (
	"gopherland/datastructures/iterator"
	"strings"
)

// Deserialize converts string data into a binary tree
func Deserialize[T string](tree string) *BinaryTreeNode[string] {
	nodeValues := strings.Split(tree, ",")
	iter := iterator.New(nodeValues)

	var dfs func(nodeData *iterator.Iterator[string]) *BinaryTreeNode[string]

	dfs = func(nodeData *iterator.Iterator[string]) *BinaryTreeNode[string] {
		data := nodeData.Next()

		if data == "nil" || data == "" {
			return nil
		}

		current := NewBinaryTreeNode(data)
		current.left = dfs(nodeData)
		current.right = dfs(nodeData)
		return current
	}

	return dfs(iter)
}
