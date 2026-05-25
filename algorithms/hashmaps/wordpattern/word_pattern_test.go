package wordpattern

import (
	"fmt"
	"testing"
)

type testCase struct {
	pattern  string
	s        string
	expected bool
}

var testCases = []testCase{
	{"abba", "dog cat cat dog", true},
	{"abba", "dog dog dog dog", false},
	{"abba", "dog cat cat fish", false},
	{"aaaa", "dog cat cat dog", false},
	{"abc", "red blue green", true},
	{"abcba", "sun moon star moon sun", true},
}

func TestWordPattern(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("wordPattern(pattern=%s, s=%s)", tc.pattern, tc.s), func(t *testing.T) {
			actual := wordPattern(tc.pattern, tc.s)
			if actual != tc.expected {
				t.Errorf("For pattern=%s and s=%s, expected %v but got %v", tc.pattern, tc.s, tc.expected, actual)
			}
		})
	}
}

func TestWordPattern2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("wordPattern2(pattern=%s, s=%s)", tc.pattern, tc.s), func(t *testing.T) {
			actual := wordPattern2(tc.pattern, tc.s)
			if actual != tc.expected {
				t.Errorf("For pattern=%s and s=%s, expected %v but got %v", tc.pattern, tc.s, tc.expected, actual)
			}
		})
	}
}

func BenchmarkWordPattern(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("wordPattern(pattern=%s, s=%s)", tc.pattern, tc.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wordPattern(tc.pattern, tc.s)
			}
		})
	}
}

func BenchmarkWordPattern2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for _, tc := range testCases {
		b.Run(fmt.Sprintf("wordPattern2(pattern=%s, s=%s)", tc.pattern, tc.s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				wordPattern2(tc.pattern, tc.s)
			}
		})
	}
}
