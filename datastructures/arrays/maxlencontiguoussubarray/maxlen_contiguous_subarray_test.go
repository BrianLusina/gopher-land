package maxlencontiguoussubarray

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaxLenContiguousSubArrayFromBinaryArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Max Length of Contiguous Sub Array in Binary Array Suite")
}

var _ = Describe("Max Length of Contiguous Sub Array in Binary Array ", func() {

	It("test one with 0, 1", func() {
		nums := []int{0, 1}
		expected := 2
		actual := findMaxLength(nums)

		Expect(actual).To(Equal(expected))
	})

	It("test two with 0, 1 ,1", func() {
		nums := []int{0, 1, 0}
		expected := 2
		actual := findMaxLength(nums)

		Expect(actual).To(Equal(expected))
	})

	It("test two with [0, 1, 0, 0, 1, 1, 0]", func() {
		nums := []int{0, 1, 0, 0, 1, 1, 0}
		expected := 6
		actual := findMaxLength(nums)

		Expect(actual).To(Equal(expected))
	})
})
