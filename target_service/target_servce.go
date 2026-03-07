package targetservice

import (
	"context"
	"errors"
	"os"
)

type ITargetService interface {
	GetSQSUrlForBot(ctx context.Context) (string, error)
}

type targetService struct {
}

func NewTargetService() ITargetService {
	return &targetService{}
}

func (t *targetService) GetSQSUrlForBot(ctx context.Context) (string, error) {
	queueURL := os.Getenv("SQS_QUEUE_URL_BOT")
	if queueURL == "" {
		return "", errors.New("SQS_QUEUE_URL_BOT is not set")
	}
	return queueURL, nil
}
