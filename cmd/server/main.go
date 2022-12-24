package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sh-valery/websocket-goroutine/pkg/services"
	"log"
	"net/http"
	"time"
)

type MessageService interface {
	GenerateMessage() (string, error)
}

var MessagePublisher MessageService

func main() {
	// init services
	MessagePublisher = services.NewMessageService()

	// init routing
	http.HandleFunc("/", serveStatic)
	http.HandleFunc("/ws", handleWebSocket)

	// run http server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/example.html")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade http connection to web socket connection
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// send messages to the websocket
	for {
		message, err := MessagePublisher.GenerateMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println(err)
		}
		MessageDelay := 3 * time.Second
		time.Sleep(MessageDelay)
	}
}
