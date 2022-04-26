package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	words := strings.Fields(question)
	if len(words) < 3 { // length check for first two words and last word
		return 0, false
	}
	if words[0] != "What" || words[1] != "is" { // first two words required
		return 0, false
	}
	operation := words[2:]
	lastIdx := len(operation) - 1
	lastWord := operation[lastIdx]
	if lastWord[len(lastWord)-1] != '?' { // trailing ? required
		return 0, false
	}
	operation[lastIdx] = lastWord[:len(lastWord)-1]
	firstNumber, err := strconv.Atoi(operation[0]) // first word after preamble must be a number
	if err != nil {
		return 0, false
	}
	for i := 1; i < len(operation); i++ { // remainder of question must alternate between operator, number
		operator := operation[i]
		i++
		switch operator {
		case "multiplied", "divided":
			if i == len(operation) || operation[i] != "by" { // length check, required word
				return 0, false
			}
			i++ // consume
		}
		if i == len(operation) { // length check for operand
			return 0, false
		}
		nextNumber, err := strconv.Atoi(operation[i]) // must be a number
		if err != nil {
			return 0, false
		}
		switch operator { // apply operator
		case "plus":
			firstNumber += nextNumber
		case "minus":
			firstNumber -= nextNumber
		case "multiplied":
			firstNumber *= nextNumber
		case "divided":
			firstNumber /= nextNumber
		default: // valid operator not found
			return 0, false
		}
	}
	return firstNumber, true
}

func AnswerWithRe(s string) (int, bool) {
	if match, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`, s); !match {
		return 0, false
	}
	operatorsRegex := regexp.MustCompile(`(plus|minus|divided|multiplied)`)
	operators := operatorsRegex.FindAllString(s, -1)

	digitsRegex := regexp.MustCompile(`-?\d+`)
	numbers := digitsRegex.FindAllString(s, -1)
	if len(operators) != len(numbers)-1 {
		return 0, false
	}
	firstNumber, _ := strconv.Atoi(numbers[0])
	for i, operator := range operators {
		nextNumber, _ := strconv.Atoi(numbers[i+1])
		switch operator {
		case "plus":
			firstNumber += nextNumber
		case "minus":
			firstNumber -= nextNumber
		case "divided":
			firstNumber /= nextNumber
		case "multiplied":
			firstNumber *= nextNumber
		}
	}
	return firstNumber, true
}
