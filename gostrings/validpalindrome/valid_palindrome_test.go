package validpalindrome

import "testing"

type testCase struct {
	phrase   string
	expected bool
}

var testCases = []testCase{
	{
		phrase:   "A man, a plan, a canal: Panama",
		expected: true,
	},
	{
		phrase:   "race a car",
		expected: false,
	},
	{
		phrase:   " ",
		expected: true,
	},
}

func TestValidPalindrome(t *testing.T) {
	for _, tc := range testCases {
		actual := isPalindrome(tc.phrase)
		if actual != tc.expected {
			t.Fatalf("isPalindrome(%s) = %v, expected %v", tc.phrase, actual, tc.expected)
		}
	}
}

func BenchmarkValidPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			isPalindrome(tt.phrase)
		}
	}
}

func TestValidPalindromeTwoPointers(t *testing.T) {
	for _, tc := range testCases {
		actual := isPalindromeTwoPointers(tc.phrase)
		if actual != tc.expected {
			t.Fatalf("isPalindrome(%s) = %v, expected %v", tc.phrase, actual, tc.expected)
		}
	}
}

func BenchmarkValidPalindromeTwoPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			isPalindromeTwoPointers(tt.phrase)
		}
	}
}
