package main

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) New(context *gin.Context) {
	urls := Urls{}

	if err := context.ShouldBindJSON(&urls); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		slog.Error(err.Error())
		return
	}

	urls.Shortend = strings.TrimSpace(urls.Shortend)
	urls.Long = strings.TrimSpace(urls.Long)
	if !strings.HasPrefix(urls.Long, "http://") && !strings.HasPrefix(urls.Long, "https://") && !strings.HasPrefix(urls.Long, "//") {
		urls.Long = "//" + urls.Long
	}

	result := s.db.Create(&urls)
	if result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: urls.shortend" {
			context.JSON(400, gin.H{"error": result.Error.Error()})
			slog.Warn("Shortend URL already exists in the database")
			return
		}
		context.JSON(400, gin.H{"error": result.Error.Error()})
		slog.Warn(result.Error.Error())
		return
	}
}

func (s *Server) Remove(context *gin.Context) {
	urls := Urls{}

	if err := context.ShouldBindJSON(&urls); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		slog.Error(err.Error())
		return
	}

	urls.Shortend = strings.TrimSpace(urls.Shortend)

	result := s.db.Delete(&urls, "shortend = ?", urls.Shortend)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		slog.Warn(result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		context.JSON(400, gin.H{"error": "Shortend URL does not exist in the database"})
		slog.Warn("Shortend URL does not exist in the database")
		return
	}

	context.JSON(200, gin.H{"message": "Shortend URL deleted successfully"})
}

func (s *Server) Get(context *gin.Context) {
	Shortend := context.Param("Shortened")
	url := Urls{}

	result := s.db.Raw("SELECT Long from urls WHERE Shortend = ?", Shortend).Scan(&url)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		slog.Warn(result.Error.Error())
		return
	}

	context.Redirect(302, url.Long)
}
