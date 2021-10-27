package kclosest

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKClosest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KClosest Suite")
}

var _ = Describe("KClosest", func() {
	It("should return [[-2, 2]] for points = [[1,3], [-2,2]] and k = 1", func() {
		points := [][]int{{1, 3}, {-2, 2}}
		k := 1
		expected := [][]int{{-2, 2}}
		Expect(kClosest(points, k)).To(Equal(expected))
	})

	It("should return [[3, 3], [-2, 4]] for points = [[3,3], [5,-1], [-2,4]] and k = 2", func() {
		points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
		k := 2
		expected := [][]int{{3, 3}, {-2, 4}}
		Expect(kClosest(points, k)).To(Equal(expected))
	})
})
