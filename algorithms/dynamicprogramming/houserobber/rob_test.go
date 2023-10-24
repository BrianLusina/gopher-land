package houserobber

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 2, 3, 1},
		expected: 4,
	},
	{
		nums:     []int{2, 7, 9, 3, 1},
		expected: 12,
	},
	{
		nums:     []int{},
		expected: 0,
	},
	{
		nums:     []int{3},
		expected: 3,
	},
	{
		nums:     []int{3, 5},
		expected: 5,
	},
}

func TestRob(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("rob(nums=%v)", tc.nums), func(t *testing.T) {
			actual := rob(tc.nums)
			if actual != tc.expected {
				t.Fail()
			}
		})
	}
}

func BenchmarkRob(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := rob(tc.nums)
			if actual != tc.expected {
				b.Fail()
			}
		}
	}
}
