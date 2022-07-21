package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden map[string][]string

func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")

	if len(rows) != 3 || rows[0] != "" {
		return nil, errors.New("diagram must have 2 rows")
	}

	if len(rows[1]) != len(rows[2]) {
		return nil, errors.New("diagram rows must be of same length")
	}

	if len(rows[1]) != 2*len(children) {
		return nil, errors.New("each diagram row must have 2 cups per child")
	}

	garden := Garden{}
	childList := make([]string, len(children))
	copy(childList, children)
	sort.Strings(childList)

	for _, child := range childList {
		garden[child] = make([]string, 0, 4)
	}

	if len(garden) != len(childList) {
		return nil, errors.New("no 2 children can have the same name")
	}

	for _, row := range rows[1:] {
		for childIdx, child := range childList {
			for cupIdx := range rows[1:] {
				var plant string
				switch row[2*childIdx+cupIdx] {
				case 'G':
					plant = "grass"
				case 'C':
					plant = "clover"
				case 'R':
					plant = "radishes"
				case 'V':
					plant = "violets"
				default:
					return nil, errors.New("plant codes must be one of G, C, R or V")
				}
				garden[child] = append(garden[child], plant)
			}
		}
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
