package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCases struct {
	list     []int
	expected []int
}

var tests = []testCases{
	{
		list:     []int{},
		expected: []int{},
	},
	{
		list:     []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func TestMergeSort(t *testing.T) {
	for _, test := range tests {
		actual := MergeSort(test.list)
		result := assert.Equal(t, test.expected, actual)
		if !result {
			t.Errorf("FAIL: \nMergeSort(%v)\nreturned %v, expected %v.", test.list, actual, test.expected)
		}
	}
}

func TestMergeSortConcurrent(t *testing.T) {
	for _, test := range tests {
		actual := MergeSortConcurrent(test.list)
		result := assert.Equal(t, test.expected, actual)
		if !result {
			t.Errorf("FAIL: \nMergeSort(%v)\nreturned %v, expected %v.", test.list, actual, test.expected)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			MergeSort(test.list)
		}
	}
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			MergeSortConcurrent(test.list)
		}
	}
}
