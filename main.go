package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("main.go")
	HandleRequest()
}

func handleWebSocket(c echo.Context) error {
	log.Println("Serving at localhost:8080...")
	websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			// 初回のメッセージを送信
			err := websocket.Message.Send(ws, "Server: Hello, Next.js!")
			if err != nil {
				c.Logger().Error(err)
			}

	for {
			// メッセージを受信する
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
					log.Fatalln(err)
			}

				// Client からのメッセージを元に返すメッセージを作成し送信する
				err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
				if err != nil {
					c.Logger().Error(err)
				}
			}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
