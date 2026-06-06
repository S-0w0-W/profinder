package handlers

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
)

type FileHandler struct {
    pool *pgxpool.Pool
}

func NewFileHandler(pool *pgxpool.Pool) *FileHandler {
    return &FileHandler{pool: pool}
}

func (h *FileHandler) Upload(c *gin.Context) {
    fmt.Println("in upload")
}

func (h *FileHandler) Test(c *gin.Context) {
    fmt.Println("in Test")
}