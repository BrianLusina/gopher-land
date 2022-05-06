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
		nodeOne := NewNode(1)
		nodeTwo := NewNode(2)
		nodeThree := NewNode(3)
		nodeFour := NewNode(4)

		nodeOne.Neighbors = map[any]*Node{nodeTwo.Data: nodeTwo, nodeFour.Data: nodeFour}
		nodeTwo.Neighbors = map[any]*Node{nodeThree.Data: nodeThree, nodeOne.Data: nodeOne}
		nodeThree.Neighbors = map[any]*Node{nodeFour.Data: nodeFour, nodeTwo.Data: nodeTwo}
		nodeFour.Neighbors = map[any]*Node{nodeOne.Data: nodeOne, nodeThree.Data: nodeThree}

		actual := CloneGraph(nodeOne)

		Expect(reflect.ValueOf(actual)).ToNot(Equal(reflect.ValueOf(nodeOne)))
	})
})
