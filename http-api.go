package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest() {
	http.HandleFunc("/homepage", homepage)
	log.Fatal(http.ListenAndServe(":7276", nil))

}

func main() {
	handleRequest()
}
