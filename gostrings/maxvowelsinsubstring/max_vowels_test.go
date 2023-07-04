package maxvowelsinsubstring

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	k        int
	expected int
}

var testCases = []testCase{
	{
		s:        "abciiidef",
		k:        3,
		expected: 3,
	},
	{
		s:        "aeiou",
		k:        2,
		expected: 2,
	},
	{
		s:        "leetcode",
		k:        3,
		expected: 2,
	},
}

func TestMaxVowels(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("maxVowels(%s, %d)", tc.s, tc.k), func(t *testing.T) {
			actual := maxVowels(tc.s, tc.k)
			if actual != tc.expected {
				t.Errorf("expected %d, go %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMaxVowels(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("maxVowels(%s, %d)", tc.s, tc.k), func(b *testing.B) {
				actual := maxVowels(tc.s, tc.k)
				if actual != tc.expected {
					b.Errorf("expected %d, go %d", tc.expected, actual)
				}
			})
		}
	}
}
