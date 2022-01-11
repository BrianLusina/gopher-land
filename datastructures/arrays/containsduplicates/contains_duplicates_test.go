package containsduplicates

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainsDuplicates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contains Duplicates Test Suite")
}

var _ = Describe("Contains Duplicates", func() {
	It("should return true for [1,2,3,1]", func() {
		nums := []int{1, 2, 3, 1}
		expected := true
		actual := ContainsDuplicate(nums)
		Expect(actual).Should(Equal(expected))
	})

	It("should return false for nums [1,2,3,4]", func() {
		cost := []int{1, 2, 3, 4}
		expected := false
		actual := ContainsDuplicate(cost)
		Expect(actual).Should(Equal(expected))
	})

	It("should return true for nums [1,1,1,3,3,4,3,2,4,2]", func() {
		cost := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		expected := true
		actual := ContainsDuplicate(cost)
		Expect(actual).Should(Equal(expected))
	})
})
