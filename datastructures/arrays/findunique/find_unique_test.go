package findunique

import (
	. "math"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFindUnique(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Find Unique Test Cases")
}

func generateTestArr(answer, mass float32, length int) []float32 {
	arr := make([]float32, length)
	answerIndex := rand.Intn(length - 1)
	for i := range arr {
		arr[i] = mass
		if i == answerIndex {
			arr[i] = answer
		}
	}

	return arr
}

var _ = Describe("FindUniq", func() {
	It("should return 2.0 for input of [1.0, 1.0, 1.0, 2.0, 1.0, 1.0]", func() {
		input := []float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}
		expected := float32(2)

		Expect(FindUniq(input)).To(Equal(expected))
	})

	It("should return 0.55 for inputs of [0, 0, 0.55, 0, 0]", func() {
		input := []float32{0, 0, 0.55, 0, 0}
		expected := float32(0.55)

		Expect(FindUniq(input)).To(Equal(expected))
	})

	It("should return 3 for input of [4, 4, 4, 3, 4, 4, 4, 4]", func() {
		input := []float32{4, 4, 4, 3, 4, 4, 4, 4}
		expected := float32(3)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 4 for input of [5, 5, 5, 5, 4, 5, 5, 5]", func() {
		input := []float32{5, 5, 5, 5, 4, 5, 5, 5}
		expected := float32(4)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 5 for input of [6, 6, 6, 6, 6, 5, 6, 6]", func() {
		input := []float32{6, 6, 6, 6, 6, 5, 6, 6}
		expected := float32(5)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 6 for input of [7, 7, 7, 7, 7, 7, 6, 7]", func() {
		input := []float32{7, 7, 7, 7, 7, 7, 6, 7}
		expected := float32(6)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 7 for input of [8, 8, 8, 8, 8, 8, 8, 7]", func() {
		input := []float32{8, 8, 8, 8, 8, 8, 8, 7}
		expected := float32(7)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 2 for input of [3, 3, 2, 3, 3, 3, 3, 3]", func() {
		input := []float32{3, 3, 2, 3, 3, 3, 3, 3}
		expected := float32(2)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 2 for input of [2, 1, 2, 2, 2, 2, 2, 2]", func() {
		input := []float32{2, 1, 2, 2, 2, 2, 2, 2}
		expected := float32(1)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should return 0 for input of [0, 1, 1, 1, 1, 1, 1, 1]", func() {
		input := []float32{0, 1, 1, 1, 1, 1, 1, 1}
		expected := float32(0)
		actual := FindUniq(input)

		Expect(actual).To(Equal(expected))
	})

	It("should work with random cases", func() {
		ans := float32(rand.Intn(100))
		mass := float32(rand.Intn(100))
		input := generateTestArr(ans, mass, 30)
		actual := FindUniq(input)

		Expect(actual).To(Equal(ans))
	})

	It("should work with some edge cases", func() {
		// Very big number
		Expect(FindUniq(generateTestArr(MaxFloat32, MaxFloat32/2, 1000))).To(Equal(float32(MaxFloat32)))
		// Negative number
		Expect(FindUniq(generateTestArr(float32(-30), 0.0, 10))).To(Equal(float32(-30)))
		// Very big array
		ans := float32(rand.Intn(100))
		mass := float32(rand.Intn(100))
		Expect(FindUniq(generateTestArr(ans, mass, MaxInt32/32))).To(Equal(ans))
	})
})
