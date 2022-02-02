package highestrank

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHighestRankNumberInArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Highest Rank Number Test Suite")
}

var _ = Describe("Highest Rank Number Tests", func() {
	It("Should return 12 for input of 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
		in := []int{12, 10, 8, 12, 7, 6, 4, 10, 12}
		actual := HighestRank(in)
		expected := 12
		Expect(actual).To(Equal(expected))
	})

	It("Should return 12 for input of: [12, 10, 8, 12, 7, 6, 4, 10, 12, 10]", func() {
		input := []int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}
		actual := HighestRank(input)
		expected := 12
		Expect(actual).To(Equal(expected))
	})

	It("Should return 3 for input of [12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]", func() {
		input := []int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}
		actual := HighestRank(input)
		expected := 3
		Expect(actual).To(Equal(expected))
	})

	It("Should return 63 for input of [224 241 63 229 252 172 132 178 222 236 183 169 63 232 110", func() {
		input := []int{224, 241, 63, 229, 252, 172, 132, 178, 222, 236, 183, 169, 63, 232, 110}
		actual := HighestRank(input)
		expected := 63
		Expect(actual).To(Equal(expected))
	})

	It("Should return 194 for input of [49 202 236 237 194 123 61 204 135 194 14 10 248 221 119", func() {
		input := []int{49, 202, 236, 237, 194, 123, 61, 204, 135, 194, 14, 10, 248, 221, 119}
		actual := HighestRank(input)
		expected := 194
		Expect(actual).To(Equal(expected))
	})

	It("Should return 200 for input of [194 194 250 188 198 224 200 31 254 54 99 118 30 239 233", func() {
		input := []int{194, 194, 250, 188, 198, 224, 200, 31, 254, 54, 99, 118, 30, 239, 233}
		actual := HighestRank(input)
		expected := 194
		Expect(actual).To(Equal(expected))
	})

	Describe("Random tests", func() {
		const testsCount = 50

		const arrayLength = 15

		for i := 1; i <= testsCount; i++ {
			rand.Seed(time.Now().UnixNano() + int64(i))

			randomNumbers := make([]int, arrayLength)

			for i := range randomNumbers {
				randomNumbers[i] = int(rand.Float64() * math.MaxUint8)
			}

			It(fmt.Sprintf("Random test %d: %v", i, randomNumbers), func() {
				randomNumbers[5] = randomNumbers[6]

				Expect(HighestRank(randomNumbers)).To(Equal(highestRank(randomNumbers)))
			})
		}
	})
})

func highestRank(nums []int) int {
	rankMap := make(map[int]int)

	for _, n := range nums {
		rankMap[n] += 1
	}

	max, highestRank := 0, 0

	for n, rank := range rankMap {
		if rank > highestRank || (rank == highestRank && n > max) {
			max, highestRank = n, rank
		}
	}

	return max
}
