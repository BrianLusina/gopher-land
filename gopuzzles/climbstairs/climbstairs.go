package climbstairs

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

func ClimbStaircase(steps int) int {
	if steps < 0 {
		return 0
	}
	if steps == 1 || steps == 0 {
		return 1
	}
	return ClimbStaircase(steps-1) + ClimbStaircase(steps-2)
}
