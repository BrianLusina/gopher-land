package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// team represents a team in a tournament.
type (
	team struct {
		Name string
		MP   int
		W    int
		D    int
		L    int
		P    int
	}

	// teamMap represents mapping of teams to their names.
	teamMap map[string]team

	// match represents a match in a tournament.
	// outcome is a string representing the outcome of the match.
	// if it's a 'win', then team1 won the match, 'draw', both teams drew, 'loss' means
	// team2 won the match.
	match struct {
		team1   string
		team2   string
		outcome string
	}
)

func assignPointsToTeams(teamOne, teamTwo team, outcome string) (team, team, error) {
	teamOne.MP++
	teamTwo.MP++
	switch outcome {
	case "win":
		teamOne.W++
		teamOne.P += 3

		teamTwo.L++
	case "loss":
		teamOne.L++

		teamTwo.W++
		teamTwo.P += 3
	case "draw":
		teamOne.D++
		teamOne.P++

		teamTwo.D++
		teamTwo.P++
	default:
		return teamOne, teamTwo, fmt.Errorf("invalid match outcome: %s", outcome)
	}
	return teamOne, teamTwo, nil
}

func assignPointsAfterMatch(match match, teams teamMap) (teamMap, error) {
	matchOutcome := match.outcome
	team1, teamOnePresent := teams[match.team1]
	team2, teamTwoPresent := teams[match.team2]
	if teamOnePresent && teamTwoPresent {
		t1, t2, err := assignPointsToTeams(team1, team2, matchOutcome)
		if err != nil {
			return teams, err
		}

		teams[match.team1] = t1
		teams[match.team2] = t2
	} else if teamOnePresent && !teamTwoPresent {
		newTeam := team{Name: match.team2}
		t1, t2, err := assignPointsToTeams(team1, newTeam, matchOutcome)
		if err != nil {
			return teams, err
		}

		teams[match.team1] = t1
		teams[match.team2] = t2
	} else if teamTwoPresent && !teamOnePresent {
		newTeam := team{Name: match.team1}

		t1, t2, err := assignPointsToTeams(newTeam, team2, matchOutcome)
		if err != nil {
			return teams, err
		}

		teams[match.team1] = t1
		teams[match.team2] = t2
	} else {
		newTeamOne := team{Name: match.team1}
		newTeamTwo := team{Name: match.team2}

		t1, t2, err := assignPointsToTeams(newTeamOne, newTeamTwo, matchOutcome)
		if err != nil {
			return teams, err
		}

		teams[match.team1] = t1
		teams[match.team2] = t2
	}
	return teams, nil
}

func readRecords(reader io.Reader) (teamMap, error) {
	var teams = teamMap{}
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	// this allows a variable number of fields in each row
	csvReader.FieldsPerRecord = -1

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) == 3 {
			team1, team2, matchOutcome := record[0], record[1], record[2]
			match := match{team1: team1, team2: team2, outcome: matchOutcome}
			teams, err = assignPointsAfterMatch(match, teams)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid record: %v", record)
		}
	}
	return teams, nil
}

func generateReport(writer io.Writer, teamMap teamMap) {
	entries := make([]team, 0, len(teamMap))

	for _, team := range teamMap {
		entries = append(entries, team)
	}

	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		switch {
		case a.P != b.P:
			return a.P > b.P
		case a.W != b.W:
			return a.W > b.W
		default:
			return a.Name < b.Name
		}
	})

	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, entry := range entries {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			entry.Name, entry.MP, entry.W, entry.D, entry.L, entry.P)
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams, err := readRecords(reader)
	if err != nil {
		return err
	}
	generateReport(writer, teams)
	return nil
}
