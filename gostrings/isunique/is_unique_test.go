package isunique

import (
	"fmt"
	"testing"
)

type testCase struct {
	input    string
	expected bool
}

// TestCases for the IsUnique function
var testCases = []testCase{
	{"", true},
	{"a", true},
	{"aa", false},
	{"abc", true},
	{"abca", false},
	{"abcdefghijklmnopqrstuvwxyz", true},
	{"abcdefghijklmnopqrstuvwxyza", false},
	{"1234567890", true},
	{"12345678901", false},
	{"abCDefGh", true},
	{"nonunique", false},
	{"abCedFghI", true},
	{"I Am Not Unique", false},
	{"heythere", false},
	{"hi", true},
}

func TestIsUnique(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsUnique(%s)", tc.input), func(t *testing.T) {
			result := IsUnique(tc.input)
			if result != tc.expected {
				t.Errorf("For input %q, expected %v but got %v", tc.input, tc.expected, result)
			}
		})
	}
}

func BenchmarkIsUnique(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode.")
	}

	for _, tc := range testCases {
		for i := 0; i < b.N; i++ {
			b.Run(fmt.Sprintf("IsUnique(%s)", tc.input), func(b *testing.B) {
				result := IsUnique(tc.input)
				if result != tc.expected {
					b.Errorf("For input %q, expected %v but got %v", tc.input, tc.expected, result)
				}
			})
		}
	}
}
