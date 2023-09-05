package binary

import (
	"fmt"
	"gopherland/datastructures/trees"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseSlice[T types.Comparable] struct {
	slice []T
	tree  *BinaryTree[T]
}

var testCaseIntSlice = []testCaseSlice[int]{
	{
		slice: []int{},
		tree:  &BinaryTree[int]{},
	},
	{
		slice: []int{1, 7, 9, 7, -8, 1, 2},
		tree: &BinaryTree[int]{
			root: &BinaryTreeNode[int]{
				TreeNode: trees.TreeNode[int]{
					Data: 1,
				},
				left:  NewBinaryTreeNode[int](7, Left(NewBinaryTreeNode(7)), Right(NewBinaryTreeNode(-8))),
				right: NewBinaryTreeNode[int](9, Left(NewBinaryTreeNode(1)), Right(NewBinaryTreeNode(2))),
			},
		},
	},
}

func TestNewBinaryTreeSlice(t *testing.T) {
	for _, tc := range testCaseIntSlice {
		t.Run(fmt.Sprintf("NewBinaryTreeFromSlice(%v)", tc.slice), func(t *testing.T) {
			actual := NewBinaryTreeFromSlice(tc.slice)

			if !assert.Equal(t, tc.tree, actual) {
				t.Fatalf("expected %v, got %v", tc.tree, actual)
			}
		})
	}
}
