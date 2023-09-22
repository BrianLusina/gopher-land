package platesbetweencandles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s        string
	queries  [][]int
	expected []int
}

var testCases = []testCase{
	{
		s:        "**|**|***|",
		queries:  [][]int{{2, 5}, {5, 9}},
		expected: []int{2, 3},
	},
	{
		s:        "***|**|*****|**||**|*",
		queries:  [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
		expected: []int{9, 0, 0, 0, 0},
	},
	{
		s:        "||**||**|*",
		queries:  [][]int{{3, 8}},
		expected: []int{2},
	},
}

func TestPlatesBetweenCandles(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("platesBetweenCandles(s=%s, queries=%v)", tc.s, tc.queries), func(t *testing.T) {
			actual := platesBetweenCandles(tc.s, tc.queries)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func BenchmarkPlatesBetweenCandles(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("platesBetweenCandles(s=%s, queries=%v)", tc.s, tc.queries), func(b *testing.B) {
				actual := platesBetweenCandles(tc.s, tc.queries)
				assert.ElementsMatch(b, tc.expected, actual)
			})
		}
	}
}
