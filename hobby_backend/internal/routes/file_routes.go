package routes

import (
	"profinder_backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

)

func FileRoutes(c *gin.Engine, pool *pgxpool.Pool) {
	FileHandler := handlers.NewFileHandler(pool)

	c.POST("/upload", FileHandler.Upload)
	c.GET("/test", FileHandler.Test)
}
