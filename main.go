package main

import (
	"fmt"
	"log"
	"net/http"

	"app/b-komachi-be/src/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Home!")
  fmt.Println("Endpoint Hit: homePage")
}
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", websocket.Endpoint)

	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}