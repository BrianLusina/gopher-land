package bandnamegenerator

import (
	"fmt"
	"strings"
)

func bandNameGenerator(word string) string {
	first, last := word[0], word[len(word)-1]
	if first == last {
		return strings.Title(word + word[1:])
	} else {
		return fmt.Sprintf("The %s", strings.Title(word))
	}
}
