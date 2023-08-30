package main

import (
	"net/http"

	"github.com/Laminator42/go-urlshortener-mongo/db"
	"github.com/Laminator42/go-urlshortener-mongo/urlshortener"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()

	router := gin.Default()

	// Root endpoint
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusAccepted, gin.H{"message": "URL Shortener"}) })

	// Health check
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Add shortener router group
	v1 := router.Group("/")
	urlshortener.UrlsRegister(v1)

	router.Run(urlshortener.AppConf.HostAddress())
}
