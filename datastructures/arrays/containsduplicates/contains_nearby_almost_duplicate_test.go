package containsduplicates

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contains Nearby Duplicate Test Suite")
}

var _ = Describe("Contains Nearby Duplicate", func() {
	It("should return true for nums = [1,2,3,1] and k = 3 and t = 0", func() {
		nums := []int{1, 2, 3, 1}
		k := 3
		t := 0
		expected := true
		actual := ContainsNearbyAlmostDuplicate(nums, k, t)
		Expect(actual).To(Equal(expected))
	})
	It("should return true for nums = [1,0,1,1] and k = 1, t=2", func() {
		nums := []int{1, 0, 1, 1}
		k := 1
		t := 2
		expected := true
		actual := ContainsNearbyAlmostDuplicate(nums, k, t)
		Expect(actual).To(Equal(expected))
	})
	It("should return false for nums = [1,5,9,1,5,9] and k = 2, t=3", func() {
		nums := []int{1, 5, 9, 1, 5, 9}
		k := 2
		t := 3
		expected := false
		actual := ContainsNearbyAlmostDuplicate(nums, k, t)
		Expect(actual).To(Equal(expected))
	})
})
