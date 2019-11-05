// Package containing convert function that takes in number input
// and checks for its factors and returns a string with either Pling, Plang, or Plong depending on whether the factors
// of the number are 3, 5 or 7 respectively
package raindrops

import "strconv"

/**
Function that takes in number input and checks for its factors and converts it into 'raindrops'
Returns Pling if 3 is a factor, Plang if 5 is a factor or Plong is 7 is a factor. It returns a concatenated string
with these words if 3, 5 or 7 are factors of the input number in the order in which they have been found.
Returns the number itself if it has no factors
*/
func Convert(input int) string {
	raindrops := ""

	if input%3 == 0 {
		raindrops += "Pling"
	}

	if input%5 == 0 {
		raindrops += "Plang"
	}

	if input%7 == 0 {
		raindrops += "Plong"
	}

	if len(raindrops) == 0 {
		raindrops = strconv.Itoa(input)
	}

	return raindrops
}
