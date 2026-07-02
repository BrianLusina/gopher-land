package threadedbinary

import (
	"fmt"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadedBinarySearchTreeInorderTraversal(t *testing.T) {
	type testCase[T types.Comparable] struct {
		data     []T
		expected []T
	}

	t.Run("Inorder Traversal", func(t *testing.T) {

		intTestCases := []testCase[int]{
			{
				data:     []int{4, 2, 6, 1, 3, 5, 7},
				expected: []int{1, 2, 3, 4, 5, 6, 7},
			},
		}

		for _, tc := range intTestCases {
			t.Run(fmt.Sprintf("should return %v", tc.expected), func(t *testing.T) {
				binaryTree := NewThreadedBinarySearchTree[int]()

				for _, v := range tc.data {
					binaryTree.Insert(v)
				}

				actual := binaryTree.InorderTraversal()

				assert.ElementsMatch(t, actual, tc.expected)
			})
		}
	})

	t.Run("Reverse Inorder Traversal", func(t *testing.T) {

		intTestCases := []testCase[int]{
			{
				data:     []int{4, 2, 6, 1, 3, 5, 7},
				expected: []int{7, 6, 5, 4, 3, 2, 1},
			},
		}

		for _, tc := range intTestCases {
			t.Run(fmt.Sprintf("should return %v", tc.expected), func(t *testing.T) {
				binaryTree := NewThreadedBinarySearchTree[int]()

				for _, v := range tc.data {
					binaryTree.Insert(v)
				}

				actual := binaryTree.ReverseInorderTraversal()

				assert.ElementsMatch(t, actual, tc.expected)
			})
		}
	})
}
