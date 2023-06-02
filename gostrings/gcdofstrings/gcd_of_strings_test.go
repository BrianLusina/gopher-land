package gcdofstrings

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	word1    string
	word2    string
	expected string
}{
	{
		word1:    "ABCABC",
		word2:    "ABC",
		expected: "ABC",
	},
	{
		word1:    "ABABAB",
		word2:    "ABAB",
		expected: "AB",
	},
	{
		word1:    "LEET",
		word2:    "CODE",
		expected: "",
	},
}

func TestGcdOfStringsBruteForce(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("should return %s from str1=%s and str2=%s", tc.expected, tc.word1, tc.word2), func(t *testing.T) {
			actual := gcdOfStringsBruteForce(tc.word1, tc.word2)
			if actual != tc.expected {
				t.Logf("Expected %s, got: %s", tc.expected, actual)
				t.Fail()
			}
		})
	}
}

func BenchmarkGcdOfStringsBruteForce(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("should return %s from str1=%s and str2=%s", tc.expected, tc.word1, tc.word2), func(b *testing.B) {
				actual := gcdOfStringsBruteForce(tc.word1, tc.word2)
				if actual != tc.expected {
					b.Logf("Expected %s, got: %s", tc.expected, actual)
					b.Fail()
				}
			})
		}
	}
}
