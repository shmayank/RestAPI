package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)

}

func main() {
	handleRequest()
}
