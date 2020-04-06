package concurrency

import "fmt"

// sendValues sends values to the specified channel that takes in integers
func sendValues(myIntChannel chan int) {
	for x := 0; x < 10; x++ {
		myIntChannel <- x
	}

	// close the channel from the receiver as it knows whether the channel is open
	close(myIntChannel)
}

func main() {
	// make a channel with
	resultChannel := make(chan int)

	// goroutine to send values to resultChannel
	go sendValues(resultChannel)

	for y := 0; y < 10; y++ {
		// receiving value using a unary operator
		// get the value and status of channel
		value, open := <- resultChannel

		// if not open, break
		if !open {
			break
		}

		fmt.Println(value)
	}

	// syntactic sugar to perform the above operation
	//for value := range resultChannel {
	//	fmt.Println(value)
	//}
}