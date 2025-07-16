package tournamentwinner

import "gopherland/pkg/utils"

func tournamentWinner(competitions [][]string, results []int) string {
	winner := ""
	teams := map[string]int{}

	pairs, err := utils.Zip(competitions, results)
	if err != nil {
		panic(err)
	}

	for _, pair := range pairs {
		homeTeam := pair.A[0]
		awayTeam := pair.A[1]

		teams[homeTeam] = 0
		teams[awayTeam] = 0

		result := pair.B

		if result == 1 {
			if points, ok := teams[homeTeam]; ok {
				teams[homeTeam] = points + 3
			}
		} else {
			if points, ok := teams[awayTeam]; ok {
				teams[awayTeam] = points + 3
			}
		}
	}

	for team, points := range teams {
		if points > teams[winner] {
			winner = team
		}
	}

	return winner
}

func tournamentWinnerV2(competitions [][]string, results []int) string {
	winner := ""
	teams := map[string]int{}

	for idx, competition := range competitions {
		homeTeam, awayTeam := competition[0], competition[1]
		competitionWinner := results[idx]

		// 0 means that the away team won
		if competitionWinner == 0 {
			if points, ok := teams[awayTeam]; ok {
				teams[awayTeam] = points + 3
			} else {
				teams[awayTeam] = points + 3
			}
		} else {
			if points, ok := teams[homeTeam]; ok {
				teams[homeTeam] = points + 3
			} else {
				teams[homeTeam] = points + 3
			}
		}
	}

	for team, points := range teams {
		if points > teams[winner] {
			winner = team
		}
	}

	return winner
}
