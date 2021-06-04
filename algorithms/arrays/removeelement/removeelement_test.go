package removeelement

import (
	"fmt"
	"testing"
)

var tests = []struct {
	given    []int
	val      int
	expected int
}{
	{[]int{3, 2, 2, 3}, 3, 2},
}

func TestRemoveElement(t *testing.T) {
	for _, test := range tests {
		actual := RemoveElement(test.given, test.val)

		if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", test.expected) {
			t.Fatalf("RemoveElement(%q, %q): expected %q, actual: %q", test.given, test.val, test.expected, actual)
		}
	}
}
