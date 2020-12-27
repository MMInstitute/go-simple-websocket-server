package main

import (
	"log"
	"message-server/internal/app/handlerFuncs"
	"net/http"

	"github.com/gorilla/websocket"
)

func checkOriginFunc(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket Server try to be listening port 8082.")
	log.Fatal(http.ListenAndServe(":8082", nil))

}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		handlerFuncs.Handler(conn, messageType, p)
	}
}
