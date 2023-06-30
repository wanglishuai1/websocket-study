package core

import (
	"github.com/gorilla/websocket"
	"sync"
)

var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map
}

func (c *ClientMapStruct) Store(key string, conn *websocket.Conn) {
	c.data.Store(key, conn)
}
