package graphs

import (
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGraphUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graph Utils Test Suite")
}

var _ = Describe("CloneGraph", func() {
	It("should clone a graph and return a new graph", func() {
		nodeOne := NewVertex(1)
		// nodeTwo := NewVertex(2)
		// nodeThree := NewVertex(3)
		// nodeFour := NewVertex(4)

		// nodeOne.Neighbors = map[any]*Node{nodeTwo.Data: nodeTwo, nodeFour.Data: nodeFour}
		// nodeTwo.Neighbors = map[any]*Node{nodeThree.Data: nodeThree, nodeOne.Data: nodeOne}
		// nodeThree.Neighbors = map[any]*Node{nodeFour.Data: nodeFour, nodeTwo.Data: nodeTwo}
		// nodeFour.Neighbors = map[any]*Node{nodeOne.Data: nodeOne, nodeThree.Data: nodeThree}

		actual := CloneGraph(nodeOne)

		Expect(reflect.ValueOf(actual)).ToNot(Equal(reflect.ValueOf(nodeOne)))
	})
})
