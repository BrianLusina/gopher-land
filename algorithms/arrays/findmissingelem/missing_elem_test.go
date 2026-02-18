package findmissingelem

import "testing"

type testCase struct {
	name     string
	numbers  []int
	expected int
}

var testCases = []testCase{
	{
		name:     "missing element is in the second position",
		numbers:  []int{0, 2, 3, 4, 5},
		expected: 1,
	},
	{
		name:     "missing element is the middle one",
		numbers:  []int{0, 1, 2, 3, 5, 6, 7},
		expected: 4,
	},
	{
		name:     "missing element is the last one one",
		numbers:  []int{0, 1, 2, 3, 4, 5, 6, 7},
		expected: 8,
	},
	{
		name:     "missing element is the first one",
		numbers:  []int{1},
		expected: 0,
	},
	{
		name:     "missing element is the second last one",
		numbers:  []int{1, 0, 2, 3, 4, 5, 6, 8, 9, 7, 11},
		expected: 10,
	},
	{
		name:     "missing element is 3",
		numbers:  []int{1, 4, 5, 6, 8, 2, 0, 7},
		expected: 3,
	},
	{
		name:     "missing element is 2",
		numbers:  []int{3, 0, 1, 4},
		expected: 2,
	},
	{
		name:     "missing element is 3",
		numbers:  []int{0, 1, 2, 4},
		expected: 3,
	},
	{
		name:     "missing element is 8",
		numbers:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		expected: 8,
	},
	{
		name:     "missing element is 2",
		numbers:  []int{0, 1},
		expected: 2,
	},
	{
		name:     "missing element is 3",
		numbers:  []int{3, 0, 1},
		expected: 2,
	},
	{
		name:     "missing element is 6",
		numbers:  []int{8, 2, 4, 5, 3, 7, 1, 0},
		expected: 6,
	},
	{
		name:     "missing element is 4",
		numbers:  []int{0, 1, 2, 3, 5},
		expected: 4,
	},
}

func TestFindMissingElement(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMissingElemXor(tc.numbers)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestFindMissingElementSumOfNTerms(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMissingElemSumOfNTerms(tc.numbers)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestFindMissingElement2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMissingElem2(tc.numbers)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkFindMissingElementXor(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			FindMissingElemXor(tc.numbers)
		}
	}
}

func BenchmarkFindMissingElementSumOfNTerms(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			FindMissingElemSumOfNTerms(tc.numbers)
		}
	}
}

func BenchmarkFindMissingElement2(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			FindMissingElem2(tc.numbers)
		}
	}
}
