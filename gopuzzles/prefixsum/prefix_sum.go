package prefixsum

func PrefixSumHelper(input, output []int, parent chan int) {
	if len(input) < 2 {
		parent <- input[0]
		output[0] = input[0] + <-parent

	} else if len(input) < 1 {
		parent <- 0
		<-parent
	} else {
		mid := len(input) / 2
		left := make(chan int)
		right := make(chan int)
		go PrefixSumHelper(input[:mid], output[:mid], left)
		go PrefixSumHelper(input[mid:], output[mid:], right)
		leftsum := <-left
		parent <- leftsum + <-right
		fromleft := <-parent
		left <- fromleft
		right <- fromleft + leftsum
		<-left
		<-right

	}
	parent <- 0
}

func PrefixSum(data []int) ([]int, int) {
	if len(data) == 0 {
		return data, 0
	}

	output := make([]int, len(data))
	parent := make(chan int)

	go PrefixSumHelper(data, output, parent)

	sum := <-parent
	fromleft := 0
	parent <- fromleft
	<-parent
	return output, sum
}
