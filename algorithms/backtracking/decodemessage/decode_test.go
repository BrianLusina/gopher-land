package decodemessage

import "testing"

type testCase struct {
	digits   string
	expected int
}

var testCases = []testCase{
	{
		digits:   "12",
		expected: 2,
	},
	{
		digits:   "123",
		expected: 3,
	},
	{
		digits:   "11223",
		expected: 8,
	},
	{
		digits:   "313",
		expected: 3,
	},
	{
		digits:   "02",
		expected: 0,
	},
}

func TestNumberOfWaysToDecodeMessage(t *testing.T) {
	for _, tc := range testCases {
		actual := numberOfWaysToDecodeMessage(tc.digits)
		if actual != tc.expected {
			t.Fatalf(" numberOfWaysToDecodeMessage(%s)=%d, expected=%d", tc.digits, actual, tc.expected)
		}
	}
}

func BenchmarkNumberOfWaysToDecodeMessage(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := numberOfWaysToDecodeMessage(tc.digits)
			if actual != tc.expected {
				b.Fatalf(" numberOfWaysToDecodeMessage(%s)=%d, expected=%d", tc.digits, actual, tc.expected)
			}
		}
	}
}
