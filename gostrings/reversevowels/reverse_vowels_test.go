package reversevowels

import "testing"

type testCase struct {
	word     string
	expected string
}

var testCases = []testCase{
	{
		word:     "hello",
		expected: "holle",
	},
	{
		word:     "leetcode",
		expected: "leotcede",
	},
}

func TestReverseVowels(t *testing.T) {
	for _, tc := range testCases {
		actual := reverseVowels(tc.word)
		if actual != tc.expected {
			t.Fatalf("reverseVowels(%s) = %s, expected = %s", tc.word, actual, tc.expected)
		}
	}
}

func BenchmarkReverseVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := reverseVowels(tc.word)
			if actual != tc.expected {
				b.Fatalf("reverseVowels(%s) = %s, expected = %s", tc.word, actual, tc.expected)
			}
		}
	}
}
