package binary

import (
	"fmt"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeTraversal(t *testing.T) {

	t.Run("Level Order Traversal", func(t *testing.T) {
		type testCase[T types.Comparable] struct {
			root     *BinaryTreeNode[T]
			expected []any
		}

		intTestCases := []testCase[int]{
			{
				root: NewBinaryTreeNode(1,
					Left(NewBinaryTreeNode(2,
						Left(NewBinaryTreeNode(4)),
						Right(NewBinaryTreeNode(5)),
					)),
					Right(NewBinaryTreeNode(3)),
				),
				expected: []any{1, 2, 3, 4, 5},
			},
		}

		for _, tc := range intTestCases {
			t.Run(fmt.Sprintf("should return %v", tc.expected), func(t *testing.T) {
				binaryTree := NewBinaryTree(tc.root)

				actual := binaryTree.LevelOrderTraversal()

				assert.ElementsMatch(t, actual, tc.expected)
			})
		}
	})
}
