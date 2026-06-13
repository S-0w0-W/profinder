package routes

import (
	"fmt"
	"profinder_backend/internal/chat"
	"profinder_backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(c *gin.Engine, h *chat.Hub) {
	ChatHandler := handlers.NewChatHandler(h)

	fmt.Println("in chatroutes")
	c.GET("/chat", ChatHandler.Connect)
	c.GET("/chat_test", ChatHandler.ChatTest)
}
