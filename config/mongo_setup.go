package config

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	clientInstance   *mongo.Client
	Once             sync.Once
	clientMongoError error
)

func GetMongoClient() *mongo.Client {
	Once.Do(func() {
		mongoURI := MONGO_URL
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI(mongoURI)
		clientInstance, clientMongoError = mongo.Connect(ctx, clientOptions)
		if clientError != nil {
			panic(clientError)
		}

		clientError = clientInstance.Ping(ctx, readpref.Primary())
		if clientError != nil {
			panic(clientError)
		}

		fmt.Println("Connected to MongoDB!")
	})

	return clientInstance
}
