package faninfanout

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var numbers [10]int
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		numbers[i] = rand.Intn(50)
	}

	channelOut := channelGenerator(numbers)

	channel1 := double(channelOut)
	channel2 := double(channelOut)

	channelIn := fanInNumChannels(channel1, channel2)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-channelIn)
	}
}

func double(inChan <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := range inChan {
			num, err := strconv.Atoi(i)
			if err != nil {
				_ = fmt.Errorf("Error converting string to int: %v", err)
			}
			channel <- fmt.Sprintf("%d * 2 = %d", num, num*2)
		}
		close(channel)
	}()

	return channel
}

func channelGenerator(numbers [10]int) <-chan string {
	channel := make(chan string)

	go func() {
		for _, number := range numbers {
			channel <- strconv.Itoa(number)
		}
		close(channel)
	}()
	return channel
}

func fanInNumChannels(chan1, chan2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			select {
			case messageOne := <-chan1:
				channel <- messageOne
			case messageTwo := <-chan2:
				channel <- messageTwo
			}
		}
	}()

	return channel
}
