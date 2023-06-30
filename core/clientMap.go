package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map
}

func (c *ClientMapStruct) Store(conn *websocket.Conn) {
	ws := NewWsClient(conn)
	c.data.Store(conn.RemoteAddr().String(), ws)
	go ws.Ping(1 * time.Second)
	go ws.ReadLoop()    //处理读循环
	go ws.HandlerLoop() //处理 总的控制循环
}
func (c *ClientMapStruct) SendAll(msg string) {
	c.data.Range(func(k, v interface{}) bool {
		ws := v.(*WsClient)
		err := ws.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return true
	})
}

func (c *ClientMapStruct) Delete(conn *websocket.Conn) {
	c.data.Delete(conn.RemoteAddr().String())
}
