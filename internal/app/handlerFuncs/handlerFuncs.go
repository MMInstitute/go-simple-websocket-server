package handlerFuncs

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/gorilla/websocket"
)

type Message struct {
	Command string          `json:"command"`
	Body    json.RawMessage `json:"body"`
}

type MessageBody struct {
	A int64 `json:a`
	B int64 `json:b`
}

var funcs map[string]interface{} = map[string]interface{}{
	"command1": command1,
	"command2": command2,
}

func Handler(conn *websocket.Conn, messageType int, p []byte) {
	var msg Message
	err := json.Unmarshal(p, &msg)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("msgType ", messageType)
	log.Println("msg.command ", msg.Command)
	log.Println("msg.body ", msg.Body)

	var body MessageBody
	err = json.Unmarshal(msg.Body, &body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("parsed msg.body ", body)

	f := reflect.ValueOf(funcs[msg.Command])
	if f == reflect.ValueOf(nil) {
		log.Println(("No Func"))
		return
	}
	params := make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(conn)
	params[1] = reflect.ValueOf(body)

	f.Call(params)
}

func command1(conn *websocket.Conn, body MessageBody) {
	log.Println(body)
	reply, err := json.Marshal(body)
	if err != nil {
		log.Println("Failed to body to json")
	}
	err = conn.WriteJSON("command 1 return is : " + string(reply))
	if err != nil {
		log.Println("Failed to send response.")
	}
}

func command2(conn *websocket.Conn, body MessageBody) {
	log.Println(body)
	reply, err := json.Marshal(body)
	if err != nil {
		log.Println("Failed to body to json")
	}
	err = conn.WriteJSON("command 2 return is : " + string(reply))
	if err != nil {
		log.Println("Failed to send response.")
	}
}
