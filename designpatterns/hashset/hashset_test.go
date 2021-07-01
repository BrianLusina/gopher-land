package hashset

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHashSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HashSet Suite")
}

var _ = Describe("HashSet", func() {
	hashset := Constructor()

	It("should add 1 to hashset", func() {
		hashset.Add(1)
		Expect(hashset.Map).To(Equal(map[int]bool{1: true}))
	})

	It("should add 2 to hashset", func() {
		hashset.Add(2)
		expected := make(map[int]bool)
		expected[1] = true
		expected[2] = true
		Expect(hashset.Map).To(Equal(expected))
	})

	It("contains 1 should return true", func() {
		actual := hashset.Contains(1)
		Expect(actual).To(Equal(true))
	})

	It("contains 3 should return false", func() {
		actual := hashset.Contains(3)
		Expect(actual).To(Equal(false))
	})

	It("should not add another 2 to hashset", func() {
		hashset.Add(2)
		expected := make(map[int]bool)
		expected[1] = true
		expected[2] = true
		Expect(hashset.Map).To(Equal(expected))
	})

	It("should remove 2 from hashset", func() {
		hashset.Remove(2)
		expected := make(map[int]bool)
		expected[1] = true
		Expect(hashset.Map).To(Equal(expected))
	})
})
