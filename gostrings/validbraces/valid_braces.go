package validbraces

var openersToClosersMap = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
	'<': '>',
}

var openers = map[rune]bool{'(': true, '{': true, '[': true, '<': true}
var closers = map[rune]bool{')': true, '}': true, ']': true, '>': true}

func ValidBraces(str string) bool {
	var stack []rune

	for _, char := range str {
		if openers[char] {
			stack = append(stack, char)
		} else if closers[char] {
			if len(stack) == 0 {
				return false
			} else {
				lastOpener := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if openersToClosersMap[lastOpener] != char {
					return false
				}
			}
		}
	}
	return len(stack) == 0
}
