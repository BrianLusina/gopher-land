package diamond

import (
	"errors"
	"fmt"
	"strings"
)

const START_INDEX = 'A'

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New(fmt.Sprintf("%d not supported, Only characters between A & Z are valid", char))
	}

	diamond := []string{}
	currentIdx := int(char - START_INDEX)

	for i := 0; i <= currentIdx; i++ {
		diamond = append(diamond, getLine(currentIdx, i))
	}

	for i := currentIdx - 1; i > -1; i-- {
		diamond = append(diamond, getLine(currentIdx, i))
	}

	return strings.Join(diamond, "\n"), nil
}

func getLine(currentIdx, current int) string {
	diff := currentIdx - current
	return spaces(diff) + alphabets(current) + spaces(diff)
}

func alphabets(position int) string {
	if position == 0 {
		return "A"
	}

	c := rune(position + START_INDEX)
	return string(c) + spaces(position*2-1) + string(c)
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}
