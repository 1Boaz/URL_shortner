package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting")
	router := gin.New()
	db := DBinit()

	server := &Server{
		db: db,
	}

	router.POST("/create", server.New)
	router.GET("/:Shortend", server.Get)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Run("localhost:8080")
}
