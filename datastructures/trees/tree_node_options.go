package trees

import "gopherland/pkg/types"

type TreeNodeOption[T types.Comparable] func(*TreeNode[T])

func TreeNodeWithKeyOption[T types.Comparable](key any) TreeNodeOption[T] {
	return func(n *TreeNode[T]) {
		n.Key = key
	}
}
