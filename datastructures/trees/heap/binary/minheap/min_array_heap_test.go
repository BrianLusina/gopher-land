package minheap

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestMinArrayHeap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MinArrayHeap Test Suite")
}

var _ = Describe("MinArrayHeap", func() {
	Context("Int Min Heap", func() {
		heap := NewMinArrayHeap[int]()
		data := []int{100, 88, 25, 87, 16, 8, 12, 86, 50, 2, 15, 3}

		When("inserting new data", func() {
			It("should add and balance the heap appropriately", func() {
				for _, d := range data {
					heap.Insert(d)
				}

				actualRoot, err := heap.RootNode()
				assert.NoError(GinkgoT(), err)

				expectedRoot := 2
				Expect(actualRoot, expectedRoot)

				actualLast, err := heap.LastNode()
				assert.NoError(GinkgoT(), err)

				expectedLast := 100
				Expect(actualLast, expectedLast)
			})
		})

		When("deleting existing data", func() {
			It("should delete only the root node and set a new root node according to the heap condition", func() {
				actualRoot, err := heap.Delete()
				assert.NoError(GinkgoT(), err)

				expectedRoot := 100
				Expect(actualRoot, expectedRoot)

				newActualRoot, err := heap.RootNode()
				assert.NoError(GinkgoT(), err)

				expectedNewRoot := 88
				Expect(newActualRoot, expectedNewRoot)

				actualLast, err := heap.LastNode()
				assert.NoError(GinkgoT(), err)

				expectedLast := 12
				Expect(actualLast, expectedLast)
			})
		})
	})
})
