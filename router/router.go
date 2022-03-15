package router

import (
	"net/http"

	"github.com/ivanauliaa/shortener-go/handler"

	"github.com/gin-gonic/gin"
)

func ServerRouter() *gin.Engine {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "URL Shortener built with Go"})
	})
	server.GET("/:code", handler.Redirect)
	server.POST("/shorten", handler.Shorten)

	return server
}
