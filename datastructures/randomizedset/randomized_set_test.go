package randomizedset

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRandomizedSet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RandomizedSet Suite")
}

var _ = Describe("RandomizedSet", func() {

	It("insert(1) -> remove(2) -> insert(2) -> getRandom() -> remove(1) -> insert(2) -> getRandom()", func() {
		randomizedSet := New[int]()

		actualStepOne := randomizedSet.Insert(1)
		Expect(actualStepOne).To(BeTrue())

		actualStepTwo := randomizedSet.Remove(2)
		Expect(actualStepTwo).To(BeFalse())

		actualStepThree := randomizedSet.Insert(2)
		Expect(actualStepThree).To(BeTrue())

		actualStepFour := randomizedSet.GetRandom()
		Expect(actualStepFour).To(BeElementOf([]int{1, 2}))

		actualStepFive := randomizedSet.Remove(1)
		Expect(actualStepFive).To(BeTrue())

		actualStepSix := randomizedSet.Insert(2)
		Expect(actualStepSix).To(BeFalse())

		actualStepSeven := randomizedSet.GetRandom()
		Expect(actualStepSeven).To(Equal(2))
	})
})
