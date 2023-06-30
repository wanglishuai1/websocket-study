package main

import (
	"net/http"
	"websocket-study/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.Echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
