package main

import (
	"fmt"
	"log"
	"net/http"
)

type HttpAddresses struct {
	Adress string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/getUsers/", getUsers)
	http.HandleFunc("/createUsers/", createUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created User", r.URL.Path[1:])
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getUsers", r.URL.Path[1:])
}
