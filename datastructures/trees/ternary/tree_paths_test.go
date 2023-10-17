package ternary

import (
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type treePathTestCase[T types.Comparable] struct {
	tree     *TenaryTree[T]
	expected []string
}

var treePathIntTestCases = []treePathTestCase[int]{
	{
		tree: &TenaryTree[int]{
			root: &TernaryTreeNode[int]{
				data: 1,
				children: []*TernaryTreeNode[int]{
					{
						data: 2,
						children: []*TernaryTreeNode[int]{
							{
								data: 3,
							},
						},
					},
					{
						data:     4,
						children: []*TernaryTreeNode[int]{},
					},
					{
						data:     5,
						children: []*TernaryTreeNode[int]{},
					},
				},
			},
		},
		expected: []string{"1->2->3", "1->4", "1->5"},
	},
}

// TestPaths tests the paths from root to leaf nodes
func TestPaths(t *testing.T) {
	for _, tc := range treePathIntTestCases {
		actual := tc.tree.Paths()
		assert.Equal(t, tc.expected, actual)
	}
}
