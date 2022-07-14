package functionaladd

func Add(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}
