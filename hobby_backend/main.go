package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"profinder_backend/internal/chat"
	db "profinder_backend/internal/migration_handler"
	"profinder_backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	pool, err := db.ConnectDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	err = db.RunMigrations(ctx, pool)

	c := gin.Default()

	rc := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	hub := chat.NewHub(rc)
	go hub.Run(ctx)

	routes.ChatRoutes(c, hub)
	routes.FileRoutes(c, pool)

	c.Run(":8090")
}
