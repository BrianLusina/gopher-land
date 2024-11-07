package collatzconjecture

import (
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("%d is not allowed", n)
	}

	if n == 1 {
		return 0, nil
	}

	steps := 0
	num := n

	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		steps++
	}

	return steps, nil
}
