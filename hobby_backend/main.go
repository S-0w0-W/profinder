package main

import (
	"context"
	db "hobby_backend/internal/migration_handler"
	"hobby_backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    ctx := context.Background()
    pool, err := db.ConnectDB(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()

    err = db.RunMigrations(ctx, pool)

    c := gin.Default()

    routes.FileRoutes(c, pool)

    // db.RunMigrations(postgres_db)

    c.Run(":8080")
}