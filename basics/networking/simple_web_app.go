package main

import (
	"io"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request)

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></html></body>`

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func logPanics(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		handler(w, r)
	}
}

func SimpleServer(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "<h1>hello, world</h1>")
	if err != nil {
		return
	}
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch request.Method {
	case "GET":
		_, err := io.WriteString(w, form)
		if err != nil {
			return
		}
	case "POST":
		err := request.ParseForm()
		if err != nil {
			return
		}
		_, err = io.WriteString(w, request.FormValue("in"))
		if err != nil {
			return
		}
	}
}
