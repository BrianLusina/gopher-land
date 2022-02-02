package maxsubarray

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaxSubarray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Max Subarray Test Suite")
}

var _ = Describe("Max Subarray Tests", func() {
	It("Should return 6 for nums = [-2,1,-3,4,-1,2,1,-5,4]", func() {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		actual := maxSubArray(nums)
		expected := 6
		Expect(actual).To(Equal(expected))
	})

	It("Should return 1 for nums = [1]", func() {
		nums := []int{1}
		actual := maxSubArray(nums)
		expected := 1
		Expect(actual).To(Equal(expected))
	})

	It("Should return 23 for nums = [5,4,-1,7,8]", func() {
		nums := []int{5, 4, -1, 7, 8}
		actual := maxSubArray(nums)
		expected := 23
		Expect(actual).To(Equal(expected))
	})
})
