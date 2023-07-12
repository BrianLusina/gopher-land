package uniqueoccurences

import (
	"fmt"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

type testCase struct {
	arr      []int
	expected bool
}

var testCases = []testCase{
	{
		arr:      []int{1, 2, 2, 1, 1, 3},
		expected: true,
	},
	{
		arr:      []int{1, 2},
		expected: false,
	},
	{
		arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
		expected: true,
	},
	{
		arr:      []int{1, 2, 4, 3, 1, 3, 3, 4},
		expected: false,
	},
	{
		arr:      []int{1, 2, 4, 3, 1, 3, 3, 4, 4, 4},
		expected: true,
	},
	{
		arr:      []int{-1, -1, -1, -3, -1, -1},
		expected: true,
	},
}

func TestValidPath(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "UniqueOccurrences Suite")
}

var _ = ginkgo.Describe("UniqueOccurrences", func() {
	for _, tc := range testCases {
		ginkgo.It(fmt.Sprintf("uniqueOccurrences(%v) should return %v", tc.arr, tc.expected), func() {
			actual := uniqueOccurrences(tc.arr)
			gomega.Expect(actual).To(gomega.Equal(tc.expected))
		})
	}
})

func BenchmarkUniqueOccurrences(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := uniqueOccurrences(tc.arr)
			if actual != tc.expected {
				b.Errorf("expected %v, got %v", tc.expected, actual)
			}
		}
	}
}
