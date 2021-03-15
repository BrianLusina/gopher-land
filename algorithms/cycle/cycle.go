package cycle

func Cycle(n int) int {
	if n%2 == 0 || n%5 == 0 {
		return -1
	}

	r := 1
	for i := 1; i < n; i++ {
		r = 10 * r % n
		if r == 1 {
			return i
		}
	}
	return -1
}
