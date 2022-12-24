package main

import (
	"github.com/sh-valery/websocket-goroutine/pkg/handlers"
	"github.com/sh-valery/websocket-goroutine/pkg/services"
	"log"
	"net/http"
)

func main() {
	// init services
	messageService := services.NewMessageService()

	// injection services to handlers
	server := handlers.MessageRepository{
		Service: messageService,
	}

	// init routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/example.html")
	})
	http.HandleFunc("/ws", server.HandleWebSocket)

	// run http server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
