package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	sendCh chan []byte
	room   *room
}

func (c *client) read() {
	for {
		if _, message, err := e.socket.ReadMessage(); err == nil {
			c.room.forwardCh <- message
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for message := range c.sendCh {
		err := c.socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.socket.Close()
}
