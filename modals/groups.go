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
	GroupMembers []string      `json:"groupMembers" bson:"groupMembers,omitempty"`
}

type GroupMembers struct {
	ID      int64    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  []string `json:"userID,omitempty" bson:"userID,omitempty"`
	GroupID int64    `json:"groupID,omitempty" bson:"groupID,omitempty"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
	var params = mux.Vars(r)

	//For Groups
	myID, _ := strconv.Atoi(params["id"])
	myCreaterID, _ := strconv.Atoi(params["createrID"])
	myName, _ := params["name"]
	myUserIDSlices := []string{params["userID"]}
	


	newGroups := Groups{
		ID:           int64(myID),
		CreaterID:    int64(myCreaterID),
		Name:         myName,
		GroupMembers: myUserIDSlices,
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

func DeleteGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
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

func AddMembersToGroup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := strconv.Atoi(params["id"])
	myUserIDSlices := []string{params["userID"]}
	var user Users

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		{"$set", bson.D{
			{"groupMembers", myUserIDSlices},
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

func GetGroupSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var collection = ConnectDB("groups")
	var group Groups
	var params = mux.Vars(r)

	myID, _ := strconv.Atoi(params["id"])

	filter := bson.M{"_id": myID}
	err := collection.FindOne(context.TODO(), filter).Decode(&group)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func GetGroupMemberSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var collection = ConnectDB("groups")
	var myGroups []Groups
	var params = mux.Vars(r)
	myUserIDSlices := []string{params["userID"]}
	cur, err := collection.Find(context.TODO(), bson.M{"groupMembers": bson.M{"$in": myUserIDSlices}})

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