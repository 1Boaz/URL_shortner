package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Starting")
	router := gin.New()
	db := DBinit()

	server := &Server{
		db: db,
	}

	router.POST("/create", server.New)
	router.POST("/remove", server.Remove)
	router.GET("/:Shortened", server.Get)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run("localhost:8080")
}
