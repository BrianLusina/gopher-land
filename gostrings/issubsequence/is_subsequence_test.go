package issubsequence

import (
	"fmt"
	"testing"
)

type testCase struct {
	s        string
	t        string
	expected bool
}

var testCases = []testCase{
	{
		s:        "abc",
		t:        "ahbgdc",
		expected: true,
	},
	{
		s:        "axc",
		t:        "ahbgdc",
		expected: false,
	},
	{
		s:        "",
		t:        "ahbgdc",
		expected: true,
	},
	{
		s:        "b",
		t:        "abc",
		expected: true,
	},
}

func TestIsSubsequence(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsSubsequence(%s, %s)", tc.s, tc.t), func(t *testing.T) {
			actual := IsSubsequence(tc.s, tc.t)

			if actual != tc.expected {
				t.Fatal(fmt.Sprintf("IsSubsequence(%s, %s) = %v, expected %v", tc.s, tc.t, actual, tc.expected))
			}
		})
	}
}
