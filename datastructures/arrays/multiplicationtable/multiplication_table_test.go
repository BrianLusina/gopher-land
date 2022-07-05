package multiplicationtable

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestMultiplicationTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MultiplicationTable Suite")
}

func reference(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			table[i][j] = (i + 1) * (j + 1)
		}
	}
	return table
}

var _ = Describe("Sample Tests", func() {
	It("Table of size 1", func() {
		size := 1
		expected := [][]int{{1}}
		actual := MultiplicationTable(size)

		Expect(actual).To(Equal(expected))
	})

	It("Table of size 2", func() {
		size := 2
		expected := [][]int{{1, 2}, {2, 4}}
		actual := MultiplicationTable(size)

		Expect(actual).To(Equal(expected))
	})

	It("Table of size 3", func() {
		size := 3
		expected := [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
		actual := MultiplicationTable(size)

		Expect(actual).To(Equal(expected))
	})

	for i := 0; i < 100; i++ {
		It(fmt.Sprintf("Random Test %d", i), func() {
			randomSize := rand.Intn(99) + 1
			expected := reference(randomSize)
			actual := MultiplicationTable(randomSize)
			Expect(actual).To(Equal(expected))
		})
	}
})
