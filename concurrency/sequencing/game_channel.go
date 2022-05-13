package sequencing

import (
	"fmt"
	"math/rand"
	"time"
)

type cookInfo struct {
	foodCooked     string
	waitForPartner chan bool
}

func cookFood(name string) <-chan cookInfo {
	cookChannel := make(chan cookInfo)
	wait := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			cookChannel <- cookInfo{fmt.Sprintf("%s %s", name, "Done"), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-wait
		}
	}()

	return cookChannel
}

func fanIn(chan1, chan2 <-chan cookInfo) <-chan cookInfo {
	outCh := make(chan cookInfo)

	go func() {
		for {
			outCh <- <-chan1
		}
	}()

	go func() {
		for {
			outCh <- <-chan2
		}
	}()

	return outCh
}

func main() {
	gameChannel := fanIn(cookFood("Player 1 : "), cookFood("Player 2 :"))

	for round := 0; round < 3; round++ {
		food1 := <-gameChannel
		fmt.Println(food1.foodCooked)

		food2 := <-gameChannel
		fmt.Println(food2.foodCooked)

		food1.waitForPartner <- true
		food2.waitForPartner <- true

		fmt.Printf("Done with round %d\n", round+1)
	}

	fmt.Println("Done with the competition")
}
