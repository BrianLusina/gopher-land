package performance

// The sumValue and sumPtr demonstrate the difference of allocations on the heap vs the stack. In Go, every goroutine has access to its own stack
// making it trivial to handle allocations on the stack and since a stack is self-healing(will pop a stack frame once a function has been executed to completion)
// it makes it much faster and improves performance. However, using a heap is different and will require a Garbage Collector(GC) to perform cleanup which
// uses up 25% of CPU memory and potentially impacts performance as it could lead to "Stop The World GC events" in an application.
// Additionally, it requires a bit of time to allocate a variable to the heaps which also adds a performance impact.
// The sumValue does not use pointers but values, so, the variables are allocated on the stack and therefore much faster. Once the execution is done, the result
// of the function is returned ann all variables in the scope of the function are removed from the stack.
// However, the sumPtr function returns a pointer which is more complex as that can not be removed from the stack as Go can not reliably tell where it is
// referenced again in the execution path. Therefore, once the sumPtr returns, the pointer returned can not be simply popped off from the stack, this means that
// it has to be allocated on the heap to be later referenced by other parts of the code.
// Using benchmark tests reveals:
//
// BenchmarkSumValue-12    	770171083	         1.560 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSumPtr-12      	51745873	        20.72 ns/op	       8 B/op	       1 allocs/op
//
// this shows that the sumPtr is much slower than sumValue due to the impact of heap allocations for pointers over value allocations on the stack.
// the benchmark tests do not show stack allocations, but rather heap allocations
//
// B/op: how many bytes per operation allocated
// allocs/op: how many allocations per operation

//go:noinline
func sumValue(x, y int) int {
	z := x + y
	return z
}

//go:noinline
func sumPtr(x, y int) *int {
	z := x + y
	return &z
}
