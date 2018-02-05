package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	socket *websocket.Conn

	send chan []byte
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, m, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- m
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for m := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, m)
		if err != nil {
			return
		}
	}
}
