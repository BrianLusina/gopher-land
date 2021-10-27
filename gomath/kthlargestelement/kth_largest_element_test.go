package kthlargestelement

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKthLargestElement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KthLargestElement Suite")
}

var _ = Describe("KthLargestElement", func() {
	It("should return 5 for nums = [3,2,1,5,6,4] and k = 2", func() {
		k := 2
		nums := []int{3, 2, 1, 5, 6, 4}
		expected := 5
		Expect(findKthLargest(nums, k)).To(Equal(expected))
	})

	It("should return 4 for nums = [3,2,3,1,2,4,5,5,6] and k = 4", func() {
		k := 4
		nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
		expected := 4
		Expect(findKthLargest(nums, k)).To(Equal(expected))
	})
})
