package graphs

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEdge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Edge Test Suite")
}

var _ = Describe("Edge", func() {
	It("should create a new edge with 0 weight and 2 nodes/vertices", func() {
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)

		edge := NewEdge(nodeOne, nodeTwo, 0)

		Expect(edge).ToNot(BeNil())
		Expect(edge.GetWeight()).To(Equal(0))
		Expect(edge.GetNodes()).To(Equal([]*Node{nodeOne, nodeTwo}))
	})

	It("should add multiple edges to existing edge to create a hyper edge", func() {
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)

		edgeOne := NewEdge(nodeOne, nodeTwo, 0)

		Expect(edgeOne).ToNot(BeNil())
		Expect(edgeOne.GetWeight()).To(Equal(0))
		Expect(edgeOne.GetNodes()).To(Equal([]*Node{nodeOne, nodeTwo}))

		nodeThree := NewNode(3)
		nodeFour := NewNode(4)

		edgeTwo := NewEdge(nodeThree, nodeFour, 0)

		Expect(edgeTwo).ToNot(BeNil())
		Expect(edgeTwo.GetWeight()).To(Equal(0))
		Expect(edgeTwo.GetNodes()).To(Equal([]*Node{nodeThree, nodeFour}))

		edgeOne.AddEdge(edgeTwo)

		Expect(edgeOne.GetWeight()).To(Equal(0))
		Expect(edgeOne.GetNodes()).To(Equal([]*Node{nodeOne, nodeTwo, nodeThree, nodeFour}))
	})
})
