package binary

import "testing"

var isFullTestCases = []struct {
	name     string
	expected bool
	node     *BinaryTreeNode[int]
}{
	{
		name:     "Node with no children returns true",
		node:     NewBinaryTreeNode(1),
		expected: true,
	},
	{
		name: "Node with 1 left node and 1 right node returns true",
		node: func() *BinaryTreeNode[int] {
			root := NewBinaryTreeNode(1)
			left := NewBinaryTreeNode(2)
			right := NewBinaryTreeNode(3)
			root.Left = left
			root.Right = right
			return root
		}(),
		expected: true,
	},
	{
		name: "Node with 1 left node 0 right node returns false",
		node: func() *BinaryTreeNode[int] {
			root := NewBinaryTreeNode(1)
			left := NewBinaryTreeNode(2)
			root.Left = left
			return root
		}(),
		expected: false,
	},
}

func TestBinaryTreeNodeIsFull(t *testing.T) {
	for _, tc := range isFullTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.node.IsFull()

			if actual != tc.expected {
				t.Fatalf("IsFull(): %v, expected: %v", actual, tc.expected)
			}
		})
	}
}

func BenchmarkBinaryTreeNodeIsFull(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range isFullTestCases {
		for i := 0; i < b.N; i++ {
			b.Run(tc.name, func(b *testing.B) {
				actual := tc.node.IsFull()

				if actual != tc.expected {
					b.Fatalf("IsFull(): %v, expected: %v", actual, tc.expected)
				}
			})
		}
	}
}
