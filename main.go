package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {}

func NewWebsocketHandler() *WebsocketHandler {
	return &WebsocketHandler{}
}

func (h *WebsocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/ws", NewWebsocketHandler().Handle)

	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}