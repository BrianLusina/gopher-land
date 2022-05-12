package fibonacci

// NthFibonacci returns the nth fibonacci number.
func NthFibonacci(n int) int {
	return <-NthFibonacciGenerator(n)
}

// NthFibonacciGenerator returns the nth fibonacci number.
func NthFibonacciGenerator(n int) chan int {
	fibChan := make(chan int)

	go func() {
		fibChan <- FibMemo(n)
		// k := 0
		// for i, j := 0, 1; k < n; k++ {
		// 	fibChan <- i
		// 	i, j = j, i+j
		// }
		close(fibChan)
	}()

	return fibChan
}
