package main

import (
	"golesson/modals"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	//Users Collection

	r.HandleFunc("/api/users/", modals.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/searchUser/{id}", modals.GetUserSearch).Methods("GET")
	r.HandleFunc("/api/users/create/{id}&{username}&{password}&{fullname}&{department}&{photourl}", modals.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/update/{id}&{password}", modals.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/delete/{id}", modals.DeleteUser).Methods("DELETE")
	r.HandleFunc("/api/users/searchUserPass/{username}&{password}", modals.GetUserSearch2).Methods("GET")

	//Groups Collection

	r.HandleFunc("/api/groups/", modals.GetGroups).Methods("GET")
	r.HandleFunc("/api/groups/create/{id}&{createrID}&{name}&{memberID}&{userID}&{groupID}", modals.CreateGroup).Methods("POST")
	r.HandleFunc("/api/groups/delete/{id}", modals.DeleteGroups).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
