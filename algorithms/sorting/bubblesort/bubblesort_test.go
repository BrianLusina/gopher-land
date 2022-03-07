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
	It("Should sort the given array in ascending order", func() {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		actual := BubbleSort(input)
		Expect(actual).To(Equal(expected))
	})
})
