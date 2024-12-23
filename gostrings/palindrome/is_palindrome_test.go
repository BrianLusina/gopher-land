package palindrome

import "testing"

type isPalindromeTestCase struct {
	phrase   string
	expected bool
}

var isPalindromeCases = []isPalindromeTestCase{
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

func TestIsPalindrome(t *testing.T) {
	for _, tc := range isPalindromeCases {
		actual := isPalindrome(tc.phrase)
		if actual != tc.expected {
			t.Fatalf("isPalindrome(%s) = %v, expected %v", tc.phrase, actual, tc.expected)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range isPalindromeCases {
			isPalindrome(tt.phrase)
		}
	}
}

func TestIsPalindromeTwoPointers(t *testing.T) {
	for _, tc := range isPalindromeCases {
		actual := isPalindromeTwoPointers(tc.phrase)
		if actual != tc.expected {
			t.Fatalf("isPalindrome(%s) = %v, expected %v", tc.phrase, actual, tc.expected)
		}
	}
}

func BenchmarkIsPalindromeTwoPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range isPalindromeCases {
			isPalindromeTwoPointers(tt.phrase)
		}
	}
}
