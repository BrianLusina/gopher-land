package countbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n        int
	expected []int
}

var testCases = []testCase{
	{
		n:        2,
		expected: []int{0, 1, 1},
	},
	{
		n:        5,
		expected: []int{0, 1, 1, 2, 1, 2},
	},
}

func TestCountBits(t *testing.T) {
	for _, tc := range testCases {
		actual := countBits(tc.n)
		assert.ElementsMatch(t, actual, tc.expected)
	}
}

func BenchmarkCountBits(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := countBits(tc.n)
			assert.ElementsMatch(b, actual, tc.expected)
		}
	}
}
