package mergestrings

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
		word1:    "abc",
		word2:    "pqr",
		expected: "apbqcr",
	},
	{
		word1:    "ab",
		word2:    "pqrs",
		expected: "apbqrs",
	},
	{
		word1:    "abcd",
		word2:    "pq",
		expected: "apbqcd",
	},
	{
		word1:    "cf",
		word2:    "eee",
		expected: "cefee",
	},
}

func TestMergeStringsAlternatively(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should merge %s and %s to form %s", tc.word1, tc.word2, tc.expected), func(t *testing.T) {
			actual := MergeAlternately(tc.word1, tc.word2)
			if actual != tc.expected {
				t.Fatalf("Expected %s to equal %s", actual, tc.expected)
			}
		})
	}
}

func BenchmarkMergeStringsAlternatively(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("Should merge %s and %s to form %s", tc.word1, tc.word2, tc.expected), func(b *testing.B) {
				actual := MergeAlternately(tc.word1, tc.word2)
				if actual != tc.expected {
					b.Fatalf("Expected %s to equal %s", actual, tc.expected)
				}
			})
		}
	}
}

func TestMergeStringsAlternatively2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Should merge %s and %s to form %s", tc.word1, tc.word2, tc.expected), func(t *testing.T) {
			actual := MergeAlternately2(tc.word1, tc.word2)
			if actual != tc.expected {
				t.Fatalf("Expected %s to equal %s", actual, tc.expected)
			}
		})
	}
}

func BenchmarkMergeStringsAlternatively2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("Should merge %s and %s to form %s", tc.word1, tc.word2, tc.expected), func(b *testing.B) {
				actual := MergeAlternately2(tc.word1, tc.word2)
				if actual != tc.expected {
					b.Fatalf("Expected %s to equal %s", actual, tc.expected)
				}
			})
		}
	}
}
