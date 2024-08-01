package repository

import (
	"context"
	"lexium-utility/config"
	"lexium-utility/datahandler"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByUsername(username string) (datahandler.User, error) {
	var user datahandler.User

	client := config.GetMongoClient()
	collection := client.Database("Lexium").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"username": username}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, nil // No matching document found
		}
		return user, err // An error occurred while retrieving the document
	}

	return user, nil
}
