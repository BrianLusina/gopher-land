package bubblesort

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
