package disjointset

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNodeDisjointSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DisjointSet Node Test Suite")
}

var _ = Describe("DisjointSet Node", func() {
	It("should create a new node and make itself the parent", func() {
		node := NewNode(1)

		Expect(node.parent).To(Equal(node))
	})

	It("should create a new node and make the rank == 0", func() {
		node := NewNode(1)

		Expect(node.rank).To(Equal(0))
	})

})
