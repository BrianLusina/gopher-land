package squaredstrings

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSquaredStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SquaredStrings Test Suite")
}

func dotestVert(a1 string, exp string) {
	var ans = Oper(VertMirror, a1)
	Expect(ans).To(Equal(exp))
}
func dotestHor(a1 string, exp string) {
	var ans = Oper(HorMirror, a1)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("SquaredStrings Tests", func() {

	Describe("Vertical Tests", func() {
		It("should handle basic cases Oper VertMirror", func() {
			dotestVert("hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu", "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw")
		})
	})

	Describe("Horizontal Tests", func() {
		It("should handle basic cases Oper HorMirror", func() {
			dotestHor("lVHt\nJVhv\nCSbg\nyeCt", "yeCt\nCSbg\nJVhv\nlVHt")
		})
	})
})
