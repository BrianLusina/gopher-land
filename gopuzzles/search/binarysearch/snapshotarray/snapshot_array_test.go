package snapshotarray

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotArray(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "SnapshotArray Suite")
}

var _ = ginkgo.Describe("Snapshot Array", func() {
	ginkgo.When("SnapshotArray of length 3", func() {
		snapshotArray := New(3)

		ginkgo.It("set(0, 5) -> set(0, 6) -> get(0, 0)", func() {
			snapshotArray.Set(0, 5)

			expectedFirstSnap := 0
			actualFirstSnap := snapshotArray.Snap()

			assert.Equal(ginkgo.GinkgoT(), expectedFirstSnap, actualFirstSnap)

			snapshotArray.Set(0, 6)

			expectedFirstSnapId := 5
			actualFirstSnapId := snapshotArray.Get(0, 0)

			assert.Equal(ginkgo.GinkgoT(), expectedFirstSnapId, actualFirstSnapId)
		})
	})

	ginkgo.When("SnapshotArray of length 1", func() {
		snapshotArray := New(1)

		ginkgo.It("set(0, 4) -> set(0, 16) -> set(0, 13) -> snap() -> get(0, 0) -> snap()", func() {
			snapshotArray.Set(0, 4)
			snapshotArray.Set(0, 16)
			snapshotArray.Set(0, 13)

			expectedFirstSnap := 0
			actualFirstSnap := snapshotArray.Snap()

			assert.Equal(ginkgo.GinkgoT(), expectedFirstSnap, actualFirstSnap)

			expectedFirstSnapId := 13
			actualFirstSnapId := snapshotArray.Get(0, 0)
			assert.Equal(ginkgo.GinkgoT(), expectedFirstSnapId, actualFirstSnapId)

			expectedSecondSnap := 1
			actualSecondSnap := snapshotArray.Snap()
			assert.Equal(ginkgo.GinkgoT(), expectedSecondSnap, actualSecondSnap)
		})
	})
})
