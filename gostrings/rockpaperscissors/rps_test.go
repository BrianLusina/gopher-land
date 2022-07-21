package rockpaperscissors

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

func referenceSolution(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	f := func(s string) int {
		switch s {
		case "scissors":
			return 0
		case "paper":
			return 1
		default:
			return 2
		}
	}
	i := f(p1) - f(p2)
	if i == -1 || i == 2 {
		return "Player 1 won!"
	} else {
		return "Player 2 won!"
	}
}

func rndPlay() string {
	switch rand.Intn(3) {
	case 0:
		return "scissors"
	case 1:
		return "paper"
	default:
		return "rock"
	}
}

func TestRockPaperScissors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rock Paper Scissors Test Cases")
}

var _ = Describe("Rock Paper Scissors", func() {
	It("should return Player 1 won! for p1=rock, p2=scissors", func() {
		p1 := "rock"
		p2 := "scissors"
		expected := "Player 1 won!"
		actual := Rps(p1, p2)
		Expect(actual).To(Equal(expected))
	})

	It("should return Player 2 won! for p1=scissors, p1=rock", func() {
		p1 := "scissors"
		p2 := "rock"
		expected := "Player 2 won!"
		actual := Rps(p1, p2)
		Expect(actual).To(Equal(expected))
	})

	It("should return Draw! for p1=rock, p2=rock", func() {
		p1 := "rock"
		p2 := "rock"
		expected := "Draw!"
		actual := Rps(p1, p2)
		Expect(actual).To(Equal(expected))
	})

	for i := 0; i < 100; i++ {
		p1, p2 := rndPlay(), rndPlay()
		actual := Rps(p1, p2)
		expected := referenceSolution(p1, p2)
		It(fmt.Sprintf("should return %s for p1=%s, p2=%s", expected, p1, p2), func() {
			Expect(actual).To(Equal(expected))
		})
	}
})
