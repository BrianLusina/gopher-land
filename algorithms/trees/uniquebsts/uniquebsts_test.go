package uniquebsts

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUniqueBinarySearchTrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unique Binary Search Trees Test Suite")
}

var _ = Describe("Unique Binary Search Trees", func() {
	It("Should return 5 for input of 3", func() {
		expected := 5
		actual := NumTrees(3)
		Expect(actual).To(Equal(expected))
	})
	It("Should return 1 for input of 1", func() {
		expected := 1
		actual := NumTrees(1)
		Expect(actual).To(Equal(expected))
	})
})
