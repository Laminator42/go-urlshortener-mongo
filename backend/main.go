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
// @BasePath		/
func main() {

	db.Init()

	router := gin.Default()

	// Add Swagger UI
	router.GET("/docs/", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/docs/index.html") })
	router.GET("/docs/:any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add urlshortener router group
	v1 := router.Group("/")
	urlshortener.UrlsRegister(v1)

	// Health check
	router.GET("/health", health)

	router.Run(urlshortener.AppConf.HostAddress())
}

type healthResponse struct {
	Status string `json:"status"`
}

// health		godoc
// @Summary		Health check
// @Description	Is the application up?
// @Tags		main
// @Accept		plain
// @Success	200	object	healthResponse	"ok"
// @Router		/health	[get]
func health(c *gin.Context) {
	c.JSON(http.StatusOK, healthResponse{Status: "ok"})
}
