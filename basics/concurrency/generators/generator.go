package main

import (
	"fmt"
)

func foo() <-chan string {
	mychannel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			mychannel <- fmt.Sprintf("%s %d", "Counter at : ", i)
		}
	}()

	return mychannel // returns the channel as returning argument
}

func generator() {
	mychannel := foo() // foo() returns a channel.

	for i := 0; i < 5; i++ {
		fmt.Printf("%q\n", <-mychannel)
	}

	fmt.Println("Done with Counter")
}
