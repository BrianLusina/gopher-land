package dynamicarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n        int
	queries  [][]int
	expected []int
}

var testCases = []testCase{
	{
		n:        2,
		queries:  [][]int{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}},
		expected: []int{7, 3},
	},
}

func TestDynamicArray(t *testing.T) {
	for _, tc := range testCases {
		actual := dynamicArray(tc.n, tc.queries)

		if !assert.Equal(t, tc.expected, actual) {
			t.Fatalf("dynamicArray(%d, %v) = %v, expected = %v", tc.n, tc.queries, actual, tc.expected)
		}
	}
}

func BenchmarkDynamicArray(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			dynamicArray(tc.n, tc.queries)
		}
	}
}
