package db

import (
	"fmt"

	"github.com/Laminator42/go-urlshortener-mongo/util"
)

type MongoConfig struct {
	mongoHost string
	mongoPort string
	mongoDb   string
	username  string
	password  string
}

func (c MongoConfig) connectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		conf.username,
		conf.password,
		conf.mongoHost,
		conf.mongoPort,
		conf.mongoDb,
	)
}

var conf = MongoConfig{
	mongoHost: util.GetEnv("MONGO_HOST", "10.100.253.55"),
	mongoPort: util.GetEnv("MONGO_PORT", "27017"),
	mongoDb:   util.GetEnv("MONGO_DB", "urlshortener"),
	username:  util.GetEnv("MONGO_USER", "urluser"),
	password:  util.GetEnv("MONGO_PASS", "urlpassword"),
}
