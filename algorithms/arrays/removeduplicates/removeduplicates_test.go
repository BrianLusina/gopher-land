package removeduplicates

import (
	"fmt"
	"testing"
)

var tests = []struct {
	expected int
	given    []int
}{
	{2, []int{1, 1, 2}},
	{5, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		actual := RemoveDuplicates(test.given)
		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("RemoveDuplicates(%q): expected %q, actual %q", test.given, test.expected, actual)
		}
	}
}
