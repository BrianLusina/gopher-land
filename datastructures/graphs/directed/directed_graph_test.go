package directed

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDirectedGraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DirectedGraph Test Suite")
}

var _ = Describe("DirectedGraph", func() {
	It("should add a new node to the graph successfully", func() {
		g := NewDirectedGraph()
		node := NewNode(1)
		g.AddNode(node)

		Expect(g.Size()).To(Equal(1))
	})

	It("should not add a new node when there is an existing node with the same value", func() {
		g := NewDirectedGraph()
		nodeOne := NewNode(1)
		nodeTwo := NewNode(1)

		errOne := g.AddNode(nodeOne)
		Expect(errOne).To(BeNil())

		errTwo := g.AddNode(nodeTwo)
		Expect(errTwo).ToNot(BeNil())

		Expect(g.Size()).To(Equal(1))
	})

	It("should add new nodes in succession", func() {
		g := NewDirectedGraph()
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)
		nodeThree := NewNode(3)

		_ = g.AddNode(nodeOne)
		_ = g.AddNode(nodeTwo)
		_ = g.AddNode(nodeThree)

		Expect(g.Size()).To(Equal(3))
	})
})
