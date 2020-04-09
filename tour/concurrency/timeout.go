package concurrency

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	responseChannel := make(chan *http.Response, 1)
	errorChannel := make(chan *error)

	go func() {
		resp, err := http.Get("http://matt.aimonetti.net/")
		if err != nil {
			errorChannel <- &err
		}

		responseChannel <- resp
	}()

	for {
		select {
		case r := <-responseChannel:
			fmt.Println("%s", r.Body)
			return
		case e := <-errorChannel:
			log.Fatal(*e)

		case <-time.After(2000 * time.Millisecond):
			fmt.Println("Timed Out")
			return
		}
	}
}
