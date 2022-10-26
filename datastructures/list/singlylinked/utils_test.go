package singlylinkedlist

import (
	"gopherland/datastructures/list"
	"reflect"
	"testing"
)

type testCase struct {
	nodeOne  *list.Node[int]
	nodeTwo  *list.Node[int]
	expected *list.Node[int]
}

var testCases = []testCase{
	{
		nodeOne: &list.Node[int]{Data: 1, Next: &list.Node[int]{Data: 2, Next: &list.Node[int]{Data: 4}}},
		nodeTwo: &list.Node[int]{Data: 1, Next: &list.Node[int]{Data: 3, Next: &list.Node[int]{Data: 4}}},
		expected: &list.Node[int]{
			Data: 1, Next: &list.Node[int]{
				Data: 1, Next: &list.Node[int]{
					Data: 2, Next: &list.Node[int]{
						Data: 3, Next: &list.Node[int]{
							Data: 4, Next: &list.Node[int]{
								Data: 4}}}}}},
	},
}

func TestMergeTwoSortedLists(t *testing.T) {
	for _, tc := range testCases {
		actual := mergeTwoSortedLists(tc.nodeOne, tc.nodeTwo)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("mergeTwoSortedLists(%v, %v) = %v, expected = %v", tc.nodeOne, tc.nodeTwo, actual, tc.expected)
		}
	}
}
