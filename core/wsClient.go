package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type WsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage
	closeChan chan int
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{
		conn:      conn,
		readChan:  make(chan *WsMessage),
		closeChan: make(chan int),
	}
}

func (w *WsClient) Ping(waitTime time.Duration) {
	for {
		time.Sleep(waitTime)
		err := w.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			ClientMap.Delete(w.conn)
			return
		}
	}
}

func (w *WsClient) ReadLoop() {
	for {
		messageType, p, err := w.conn.ReadMessage()
		if err != nil {
			ClientMap.Delete(w.conn)
			w.closeChan <- 1 //读取有错误，关闭所有的goroutine
			return

		}
		w.readChan <- NewWsMessage(messageType, p)

	}
}
func (w *WsClient) HandlerLoop() {
loop:
	for {
		select {
		case msg := <-w.readChan: //读取到消息
			fmt.Println(string(msg.MessageData))
		case <-w.closeChan: //一旦读取到数据，关闭所有的goroutine
			fmt.Println("close")
			break loop
		}
	}
}
