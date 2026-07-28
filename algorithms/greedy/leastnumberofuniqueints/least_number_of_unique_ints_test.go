package leastnumberofuniqueints

import (
	"fmt"
	"testing"
)

type testCase struct {
	arr      []int
	k        int
	expected int
}

var testCases = []testCase{
	{arr: []int{5, 5, 4}, k: 1, expected: 1},
	{arr: []int{4, 3, 1, 1, 3, 3, 2}, k: 3, expected: 2},
	{arr: []int{1, 2, 3, 4, 5}, k: 5, expected: 0},
	{arr: []int{1, 2, 2, 3, 3}, k: 2, expected: 2},
	{arr: []int{1}, k: 0, expected: 1},
	{arr: []int{2, 4, 1, 8, 3, 5, 1, 3}, k: 3, expected: 3},
	{arr: []int{1, 1, 1, 2, 2, 3, 3, 3, 3}, k: 4, expected: 2},
	{arr: []int{5, 5, 5, 5, 6, 7, 8}, k: 2, expected: 2},
}

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("findLeastNumOfUniqueInts(arr: %v, k: %d)", tc.arr, tc.k), func(t *testing.T) {
			result := findLeastNumOfUniqueInts(tc.arr, tc.k)
			if result != tc.expected {
				t.Errorf("For arr: %v and k: %d, expected: %d, but got: %d", tc.arr, tc.k, tc.expected, result)
			}
		})
	}
}
