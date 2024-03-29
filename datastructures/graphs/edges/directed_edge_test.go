package edges

import (
	"gopherland/datastructures/graphs"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirectedEdge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DirectedEdge Test Suite")
}

var _ = Describe("DirectedEdge", func() {
	It("should create a new  edge with 0 weight and 2 nodes/vertices", func() {
		source := graphs.NewVertex(1)
		destination := graphs.NewVertex(2)

		edge := NewDirectedEdge(source, destination, 0)

		Expect(edge).ToNot(BeNil())
		Expect(edge.GetWeight()).To(Equal(0))
		Expect(edge.GetNodes()).To(Equal([]*graphs.Vertex[int]{source, destination}))

		Expect(source.GetDegree()).To(Equal(1))
		Expect(destination.GetDegree()).To(Equal(1))
	})

	// It("should add multiple edges to existing edge to create a hyper edge", func() {
	// 	nodeOne := NewNode(1)
	// 	nodeTwo := NewNode(2)

	// 	edgeOne := NewEdge(nodeOne, nodeTwo, 0)

	// 	Expect(edgeOne).ToNot(BeNil())
	// 	Expect(edgeOne.GetWeight()).To(Equal(0))
	// 	Expect(edgeOne.GetNodes()).To(Equal([]*Node{nodeOne, nodeTwo}))
	// 	Expect(nodeOne.Degree).To(Equal(1))
	// 	Expect(nodeTwo.Degree).To(Equal(1))

	// 	nodeThree := NewNode(3)
	// 	nodeFour := NewNode(4)

	// 	edgeTwo := NewEdge(nodeThree, nodeFour, 0)

	// 	Expect(edgeTwo).ToNot(BeNil())
	// 	Expect(edgeTwo.GetWeight()).To(Equal(0))
	// 	Expect(edgeTwo.GetNodes()).To(Equal([]*Node{nodeThree, nodeFour}))
	// 	Expect(nodeThree.Degree).To(Equal(1))
	// 	Expect(nodeFour.Degree).To(Equal(1))

	// 	edgeOne.AddEdge(edgeTwo)

	// 	Expect(edgeOne.GetWeight()).To(Equal(0))
	// 	Expect(edgeOne.GetNodes()).To(Equal([]*Node{nodeOne, nodeTwo, nodeThree, nodeFour}))
	// })
})
