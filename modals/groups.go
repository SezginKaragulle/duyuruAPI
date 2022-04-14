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
)

type Groups struct {
	ID           int64  `json:"_id,omitempty" bson:"_id,omitempty"`
	CreaterID    int64  `json:"createrID,omitempty" bson:"createrID,omitempty"`
	Name         string `json:"name" bson:"name,omitempty"`
	GroupMembers []int  `json:"groupMembers" bson:"groupMembers,omitempty"`
}

var collection_Groups = ConnectDB("groups")

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)

	myID, _ := strconv.Atoi(params["id"])
	myCreaterID, _ := strconv.Atoi(params["createrID"])
	myName, _ := params["name"]
	myUserIDSlices := []int{}
	paramUserID := strings.Split(params["userID"], ",")

	for _, myUserID := range paramUserID {

		myUsers, _ := strconv.Atoi(myUserID)
		myUserIDSlices = append(myUserIDSlices, myUsers)
	}

	newGroups := Groups{
		ID:           int64(myID),
		CreaterID:    int64(myCreaterID),
		Name:         myName,
		GroupMembers: myUserIDSlices,
	}
	_ = json.NewDecoder(r.Body).Decode(&newGroups)

	result, err := collection_Groups.InsertOne(context.TODO(), newGroups)
	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var myGroups []Groups
	cur, err := collection_Groups.Find(context.TODO(), bson.M{})

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

	
	var params = mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	filter := bson.M{"_id": id}
	deleteResult, err := collection_Groups.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func ArrangeMembersOfGroup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)
	var user Users
	id, _ := strconv.Atoi(params["id"])
	myUserIDSlices := []int{}
	paramUserID := strings.Split(params["userID"], ",")

	for _, myUserID := range paramUserID {

		myUserIDList, _ := strconv.Atoi(myUserID)
		myUserIDSlices = append(myUserIDSlices, myUserIDList)
	}

	filter := bson.M{"_id": id}
	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		{"$set", bson.D{
			{"groupMembers", myUserIDSlices},
		}},
	}

	err := collection_Groups.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.ID = int64(id)

	json.NewEncoder(w).Encode(user)

}

func GetGroupSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	var myGroups []Groups
	var params = mux.Vars(r)
	myID, _ := strconv.Atoi(params["id"])
	cur, err := collection_Groups.Find(context.TODO(), bson.M{"_id": myID})

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

func GetGroupMemberSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var myGroups []Groups
	var params = mux.Vars(r)
	myUserIDSlices := []int{}
	paramUserID := strings.Split(params["userID"], ",")

	for _, myUserID := range paramUserID {

		myUserIDList, _ := strconv.Atoi(myUserID)
		myUserIDSlices = append(myUserIDSlices, myUserIDList)
	}
	
	cur, err := collection_Groups.Find(context.TODO(), bson.M{"groupMembers": bson.M{"$in": myUserIDSlices}})

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
