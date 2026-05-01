package sortitemsbygroup

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n           int
	m           int
	group       []int
	beforeItems [][]int
	expected    []int
}

var testCases = []testCase{
	{
		n:           8,
		m:           2,
		group:       []int{-1, -1, 1, 0, 0, 1, 0, -1},
		beforeItems: [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}},
		expected:    []int{6, 3, 4, 5, 2, 0, 7, 1},
	},
}

func TestSortItemsByGroups(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("sortItems(%d, %d, %v, %v)", tc.n, tc.m, tc.group, tc.beforeItems), func(t *testing.T) {
			actual := sortItems(tc.n, tc.m, tc.group, tc.beforeItems)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkSortItemsByGroups(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			sortItems(tc.n, tc.m, tc.group, tc.beforeItems)
		}
	}
}
