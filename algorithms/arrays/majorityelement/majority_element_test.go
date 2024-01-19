package majorityelement

import "testing"

type testCase struct {
	nums     []int
	expected int
}

var testCases = []testCase{
	{
		nums:     []int{3, 2, 3},
		expected: 3,
	},
	{
		nums:     []int{2, 2, 1, 1, 1, 2, 2},
		expected: 2,
	},
}

func TestMajorityElementWithSorting(t *testing.T) {
	for _, tc := range testCases {
		actual := majorityElementWithSorting(tc.nums)
		if actual != tc.expected {
			t.Fatalf("majorityElementWithSorting(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkMajorityElementWithSorting(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := majorityElementWithSorting(tc.nums)
			if actual != tc.expected {
				b.Fatalf("majorityElementWithSorting(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
			}
		}
	}
}

func TestMajorityElementWithMap(t *testing.T) {
	for _, tc := range testCases {
		actual := majorityElementWithMap(tc.nums)
		if actual != tc.expected {
			t.Fatalf("majorityElementWithMap(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkMajorityElementWithMap(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := majorityElementWithMap(tc.nums)
			if actual != tc.expected {
				b.Fatalf("majorityElementWithMap(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
			}
		}
	}
}

func TestMajorityElementWithVoting(t *testing.T) {
	for _, tc := range testCases {
		actual := majorityElementWithVoting(tc.nums)
		if actual != tc.expected {
			t.Fatalf("majorityElementWithVoting(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
		}
	}
}

func BenchmarkMajorityElementWithVoting(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := majorityElementWithVoting(tc.nums)
			if actual != tc.expected {
				b.Fatalf("majorityElementWithVoting(nums=%v)=%d, expected=%d", tc.nums, actual, tc.expected)
			}
		}
	}
}
