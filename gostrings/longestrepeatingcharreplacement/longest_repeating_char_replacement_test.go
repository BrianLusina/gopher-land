package longestrepeatingcharreplacement

import "testing"

type testCase struct {
	input    string
	k        int
	expected int
}

var testCases = []testCase{
	{
		input:    "ABAB",
		k:        2,
		expected: 4,
	},
	{
		input:    "AABABBA",
		k:        1,
		expected: 4,
	},
	{
		input:    "ABCCDC",
		k:        1,
		expected: 4,
	},
}

func TestLongestRepeatingCharacterReplacement(t *testing.T) {
	for _, tc := range testCases {
		actual := characterReplacement(tc.input, tc.k)
		if actual != tc.expected {

		}
	}
}

func BenchmarkLongestRepeatingCharacterReplacement(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark")
	}

	for _, tc := range testCases {
		for i := 0; i <= b.N; i++ {
			characterReplacement(tc.input, tc.k)
		}
	}
}
