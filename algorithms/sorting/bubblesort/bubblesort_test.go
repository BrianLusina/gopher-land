package bubblesort

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BubbleSort Suite")
}

var _ = Describe("BubbleSort", func() {
	It("Should sort a given int slice in ascending order", func() {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		actual := BubbleSort(input)
		Expect(actual).To(Equal(expected))
	})

	It("Should sort a given float slice in ascending order", func() {
		input := []float64{5.0, 4.3, 3.3, 2.2, 1.1}
		expected := []float64{1.1, 2.2, 3.3, 4.3, 5.0}
		actual := BubbleSort(input)
		Expect(actual).To(Equal(expected))
	})
})

type testCase struct {
	input    []int
	expected []int
}

var testCases = []testCase{
	{
		input:    []int{3, 4, 1, 8, 3, 5, 3, 5},
		expected: []int{1, 3, 3, 3, 4, 5, 5, 8},
	},
}

func TestBubbleSort1(t *testing.T) {
	for _, tc := range testCases {
		actual := BubbleSort(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestBubbleSort2(t *testing.T) {
	for _, tc := range testCases {
		actual := bubbleSort(tc.input)
		assert.Equal(t, tc.expected, actual)
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			BubbleSort(tc.input)
		}
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			bubbleSort(tc.input)
		}
	}
}
