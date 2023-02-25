package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// sending a message every second
	for {
		msg := message{
			Type: "transaction",
			Data: "0x123456789",
		}
		err = conn.WriteJSON(msg)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println("Transaction sent to client")
		time.Sleep(time.Second)
	}
}
