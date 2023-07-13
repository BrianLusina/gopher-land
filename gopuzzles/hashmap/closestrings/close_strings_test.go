package closestrings

import (
	"fmt"
	"testing"
)

type testCase struct {
	word1    string
	word2    string
	expected bool
}

var testCases = []testCase{
	{
		word1:    "abc",
		word2:    "bca",
		expected: true,
	},
	{
		word1:    "a",
		word2:    "aa",
		expected: false,
	},
	{
		word1:    "cabbba",
		word2:    "abbccc",
		expected: true,
	},
}

func TestCloseStrings(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("closeStrings(%s, %s)", tc.word1, tc.word2), func(t *testing.T) {
			actual := closeStrings(tc.word1, tc.word2)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkCloseStrings(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("closeStrings(%s, %s)", tc.word1, tc.word2), func(b *testing.B) {
				actual := closeStrings(tc.word1, tc.word2)
				if actual != tc.expected {
					b.Errorf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
