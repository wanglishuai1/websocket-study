package main

import (
	"net/http"
	"websocket-study/core"
	"websocket-study/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/sendall", func(w http.ResponseWriter, req *http.Request) {
		msg := req.URL.Query().Get("msg")
		core.ClientMap.SendAll(msg)
		w.Write([]byte("ok"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
