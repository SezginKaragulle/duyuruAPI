package modals

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Users struct {
	UserName string
	Password string
}

const databaseURL = "mongodb+srv://solmaz:tlV37y9TbAxxGXlF@cluster0.iup4a.mongodb.net/duyuruDB?retryWrites=true&w=majority"

//MongoDB bağlantısını test amaçlı fonksiyondur.
func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI(databaseURL)
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

//users collectionundaki verileri çekmek için çalışan bir fonksiyondur.
func UsersList() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	users := client.Database("duyuruDB").Collection("users")
	cur, err := users.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	defer cur.Close(ctx)
	var myUserList []Users
	for cur.Next(ctx) {
		var result Users
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Error : " + err.Error())
		}
		myUserList = append(myUserList, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal("Error : " + err.Error())
	}
	fmt.Println(myUserList)
}

//Users collectiona kullanıcı ekleme için çalışan bir fonksiyondur.
func UsersAdd(userName string, password string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	users := client.Database("duyuruDB").Collection("users")
	newUser := Users{
		UserName: userName,
		Password: password,
	}
	res, err := users.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	id := res.InsertedID
	fmt.Println("User Added... ", id)
}

//MongoDB ' de users collection içerisinde userName ile aratmak için çalışan bir fonksiyondur.
func UsersSearch(userName string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	users := client.Database("duyuruDB").Collection("users")
	var result Users
	err = users.FindOne(ctx, bson.M{"username": userName}).Decode(&result)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	out, err := json.Marshal(&result)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	fmt.Println(string(out))
}

//Users collection içerisinde kullanıcı güncelleme işlemi için çalışan bir fonksiyon.
func UsersUpdate(_userName string, _password string, _newUserName string, _newPassWord string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	users := client.Database("duyuruDB").Collection("users")
	changingUser := Users{
		UserName: _userName,
		Password: _password,
	}
	newUserUpdate := Users{
		UserName: _newUserName,
		Password: _newPassWord,
	}
	var myFilter bson.M
	bytes, err := bson.Marshal(changingUser)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	bson.Unmarshal(bytes, &myFilter)
	var usr bson.M
	bytes, err = bson.Marshal(newUserUpdate)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	bson.Unmarshal(bytes, &usr)
	update := bson.D{
		{"$set", usr},
	}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	_, err = users.UpdateOne(ctx, myFilter, update)
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	fmt.Println("User Updated...")

}

//Users collection içerisinde kullanıcı silme işlemi için çalışan bir fonksiyon.
func UsersDelete(_userName string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURL))
	users := client.Database("duyuruDB").Collection("users")
	_, err = users.DeleteOne(ctx, bson.D{{"username", _userName}})
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}
	fmt.Println("User Deleted...")
}
