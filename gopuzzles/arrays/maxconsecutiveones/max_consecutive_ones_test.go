package maxconsecutiveones

import (
	"fmt"
	"testing"
)

type testCase struct {
	nums     []int
	expected int
}

type longestOnesTestCase struct {
	testCase
	k int
}

var longestOnesTestCases = []longestOnesTestCase{
	{
		testCase: testCase{
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			expected: 6,
		},
		k: 2,
	},
	{
		testCase: testCase{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			expected: 10,
		},
		k: 3,
	},
}

func TestLongestOnes(t *testing.T) {
	for _, tc := range longestOnesTestCases {
		t.Run(fmt.Sprintf("longestOnes(%v, %d)", tc.nums, tc.k), func(t *testing.T) {
			actual := longestOnes(tc.nums, tc.k)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkLongestOnes(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range longestOnesTestCases {
			b.Run(fmt.Sprintf("longestOnes(%v, %d)", tc.nums, tc.k), func(b *testing.B) {
				actual := longestOnes(tc.nums, tc.k)
				if actual != tc.expected {
					b.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}

var maxConsecutiveOnesTestCases = []testCase{
	{
		nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		expected: 4,
	},
	{
		nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		expected: 4,
	},
}

func TestMaxConsecutiveOnes(t *testing.T) {
	for _, tc := range maxConsecutiveOnesTestCases {
		t.Run(fmt.Sprintf("findMaxConsecutiveOnes(%v)", tc.nums), func(t *testing.T) {
			actual := findMaxConsecutiveOnes(tc.nums)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMaxConsecutiveOnes(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range maxConsecutiveOnesTestCases {
			b.Run(fmt.Sprintf("findMaxConsecutiveOnes(%v)", tc.nums), func(b *testing.B) {
				actual := findMaxConsecutiveOnes(tc.nums)
				if actual != tc.expected {
					b.Errorf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
