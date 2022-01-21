package reshapematrix

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReshapeMatrix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reshape Matrix Test Suite")
}

var _ = Describe("Reshape Matrix", func() {
	It("Should return [[1,2,3,4]] for input[[1,2],[3,4]] and r=1 and c=4", func() {
		matrix := [][]int{{1, 2}, {3, 4}}
		r := 1
		c := 4
		expected := [][]int{{1, 2, 3, 4}}
		actual := matrixReshape(matrix, r, c)
		Expect(actual).To(Equal(expected))
	})

	It("Should return [[1,2],[3,4]] for input[[1,2],[3,4]] and r=2 and c=4", func() {
		matrix := [][]int{{1, 2}, {3, 4}}
		r := 2
		c := 4
		expected := [][]int{{1, 2}, {3, 4}}
		actual := matrixReshape(matrix, r, c)
		Expect(actual).To(Equal(expected))
	})
})
