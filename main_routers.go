package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description1", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description2", Content: "Article Content"},
	}
	handleRequests()
}
