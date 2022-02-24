package containsduplicates

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contains Nearby Duplicate Test Suite")
}

var _ = Describe("Contains Nearby Duplicate", func() {
	It("should return true for nums = [1,2,3,1] and k = 3", func() {
		nums := []int{1, 2, 3, 1}
		k := 3
		expected := true
		actual := ContainsNearbyDuplicate(nums, k)
		Expect(actual).To(Equal(expected))
	})
	It("should return true for nums = [1,0,1,1] and k = 1", func() {
		nums := []int{1, 0, 1, 1}
		k := 1
		expected := true
		actual := ContainsNearbyDuplicate(nums, k)
		Expect(actual).To(Equal(expected))
	})
	It("should return false for nums = [1,2,3,1,2,3] and k = 2", func() {
		nums := []int{1, 2, 3, 1, 2, 3}
		k := 2
		expected := false
		actual := ContainsNearbyDuplicate(nums, k)
		Expect(actual).To(Equal(expected))
	})
})
