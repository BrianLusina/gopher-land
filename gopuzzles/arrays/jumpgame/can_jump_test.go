package jumpgame

import "testing"

type testCase struct {
	nums     []int
	expected bool
}

var testCases = []testCase{
	{
		nums:     []int{2, 3, 1, 1, 4},
		expected: true,
	},
	{
		nums:     []int{3, 2, 1, 0, 4},
		expected: false,
	},
}

func TestCanJump(t *testing.T) {
	for _, tc := range testCases {
		actual := canJump(tc.nums)
		if actual != tc.expected {
			t.Fatalf("canJump(%v) = %v, should be %v", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkCanJump(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := canJump(tc.nums)
			if actual != tc.expected {
				b.Fatalf("canJump(%v) = %v, should be %v", tc.nums, actual, tc.expected)
			}
		}
	}
}
