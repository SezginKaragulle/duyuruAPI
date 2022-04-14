package modals

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Posts struct {
	ID           int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	AuthorID     int64  `json:"authorID,omitempty" bson:"authorID,omitempty"`
	Topic        string `json:"topic" bson:"topic,omitempty"`
	Content      string `json:"content" bson:"content,omitempty"`
	Publish_Time string `json:"publishTime" bson:"publishTime,omitempty"`
	Receivers    []int  `json:"receivers" bson:"receivers,omitempty"`
	Seen_Users   []int  `json:"seenUsers" bson:"seenUsers,omitempty"`
}

var collection_Posts = ConnectDB("posts")

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)
	myID, _ := strconv.Atoi(params["id"])
	myAuthorID, _ := strconv.Atoi(params["authorID"])
	myTopic, _ := params["topic"]
	myContent, _ := params["content"]
	myPublishTime := time.Now().Format("2006-01-02 15:04:05")
	myReceivers := []int{}
	mySeenUsers := []int{}
	paramReceiversID := strings.Split(params["receivers"], ",")

	for _, myReceiverID := range paramReceiversID {

		myReceiver, _ := strconv.Atoi(myReceiverID)
		myReceivers = append(myReceivers, myReceiver)
	}

	newPosts := Posts{
		ID:           int64(myID),
		AuthorID:     int64(myAuthorID),
		Topic:        myTopic,
		Content:      myContent,
		Publish_Time: myPublishTime,
		Receivers:    myReceivers,
		Seen_Users:   mySeenUsers,
	}
	
	_ = json.NewDecoder(r.Body).Decode(&newPosts)

	result, err := collection_Posts.InsertOne(context.TODO(), newPosts)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var myPosts []Posts
	cur, err := collection_Posts.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var myPost Posts
		err := cur.Decode(&myPost)
		if err != nil {
			log.Fatal(err)
		}
		myPosts = append(myPosts, myPost)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(myPosts)
}

func GetPostReceivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var myPosts []Posts
	var params = mux.Vars(r)
	myReceivers := []int{}
	paramReceiversID := strings.Split(params["receivers"], ",")

	for _, myReceiverID := range paramReceiversID {

		myReceiver, _ := strconv.Atoi(myReceiverID)
		myReceivers = append(myReceivers, myReceiver)
	}

	cur, err := collection_Posts.Find(context.TODO(), bson.M{"receivers": bson.M{"$in": myReceivers}})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var myPost Posts
		err := cur.Decode(&myPost)
		if err != nil {
			log.Fatal(err)
		}
		myPosts = append(myPosts, myPost)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(myPosts)
}

func ArrangeSeenUsersOfPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)
	var post Posts
	id, _ := strconv.Atoi(params["id"])
	mySeenUsers := []int{}
	paramSeenUserID := strings.Split(params["seenUsers"], ",")

	for _, mySeenUser := range paramSeenUserID {

		mySeenUsersList, _ := strconv.Atoi(mySeenUser)
		mySeenUsers = append(mySeenUsers, mySeenUsersList)
	}

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&post)

	update := bson.D{
		{"$set", bson.D{
			{"seenUsers", mySeenUsers},
		}},
	}

	err := collection_Posts.FindOneAndUpdate(context.TODO(), filter, update).Decode(&post)

	if err != nil {
		log.Fatal(err)
		return
	}

	post.ID = int64(id)

	json.NewEncoder(w).Encode(post)
}

func DeletePosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	filter := bson.M{"_id": id}
	deleteResult, err := collection_Posts.DeleteOne(context.TODO(), filter)
	
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}