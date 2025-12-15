package longesthappystrings

import (
	"fmt"
	"testing"
)

type testCase struct {
	a              int
	b              int
	c              int
	expected       string
	expectedLength int
}

var testCases = []testCase{
	{
		a:              7,
		b:              1,
		c:              0,
		expected:       "aabaa",
		expectedLength: 5,
	},
	{
		a:              2,
		b:              2,
		c:              2,
		expected:       "abcabc",
		expectedLength: 6,
	},
	{
		a:              0,
		b:              5,
		c:              5,
		expected:       "bcbcbcbcbc",
		expectedLength: 10,
	},
	{
		a:              6,
		b:              3,
		c:              0,
		expected:       "aabaabaab",
		expectedLength: 9,
	},
	{
		a:              3,
		b:              3,
		c:              1,
		expected:       "abababc",
		expectedLength: 7,
	},
	{
		a:              2,
		b:              2,
		c:              1,
		expected:       "ababc",
		expectedLength: 5,
	},
	{
		a:              5,
		b:              1,
		c:              0,
		expected:       "aabaa",
		expectedLength: 5,
	},
	{
		a:              7,
		b:              2,
		c:              0,
		expected:       "aabaabaa",
		expectedLength: 8,
	},
	{
		a:              1,
		b:              1,
		c:              7,
		expected:       "ccaccbcc",
		expectedLength: 8,
	},
}

func TestLongestDiverseString(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("longestDiverseString(a=%d,b=%d,c=%d)", tc.a, tc.b, tc.c), func(t *testing.T) {
			result := longestDiverseString(tc.a, tc.b, tc.c)
			if valid, reason := isValidHappyString(result, tc.a, tc.b, tc.c); !valid {
				t.Errorf("Invalid happy string: %s. Got: %s", reason, result)
			}
			if len(result) != tc.expectedLength {
				t.Errorf("Expected length %d but got %d. String: %s", tc.expectedLength, len(result), result)
			}
		})
	}
}

func BenchmarkLongestDiverseString(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			longestDiverseString(tc.a, tc.b, tc.c)
		}
	}
}
