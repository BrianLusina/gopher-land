package rockpaperscissors

var table = map[string]string{
	"scissors": "paper",
	"paper":    "rock",
	"rock":     "scissors",
}

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}

	result := table[p1] == p2
	if result {
		return "Player 1 won!"
	}
	return "Player 2 won!"

}
