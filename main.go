package main

import (
	"fmt"
	"net/http"
	"log"
	"golang.org/x/net/websocket"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
  fmt.Println("Endpoint Hit: homePage")
}

func HandleRequest() {
	http.HandleFunc("/",homepage)
	http.Handle("/ws", websocket.Handler(msgHandler))
	http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("main.go")
	HandleRequest()
}
func msgHandler(ws *websocket.Conn) {
	defer ws.Close()

	// 初回のメッセージを送信
	err := websocket.Message.Send(ws, "こんにちは！ :)")
	if err != nil {
			log.Fatalln(err)
	}

	for {
			// メッセージを受信する
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
					log.Fatalln(err)
			}

			// メッセージを返信する
			err := websocket.Message.Send(ws, fmt.Sprintf(`%q というメッセージを受け取りました。`, msg))
			if err != nil {
					log.Fatalln(err)
			}
	}
}