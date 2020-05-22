// Package proverb contains helper functions to create proverbs from rhymes
package proverb

import "fmt"

// Proverb returns a proverb given an array of rhymes
func Proverb(rhyme []string) (proverb []string) {
	if len(rhyme) == 0 {
		return proverb
	}

	if len(rhyme) > 1 {
		for index := 0; index < len(rhyme)-1; index++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index], rhyme[index+1]))
		}
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
