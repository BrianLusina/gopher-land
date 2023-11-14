package generatepermutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	letters  string
	expected []string
}

var testCases = []testCase{
	{
		letters:  "ab",
		expected: []string{"ab", "ba"},
	},
	{
		letters:  "abc",
		expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
	},
}

func TestGeneratePermutations(t *testing.T) {
	for _, tc := range testCases {
		actual := generatePermutations(tc.letters)
		assert.ElementsMatch(t, tc.expected, actual)
	}
}

func BenchmarkGeneratePermutations(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := generatePermutations(tc.letters)
			assert.ElementsMatch(b, tc.expected, actual)
		}
	}
}
