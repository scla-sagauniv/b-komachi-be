package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WsJsonResponse struct {
	Message string `json:"message"`
}
// コネクション保持
type WebSocketConnection struct {
	*websocket.Conn
}
// ペイロード保持
type WsPayload struct {
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}

var (
	// ペイロードチャネルを作成
	wsChan = make(chan WsPayload)

	// keyはコネクション情報, valueにはユーザー名を入れる
	clients = make(map[WebSocketConnection]string)
)
// 受け取ったmessageをchannelに送信
func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()
	
	var payload WsPayload
	
	for {
		log.Println("listen message!")
		err := conn.ReadJSON(&payload)
		log.Println(payload.Message)
		log.Println(err)
		if err == nil {
			payload.Conn = *conn
			wsChan <- payload
			log.Println("send to goroutine1")
		}
	}
}
// channelをlistenしてブロードキャスト
func ListenToWsChannel() {
	
	var response WsJsonResponse
	for {
		e := <-wsChan
		log.Println("send to channel")
		response.Message = e.Message
		broadcastToAll(response)
	}
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(&response.Message)
		log.Println("send to goroutine2")
		if err != nil {
			_ = client.Close()
			delete(clients, client)
		}
	}
}
// コネクションの設定
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
// エンドポイントの設定
func Endpoint(w http.ResponseWriter, r *http.Request){
	// 受け取ったリクエストをWebソケット用のリクエストにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)

	// コネクション情報を格納
	conn := WebSocketConnection{Conn: ws}
	// ブラウザが読み込まれた時に一度だけ呼び出される
	clients[conn] = "user1"

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Connecting")

	// goroutineで呼び出し
	go ListenForWs(&conn) 
	go ListenToWsChannel()
}
