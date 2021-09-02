package findifpathexists

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidPath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FindIfPathExistsInGraph Suite")
}

var _ = Describe("FindIfPathExistsInGraph", func() {
	It("should return true for n = 3, edges = [[0,1],[1,2],[2,0]], start = 0, end = 2", func() {
		n := 3
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		start := 0
		end := 2

		actual := validPath(n, edges, start, end)
		Expect(actual).To(Equal(true))
	})

	It("should return false for n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], start = 0, end = 5", func() {
		n := 6
		edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
		start := 0
		end := 5

		actual := validPath(n, edges, start, end)
		Expect(actual).To(Equal(false))
	})

	It("should return true for n = 1, edges = [], start = 0, end = 0", func() {
		n := 0
		edges := [][]int{}
		start := 0
		end := 0

		actual := validPath(n, edges, start, end)
		Expect(actual).To(Equal(true))
	})
})
