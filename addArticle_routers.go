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
var data []string = []string{}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func addArticles(w http.ResponseWriter, r *http.Request) {
	routeVariable := mux.Vars(r)["items"]
	data = append(data, routeVariable)

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/homepage", homepage)
	myRouter.HandleFunc("/articles/{items}", addArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description1", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description2", Content: "Article Content"},
	}
	handleRequests()
}
