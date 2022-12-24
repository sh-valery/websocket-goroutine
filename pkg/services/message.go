package services

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type messageService struct{}

func NewMessageService() *messageService {
	return &messageService{}
}

func (s *messageService) GetMessageChannel() (chan string, error) {
	messageChannel := make(chan string)
	done := make(chan bool)

	// run fakeMessageGenerator in a goroutine, change to data that you want to send
	go s.fakeMessageGenerator(messageChannel, done)
	return messageChannel, nil
}

func (s *messageService) fakeMessageGenerator(messageChan chan string, done chan bool) {
	fmt.Println("goroutine started")
	messageChan <- fmt.Sprintf("Connection established at %s, waiting for a message", time.Now().String())

	for {
		select {
		// generate message every 5 seconds
		case <-time.After(5 * time.Second):
			messageID, err := uuid.NewRandom()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("send message to the channel")
			messageChan <- fmt.Sprintf(" MessageID: %s, Current time: %s", messageID.String(), time.Now().String())

		// close the message channel if received signal from done channel
		case <-done:
			fmt.Println("goroutine done")
			close(messageChan)
			return
		}
	}
}
