package functionaladd

import (
	"fmt"
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFunctionalAdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Functional Add Test Cases")
}

var _ = Describe("FunctionalAdd", func() {
	It("should reeturn 4 for Add(1)(3)", func() {
		inputOne := 1
		inputTwo := 3
		expected := 4
		actual := Add(inputOne)(inputTwo)

		Expect(actual).To(Equal(expected))
	})

	It("make sure Add() is pure", func() {
		addThree := Add(3)
		Expect(addThree(5)).To(Equal(8))
		_ = Add(4)
		Expect(addThree(5)).To(Equal(8))
	})

	for i := 0; i < 100; i++ {
		It(fmt.Sprintf("Random calculation %d", i), func() {
			num1 := rand.Intn(1001) - 500
			num2 := rand.Intn(1001) - 500
			Expect(Add(num1)(num2)).To(Equal(num1 + num2))
		})
	}
})
