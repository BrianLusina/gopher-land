package disjointset

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDisjointSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DisjointSet Test Suite")
}

var _ = Describe("DisjointSet", func() {
	It("should create a new disjoint set with the given nodes", func() {
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)
		nodeThree := NewNode(3)
		nodeFour := NewNode(4)
		nodeFive := NewNode(5)

		nodes := []Node[int]{*nodeOne, *nodeTwo, *nodeThree, *nodeFour, *nodeFive}

		disjointSet := NewDisjointSet(nodes...)

		Expect(disjointSet.Find(nodeOne)).ToNot(BeNil())
		Expect(disjointSet.Find(nodeOne)).To(Equal(nodeOne))

		Expect(disjointSet.Find(nodeTwo)).ToNot(BeNil())
		Expect(disjointSet.Find(nodeTwo)).To(Equal(nodeTwo))

		Expect(disjointSet.Find(nodeThree)).ToNot(BeNil())
		Expect(disjointSet.Find(nodeThree)).To(Equal(nodeThree))

		Expect(disjointSet.Find(nodeFour)).ToNot(BeNil())
		Expect(disjointSet.Find(nodeFour)).To(Equal(nodeFour))

		Expect(disjointSet.Find(nodeFive)).ToNot(BeNil())
		Expect(disjointSet.Find(nodeFive)).To(Equal(nodeFive))

		Expect(disjointSet.size).To(Equal(len(nodes)))
	})

	It("should create a union between nodes", func() {
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)
		nodeThree := NewNode(3)
		nodeFour := NewNode(4)
		nodeFive := NewNode(5)

		nodes := []Node[int]{*nodeOne, *nodeTwo, *nodeThree, *nodeFour, *nodeFive}

		disjointSet := NewDisjointSet(nodes...)

		Expect(disjointSet.size).To(Equal(len(nodes)))

		disjointSet.Union(nodeOne, nodeTwo)
		Expect(disjointSet.Connected(nodeOne, nodeTwo)).To(BeTrue())
	})

})
