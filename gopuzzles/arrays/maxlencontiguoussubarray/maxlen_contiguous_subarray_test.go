package maxlencontiguoussubarray

import (
	"gopherland/math/utils"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaxLenContiguousSubArrayFromBinaryArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Max Length of Contiguous Sub Array in Binary Array Suite")
}

func randArr(minLen, maxLen int) []int {
	k := 1.0 / float64(1+rand.Intn(3))
	arr := make([]int, minLen+rand.Intn(maxLen-minLen+1))
	for i, _ := range arr {
		if rand.Float64() > k {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}
	return arr
}

func referenceSolution(a []int) int {
	f := func(n int) int {
		if n == 1 {
			return 1
		} else {
			return -1
		}
	}
	h := make(map[int]int)
	var s, m int
	for i, n := range a {
		s += f(n)
		if s == 0 {
			m = i + 1
		} else {
			if v, ok := h[s]; ok {
				m = utils.Max(m, i-v)
			} else {
				h[s] = i
			}
		}
	}
	return m
}

var _ = Describe("Max Length of Contiguous Sub Array in Binary Array ", func() {

	It("test one with 0, 1", func() {
		nums := []int{0, 1}
		expected := 2
		actual := FindMaxLengthOfBinarySubArray(nums)

		Expect(actual).To(Equal(expected))
	})

	It("test two with 0, 1 ,1", func() {
		nums := []int{0, 1, 0}
		expected := 2
		actual := FindMaxLengthOfBinarySubArray(nums)

		Expect(actual).To(Equal(expected))
	})

	It("test two with [0, 1, 0, 0, 1, 1, 0]", func() {
		nums := []int{0, 1, 0, 0, 1, 1, 0}
		expected := 6
		actual := FindMaxLengthOfBinarySubArray(nums)

		Expect(actual).To(Equal(expected))
	})

	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 100; i++ {
			arr := randArr(10, 20)
			expected := referenceSolution(arr)
			actual := FindMaxLengthOfBinarySubArray(arr)

			Expect(actual).To(Equal(expected), "With a = %v", arr)
		}

		for i := 0; i < 100; i++ {
			arr := randArr(50, 100)
			expected := referenceSolution(arr)
			actual := FindMaxLengthOfBinarySubArray(arr)
			Expect(actual).To(Equal(expected), "With a = %v", arr)
		}

		for i := 0; i < 100; i++ {
			arr := randArr(50000, 100000)
			expected := referenceSolution(arr)
			actual := FindMaxLengthOfBinarySubArray(arr)
			Expect(actual).To(Equal(expected), "With a = %v", arr)
		}
	})
})
