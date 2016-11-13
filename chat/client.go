package chat

import (
	"github.com/gorilla/websocket"
	"log"
)

type client struct {
	socket *websocket.Conn
	sendCh chan []byte
	room   *room
}

func (c *client) read() {
	log.Println("client read")

	for {
		if _, message, err := c.socket.ReadMessage(); err == nil {
			log.Println("送信するメッセージ:", message)

			// TODO: メソッド化
			c.room.forwardCh <- message
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	log.Println("client write")

	for message := range c.sendCh {
		err := c.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.socket.Close()
}
