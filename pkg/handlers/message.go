package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//go:generate mockgen -source=message.go -destination=../services/mocks/message.go -package=mocks
type MessageService interface {
	GetMessageChannel() (chan string, chan bool)
}

type MessageRepository struct {
	Service MessageService
}

func (h *MessageRepository) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
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
	messageChan, done := h.Service.GetMessageChannel()

	for m := range messageChan {
		err = conn.WriteMessage(websocket.TextMessage, []byte(m))
		if err != nil { // client disconnected
			fmt.Println(err)
			done <- true
		}
	}
}
