package shellsort

import (
	"fmt"
	"gopherland/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T types.Comparable] struct {
	collection []T
	expected   []T
}

var intTestCases = []testCase[int]{
	{
		collection: []int{26, 17, 20, 11, 23, 21, 13, 18, 24, 14, 12, 22, 16, 15, 19, 25},
		expected:   []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26},
	},
}

func TestShellSortInts(t *testing.T) {
	for _, tc := range intTestCases {
		t.Run(fmt.Sprintf("ShellSort(%v)", tc.collection), func(t *testing.T) {
			actual := ShellSort(tc.collection)

			if !assert.ElementsMatch(t, tc.expected, actual) {
				t.Fatalf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func BenchmarkShellSortInts(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range intTestCases {
			b.Run(fmt.Sprintf("ShellSort(%v)", tc.collection), func(b *testing.B) {
				actual := ShellSort(tc.collection)

				if !assert.ElementsMatch(b, tc.expected, actual) {
					b.Fatalf("expected %v, got %v", tc.expected, actual)
				}
			})
		}
	}
}
