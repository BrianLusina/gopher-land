package pascal

func Triangle(n int) [][]int {
	result := [][]int{}

	if n < 1 {
		return result
	}

	for x := 0; x < n; x++ {
		output := []int{}
		for y := 0; y <= x; y++ {
			bin := binomial(x, y)
			output = append(output, bin)
		}
		result = append(result, output)
	}

	return result
}

func binomial(x, y int) int {
	return factorial(x) / factorial(y) / factorial(x-y)
}

func factorial(n int) int {
	result := 1

	for n >= 1 {
		result *= n
		n -= 1
	}

	return result
}
