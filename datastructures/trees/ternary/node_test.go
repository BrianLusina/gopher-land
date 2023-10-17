package ternary

import (
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type newNodeTestCase[T types.Comparable] struct {
	data     T
	children []*TernaryTreeNode[T]
	expected *TernaryTreeNode[T]
}

var newIntNodeTestCases = []newNodeTestCase[int]{
	{
		data:     1,
		children: []*TernaryTreeNode[int]{},
		expected: &TernaryTreeNode[int]{
			data:     1,
			children: []*TernaryTreeNode[int]{},
		},
	},
}

// TestNewTernaryNode tests that a new Ternary Node can be created with the given arguments
func TestIntNewTernaryNode(t *testing.T) {

	for _, tc := range newIntNodeTestCases {
		actual := NewTernaryTreeNode(tc.data, Children(tc.children))
		assert.Equal(t, tc.expected, actual)
	}
}
