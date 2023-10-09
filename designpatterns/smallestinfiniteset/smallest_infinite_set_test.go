package smallestinfiniteset

import (
	"gopherland/pkg/utils"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestSmallestInfiniteSet(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SmallestInfiniteSet Test Suite")
}

var _ = ginkgo.Describe("SmallestInfiniteSet", func() {
	ginkgo.When("Operations are AddBack(2) - PopSmallest() - PopSmallest() - PopSmallest() - AddBack(1) - PopSmallest() - PopSmallest() - PopSmallest()", func() {
		smallestInfiniteSet := New()

		ginkgo.It("should perform operations in sequence correctly", func() {
			// 2 is already in the set, so no change is made.
			smallestInfiniteSet.AddBack(2)

			// return 1, since 1 is the smallest number, and remove it from the set.
			actual1 := smallestInfiniteSet.PopSmallest()
			expected1 := 1
			assert.Equal(ginkgo.GinkgoT(), expected1, actual1)

			// return 2, and remove it from the set.
			actual2 := smallestInfiniteSet.PopSmallest()
			expected2 := 2
			assert.Equal(ginkgo.GinkgoT(), expected2, actual2)

			// return 3, and remove it from the set.
			actual3 := smallestInfiniteSet.PopSmallest()
			expected3 := 3
			assert.Equal(ginkgo.GinkgoT(), expected3, actual3)

			// 1 is added back to set
			smallestInfiniteSet.AddBack(1)

			// return 1, since 1 was added back to the set and is the smallest number, and remove it from the set.
			actual1_1 := smallestInfiniteSet.PopSmallest()
			expected1_1 := 1
			assert.Equal(ginkgo.GinkgoT(), expected1_1, actual1_1)

			// return 4, and remove it from the set.
			smallestInfiniteSet.PopSmallest()
		})
	})

	ginkgo.When("PopSmallest is called 1000 times", func() {
		smallestInfiniteSet := New()

		ginkgo.It("should perform operations in sequence correctly", func() {
			intRange := utils.Range(1, 1001, 1)

			for _, x := range intRange {
				actual := smallestInfiniteSet.PopSmallest()
				assert.Equal(ginkgo.GinkgoT(), x, actual)
			}
		})
	})
})
