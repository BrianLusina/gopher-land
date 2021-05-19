package fibonacci

func Tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	first, second, third := 0, 1, 1

	for i := 3; i <= n; i++ {
		fourth := first + second + third
		first = second
		second = third
		third = fourth
	}

	return third
}
