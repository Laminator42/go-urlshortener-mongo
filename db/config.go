package db

import "github.com/Laminator42/go-urlshortener-mongo/util"

type MongoConfig struct {
	mongoHost string
	mongoPort string
	mongoDb   string
	username  string
	password  string
}

var conf = MongoConfig{
	mongoHost: util.GetEnv("MONGO_HOST", "10.100.253.55"),
	mongoPort: util.GetEnv("MONGO_PORT", "27017"),
	mongoDb:   util.GetEnv("MONGO_DB", "urlshortener"),
	username:  util.GetEnv("MONGO_USER", "urluser"),
	password:  util.GetEnv("MONGO_PASS", "urlpassword"),
}
