package main

import (
	"net/http"

	"github.com/Laminator42/go-urlshortener-mongo/db"
	"github.com/Laminator42/go-urlshortener-mongo/urlshortener"
	"github.com/gin-gonic/gin"

	// Swagger imports
	_ "github.com/Laminator42/go-urlshortener-mongo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title URL Shortener API
// @version			1.0
// @description		An open source URL shortener API in Go using Gin framework and MongoDB as backend.
// @contact.name	Jannik Bach
// @contact.url
// @contact.email
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {

	db.Init()

	router := gin.Default()

	// Add Swagger router
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Root endpoint
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusAccepted, gin.H{"message": "URL Shortener"}) })

	// Health check
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })

	// Add urlshortener router group
	v1 := router.Group("/")
	urlshortener.UrlsRegister(v1)

	router.Run(urlshortener.AppConf.HostAddress())
}
