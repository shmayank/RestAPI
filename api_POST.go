package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func addArticle(w http.ResponseWriter, r *http.Request) {

	var newItem Article
	json.NewDecoder(r.Body).Decode(&newItem)
	Articles = append(Articles, newItem)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Articles)
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/add", addArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description1", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description2", Content: "Article Content"},
	}
	handleRequests()
}
