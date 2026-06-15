package chat

import (
  "github.com/gorilla/websocket"
)

type Client struct {
	WebSocketConnection *websocket.Conn
	Hub *Hub
	ReadQueue chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn) *Client{
	return &Client{
		WebSocketConnection: conn,
		Hub: hub,
		ReadQueue: make(chan []byte),
	}
}

func (c *Client) SendMsg(){
	defer func() {
		c.Hub.Unregister <- c
		c.WebSocketConnection.Close()
	}()

	for {
		_, msg, err := c.WebSocketConnection.ReadMessage()
		println("in SendMsg loop")
		if err != nil {
			break
		}
		c.Hub.MsgQueue <- msg
	}
}

func (c *Client) GetMsg(){
	defer func() {
		c.Hub.Unregister <- c
		c.WebSocketConnection.Close()
	}()

	for msg := range c.ReadQueue{
		c.WebSocketConnection.WriteMessage(websocket.TextMessage, msg)
	}
}