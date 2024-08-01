package repository

import (
	"context"
	"lexium-utility/config"
	"lexium-utility/datahandler"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetScheamList(itr string) ([]string, error) {

	client := config.GetMongoClient()
	collection := client.Database("Lexium").Collection("Schema_" + itr)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	fieldName := "schemaName"

	// Perform distinct query to get unique schemaName values
	distinctValues, err := collection.Distinct(ctx, fieldName, filter)
	if err != nil {
		return nil, err
	}

	// Convert distinctValues to a slice of strings
	var schemaNames []string
	for _, value := range distinctValues {
		if schemaName, ok := value.(string); ok {
			schemaNames = append(schemaNames, schemaName)
		}
	}
	return schemaNames, nil
}

func ReadCollection(itr, schema string) ([]datahandler.Element, error) {
	client := config.GetMongoClient()
	collection := client.Database("Lexium").Collection("Elements_" + itr + "_" + schema)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find all documents
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var Elements []datahandler.Element
	for cursor.Next(ctx) {
		var element datahandler.Element
		err := cursor.Decode(&element)
		if err != nil {
			return nil, err
		}
		Elements = append(Elements, element)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return Elements, nil
}
func ReadElement(itr, schema, elementId string) (*datahandler.Element, error) {
	client := config.GetMongoClient()
	collection := client.Database("Lexium").Collection("Elements_" + itr + "_" + schema)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var element datahandler.Element
	err := collection.FindOne(ctx, bson.D{{"elementID", elementId}}).Decode(&element)
	if err != nil {
		return nil, err
	}
	return &element, nil
}
