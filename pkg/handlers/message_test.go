package handlers

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/sh-valery/websocket-goroutine/pkg/services/mocks"
	"net/http"
	"testing"
)

func TestMessageRepository_HandleWebSocket(t *testing.T) {
	// prepare mock
	mockService := mocks.NewMockMessageService(gomock.NewController(t))
	messagePipe := make(chan string, 1) // non blocking call
	testMessage := "test message"
	messagePipe <- testMessage
	mockService.EXPECT().GetMessageChannel().Return(messagePipe, make(chan bool))
	server := MessageRepository{
		Service: mockService,
	}

	// run server
	http.HandleFunc("/ws", server.HandleWebSocket)
	go http.ListenAndServe(":8080", nil) // nolint:errcheck

	// connect to server
	ws, resp, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		t.Error(err)
	}

	// validate status code
	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Errorf("unexpected status code, want: %d got %d", http.StatusSwitchingProtocols, resp.StatusCode)
	}

	// validate message
	_, message, _ := ws.ReadMessage()
	if string(message) != testMessage {
		t.Error("message not equal, want: ", testMessage, " got: ", message)
	}
}
