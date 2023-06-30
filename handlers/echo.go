package handlers

import (
	"fmt"
	"net/http"
	"websocket-study/core"
)

func Echo(w http.ResponseWriter, req *http.Request) {
	client, err := core.Upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println(err)
	}
	core.ClientMap.Store(client)

}
