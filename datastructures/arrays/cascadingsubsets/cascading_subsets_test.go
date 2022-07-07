package cascadingsubsets

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func referenceSolution(arr []int, n int) (o [][]int) {
	for i := 0; i < len(arr)-n+1; i++ {
		var a []int
		for j := i; j < i+n; j++ {
			a = append(a, arr[j])
		}
		o = append(o, a)
	}
	return
}

func TestCascadingSubsets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cascading Subsets Suite")
}

var _ = Describe("CascadingSubsets", func() {
	It("should return [[3], [5], [8], [13]] from [3,5,8,13] with size of 1", func() {
		size := 1
		arr := []int{3, 5, 8, 13}
		expected := [][]int{{3}, {5}, {8}, {13}}
		Expect(EachCons(arr, size)).To(Equal(expected))
	})

	It("should return [[3, 5], [5, 8], [8, 13]] from [3,5,8,13] with size of 2", func() {
		size := 2
		arr := []int{3, 5, 8, 13}
		expected := [][]int{{3, 5}, {5, 8}, {8, 13}}
		Expect(EachCons(arr, size)).To(Equal(expected))

		Expect(EachCons(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
		Expect(EachCons([]int{}, 3)).To(BeEmpty())
	})

	It("should return [[3, 5, 8], [5, 8, 13]] from [3,5,8,13] with size of 3", func() {
		size := 3
		arr := []int{3, 5, 8, 13}
		expected := [][]int{{3, 5, 8}, {5, 8, 13}}
		Expect(EachCons(arr, size)).To(Equal(expected))
	})

	It("should return [] from [] with size of 3", func() {
		size := 3
		arr := []int{}
		expected := [][]int{}
		Expect(EachCons(arr, size)).To(Equal(expected))
	})

	for i := 0; i < 100; i++ {
		It(fmt.Sprintf("Random test %d", i), func() {
			l, n := rand.Intn(11), 1+rand.Intn(10)
			arr := make([]int, l)
			for j := 0; j < l; j++ {
				arr[j] = rand.Intn(21)
			}
			expected, actual := referenceSolution(arr, n), EachCons(arr, n)

			if len(expected) > 0 {
				Expect(actual).To(Equal(expected))
			} else {
				Expect(actual).To(BeEmpty())
			}
		})
	}
})
