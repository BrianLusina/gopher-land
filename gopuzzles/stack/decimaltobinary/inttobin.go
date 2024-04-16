package decimaltobinary

import (
	dynamicstack "gopherland/datastructures/stack/dynamic"
	"gopherland/math/utils"
	"strconv"
)

// convertIntToBin Converts a decimal integer to binary using a stack data structure. Uses the division by 2 method to perform the
// conversion. Each remainder from the division is stored on the stack and popped off until we have the binary
// number
func convertIntToBin(num int) int {
	stack := dynamicstack.New[int]()
	current := num

	for current != 0 {
		a, rem := utils.DivMod(current, 2)
		current = a
		stack.Push(rem)
	}

	binNum := ""
	for !stack.IsEmpty() {
		binary, _ := stack.Pop()
		binNum += strconv.Itoa(binary)
	}

	n, err := strconv.Atoi(binNum)
	if err != nil {
		panic("failed to convert to binary")
	}

	return n
}
