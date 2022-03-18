package main

import (
	"fmt"
	"golesson/modals"
	"net/http"
)

func main() {
	// http.HandleFunc("/getUsers/", getUsers)
	// http.HandleFunc("/createUsers/", createUser)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// modals.ConnectMongoDB()
	// modals.AddUsers("enes","35789654")
	modals.ConnectMongoDB()

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created User", r.URL.Path[1:])
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getUsers", r.URL.Path[1:])
}
