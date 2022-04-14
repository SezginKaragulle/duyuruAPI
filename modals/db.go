package modals

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseURL = "mongodb+srv://solmaz:OmdYWoX8myGOcvEL@cluster0.iup4a.mongodb.net/duyuruDB?retryWrites=true&w=majority"

//var collection = ConnectDB("users")

func ConnectDB(collection_name string) *mongo.Collection {

	clientOptions := options.Client().ApplyURI(databaseURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("duyuruDB").Collection(collection_name)

	return collection
}
