package urlshortener

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type shortenBody struct {
	LongUrl string `json:"longUrl"`
}

type urlDocument struct {
	ID        primitive.ObjectID `bson:"_id"`
	UrlKey    string             `bson:"urlKey"`
	LongUrl   string             `bson:"longUrl"`
	ShortUrl  string             `bson:"shortUrl"`
	CreatedAt time.Time          `bson:"createdAt"`
	ExpiresAt time.Time          `bson:"expiresAt"`
}
