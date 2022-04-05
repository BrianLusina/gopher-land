package main

import (
	"log"
	"os"
	"os/exec"
)

func ultraCrazyFunction() {
	prog := os.Args[0]
	e := recover()
	if e != nil {
		// In case of panic, try to restart the entire application:
		cmd := exec.Command(prog)

		err := cmd.Run()
		// If that fails, log the error
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	defer ultraCrazyFunction()
}
