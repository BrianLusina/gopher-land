package reversestring

import dynamicstack "gopherland/datastructures/stack/dynamic"

// reverseString Reverses an input string using a stack data structure
// Complexity:
// Where n is the length of the input string
// Time: O(n), iterate through the whole string
// Space: O(n), a stack is used to store the characters in the input string.
// Args:
// 	text (str): string to reverse
// Return:
// 	str: reversed string
func reverseString(text string) string {
	stack := dynamicstack.New[rune]()

	// # push all the characters onto the stack
	for _, char := range text {
		stack.Push(char)
	}

	// # initialize an empty string which will be the result
	reversedString := ""

	// # iterate over all the items in the stack popping of items from the top of the stack and adding them to the result
	// # string.
	for !stack.IsEmpty() {
		value, err := stack.Pop()
		if err != nil {
			panic("failed to pop off from stack")
		}
		reversedString += string(value)
	}

	return reversedString
}
