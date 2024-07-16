package binary

import (
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T types.Comparable] struct {
	root        *BinaryTreeNode[T]
	expected    any
	description string
}

var intTestCases = []testCase[int]{
	{
		root:        nil,
		expected:    0,
		description: "should return 0 for no root",
	},
	{
		root:        NewBinaryTreeNode(1),
		expected:    1,
		description: "should return 1 if the binary tree has a root, but no left nor right subtrees",
	},
	{
		root: NewBinaryTreeNode(3,
			Left(NewBinaryTreeNode(9)),
			Right(NewBinaryTreeNode(20,
				Left(NewBinaryTreeNode(15)),
				Right(NewBinaryTreeNode(7))),
			),
		),
		expected:    3,
		description: "should return 3 for tree of 3,9,20,null,null,15,7",
	},
	{
		root:        NewBinaryTreeNode(1, Right(NewBinaryTreeNode(2))),
		expected:    2,
		description: "should return 1 for tree of 1,null,2",
	},
}

func TestBinaryTreeHeight(t *testing.T) {

	for _, tc := range intTestCases {
		t.Run(tc.description, func(t *testing.T) {
			tree := NewBinaryTree[int](tc.root)
			actual := tree.Height()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkBinaryTreeHeight(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range intTestCases {
			b.Run(tc.description, func(t *testing.B) {
				tree := NewBinaryTree[int](tc.root)
				actual := tree.Height()
				assert.Equal(t, tc.expected, actual)
			})
		}
	}
}
