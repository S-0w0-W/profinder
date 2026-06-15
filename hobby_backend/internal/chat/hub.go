package chat

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Hub struct {
	Clients    map[*Client]bool
	Unregister chan *Client
	Register   chan *Client
	MsgQueue   chan []byte
	RC         *redis.Client
}

func NewHub(rc *redis.Client) *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Unregister: make(chan *Client),
		Register:   make(chan *Client),
		MsgQueue:   make(chan []byte),
		RC:         rc,
	}
}

func (h *Hub) Run(ctx context.Context){
	sub := h.RC.Subscribe(ctx, "chat")

	go func() {
		for {
			msg, ok := <-sub.Channel()
			if !ok {
				break
			}
			for client := range h.Clients {
				client.ReadQueue <- []byte(msg.Payload)
			}
		}
	}()

	// go func() {
	for {
		println("in hub run loop")
		select{
			case newClient := <-h.Register:
				h.Clients[newClient] = true
			case client := <-h.Unregister:
				delete(h.Clients, client)
			case msg := <-h.MsgQueue:
				h.RC.Publish(ctx, "chat", msg)
		}
	}
	// }()
}