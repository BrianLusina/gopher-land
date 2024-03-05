package candies

import "testing"

type testCase struct {
	ratings  []int
	expected int
}

var testCases = []testCase{
	{
		ratings:  []int{1, 0, 2},
		expected: 5,
	},
	{
		ratings:  []int{1, 2, 2},
		expected: 4,
	},
}

func TestCandy(t *testing.T) {
	for _, tc := range testCases {
		actual := candy(tc.ratings)
		if actual != tc.expected {
			t.Fatalf("candy(%v) = %d, expected = %d", tc.ratings, actual, tc.expected)
		}
	}
}

func BenchmarkCandy(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := candy(tc.ratings)
			if actual != tc.expected {
				b.Fatalf("candy(%v) = %d, expected = %d", tc.ratings, actual, tc.expected)
			}
		}
	}
}
