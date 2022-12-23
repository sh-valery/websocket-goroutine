package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type MessageService interface {
	GenerateMessage() (string, error)
}

type messageService struct{}

func (s *messageService) GenerateMessage() (string, error) {
	t := time.Now()
	MessageID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// fill in data that you want to send
	return fmt.Sprintf("Current time: %s, MessageID: %s", t.String(), MessageID.String()), nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	service := &messageService{}
	for {
		message, err := service.GenerateMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.WriteMessage(websocket.TextMessage, []byte(message))
		time.Sleep(3 * time.Second)
	}
}
