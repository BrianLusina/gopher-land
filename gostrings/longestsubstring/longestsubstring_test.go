package longestsubstring

import (
	"log"
	"testing"
)

type testCase struct {
	input    string
	expected int
}

var testCases []testCase = []testCase{
	{
		input: "abcabcbb",
		expected: 3,
	},
	{
		input: "bbbbb",
		expected: 1,
	},
	{
		input: "pwwkew",
		expected: 3,
	},
}

func TestLongestSubstring(t *testing.T) {
	for _, tc := range testCases {
		actual := lengthOfLongestSubstring(tc.input)
		if actual != tc.expected {
			log.Fatalf("lengthOfLongestSubstring(%s) = %d, expected = %d", tc.input, actual, tc.expected)
		}
	}
}
