package palindrome

import "testing"

type isPalindromePermutationTestCase struct {
	phrase   string
	expected bool
}

var isPalindromePermutationCases = []isPalindromePermutationTestCase{
	{
		phrase:   "Tact Coa",
		expected: true,
	},
	{
		phrase:   "This is not a palindrome permutation",
		expected: false,
	},
	{
		phrase:   "taco cat",
		expected: true,
	},
}

func TestIsPalindromePermutation(t *testing.T) {
	for _, tc := range isPalindromePermutationCases {
		actual := isPalindromePermutation(tc.phrase)
		if actual != tc.expected {
			t.Fatalf("isPalindromePermutation(%s) = %v, expected %v", tc.phrase, actual, tc.expected)
		}
	}
}

func BenchmarkIsPalindromePermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range isPalindromePermutationCases {
			_ = isPalindromePermutation(tt.phrase)
		}
	}
}
