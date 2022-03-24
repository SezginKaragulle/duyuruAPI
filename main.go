package main

import (
	"golesson/modals"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/users", modals.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", modals.GetUserSearch).Methods("GET")
	r.HandleFunc("/api/users/createUser", modals.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/update/{id}", modals.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/delete/{id}", modals.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
