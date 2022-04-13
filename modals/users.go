package modals

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID         int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName   string `json:"username,omitempty" bson:"username,omitempty"`
	Password   string `json:"password" bson:"password,omitempty"`
	FullName   string `json:"fullname" bson:"fullname,omitempty"`
	Department string `json:"department" bson:"department,omitempty"`
	PhotoURl   string `json:"photourl" bson:"photourl,omitempty"`
	Bookmarks  []int  `json:"bookmarks" bson:"bookmarks,omitempty"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collection = ConnectDB("users")
	var myUsers []Users

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var myUser Users
		err := cur.Decode(&myUser)
		if err != nil {
			log.Fatal(err)
		}
		myUsers = append(myUsers, myUser)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(myUsers)
}

func GetUserSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var user Users
	var params = mux.Vars(r)

	myID, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": myID}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var params = mux.Vars(r)

	myID, _ := strconv.Atoi(params["id"])
	myUserName, _ := params["username"]
	myPassword, _ := params["password"]
	myFullName, _ := params["fullname"]
	myDepartment, _ := params["department"]
	myPhotoUrl, _ := params["photourl"]
	myBookmarks := []int{}

	newUser := Users{
		ID:         int64(myID),
		UserName:   myUserName,
		Password:   myPassword,
		FullName:   myFullName,
		Department: myDepartment,
		PhotoURl:   myPhotoUrl,
		Bookmarks:  myBookmarks,
	}
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	result, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := strconv.Atoi(params["id"])
	myPassword, _ := (params["password"])
	var user Users

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		{"$set", bson.D{
			{"password", myPassword},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.ID = int64(id)

	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func GetUserSearch2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var user Users
	var params = mux.Vars(r)

	myUserName, _ := params["username"]
	myPassword, _ := params["password"]

	filter := bson.M{"username": myUserName, "password": myPassword}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func AddBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("users")
	var params = mux.Vars(r)

	myBookmarkIDSlices := []int{}
	paramUserID := strings.Split(params["bookmarkID"], ",")
	for _, myBookmarkID := range paramUserID {

		myBookmarks,_:=strconv.Atoi(myBookmarkID)
		myBookmarkIDSlices = append(myBookmarkIDSlices,myBookmarks)
	}
	
	
	id, _ := strconv.Atoi(params["id"])
	
	var user Users

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		{"$set", bson.D{
			{"bookmarks", myBookmarkIDSlices},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.ID = int64(id)

	json.NewEncoder(w).Encode(user)

}
