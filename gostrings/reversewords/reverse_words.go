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
	trimmedPhrase := strings.TrimSpace(phrase)
	phraseList := strings.Fields(trimmedPhrase)

	reversedCharacters := reverseCharacters(phraseList, 0, len(phraseList)-1)

	// this gives us the right word order
	// but with each word backwards
	// now we'll make the words forward again
	// by reversing each word's characters
	// we hold the index of the /start/ of the current word
	// as we look for the /end/ of the current word
	currentWordStartIndex := 0

	for i := 0; i < len(reversedCharacters); i++ {
		// found the end of the current word!
		if (i == len(reversedCharacters)) || reversedCharacters[i] == " " {
			reverseCharacters(reversedCharacters, currentWordStartIndex, i-1)

			// if we haven't exhausted the string our
			// next word's start is one character ahead
			currentWordStartIndex = i + 1
		}
	}

	return strings.Join(phraseList, " ")
}

func reverseCharacters(messageList []string, frontIndex int, backIndex int) []string {
	// walk towards the middle, from both sides
	for frontIndex < backIndex {
		// swap the front char and back char
		messageList[frontIndex], messageList[backIndex] = messageList[backIndex], messageList[frontIndex]

		frontIndex += 1
		backIndex -= 1
	}

	return messageList
}
