package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func Init() {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		conf.username,
		conf.password,
		conf.mongoHost,
		conf.mongoPort,
		conf.mongoDb,
	)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB.", err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal("Failed to reach MongoDB.", err)
	}

	// TODO: Use this query if not executed? db.urls.createIndex( { "expiresAt": 1 }, { expireAfterSeconds: 0 } )
	Collection = client.Database(conf.mongoDb).Collection("urlmapping")
	log.Print("Database connected.")
}
