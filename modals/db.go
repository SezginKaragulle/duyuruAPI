package modals

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
	UserName string
	Password string
}

const uri = "mongodb+srv://solmaz:tlV37y9TbAxxGXlF@cluster0.iup4a.mongodb.net/duyuruDB?retryWrites=true&w=majority"

func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

func AddUsers(userName string, userPassword string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("duyuruDB").Collection("users")
	ruan := Users{userName, userPassword}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Eklendi: ", insertResult.InsertedID)

}
