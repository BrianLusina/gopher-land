package topkfrequent

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	numbers  []int
	k        int
	expected []int
}{
	{
		numbers:  []int{1, 1, 1, 2, 2, 3},
		k:        2,
		expected: []int{1, 2},
	},
	{
		numbers:  []int{1},
		k:        1,
		expected: []int{1},
	},
}

func TestTopKFrequent(t *testing.T) {
	for _, tc := range testCases {
		actual := topKFrequent(tc.numbers, tc.k)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("topKFrequent(%v, %d) = %v, expected = %v", tc.numbers, tc.k, actual, tc.expected)
		} else {
			t.Logf("PASS: Numbers = %v, output: %v", tc.numbers, actual)
		}
	}
}
