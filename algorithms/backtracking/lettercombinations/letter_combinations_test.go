package lettercombinations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n        int
	expected []string
}

var testCases = []testCase{
	{
		n:        1,
		expected: []string{"a", "b"},
	},
	{
		n:        2,
		expected: []string{"aa", "ab", "ba", "bb"},
	},
	{
		n:        3,
		expected: []string{"aaa", "aab", "aba", "abb", "baa", "bab", "bba", "bbb"},
	},
	{
		n:        0,
		expected: []string{},
	},
}

func TestLetterCombination(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("letterCombination(n=%d)", tc.n), func(t *testing.T) {
			actual := letterCombinations(tc.n)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}

func BenchmarkLetterCombination(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := letterCombinations(tc.n)
			assert.ElementsMatch(b, tc.expected, actual)
		}
	}
}
