package permutation

import (
	"fmt"
	"testing"
)

type testCase struct {
	str1     string
	str2     string
	expected bool
}

var testCases = []testCase{
	{"", "", true},
	{"a", "a", true},
	{"a", "b", false},
	{"ab", "ba", true},
	{"abc", "bca", true},
	{"abc", "bcd", false},
	{"abc", "abcd", false},
	{"abc", "ab", false},
	{"abc", "ac", false},
	{"abc", "bc", false},
	{"abc", "cd", false},
	{"abc", "abd", false},
	{"abc", "acb", true},
	{"abc", "bac", true},
	{"abc", "cab", true},
	{"google", "ooggle", true},
	{"google", "oogleg", true},
	{"not", "top", false},
}

func TestCheckPermutationWithSorting(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("CheckPermutationWithSorting(%s, %s)", tc.str1, tc.str2), func(t *testing.T) {
			if result := CheckPermutationWithSorting(tc.str1, tc.str2); result != tc.expected {
				t.Errorf("CheckPermutationWithSorting(%s, %s) = %v, expected %v", tc.str1, tc.str2, result, tc.expected)
			}
		})
	}
}

func BenchmarkCheckPermutationWithSorting(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("CheckPermutationWithSorting(%s, %s)", tc.str1, tc.str2), func(b *testing.B) {
				if result := CheckPermutationWithSorting(tc.str1, tc.str2); result != tc.expected {
					b.Errorf("CheckPermutationWithSorting(%s, %s) = %v, expected %v", tc.str1, tc.str2, result, tc.expected)
				}
			})
		}
	}
}

func TestCheckPermutationWithMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("CheckPermutationWithMap(%s, %s)", tc.str1, tc.str2), func(t *testing.T) {
			if result := CheckPermutationWithMap(tc.str1, tc.str2); result != tc.expected {
				t.Errorf("CheckPermutationWithMap(%s, %s) = %v, expected %v", tc.str1, tc.str2, result, tc.expected)
			}
		})
	}
}

func BenchmarkCheckPermutationWithMap(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("CheckPermutationWithMap(%s, %s)", tc.str1, tc.str2), func(b *testing.B) {
				if result := CheckPermutationWithMap(tc.str1, tc.str2); result != tc.expected {
					b.Errorf("CheckPermutationWithMap(%s, %s) = %v, expected %v", tc.str1, tc.str2, result, tc.expected)
				}
			})
		}
	}
}
