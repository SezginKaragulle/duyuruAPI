package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	// http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/view/", viewHandler)
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
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}
