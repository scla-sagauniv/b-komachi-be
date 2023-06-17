package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
  fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	http.HandleFunc("/",homepage)
	http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("main.go")
	HandleRequest()
}