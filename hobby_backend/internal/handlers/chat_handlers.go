package handlers

import (
	"fmt"
	"net/http"
	"profinder_backend/internal/chat"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for development; wrap with strict logic in production
		return true
	},
}

type ChatHandler struct {
	hub *chat.Hub
}

func NewChatHandler(h *chat.Hub) *ChatHandler {
	return &ChatHandler{hub: h}
}

func (ch *ChatHandler) PublicChat(c *gin.Context) {
	fmt.Println("in PublicChat")
	fmt.Println("Headers:", c.Request.Header) // add this
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	client := chat.NewClient(ch.hub, conn)
	ch.hub.Register <-client

	go client.GetMsg()
	go client.SendMsg()

	println("after upgrade")

	// c.JSON(http.StatusOK, gin.H{})
}

func (ch *ChatHandler) ChatTest(c *gin.Context) {
	fmt.Println("in ChatTest")
}
