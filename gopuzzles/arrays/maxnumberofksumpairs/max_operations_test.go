package maxnumberofksumpairs

import "testing"

type testCase struct {
	nums     []int
	k        int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{1, 2, 3, 4},
		k:        5,
		expected: 2,
	},
	{
		nums:     []int{3, 1, 3, 4, 3},
		k:        6,
		expected: 1,
	},
	{
		nums:     []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4},
		k:        2,
		expected: 2,
	},
}

func TestMaxOperations(t *testing.T) {
	for _, tc := range testCases {
		actual := maxOperations(tc.nums, tc.k)
		if actual != tc.expected {
			t.Fatalf("maxOperations(%v, %d)=%d, expected %d", tc.nums, tc.k, actual, tc.expected)
		}
	}
}

func TestMaxOperationsWithMap(t *testing.T) {
	for _, tc := range testCases {
		actual := maxOperationsWithMap(tc.nums, tc.k)
		if actual != tc.expected {
			t.Fatalf("maxOperations(%v, %d)=%d, expected %d", tc.nums, tc.k, actual, tc.expected)
		}
	}
}

func BenchmarkMaxOperations(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := maxOperations(tc.nums, tc.k)
			if actual != tc.expected {
				b.Fatalf("maxOperations(%v, %d)=%d, expected %d", tc.nums, tc.k, actual, tc.expected)
			}
		}
	}
}

func BenchmarkMaxOperationsWithMap(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := maxOperationsWithMap(tc.nums, tc.k)
			if actual != tc.expected {
				b.Fatalf("maxOperations(%v, %d)=%d, expected %d", tc.nums, tc.k, actual, tc.expected)
			}
		}
	}
}
