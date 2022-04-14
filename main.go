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
	r.HandleFunc("/api/users/bookmarkAdd/{id}&{bookmarkID}", modals.AddBookmarks).Methods("PUT")

	//Groups Collection

	r.HandleFunc("/api/groups/", modals.GetGroups).Methods("GET")
	r.HandleFunc("/api/groups/create/{id}&{createrID}&{name}&{userID}", modals.CreateGroup).Methods("POST")
	r.HandleFunc("/api/groups/delete/{id}", modals.DeleteGroups).Methods("DELETE")
	r.HandleFunc("/api/groups/arrangeMembers/{id}&{userID}", modals.ArrangeMembersOfGroup).Methods("PUT")
	r.HandleFunc("/api/groups/searchGroup/{id}", modals.GetGroupSearch).Methods("GET")
	r.HandleFunc("/api/groups/searchGroupMember/{userID}", modals.GetGroupMemberSearch).Methods("GET")

	//Posts Collections

	r.HandleFunc("/api/posts/", modals.GetPosts).Methods("GET")
	r.HandleFunc("/api/posts/create/{id}&{authorID}&{topic}&{content}&{receivers}", modals.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts/postReceivers/{receivers}", modals.GetPostReceivers).Methods("GET")
	r.HandleFunc("/api/posts/arrangeSeenUsers/{id}&{seenUsers}", modals.ArrangeSeenUsersOfPosts).Methods("PUT")
	r.HandleFunc("/api/posts/delete/{id}", modals.DeletePosts).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
