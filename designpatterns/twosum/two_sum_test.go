package twosum

import "testing"

type testCase struct {
	name     string
	numbers  []int
	find     int
	expected bool
}

var testCases = []testCase{
	{
		name:     "Test Case 1",
		numbers:  []int{1, 2, 3, 4, 5},
		find:     7,
		expected: true,
	},
	{
		name:     "Test Case 3",
		numbers:  []int{1, 2, 3, 4, 5},
		find:     5,
		expected: true,
	},
}

func TestTwoSum(t *testing.T) {
	twoSum := New()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, number := range tc.numbers {
				twoSum.Add(number)
			}
			result := twoSum.Find(tc.find)
			if result != tc.expected {
				t.Errorf("Find(%d) = %v, want %v", tc.find, result, tc.expected)
			}
		})
	}
}
