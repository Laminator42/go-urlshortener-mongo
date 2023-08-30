package urlshortener

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Laminator42/go-urlshortener-mongo/common"
	"github.com/Laminator42/go-urlshortener-mongo/db"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UrlsRegister(router *gin.RouterGroup) {
	router.GET("/:key", redirect)
	router.POST("/shorten", shorten)
}

// shorten		godoc
// @Summary		Shorten any URL
// @Description	Generates unique short ID and writes mapping of short URL and long URL into Mongo database. Responds with short URL, expiration date and database id
// @Tags		urlshortener
// @Accept		json
// @Produce		json
// @Param		longUrl body	string	true	"LongURL to be shortened"
// @Success	201	object	shortenResponse	"Short URL, expiration date, and database ID"
// @Failure	400	object	common.ErrorResponse	"Bad request or invalid input"
// @Failure	500 object	common.ErrorResponse	"Internal server error"
// @Router		/shorten	[post]
func shorten(c *gin.Context) {
	// Validate input schema
	var body shortenBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
		return
	}

	// Validate if given URL is a valid URL
	_, urlErr := url.ParseRequestURI(body.LongUrl)
	if urlErr != nil {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: urlErr.Error()})
		return
	}

	// Generate a new id for the url
	urlKey, idErr := shortid.Generate()
	if idErr != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: idErr.Error()})
		return
	}

	// Check if already in database
	// TODO: What happens on collisions? --> should just generate a new one?
	var result bson.M
	queryErr := db.Collection.FindOne(db.Ctx, bson.D{{Key: "urlKey", Value: urlKey}}).Decode(&result)
	if queryErr != nil {
		if queryErr != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: queryErr.Error()})
			return
		}
	}

	// TODO: Comment
	if len(result) > 0 {
		c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: fmt.Sprintf("Key already in use: %s", urlKey)})
		return
	}

	var date = time.Now()
	var expires = date.AddDate(0, 0, 5)
	var newUrl = AppConf.BaseUrl() + urlKey
	var docId = primitive.NewObjectID()

	newDoc := &urlDocument{
		ID:        docId,
		UrlKey:    urlKey,
		LongUrl:   body.LongUrl,
		ShortUrl:  newUrl,
		CreatedAt: time.Now(),
		ExpiresAt: expires,
	}

	_, insertErr := db.Collection.InsertOne(db.Ctx, newDoc)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: err.Error()})
		return
	}

	// TODO: ChatGPT recommended using docId.Hex(). Reasearch this
	response := shortenResponse{
		NewUrl:  newUrl,
		Expires: expires.Format("2006-01-02 15:04:05"),
		DbID:    docId.Hex(),
	}

	c.IndentedJSON(http.StatusCreated, response)
}

// redirect		godoc
// @Summary		Follow short URL to destination
// @Description	URL mapping lookup and redirects to destination
// @Tags		urlshortener
// @Accept		plain
// @Param		key	path	string	true	"Short URL key"
// @Success	301	string	string	"Redirect to the destination URL"
// @Failure	400	object	common.ErrorResponse	"Invalid key or URL not found"
// @Failure	500 object	common.ErrorResponse	"Internal server error"
// @Router		/{key}	[get]
func redirect(c *gin.Context) {
	key := c.Param("key")
	var result bson.M
	queryErr := db.Collection.FindOne(db.Ctx, bson.D{{Key: "urlKey", Value: key}}).Decode(&result)
	if queryErr != nil {
		if queryErr == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: fmt.Sprintf("No URL with code: %s", key)})
			return
		} else {
			c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: queryErr.Error()})
			return
		}
	}
	log.Print(result["longUrl"])
	var longUrl = fmt.Sprint(result["longUrl"])
	c.Redirect(http.StatusPermanentRedirect, longUrl)
}
