package decimaltobinary

import "testing"

type testCase struct {
	num      int
	expected int
}

var testCases = []testCase{
	{
		num:      56,
		expected: 111000,
	},
	{
		num:      2,
		expected: 10,
	},
	{
		num:      32,
		expected: 100000,
	},
	{
		num:      10,
		expected: 1010,
	},
}

func TestIntToBin(t *testing.T) {
	for _, tc := range testCases {
		actual := convertIntToBin(tc.num)
		if actual != tc.expected {
			t.Fatalf("convertIntToBin(%d) = %d, expected = %d", tc.num, actual, tc.expected)
		}
	}
}

func BenchmarkIntToBin(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := convertIntToBin(tc.num)
			if actual != tc.expected {
				b.Fatalf("convertIntToBin(%d) = %d, expected = %d", tc.num, actual, tc.expected)
			}
		}
	}
}
