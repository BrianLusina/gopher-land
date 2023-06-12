package reversewords

import (
	"strings"
)

func reverseWords(phrase string) string {
	words := strings.Split(phrase, ".")
	output := make([]string, len(words))

	for i := range words {
		word := words[len(words)-1-i]
		output[i] = word
	}

	return strings.Join(output, ".")
}

// reverseWordsInString reverses words in a given string removing white spaces if any from both ends and trimming white space from the middle of the string
func reverseWordsInString(phrase string) string {
	phraseList := strings.Fields(phrase)

	reverseCharacters := func(messageList []string, frontIndex int, backIndex int) []string {
		// walk towards the middle, from both sides
		for frontIndex < backIndex {
			// swap the front char and back char
			messageList[frontIndex], messageList[backIndex] = messageList[backIndex], messageList[frontIndex]

			frontIndex += 1
			backIndex -= 1
		}

		return messageList
	}

	reverseCharacters(phraseList, 0, len(phraseList)-1)

	return strings.Join(phraseList, " ")
}
