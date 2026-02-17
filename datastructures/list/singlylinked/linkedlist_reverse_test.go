package singlylinkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type reverseKGroupTestCase struct {
	values   []int
	k        int
	expected []int
}

var reverseKGroupTestCases = []reverseKGroupTestCase{
	{
		values:   []int{1, 2, 3, 4, 5, 6},
		k:        4,
		expected: []int{4, 3, 2, 1, 5, 6},
	},
	{
		values:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		k:        3,
		expected: []int{7, 8, 9, 4, 5, 6, 1, 2, 3},
	},
	{
		values:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		k:        0,
		expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
	},
	{
		values:   []int{1, 2, 3, 4, 5},
		k:        2,
		expected: []int{2, 1, 4, 3, 5},
	},
	{
		values:   []int{1, 2, 3, 4, 5},
		k:        3,
		expected: []int{3, 2, 1, 4, 5},
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6},
		k:        2,
		expected: []int{2, 1, 4, 3, 6, 5},
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7},
		k:        3,
		expected: []int{3, 2, 1, 6, 5, 4, 7},
	},
	{
		values:   []int{8, 11, 4, 12, 0},
		k:        1,
		expected: []int{8, 11, 4, 12, 0},
	},
	{
		values:   []int{3, 4, 5, 6, 2, 8, 7, 7},
		k:        3,
		expected: []int{5, 4, 3, 8, 2, 6, 7, 7},
	},
	{
		values:   []int{1, 2, 3, 4, 5, 6, 7},
		k:        7,
		expected: []int{7, 6, 5, 4, 3, 2, 1},
	},
}

func TestReverseKGroup(t *testing.T) {
	for _, tc := range reverseKGroupTestCases {
		t.Run(fmt.Sprintf("should reverse a linked list %v in groups of k %d", tc.values, tc.k), func(t *testing.T) {
			sll := New[int]()
			for _, v := range tc.values {
				sll.Append(v)
			}

			actual := sll.ReverseGroups(tc.k)
			actualData := []int{}
			current := actual
			for current != nil {
				actualData = append(actualData, current.Data)
				current = current.Next
			}

			assert.Equal(t, tc.expected, actualData)
		})
	}
}

func TestReverseKGroupUtil(t *testing.T) {
	for _, tc := range reverseKGroupTestCases {
		t.Run(fmt.Sprintf("should reverse a linked list %v in groups of k %d", tc.values, tc.k), func(t *testing.T) {
			sll := New[int]()
			for _, v := range tc.values {
				sll.Append(v)
			}

			actual := ReverseGroups(sll.Head, tc.k)

			actualData := []int{}
			current := actual
			for current != nil {
				actualData = append(actualData, current.Data)
				current = current.Next
			}

			assert.Equal(t, tc.expected, actualData)
		})
	}
}

func BenchmarkReverseKGroup(b *testing.B) {
	for _, tc := range reverseKGroupTestCases {
		b.Run(fmt.Sprintf("k=%d/n=%d", tc.k, len(tc.values)), func(b *testing.B) {
			for b.Loop() {
				sll := New[int]()
				for _, v := range tc.values {
					sll.Append(v)
				}
				sll.ReverseGroups(tc.k)
			}
		})
	}
}

func BenchmarkReverseKGroupUtil(b *testing.B) {
	for b.Loop() {
		for _, tc := range reverseKGroupTestCases {
			sll := New[int]()
			b.Run("should reverse a linked list of 1000 nodes in groups of 10", func(b *testing.B) {

				for j := range tc.values {
					sll.Append(j)
				}

				ReverseGroups(sll.Head, tc.k)
			})
		}
	}
}
