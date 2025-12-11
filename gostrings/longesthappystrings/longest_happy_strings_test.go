package longesthappystrings

import (
	"fmt"
	"testing"
)

type testCase struct {
	a        int
	b        int
	c        int
	expected string
}

var testCases = []testCase{
	{
		a:        7,
		b:        1,
		c:        0,
		expected: "aabaa",
	},
	{
		a:        2,
		b:        2,
		c:        2,
		expected: "abcabc",
	},
	{
		a:        0,
		b:        5,
		c:        5,
		expected: "bcbcbcbcbc",
	},
	{
		a:        6,
		b:        3,
		c:        0,
		expected: "aabaabaab",
	},
	{
		a:        3,
		b:        3,
		c:        1,
		expected: "abababc",
	},
	{
		a:        2,
		b:        2,
		c:        1,
		expected: "ababc",
	},
	{
		a:        5,
		b:        1,
		c:        0,
		expected: "aabaa",
	},
	{
		a:        7,
		b:        2,
		c:        0,
		expected: "aabaabaa",
	},
	{
		a:        1,
		b:        1,
		c:        7,
		expected: "ccaccbcc",
	},
}

func TestLongestDiverseString(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("longestDiverseString(a=%d,b=%d,c=%d)", tc.a, tc.b, tc.c), func(t *testing.T) {
			result := longestDiverseString(tc.a, tc.b, tc.c)
			if result != tc.expected {
				t.Errorf("For a=%d, b=%d, c=%d, expected %s but got %s", tc.a, tc.b, tc.c, tc.expected, result)
			}
		})
	}
}

func BenchmarkLongestDiverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			longestDiverseString(tc.a, tc.b, tc.c)
		}
	}
}
