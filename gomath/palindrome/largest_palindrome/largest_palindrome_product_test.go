package palindromeproduct

import "testing"

type testCase struct {
	description string
	input       int
	expected    int
}

var testCases = []testCase{
	{
		description: "largest palindrome product of 2 n-digit numbers",
		input:       2,
		expected:    9009,
	},
	{
		description: "largest palindrome product of 3 n-digit numbers",
		input:       3,
		expected:    906609,
	},
	{
		description: "largest palindrome product of 4 n-digit numbers",
		input:       4,
		expected:    99000099,
	},
}

func TestLargestPalindromeProduct(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			actual := largestPalindromeProduct(test.input)
			if actual != test.expected {
				t.Fatalf("expected %d, got %d", test.expected, actual)
			}

		})
	}
}

func BenchmarkMaxAreaOfIsland(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			largestPalindromeProduct(test.input)
		}
	}
}
