package jumpgame

import "testing"

type jumpTestCase struct {
	nums     []int
	expected int
}

var jumpTestCases = []jumpTestCase{
	{
		nums:     []int{2, 3, 1, 1, 4},
		expected: 2,
	},
	{
		nums:     []int{2, 3, 0, 1, 4},
		expected: 2,
	},
}

func TestJump(t *testing.T) {
	for _, tc := range jumpTestCases {
		actual := jump(tc.nums)
		if actual != tc.expected {
			t.Fatalf("jump(%v) = %v, should be %v", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkJump(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range jumpTestCases {
			actual := jump(tc.nums)
			if actual != tc.expected {
				b.Fatalf("jump(%v) = %v, should be %v", tc.nums, actual, tc.expected)
			}
		}
	}
}
