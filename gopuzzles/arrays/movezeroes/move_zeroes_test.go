package movezeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	expected []int
}

var testCases = []testCase{
	{
		nums:     []int{0, 1, 0, 3, 12},
		expected: []int{1, 3, 12, 0, 0},
	},
	{
		nums:     []int{0},
		expected: []int{0},
	},
	{
		nums:     []int{1, 0},
		expected: []int{1, 0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, tc := range testCases {
		moveZeroes(tc.nums)
		if !assert.ElementsMatch(t, tc.nums, tc.expected) {
			t.Fatalf("moveZeroes(%v) should be %v", tc.nums, tc.expected)
		}
	}
}

func TestMoveZeroesOne(t *testing.T) {
	for _, tc := range testCases {
		moveZeroesOne(tc.nums)
		if !assert.ElementsMatch(t, tc.nums, tc.expected) {
			t.Fatalf("moveZeroesOne(%v) should be %v", tc.nums, tc.expected)
		}
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			moveZeroes(tc.nums)
			if !assert.ElementsMatch(b, tc.nums, tc.expected) {
				b.Fatalf("moveZeroes(%v) should be %v", tc.nums, tc.expected)
			}
		}
	}
}

func BenchmarkMoveZeroesOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			moveZeroesOne(tc.nums)
			if !assert.ElementsMatch(b, tc.nums, tc.expected) {
				b.Fatalf("moveZeroesOne(%v) should be %v", tc.nums, tc.expected)
			}
		}
	}
}
