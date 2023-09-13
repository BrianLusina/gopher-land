package kokoeatingbananas

import "testing"

type testCase struct {
	piles    []int
	h        int
	expected int
}

var testCases = []testCase{
	{
		piles:    []int{3, 6, 7, 11},
		h:        8,
		expected: 4,
	},
	{
		piles:    []int{30, 11, 23, 4, 20},
		h:        5,
		expected: 30,
	},
	{
		piles:    []int{30, 11, 23, 4, 20},
		h:        6,
		expected: 23,
	},
}

func TestMinEatingSpeed(t *testing.T) {
	for _, tc := range testCases {
		actual := minEatingSpeed(tc.piles, tc.h)
		if actual != tc.expected {
			t.Fatalf("minEatingSpeed(%v, %d) = %d expected %d", tc.piles, tc.h, actual, tc.expected)
		}
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := minEatingSpeed(tc.piles, tc.h)
			if actual != tc.expected {
				b.Fatalf("minEatingSpeed(%v, %d) = %d expected %d", tc.piles, tc.h, actual, tc.expected)
			}
		}
	}
}
