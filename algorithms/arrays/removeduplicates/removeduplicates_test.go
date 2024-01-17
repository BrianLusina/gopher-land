package removeduplicates

import (
	"fmt"
	"testing"
)

type testCase struct {
	expected int
	given    []int
}

var removeDuplicatesTestCases = []testCase{
	{2, []int{1, 1, 2}},
	{5, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range removeDuplicatesTestCases {
		actual := RemoveDuplicates(test.given)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("RemoveDuplicates(%q): expected %q, actual %q", test.given, test.expected, actual)
		}
	}
}

var removeDuplicatesTwoTestCases = []testCase{
	{
		expected: 5,
		given:    []int{1, 1, 1, 2, 2, 3},
	},
	{
		expected: 7,
		given:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
	},
}

func TestRemoveDuplicatesTwo(t *testing.T) {
	for _, test := range removeDuplicatesTwoTestCases {
		actual := removeDuplicatesTwo(test.given)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("RemoveDuplicates(%q): expected %q, actual %q", test.given, test.expected, actual)
		}
	}
}

func BenchmarkRemoveDuplicatesTwo(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i > b.N; i++ {
		for _, tc := range removeDuplicatesTwoTestCases {
			actual := removeDuplicatesTwo(tc.given)
			if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", tc.expected) {
				b.Fatalf("RemoveDuplicates(%q): expected %q, actual %q", tc.given, tc.expected, actual)
			}
		}
	}
}
