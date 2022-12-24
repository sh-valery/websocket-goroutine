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

func (s *messageService) GenerateMessage() (string, error) {
	t := time.Now()
	MessageID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// fill in data that you want to send
	return fmt.Sprintf("Current time: %s, MessageID: %s", t.String(), MessageID.String()), nil
}
