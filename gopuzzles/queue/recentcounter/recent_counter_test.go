package recentcounter

import (
	"fmt"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

type testCall struct {
	t        int
	expected int
}

type testCase struct {
	calls []testCall
}

var testCases = []testCase{
	{
		calls: []testCall{
			{
				t:        1,
				expected: 1,
			},
			{
				t:        100,
				expected: 2,
			},
			{
				t:        3001,
				expected: 3,
			},
			{
				t:        3002,
				expected: 3,
			},
		},
	},
}

func TestRecentCounter(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "RecentCounter Suite")
}

var _ = ginkgo.Describe("RecentCounter", func() {

	for _, tc := range testCases {
		ginkgo.It(fmt.Sprintf("when calls are %v", tc.calls), func() {
			recentCounter := New()

			for _, call := range tc.calls {
				actual := recentCounter.Ping(call.t)
				gomega.Expect(actual).To(gomega.Equal(call.expected))
			}
		})
	}
})

func BenchmarkRecentCounter(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			recentCounter := New()

			for _, call := range tc.calls {
				actual := recentCounter.Ping(call.t)
				if actual != call.expected {
					b.Errorf("expected %v, got %v", call.expected, actual)
				}
			}
		}
	}
}
