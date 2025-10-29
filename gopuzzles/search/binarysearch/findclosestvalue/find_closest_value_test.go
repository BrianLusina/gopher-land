package findclosestvalue

import (
	"gopherland/datastructures/trees/binary"
	"testing"
)

type testCase struct {
	name     string
	target   int
	expected int
	node     *binary.BinaryTreeNode[int]
}

var testCases = []testCase{
	{
		name:     "closest value is present in the tree",
		target:   7,
		expected: 5,
		node: func() *binary.BinaryTreeNode[int] {
			root := binary.NewBinaryTreeNode(10)
			root.SetLeft(binary.NewBinaryTreeNode(5))
			root.SetRight(binary.NewBinaryTreeNode(15))
			return root
		}(),
	},
	{
		name:     "closest value is not present in the tree",
		target:   8,
		expected: 10,
		node: func() *binary.BinaryTreeNode[int] {
			root := binary.NewBinaryTreeNode(10)
			root.SetLeft(binary.NewBinaryTreeNode(5))
			root.SetRight(binary.NewBinaryTreeNode(15))
			return root
		}(),
	},
	{
		name:     "single node tree",
		target:   3,
		expected: 10,
		node:     binary.NewBinaryTreeNode(10),
	},
	{
		name:     "empty tree",
		target:   3,
		expected: 0,
		node:     nil,
	},
}

func TestFindClosestValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findClosestValueInBst(tc.node, tc.target)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkFindClosestValue(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			findClosestValueInBst(tc.node, tc.target)
		}
	}
}
