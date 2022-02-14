package twelve

import (
	"fmt"
	"strings"
)

var wording = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var gifts = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"second":   "two Turtle Doves",
	"third":    "three French Hens",
	"fourth":   "four Calling Birds",
	"fifth":    "five Gold Rings",
	"sixth":    "six Geese-a-Laying",
	"seventh":  "seven Swans-a-Swimming",
	"eighth":   "eight Maids-a-Milking",
	"ninth":    "nine Ladies Dancing",
	"tenth":    "ten Lords-a-Leaping",
	"eleventh": "eleven Pipers Piping",
	"twelfth":  "twelve Drummers Drumming",
}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	var verse string
	for j := i; j > 0; j-- {
		if i != 1 && j == 1 {
			verse = fmt.Sprintf("%s, and %s", verse, gifts[wording[j]])
		} else if j != i {
			verse = fmt.Sprintf("%s, %s", verse, gifts[wording[j]])
		} else {
			verse = fmt.Sprintf("%s: %s", verse, gifts[wording[j]])
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me%s", wording[i], verse)
}

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return strings.TrimRight(song, "\n")
}
