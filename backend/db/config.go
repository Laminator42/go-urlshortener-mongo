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
	mongoHost: util.GetEnv("BACKEND_DB_HOST", "mongodb"),
	mongoPort: util.GetEnv("BACKEND_DB_PORT", "27017"),
	mongoDb:   util.GetEnv("BACKEND_DB_DATABASE", "urlshortener"),
	username:  util.GetEnv("BACKEND_DB_USERNAME", "urluser"),
	password:  util.GetEnv("BACKEND_DB_PASSWORD", "urlpassword"),
}
