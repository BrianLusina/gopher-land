package browserhistory

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "github.com/stretchr/testify/assert"
)

func TestBrowserHistory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BrowserHistory Suite")
}

var _ = Describe("Browser History Tests", func() {
	It("should handle visit('google.com') -> visit('facebook.com') -> visit('youtube.com') -> back(1) -> back(1) -> forward(1) -> visit('linkedin.com') -> forward(2) -> back(2) -> back(7)", func() {
		browserhistory := Constructor("leetcode.com")
		browserhistory.Visit("google.com")
		browserhistory.Visit("facebook.com")
		browserhistory.Visit("youtube.com")

		// move back to facebook.com from youtube.com
		Expect(browserhistory.Back(1)).To(Equal("facebook.com"))

		// move back to google.com from facebook.com
		Expect(browserhistory.Back(1)).To(Equal("google.com"))

		// move forward to facebook.com from google.com
		Expect(browserhistory.Forward(1)).To(Equal("facebook.com"))

		browserhistory.Visit("linkedin.com")

		// move forward by 2 steps is not possible
		Expect(browserhistory.Forward(2)).To(Equal("linkedin.com"))

		// move back by 2 steps to google.com
		Expect(browserhistory.Back(2)).To(Equal("google.com"))

		// move back by 7 steps to leetcode.com
		Expect(browserhistory.Back(7)).To(Equal("leetcode.com"))
	})
})
