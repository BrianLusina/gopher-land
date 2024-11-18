package resistorcolor

import "fmt"

var colorsToCodes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{}
	for color := range colorsToCodes {
		colors = append(colors, color)
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	v, ok := colorsToCodes[color]
	if !ok {
		panic(fmt.Sprintf("color %s does not exist", color))
	}
	return v
}
