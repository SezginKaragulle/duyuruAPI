package modals

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Groups struct {
	ID           int64         `json:"_id,omitempty" bson:"_id,omitempty"`
	CreaterID    int64         `json:"createrID,omitempty" bson:"createrID,omitempty"`
	Name         string        `json:"name" bson:"name,omitempty"`
	GroupMembers *GroupMembers `json:"groupMembers" bson:"groupMembers,omitempty"`
}

type GroupMembers struct {
	ID      int64 `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  int64 `json:"userID,omitempty" bson:"userID,omitempty"`
	GroupID int64 `json:"groupID,omitempty" bson:"groupID,omitempty"`
}

type GroupMembers2 struct {
	ID      int64 `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  int64 `json:"userID,omitempty" bson:"userID,omitempty"`
	GroupID int64 `json:"groupID,omitempty" bson:"groupID,omitempty"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
	var params = mux.Vars(r)

	myID, _ := strconv.Atoi(params["id"])
	myCreaterID, _ := strconv.Atoi(params["createrID"])
	myName, _ := params["name"]

	myGroupMemberID, _ := strconv.Atoi(params["memberID"])
	myUserID, _ := strconv.Atoi(params["userID"])
	myGroupID, _ := strconv.Atoi(params["groupID"])

	newMyGroupMembers := GroupMembers{
		ID:      int64(myGroupMemberID),
		UserID:  int64(myUserID),
		GroupID: int64(myGroupID),
	}

	newGroups := Groups{
		ID:           int64(myID),
		CreaterID:    int64(myCreaterID),
		Name:         myName,
		GroupMembers: &newMyGroupMembers,
	}
	_ = json.NewDecoder(r.Body).Decode(&newGroups)

	result, err := collection.InsertOne(context.TODO(), newGroups)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collection = ConnectDB("groups")
	var myGroups []Groups

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var myGroup Groups
		err := cur.Decode(&myGroup)
		if err != nil {
			log.Fatal(err)
		}
		myGroups = append(myGroups, myGroup)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(myGroups)
}

//CreateGroup2 for test

func CreateGroup2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
	var params = mux.Vars(r)

	myID, _ := strconv.Atoi(params["id"])
	myCreaterID, _ := strconv.Atoi(params["createrID"])
	myName, _ := params["name"]

	myGroupMemberID, _ := strconv.Atoi(params["memberID"])
	myUserID, _ := strconv.Atoi(params["userID"])
	myGroupID, _ := strconv.Atoi(params["groupID"])

	newMyGroupMembers := GroupMembers{
		ID:      int64(myGroupMemberID),
		UserID:  int64(myUserID),
		GroupID: int64(myGroupID),
	}

	newGroups := Groups{
		ID:           int64(myID),
		CreaterID:    int64(myCreaterID),
		Name:         myName,
		GroupMembers: &newMyGroupMembers,
	}
	_ = json.NewDecoder(r.Body).Decode(&newGroups)

	result, err := collection.InsertOne(context.TODO(), newGroups)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}
