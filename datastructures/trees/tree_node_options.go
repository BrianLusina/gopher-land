package trees

import "gopherland/pkg/types"

type TreeNodeOption[T types.Comparable] func(*TreeNode[T])

func WithKeyOption[T types.Comparable](key any) TreeNodeOption[T] {
	return func(n *TreeNode[T]) {
		n.Key = key
	}
}

func WithDepthOption[T types.Comparable](depth int) TreeNodeOption[T] {
	return func(n *TreeNode[T]) {
		n.Depth = depth
	}
}

func WithDegreeOption[T types.Comparable](degree int) TreeNodeOption[T] {
	return func(n *TreeNode[T]) {
		n.Degree = degree
	}
}

func WithHeightOption[T types.Comparable](height int) TreeNodeOption[T] {
	return func(n *TreeNode[T]) {
		n.Height = height
	}
}
