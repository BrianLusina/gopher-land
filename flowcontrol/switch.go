package main

import (
	"fmt"
	"runtime"
	"time"
)

func whenIsSaturday() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("in 2 days")
	default:
		fmt.Println("Too far away")
	}
}

// switch can be written with no condition and can be used to evaluate
// long if else chains
func switchWithNoCondition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	whenIsSaturday()
	switchWithNoCondition()
}

