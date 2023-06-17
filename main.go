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
	e := echo.New() // インスタンスを作成
	e.Use(middleware.Logger()) // ミドルウェアを設定

	e.GET("/", func(c echo.Context) error { // ルートを設定
		return c.String(http.StatusOK, "Hello, World!") // 出力
	})

	e.Logger.Fatal(e.Start(":8080")) // サーバーをポート番号8080で起動
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
				// Client からのメッセージを読み込む
				msg := ""
				err := websocket.Message.Receive(ws, &msg)
				if err != nil {
					log.Println(fmt.Errorf("read %s", err))
					c.Logger().Error(err)
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
